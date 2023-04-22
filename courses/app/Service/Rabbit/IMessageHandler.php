<?php

namespace App\Service\Rabbit;

interface IMessageHandler
{
    public function handle(array $payload): HandlerDto;
}
