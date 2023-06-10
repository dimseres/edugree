<?php

namespace App\Integration\Requests;

use Illuminate\Foundation\Http\FormRequest;

class SurveyCreateRequest extends FormRequest
{
    public function rules()
    {
        return [
            'creator' => ['required', 'array'],
            'creator.id' => ['required', 'numeric'],
            'creator.email' => ['required', 'email'],
            'creator.name' => ['required', 'string'],
            'creator.organization' => ['required', 'array'],
            'creator.organization.id' => ['required', 'numeric'],
            'creator.organization.name' => ['required', 'string'],
            'creator.organization.uuid' => ['required', 'string'],
            'survey' => ['required', 'array'],
            'survey.title' => ['required', 'string'],
            'survey.type' => ['required', 'numeric'],
            'survey.description' => ['string', 'nullable'],
        ];
    }
}
