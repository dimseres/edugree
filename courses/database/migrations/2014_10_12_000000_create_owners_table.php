<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration {
    /**
     * Run the migrations.
     */
    protected $connection = 'pgsql';
    public function up(): void
    {
        if (Schema::getConnection()->getDatabaseName() === env("DB_DATABASE")) {
            if (!Schema::hasTable('owners')) {
                Schema::connection('pgsql')->create('owners', function (Blueprint $table) {
                    $table->unsignedBigInteger('id')->primary()->unique();
                    $table->string('name');
                    $table->string('email')->unique();
                    $table->string('phone')->unique();
                    $table->text('bio')->nullable();
                    $table->rememberToken();
                    $table->timestamps();
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
            Schema::dropIfExists('users');
        }
    }
};
