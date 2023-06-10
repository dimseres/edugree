<?php

namespace App\Integration\Requests;

use Illuminate\Foundation\Http\FormRequest;

class SurveyUpdateRequest extends FormRequest
{
    public function rules()
    {
        return [
            'poll' => ['required', 'array'],
            'poll.title' => ['string'],
            'poll.description' => ['string'],
            'poll.creator_id' => ['numeric'],
            'poll.type' => ['numeric'],
        ];
    }
}
