<?php

namespace Database\Seeders;
use App\Models\User;
use Spatie\Permission\Models\Role;
use Spatie\Permission\Models\Permission;

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
        ]);
        $admin = Role::create([
            'name' => 'admin',
        ]);
        $moderator = Role::create([
            'name' => 'moderator',
        ]);
        $teacher = Role::create([
            'name' => 'teacher',
        ]);
        $student = Role::create([
            'name' => 'student',
        ]);
        $editCourse = Permission::create(['name' => 'edit course']);
        $deleteCourse = Permission::create(['name' => 'delete course']);

        if (getenv("APP_ENV") === 'development') {
            $user = User::query()->where('email', 'test@example.com')->first();
            $user?->assignRole($owner);
        }
    }
}
