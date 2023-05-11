<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Module extends Model
{
    use HasFactory;
    protected $connection = 'tenant';

    protected $guarded = [];
    public function courses() {
        return $this->belongsTo(Course::class);
    }

    public function units() {
        return $this->hasMany(Unit::class);
    }
}
