<?php

namespace App\Http\Middleware;

use Closure;
use Illuminate\Http\Request;
use Symfony\Component\HttpFoundation\Response;

class ApiGatewayMiddleware
{
    /**
     * Handle an incoming request.
     *
     * @param  \Closure(\Illuminate\Http\Request): (\Symfony\Component\HttpFoundation\Response)  $next
     */
    public function handle(Request $request, Closure $next): Response
    {
        $gatewayKey = $request->header('X-ACCESS-KEY');
        $guidKey = $request->header('X-REQUEST-ID');
        if ($gatewayKey !== env('COURSE_KEY') || !$guidKey) {
            return \response()->json([
                "error" => true,
                "message" => "access denied"
            ], 403);
        }

        return $next($request);
    }
}
