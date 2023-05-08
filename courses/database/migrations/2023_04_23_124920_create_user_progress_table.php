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
            Schema::connection('tenant')->create('user_progress', function (Blueprint $table) {
                $table->id();
                $table->foreignId('user_course_id')->constrained('user_courses');
                $table->morphs('progress');
                $table->smallInteger('status');
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
            Schema::dropIfExists('user_progress');
        }
    }
};
