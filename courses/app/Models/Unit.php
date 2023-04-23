<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Unit extends Model
{
    use HasFactory;

    public function module() {
        $this->belongsTo(Module::class);
    }

    public function steps() {
        $this->hasMany(Step::class);
    }
}
