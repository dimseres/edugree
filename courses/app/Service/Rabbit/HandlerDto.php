<?php

namespace App\Service\Rabbit;

class HandlerDto
{
    public function __construct(
        readonly bool $success,
        readonly string $message,
        readonly mixed $payload = null,
    ) {}
}
