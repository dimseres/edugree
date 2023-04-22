<?php

namespace App\Service\Rabbit;

use App\Service\Rabbit\Handlers\CreateOrganizationHandler;

class MessageHandlerService
{
    private $messages = [
        'organization.create' => CreateOrganizationHandler::class,
        'member.add' => CreateOrganizationHandler::class,
    ];

    public function __construct()
    {

    }
    public function handleMessage($type, array $payload): bool {
        if ($this->messages[$type]) {
            return (new $this->messages[$type])->handle($payload)->success;
        }

        return false;
    }
}
