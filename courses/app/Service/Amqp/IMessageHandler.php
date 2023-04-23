<?php

namespace App\Service\Amqp;

interface IMessageHandler
{
    public function handle(array $payload): HandlerDto;
}
