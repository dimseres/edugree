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
Route::middleware('auth:api')->group(function () {
//    Route::get("/", function () {
//        return [];
//    });
    Route::get("/courses/my", [\App\Http\Controllers\CoursePublicController::class, 'userCourses']);
    Route::post("/courses/{course_id}/join", [\App\Http\Controllers\CoursePublicController::class, 'userCourses']);
    Route::resource("/edit/courses", \App\Http\Controllers\CourseEditController::class);
    Route::resource("/edit/courses.modules", \App\Http\Controllers\ModuleEditController::class);
    Route::resource("/edit/modules.units", \App\Http\Controllers\UnitEditController::class);
    Route::resource("/edit/units.steps", \App\Http\Controllers\StepEditController::class);
    Route::resource("/knowledge", \App\Http\Controllers\KnowledgeController::class);
});
