<?php

namespace App\Http\Requests\Courses;

use Illuminate\Foundation\Http\FormRequest;

class CreateUnitRequest extends FormRequest
{
    public function rules()
    {
        return [
            'title' => ['required'],
            'position' => ['required'],
            'description' => ['string'],
        ];
    }
}
