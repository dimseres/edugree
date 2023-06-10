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
    private function incrementPositions(int $insertedPos, int $pollId)
    {
        $units = Question::query()
            ->where("position", '>=', $insertedPos)
            ->where("poll_id", $pollId)->get();

        foreach ($units as $unit) {
            $unit->position += 1;
            $unit->save();
        }

        return $units;
    }
    public function index($poll_id)
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

            $position = $body['position'];
            $maxPos = Question::query()->max('position');
            $position = $position <= 0 ? 1 : $position;
            if ($position > $maxPos + 1) {
                $position = $maxPos + 1;
            }
            $questions = $this->incrementPositions($position, $poll->id);

            $question = Question::query()->create([
                'title' => $body['title'],
                'position' => count($questions) == 0 && $position <= 1 ? 1 : $position,
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

    public function update(QuestionCreateRequest $request, Poll $poll, Question $question)
    {
        try {
            DB::beginTransaction();

            $body = $request->toArray();

            $position = $body['position'];
            $maxPos = Question::query()->max('position');
            $position = $position <= 0 ? 1 : $position;
            if ($position > $maxPos + 1) {
                $position = $maxPos + 1;
            }


            $position = $position <= 0 ? 1 : $position;
            $maxPos = Question::query()->count('id');

            if ($position > $maxPos) {
                $position = $maxPos;
            }

            if ($question->position !== $position) {
                $replaced = Question::query()->where('poll_id', $poll->id)->where('position', $position)->update([
                    'position' => $question->position
                ]);
            }

            $question = $question->update([
                'title' => $body['title'],
                'position' => $body['position'],
                'type' => $body['type'],
                'content' => $body['content'],
            ]);

            DB::commit();
            return response()->json([
                'error' => false,
                'updated' => $question
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
