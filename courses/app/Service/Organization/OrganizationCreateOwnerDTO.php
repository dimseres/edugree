<?php

namespace App\Service\Organization;

readonly class OrganizationCreateOwnerDTO
{
    public function __construct(
        string $id,
        string $name,
        string $email,
        string $phone,
        string|null $bio
    ){}
}
