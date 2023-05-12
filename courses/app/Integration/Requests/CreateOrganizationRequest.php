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
            'tenant_uuid' => ['required'],
            'user' => ['required', 'array'],
            'user.id' => ['required', 'numeric'],
            'user.name' => ['required','string'],
            'user.email' => ['required', 'email'],
            'user.phone' => ['required'],
            'user.avatar' => ['string'],
        ];
    }
}
