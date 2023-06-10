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
Route::middleware('auth:testauth')->group(function () {
   Route::post('/polls/{poll}', [\App\Api\Controllers\PollController::class, 'updatePoll']);
   Route::get('/polls/{poll}', [\App\Api\Controllers\PollController::class, 'updatePoll']);
   Route::delete('/polls/{poll}', [\App\Api\Controllers\PollController::class, 'deletePoll']);
   Route::patch('/polls/{poll}/restore', [\App\Api\Controllers\PollController::class, 'restorePoll']);

   Route::resource('/polls.questions', \App\Api\Controllers\QuestionController::class);
});
