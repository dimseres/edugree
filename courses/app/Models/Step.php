<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Step extends Model
{
    use HasFactory;
    protected $connection = 'tenant';

    protected $guarded = [];
    public function unit() {
        return $this->belongsTo(Unit::class);
    }

    public function entity() {
        return $this->morphTo();
    }
}
