<?php

namespace App\Service\Tenant;

use App\Models\Organization;
use App\Models\Owner;
use App\Models\User;
use Illuminate\Console\Application;
use Illuminate\Support\Facades\App;
use Illuminate\Support\Facades\Artisan;
use Illuminate\Support\Facades\Config;
use Illuminate\Support\Facades\DB;
use Symfony\Component\Console\Output\OutputInterface;

class TenantInitService
{
    public function switchConnection(string $databaseName) {
        Config::set('database.connections.tenant.database', $databaseName);
        DB::purge('tenant');
        DB::reconnect('tenant');
    }

    public function runMigration(string $databaseName) {
        $output = null;
        $artisan = App::make(Artisan::class);
        Artisan::call("tenant:migrate", ['--database' => 'tenant', '--tenant' => $databaseName], $output);
        return $output;
    }

    public function runSeeders() {

    }

    public function setInitialData(Owner $owner, Organization $organization) {
        // превращаем owner в пользователя
        $user = User::query()->create($owner->toArray());
        $organization = User::query()->create($organization->toArray());
    }
}
