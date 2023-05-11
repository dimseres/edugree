<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Unit extends Model
{
    use HasFactory;
    protected $connection = 'tenant';

    protected $guarded = [];
    public function module() {
        return $this->belongsTo(Module::class);
    }

    public function steps() {
        return $this->hasMany(Step::class);
    }
}
