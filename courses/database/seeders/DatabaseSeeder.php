<?php

namespace Database\Seeders;

// use Illuminate\Database\Console\Seeds\WithoutModelEvents;
use App\Models\Role;
use App\Models\User;
use Illuminate\Database\Seeder;
use phpseclib3\Crypt\Hash;

class DatabaseSeeder extends Seeder
{
    /**
     * Seed the application's database.
     */
    public function run(): void
    {
//        if (getenv("APP_ENV") === "development") {
//
//        }
        (new RolesPermissionsSeeder())->run();

        if (getenv("APP_ENV") === "development") {
            $ownerUser = \App\Models\User::factory()->create([
                'id' => 1,
                'name' => 'Test User',
                'email' => 'test1@example.com',
            ]);

            $adminUser = \App\Models\User::factory()->create([
                'id' => 2,
                'name' => 'Test User',
                'email' => 'test2@example.com',
            ]);

            $teacherUser = \App\Models\User::factory()->create([
                'id' => 3,
                'name' => 'Test User',
                'email' => 'test3@example.com',
            ]);

            $moderUser = \App\Models\User::factory()->create([
                'id' => 4,
                'name' => 'Test User',
                'email' => 'test4@example.com',
            ]);

            $studentUser = \App\Models\User::factory()->create([
                'id' => 5,
                'name' => 'Test User',
                'email' => 'test5@example.com',
            ]);



            $owner = Role::query()->where('name', User::ROLE_OWNER)->first();
            $admin = Role::query()->where('name', User::ROLE_ADMIN)->first();
            $moder = Role::query()->where('name', User::ROLE_MODERATOR)->first();
            $teacher = Role::query()->where('name', User::ROLE_TEACHER)->first();
            $student = Role::query()->where('name', User::ROLE_STUDENT)->first();

            $ownerUser->assignRole($owner);
            $adminUser->assignRole($admin);
            $moderUser->assignRole($moder);
            $teacherUser->assignRole($teacher);
            $studentUser->assignRole($student);
        }
    }
}
