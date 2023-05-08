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
            Schema::connection('tenant')->create('step_knowledge', function (Blueprint $table) {
                $table->foreignId('step_id')->constrained('steps');
                $table->foreignId('knowledge_id')->constrained('knowledge');
            });
        }

    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        if (Schema::getConnection()->getDatabaseName() !== env("DB_DATABASE")) {
            Schema::dropIfExists('step_knowledge');
        }
    }
};
