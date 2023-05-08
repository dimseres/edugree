<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration
{
    /**
     * Run the migrations.
     */
//    protected $connection = 'tenant';
    public function up(): void
    {
        if (Schema::getConnection()->getDatabaseName() === env("DB_DATABASE")) {
            Schema::connection('tenant')->create('failed_jobs', function (Blueprint $table) {
                $table->id();
                $table->string('uuid')->unique();
                $table->text('connection');
                $table->text('queue');
                $table->longText('payload');
                $table->longText('exception');
                $table->timestamp('failed_at')->useCurrent();
            });
        }
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        if (Schema::getConnection()->getDatabaseName() === env("DB_DATABASE")) {
            Schema::dropIfExists('failed_jobs');
        }
    }
};
