<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Knowledge extends Model
{
    use HasFactory;
    protected $connection = 'tenant';
    public function module() {
        $this->belongsTo(Module::class);
    }

    public function steps() {
        $this->morphMany(Step::class, 'entity');
    }
}
