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
            'name' => Role::ROLE_OWNER,
            'guard_name' => 'api'
        ]);
        $admin = Role::create([
            'name' => Role::ROLE_ADMINISTRATOR,
            'guard_name' => 'api'
        ]);
        $moderator = Role::create([
            'name' => Role::ROLE_MODERATOR,
            'guard_name' => 'api'
        ]);
        $teacher = Role::create([
            'name' => Role::ROLE_TEACHER,
            'guard_name' => 'api'
        ]);
        $student = Role::create([
            'name' => Role::ROLE_STUDENT,
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
