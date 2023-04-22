<?php

namespace App\Service\Rabbit\Handlers;

use App\Service\Rabbit\HandlerDto;
use App\Service\Rabbit\IMessageHandler;

class CreateOrganizationHandler implements IMessageHandler
{

    public function handle(array $payload): HandlerDto
    {
        return new HandlerDto(true, "CreateOrganization");
    }
}
