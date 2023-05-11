<?php

namespace App\Models;

// use Illuminate\Contracts\Auth\MustVerifyEmail;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Foundation\Auth\User as Authenticatable;
use Illuminate\Notifications\Notifiable;
use Laravel\Sanctum\HasApiTokens;
use Spatie\Permission\Traits\HasRoles;

class User extends Authenticatable
{
    use HasApiTokens, HasFactory, Notifiable, HasRoles;
    const ROLE_OWNER = 'owner';
    const ROLE_ADMIN = 'administrator';
    const ROLE_MODERATOR = 'moderator';
    const ROLE_TEACHER = 'teacher';
    const ROLE_STUDENT = 'student';

    protected $guard_name = 'api';

    protected $connection = 'tenant';
    protected $fillable = [
        'id',
        'name',
        'email',
        'password',
        'phone',
    ];

    protected $hidden = [
        'password',
        'remember_token',
    ];

    public function courses() {
        return $this->belongsToMany(Course::class, 'user_courses');
    }
}
