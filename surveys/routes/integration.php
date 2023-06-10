<?php

use Illuminate\Support\Facades\Route;

Route::middleware('integration')->group(function () {
    Route::post("/survey/create", [\App\Integration\Controllers\SurveyController::class, 'createSurvey']);
    Route::delete("/survey/{poll}", [\App\Integration\Controllers\SurveyController::class, 'deleteSurvey']);
    Route::patch("/survey/{poll}/restore", [\App\Integration\Controllers\SurveyController::class, 'restoreSurvey']);
    Route::post("/survey/{poll}/attach/editors", [\App\Integration\Controllers\SurveyController::class, 'attachEditors']);
});
