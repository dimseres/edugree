<?php

namespace App\Console\Commands;

use App\Jobs\TenantMigrationJob;
use App\Models\Organization;
use App\Service\Tenant\TenantInitService;
use Illuminate\Console\Command;
use Illuminate\Contracts\Events\Dispatcher;
use Illuminate\Database\Console\Migrations\BaseCommand;
use Illuminate\Database\Migrations\Migrator;
use Illuminate\Support\Facades\Log;

class MigrateCommand extends \Illuminate\Database\Console\Migrations\MigrateCommand
{
    protected $migrator;

    protected $signature = 'tenant:migrate {--database= : The database connection to use}
                {--force : Force the operation to run when in production}
                {--path=* : The path(s) to the migrations files to be executed}
                {--realpath : Indicate any provided migration file paths are pre-resolved absolute paths}
                {--schema-path= : The path to a schema dump file}
                {--pretend : Dump the SQL queries that would be run}
                {--seed : Indicates if the seed task should be re-run}
                {--seeder= : The class name of the root seeder}
                {--step : Force the migrations to be run so they can be rolled back individually}
                {--tenant= : Tenant Database Name for Change DB connection}';

//    public function handle()
//    {
//        $service = new TenantInitService();
//        $service->switchConnection($this->option('tenant'));
//        $this->migrator = app("migrator");
//        $this->migrator->run($this->getMigrationPaths(), [
//            'pretend' => $this->option('pretend'),
//            'step' => $this->option('step'),
//        ]);
//        parent::__construct($migrator);
//        return parent::handle();
//    }
    public function handle()
    {
        if (!$this->option('tenant')) {
            $orgs = Organization::query()->chunk(20, function ($organizations) {
                foreach ($organizations as $organization) {
                    $this->line("Dispatch tenant migrations for ". $organization->db_name);
                    TenantMigrationJob::dispatch($organization->toArray());
                }
            });
        } else {
            $service = new TenantInitService();
            $service->switchConnection($this->option('tenant'));
            Log::info("Start migration for". $this->option('tenant'));
            $this->line("Start migration for". $this->option('tenant'));
            return parent::handle(); // TODO: Change the autogenerated stub
        }

    }
}
