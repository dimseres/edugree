<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class UserCourse extends Model
{
    use HasFactory;

    protected $connection = 'tenant';
    public function modules() {
        $this->hasMany(Module::class);
    }

//    public function userCourses() {
//        $this->()
//    }
}
