<?php

namespace App\Http\Requests\Courses;

use Illuminate\Foundation\Http\FormRequest;

class CreateKnowledgeRequest extends FormRequest
{
    public function rules()
    {
        return [
            'title' => ['required'],
            'content' => ['string'],
            'description' => ['string'],
            'step_id' => ['numeric'],
        ];
    }
}
