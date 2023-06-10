<?php

namespace App\Api\Requests;

use Illuminate\Foundation\Http\FormRequest;


class PollUpdateRequest extends FormRequest
{
    public function rules()
    {
        return [
            'title' => ['string'],
            'description' => ['string'],
            'creator_id' => ['numeric'],
            'type' => ['numeric'],
        ];
    }
}
