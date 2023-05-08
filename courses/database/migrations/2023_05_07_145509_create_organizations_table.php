<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration
{
    /**
     * Run the migrations.
     */
    protected $connection = 'pgsql';
    public function up(): void
    {
        if (Schema::getConnection()->getDatabaseName() === env("DB_DATABASE")) {
            if (!Schema::hasTable('organizations')) {
                Schema::connection('pgsql')->create('organizations', function (Blueprint $table) {
                    $table->unsignedBigInteger('id');
                    $table->unsignedBigInteger('owner_id');
                    $table->string('name');
                    $table->string('domain')->index()->unique();
                    $table->string('db_name')->index()->unique();
                    $table->timestamps();

                    $table->foreign("owner_id")->on("owners")->references("id");
                });
            }
        }
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        if (Schema::getConnection()->getDatabaseName() === env("DB_DATABASE")) {
            Schema::dropIfExists('organizations');
        }
    }
};
