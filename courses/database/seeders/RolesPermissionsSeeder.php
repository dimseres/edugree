<?php

namespace Database\Seeders;

use App\Models\User;
use App\Models\Role;
use App\Models\Permission;

use Illuminate\Database\Console\Seeds\WithoutModelEvents;
use Illuminate\Database\Seeder;

class RolesPermissionsSeeder extends Seeder
{
    /**
     * Run the database seeds.
     */
    public function run(): void
    {
        $owner = Role::create([
            'name' => 'owner',
            'guard_name' => 'api'
        ]);
        $admin = Role::create([
            'name' => 'admin',
            'guard_name' => 'api'
        ]);
        $moderator = Role::create([
            'name' => 'moderator',
            'guard_name' => 'api'
        ]);
        $teacher = Role::create([
            'name' => 'teacher',
            'guard_name' => 'api'
        ]);
        $student = Role::create([
            'name' => 'student',
            'guard_name' => 'api'
        ]);
        Permission::query()->insert([
            ['name' => 'create course', 'guard_name' => 'api'],
            ['name' => 'edit course', 'guard_name' => 'api'],
            ['name' => 'delete course', 'guard_name' => 'api'],
            ['name' => 'create test', 'guard_name' => 'api'],
            ['name' => 'delete test', 'guard_name' => 'api'],
            ['name' => 'take test', 'guard_name' => 'api'],
        ]);

        if (getenv("APP_ENV") === 'development') {
            $user = User::query()->where('email', 'test@example.com')->first();
            $user?->assignRole($owner);
        }
    }
}
