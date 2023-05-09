<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Module extends Model
{
    use HasFactory;
    protected $connection = 'tenant';
    public function courses() {
        $this->belongsTo(Course::class);
    }

    public function units() {
        $this->hasMany(Unit::class);
    }
}
