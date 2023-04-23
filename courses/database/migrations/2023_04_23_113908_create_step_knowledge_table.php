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
        Schema::create('step_knowledge', function (Blueprint $table) {
            $table->foreignId('step_id')->constrained('steps');
            $table->foreignId('knowledge_id')->constrained('knowledge');
        });
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        Schema::dropIfExists('step_knowledge');
    }
};
