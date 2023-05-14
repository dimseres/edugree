<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Role extends \Spatie\Permission\Models\Role
{
    protected $connection = 'tenant';

    const ROLE_OWNER = 'owner';
    const ROLE_ADMINISTRATOR = 'administrator';
    const ROLE_MODERATOR = 'moderator';
    const ROLE_TEACHER = 'teacher';
    const ROLE_STUDENT = 'student';
}
