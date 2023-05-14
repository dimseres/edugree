<?php

namespace App\Providers;

use App\Console\Commands\MigrateCommand;
use Firebase\JWT\JWT;
use Illuminate\Database\ConnectionResolverInterface;
use Illuminate\Database\Migrations\DatabaseMigrationRepository;
use Illuminate\Database\Migrations\MigrationRepositoryInterface;
use Illuminate\Http\Request;
use Illuminate\Routing\Route;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\ServiceProvider;

class AppServiceProvider extends ServiceProvider
{
    /**
     * Register any application services.
     */
    public function register(): void
    {
        $this->app->singleton(MigrationRepositoryInterface::class, function ($app) {
            return new DatabaseMigrationRepository($app['db'], 'migrations');
        });
    }

    /**
     * Bootstrap any application services.
     */
    public function boot(): void
    {
    }
}
