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
Route::middleware('integration')->group(function () {
    Route::post("/organization/create", [\App\Integration\Controllers\OrganizationController::class, 'createOrganization']);
    Route::patch("/organization/update", [\App\Integration\Controllers\OrganizationController::class, 'updateOrganization']);
    Route::delete("/organization/delete", [\App\Integration\Controllers\OrganizationController::class, 'deleteOrganization']);
    Route::post("/organization/member/add", [\App\Integration\Controllers\OrganizationController::class, 'createTenantUser']);
//     Route::delete("/organization/member/add", [\App\Integration\Controllers\OrganizationController::class, 'deleteOrganization']);
});
