<?php

namespace App\Providers;

// use Illuminate\Support\Facades\Gate;
use Firebase\JWT\JWT;
use Firebase\JWT\Key;
use Firebase\JWT\SignatureInvalidException;
use Illuminate\Foundation\Support\Providers\AuthServiceProvider as ServiceProvider;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Facades\Log;
use PHPUnit\TextUI\Application;

class AuthServiceProvider extends ServiceProvider
{
    /**
     * The model to policy mappings for the application.
     *
     * @var array<class-string, class-string>
     */
    protected $policies = [
        // 'App\Models\Model' => 'App\Policies\ModelPolicy',
    ];

    /**
     * Register any authentication / authorization services.
     */
    public function boot(): void
    {
        Auth::viaRequest('jwt', function (Request $request) {
            try {
                $guid = $request->header('X-REQUEST-ID');
                $tokenPayload = JWT::decode($request->bearerToken(), new Key($guid.env("GATEWAY_KEY"), 'HS256'));
                dd($tokenPayload);
                dd(base64_decode($tokenPayload));
            } catch (\Exception $exception) {
                Log::error($exception);
                return null;
            }
        });
    }
}
