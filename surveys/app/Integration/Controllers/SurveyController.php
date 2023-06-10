<?php

namespace App\Integration\Controllers;

use App\Integration\Requests\SurveyAttachEditorsRequest;
use App\Integration\Requests\SurveyCreateRequest;
use App\Integration\Requests\SurveyUpdateRequest;
use App\Models\Organization;
use App\Models\Poll;
use App\Models\User;
use Illuminate\Routing\Controller as BaseController;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Str;

class SurveyController extends BaseController
{
    public function createSurvey(SurveyCreateRequest $request) {
        $body = $request->toArray();
        try {
            DB::beginTransaction();
            $user = User::query()->firstOrCreate([
                "id" => $body['creator']['id'],
            ], [
                "id" => $body['creator']['id'],
                "email" => $body['creator']['email'],
                "name" => $body['creator']['name'],
            ]);

            $organization = Organization::query()->firstOrCreate([
                ...$body['creator']['organization']
            ]);

            $survey = Poll::query()->create([
                'creator_id' => $user->id,
                'organization_id' => $organization->id,
                'title' => $body['survey']['title'],
                'description' => $body['survey']['description'],
                'slug' => Str::uuid()->toString(),
            ]);
            $survey->editors()->attach($user);

            DB::commit();

            return response()->json([
                'error' => false,
                'user' => $user,
                'survey' => $survey
            ]);
        } catch (\Exception $exception) {
            DB::rollBack();
            return response()->json([
                'error' => true,
                'message' => $exception->getMessage()
            ], 500);
        }
    }

    public function attachEditors(Poll $poll, SurveyAttachEditorsRequest $request) {
        try {
            DB::beginTransaction();
            $body = $request->toArray();

            $userIds = [];
            foreach ($body['editors'] as $editor) {
                $user = User::query()->firstOrCreate([
                    'id' => $editor['id'],
                    'name' => $editor['name'],
                    'email' => $editor['email'],
                ]);
                $userIds[] = $user->id;
            }

            $poll->editors()->syncWithoutDetaching($userIds);
            DB::commit();

            return response()->json([
                'error' => false,
                'message' => 'attached',
            ]);
        } catch (\Exception $exception) {
            DB::rollBack();
            return response()->json([
                'error' => true,
                'message' => $exception->getMessage()
            ], 500);
        }
    }

    public function deleteSurvey(Poll $poll)
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

    public function restoreSurvey($poll_id)
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
