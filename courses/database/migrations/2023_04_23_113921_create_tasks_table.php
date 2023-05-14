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
        if (Schema::getConnection()->getDatabaseName() !== env("DB_DATABASE")) {
            Schema::connection('tenant')->create('tasks', function (Blueprint $table) {
                $table->id();
                $table->string('title');
                $table->integer('cost')->default(1);
                $table->text('description')->nullable();
                $table->morphs('entity');
                $table->timestamps();
            });
        }
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        if (Schema::getConnection()->getDatabaseName() !== env("DB_DATABASE")) {
            Schema::dropIfExists('tasks');
        }
    }
};
