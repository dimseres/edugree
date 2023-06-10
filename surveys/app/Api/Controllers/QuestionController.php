<?php

namespace App\Api\Controllers;
use App\Api\Requests\PollUpdateRequest;
use App\Api\Requests\QuestionCreateRequest;
use App\Models\Poll;
use App\Models\Question;
use App\Models\User;
use Illuminate\Routing\Controller as BaseController;
use Illuminate\Support\Facades\DB;

class QuestionController extends BaseController
{
    public function index($poll_id, PollUpdateRequest $request)
    {
        $poll = Poll::query()->where('id', $poll_id)->with(['questions', 'questions.answers'])->first();
        if (!$poll) {
            return response()->json([
                'error' => true,
                'message' => 'poll not found',
            ],404);
        }

        return $poll;
    }

    public function store(Poll $poll, QuestionCreateRequest $request)
    {
        try {
            DB::beginTransaction();

            $body = $request->toArray();
            $question = Question::query()->create([
                'title' => $body['title'],
                'position' => $body['position'],
                'type' => $body['type'],
                'content' => $body['content'],
                'poll_id' => $poll->id,
            ]);

            DB::commit();
            return response()->json([
                'error' => false,
                'message' => $question
            ]);
        } catch (\Exception $exception) {
            DB::rollBack();
            return response()->json([
                'error' => true,
                'message' => $exception->getMessage()
            ], 500);
        }
    }
}
