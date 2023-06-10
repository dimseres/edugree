<?php

namespace App\Integration\Requests;

use Illuminate\Foundation\Http\FormRequest;

class SurveyAttachEditorsRequest extends FormRequest
{
    public function rules()
    {
        return [
            'editors' => ['required', 'array'],
            'editors.*.id' => ['required', 'numeric'],
            'editors.*.email' => ['required', 'email'],
            'editors.*.name' => ['required', 'string'],
        ];
    }
}
