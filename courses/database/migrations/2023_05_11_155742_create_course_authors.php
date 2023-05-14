<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration
{
    /**
     * Run the migrations.
     */
    public function up(): void
    {
        if (Schema::getConnection()->getDatabaseName() !== env("DB_DATABASE")) {
            Schema::connection('tenant')->create('course_authors', function (Blueprint $table) {
                $table->id();
                $table->foreignId("user_id")->constrained('users');
                $table->foreignId("course_id")->constrained('courses');
                $table->foreignId("owner_id")->constrained('users');
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
            Schema::dropIfExists('course_authors');
        }
    }
};
