<?php

namespace App\Api\Controllers;
use App\Api\Requests\PollUpdateRequest;
use App\Models\Poll;
use App\Models\User;
use Illuminate\Routing\Controller as BaseController;
use Illuminate\Support\Facades\DB;

class PollController extends BaseController
{
    public function getPollData($poll_id) {

        throw new \Exception("NOT IMPLEMENTED");
    }
    public function updatePoll(Poll $poll, PollUpdateRequest $request)
    {
        try {
            $body = $request->toArray();
            if (empty($body)) {
                return response()->json([
                    'error' => false,
                    'message' => 'nothing to update'
                ]);
            }
            DB::beginTransaction();

            foreach ($body as $key => $value) {
                if ($key === 'creator_id') {
                    $user = User::query()->findOrFail($value);
                }
                $poll->$key = $value;
            }
            $poll->save();

            DB::commit();
            return response()->json([
                'error' => false,
                'message' => 'updated'
            ]);
        } catch (\Exception $exception) {
            DB::rollBack();
            return response()->json([
                'error' => true,
                'message' => $exception->getMessage()
            ], 500);
        }
    }

    public function deletePoll(Poll $poll)
    {
        try {
            DB::beginTransaction();
            $poll->delete();
            DB::commit();
            return response()->json([
                'error' => false,
                'message' => 'deleted'
            ]);
        } catch (\Exception $exception) {
            DB::rollBack();
            return response()->json([
                'error' => true,
                'message' => $exception->getMessage()
            ], 500);
        }
    }

    public function restorePoll($poll_id)
    {
        try {
            DB::beginTransaction();
            Poll::withTrashed()->findOrFail($poll_id)->restore();
            DB::commit();
            return response()->json([
                'error' => false,
                'message' => 'restore'
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
