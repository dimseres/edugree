<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider and all of them will
| be assigned to the "api" middleware group. Make something great!
|
*/

//Route::middleware('auth:sanctum')->get('/user', function (Request $request) {
//    return $request->user();
//});
Route::middleware('auth:testauth')->group(function () {
//    Route::get("/", function () {
//        return [];
//    });
    Route::get("/courses/my", [\App\Http\Controllers\CourseController::class, 'userCourses']);
    Route::resource("courses", \App\Http\Controllers\CourseController::class);
    Route::resource("courses.modules", \App\Http\Controllers\ModuleController::class);
});
