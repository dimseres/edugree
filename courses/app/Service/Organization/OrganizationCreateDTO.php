<?php

namespace App\Service\Organization;

readonly class OrganizationCreateDTO
{
    public function __construct(
        string $organizationName,
        string $organizationDomain,
        OrganizationCreateOwnerDTO $owner
    ){}
}
