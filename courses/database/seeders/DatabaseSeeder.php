<?php

namespace Database\Seeders;

// use Illuminate\Database\Console\Seeds\WithoutModelEvents;
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
//            \App\Models\User::factory()->create([
//                'id' => 1,
//                'name' => 'Test User',
//                'email' => 'test@example.com',
//            ]);
//        }
        (new RolesPermissionsSeeder())->run();
    }
}
