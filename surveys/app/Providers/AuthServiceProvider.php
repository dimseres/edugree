<?php

namespace App\Providers;

// use Illuminate\Support\Facades\Gate;

use Illuminate\Foundation\Support\Providers\AuthServiceProvider as ServiceProvider;
use Illuminate\Support\Facades\Auth;

class AuthServiceProvider extends ServiceProvider
{
    /**
     * The model to policy mappings for the application.
     *
     * @var array<class-string, class-string>
     */
    protected $policies = [
        //
    ];

    /**
     * Register any authentication / authorization services.
     */
    public function boot(): void
    {
        Auth::viaRequest('jwt', function (Request $request) {
            try {
                $guid = $request->header('X-REQUEST-ID');
                if (!$request->bearerToken() && !$guid) {
                    throw new \Exception("token missed");
                }
                $tokenPayload = JWT::decode($request->bearerToken(), new Key($guid.env("GATEWAY_KEY"), 'HS256'));
                $user = User::query()->find($tokenPayload->data->user_id);
                return $user;

            } catch (\Exception $exception) {
                Log::error($exception);
                return null;
            }
        });
    }
}
