<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;

class Course extends Model
{
    use HasFactory, SoftDeletes;

    protected $connection = 'tenant';

    protected $guarded = [];

    const COURSE_EDITED = 1;
    const COURSE_PUBLISHED = 2;
    public function modules() {
        return $this->hasMany(Module::class);
    }

    public function userCourses() {
        return $this->belongsToMany(User::class, 'user_courses');
    }

    public function courseAuthors() {
        return $this->belongsToMany(User::class, 'course_authors');
    }
}
