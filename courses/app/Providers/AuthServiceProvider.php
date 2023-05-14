<?php

namespace App\Providers;

// use Illuminate\Support\Facades\Gate;
use App\Service\Tenant\TenantInitService;
use Firebase\JWT\JWT;
use Firebase\JWT\Key;
use Firebase\JWT\SignatureInvalidException;
use App\Models\User;
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
                $service = new TenantInitService();
                if (!$request->bearerToken() && !$guid) {
                    throw new \Exception("token missed");
                }
                $tokenPayload = JWT::decode($request->bearerToken(), new Key($guid.env("GATEWAY_KEY"), 'HS256'));
                $service->switchConnection($service->formatTenantDbName($tokenPayload->data->membership->tenant_uuid));
                $user = User::query()->find($tokenPayload->data->user_id);
                return $user;

            } catch (\Exception $exception) {
                Log::error($exception);
                return null;
            }
        });

        Auth::viaRequest('testauth', function (Request $request) {
            try {
                $email = $request->header('X-REQUEST-ID');
                $user = \App\Models\User::query()->where('email', $email)->first();
                return $user;
            } catch (\Exception $exception) {
                Log::error($exception);
                return null;
            }
        });
    }
}
