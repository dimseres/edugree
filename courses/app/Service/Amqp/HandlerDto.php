<?php

namespace App\Service\Amqp;

readonly class HandlerDto
{
    public function __construct(
        public bool   $success,
        public string $message,
        public mixed  $payload = null,
    ) {}
}
