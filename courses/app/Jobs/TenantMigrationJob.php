<?php

namespace App\Jobs;

use App\Console\Commands\MigrateCommand;
use Illuminate\Bus\Queueable;
use Illuminate\Contracts\Queue\ShouldBeUnique;
use Illuminate\Contracts\Queue\ShouldQueue;
use Illuminate\Foundation\Bus\Dispatchable;
use Illuminate\Queue\InteractsWithQueue;
use Illuminate\Queue\SerializesModels;
use Illuminate\Support\Facades\Artisan;

class TenantMigrationJob implements ShouldQueue
{
    use Dispatchable, InteractsWithQueue, Queueable, SerializesModels;

    protected array $tenantData = [];
    public function __construct($tenantData)
    {
        $this->tenantData = $tenantData;
    }

    /**
     * Execute the job.
     */
    public function handle(): void
    {
        $db_name = $this->tenantData['db_name'];
        Artisan::call("tenant:migrate", [
            '--tenant' => $db_name
        ]);
    }
}
