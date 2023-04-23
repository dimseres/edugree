<?php

namespace App\Service\Amqp;

use App\Service\Amqp\Handlers\CreateOrganizationHandler;

class MessageHandlerService
{
    // USE ONLY JOBS FOR HANDLE MESSAGE
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
