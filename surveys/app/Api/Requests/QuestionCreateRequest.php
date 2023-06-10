<?php

namespace App\Api\Requests;

use Illuminate\Foundation\Http\FormRequest;


class QuestionCreateRequest extends FormRequest
{
    public function rules()
    {
        return [
            'title' => ['required', 'string'],
            'content' => ['required', 'string'],
            'type' => ['required', 'numeric'],
            'position' => ['required', 'numeric'],
        ];
    }
}
