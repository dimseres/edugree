<?php

namespace App\Integration\Controllers;

use App\Integration\Requests\CreateOrganizationRequest;
use App\Models\Organization;
use App\Models\Owner;
use App\Models\User;
use App\Service\Tenant\TenantInitService;
use Illuminate\Support\Facades\DB;

class OrganizationController
{
    public function createOrganization(CreateOrganizationRequest $request, TenantInitService $service)
    {
        $orgName = $request->input('name');
        $orgDomain = $request->input('domain');
        $orgUuid = str_replace('-', '', $request->input('tenant_uuid'));
        $orgId = $request->input('id');

        $owner = $request->input('owner');

        $dbName = "tenant_".$orgUuid;

        try {
            $serviceOrganization = Organization::query()
                ->where("name", $orgName)
                ->where("domain", $orgDomain)
                ->where("db_name", $dbName)
                ->first();

            if ($serviceOrganization) {
                return [
                    "error" => true,
                    "message" => "organization already exist"
                ];
            }

            try {
                DB::beginTransaction();
                $user = Owner::query()->create($owner);
                $organization = Organization::query()->create([
                    'id' => $orgId,
                    'owner_id' => $user->id,
                    'name' => $orgName,
                    'domain' => $orgDomain,
                    'db_name' => $dbName
                ]);
                DB::commit();
            } catch (\Exception $exception) {
                DB::rollBack();
                return [
                    "error" => true,
                    "message" => $exception->getMessage()
                ];
            }

            DB::statement("CREATE DATABASE {$dbName}");
            $service->switchConnection($dbName);
            $service->runMigration($dbName);
            $service->runSeeders();
            $service->setInitialData($owner, $organization);

            return [
                "error" => false,
                "message" => 'created',
            ];
        } catch (\Exception $exception) {
            $service->switchConnection("forge");
            DB::statement("DROP DATABASE {$dbName}");
            Organization::query()->where('id', $orgId)->delete();
            return [
                "error" => true,
                "message" => $exception->getMessage()
            ];
        }

    }

    public function createTenantUser(CreateOrganizationRequest $request, TenantInitService $service)
    {
        $orgDomain = $request->input('domain');
        $dbName = "tenant_".$orgDomain;
        $userData = $request->input('user');

        $service->switchConnection($dbName);
//
        try {
            DB::beginTransaction();
            $user = User::query()->create([
                'id' => $userData['id'],
                'name' => $userData['name'],
                'email' => $userData['email'],
                'phone' => $userData['phone'],
            ]);


        } catch (\Exception $exception) {
            DB::rollBack();
            return [
                'error' => true,
                'message' => $exception->getMessage()
            ];
        }
    }
}
