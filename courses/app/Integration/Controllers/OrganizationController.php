<?php

namespace App\Integration\Controllers;

use App\Integration\Requests\CreateOrganizationRequest;
use App\Integration\Requests\CreateTenantUserRequest;
use App\Models\Organization;
use App\Models\Owner;
use App\Models\Role;
use App\Models\User;
use App\Service\Tenant\TenantInitService;
use Illuminate\Support\Facades\DB;

class OrganizationController
{
    public function createOrganization(CreateOrganizationRequest $request, TenantInitService $service)
    {
        $orgName = $request->input('name');
        $orgDomain = $request->input('domain');
        $tenantUuid = $request->input('tenant_uuid');
        $orgId = $request->input('id');

        $owner = $request->input('user');

        $dbName = $service->formatTenantDbName($tenantUuid);

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
                $user = Owner::query()->updateOrCreate([
                    'id' => $owner['id']
                ], $owner);
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
                return response()->json([
                    "error" => true,
                    "message" => $exception->getMessage()
                ], 500);
            }

            DB::statement("CREATE DATABASE {$dbName}");
            $service->switchConnection($dbName);
            $service->runMigration($dbName);
            $service->runSeeders();
            $service->setInitialData($owner, $organization);

            return response()->json([
                "error" => false,
                "message" => 'created',
            ]);
        } catch (\Exception $exception) {
            DB::reconnect(env('DB_CONNECTION'));
            DB::statement("DROP DATABASE {$dbName}");
            Organization::query()->where('id', $orgId)->delete();
            return response()->json([
                "error" => true,
                "message" => $exception->getMessage()
            ], 500);
        }

    }

    public function createTenantUser(CreateTenantUserRequest $request, TenantInitService $service)
    {
        $orgDomain = $request->input('domain');
        $tenantUuid = $request->input('tenant_uuid');
        $dbName = $service->formatTenantDbName($tenantUuid);
        $userData = $request->input('user');

        $service->switchConnection($dbName);
//
        try {
            DB::beginTransaction();
            $user = User::query()->updateOrCreate([
                'id' => $userData['id'],
                'name' => $userData['name'],
                'email' => $userData['email'],
                'phone' => $userData['phone'],
            ]);

            $role = Role::query()->where('name', $userData['role'])->first();
            $user->assignRole($role);


            DB::commit();
            return response()->json([
                'error' => false,
                'data' => $user
            ]);
        } catch (\Exception $exception) {
            DB::rollBack();
            return response()->json([
                'error' => true,
                'message' => $exception->getMessage()
            ], 500);
        }
    }
}
