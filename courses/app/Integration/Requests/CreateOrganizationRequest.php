<?php

namespace App\Integration\Requests;

use Illuminate\Foundation\Http\FormRequest;

class CreateOrganizationRequest extends FormRequest
{

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array<string, \Illuminate\Contracts\Validation\Rule|array|string>
     */
    public function rules(): array
    {
        return [
            'id' => ['required', 'numeric'],
            'name' => ['required'],
            'domain' => ['required'],
            'tenantUuid' => ['required'],
            'owner' => ['required', 'array'],
            'owner.id' => ['required', 'numeric'],
            'owner.name' => ['required','string'],
            'owner.email' => ['required', 'email'],
            'owner.phone' => ['required'],
        ];
    }
}
