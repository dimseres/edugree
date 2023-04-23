<?php

namespace App\Service\Amqp\Handlers;

use App\Service\Amqp\HandlerDto;
use App\Service\Amqp\IMessageHandler;

class CreateOrganizationHandler implements IMessageHandler
{

    public function handle(array $payload): HandlerDto
    {
        return new HandlerDto(true, "CreateOrganization");
    }
}
