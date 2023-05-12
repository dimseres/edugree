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
        $this->belongsTo(Unit::class);
    }

    public function content() {
        $this->morphTo();
    }
}
