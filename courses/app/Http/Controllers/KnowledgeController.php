<?php

namespace App\Http\Controllers;

use App\Http\Requests\Courses\CreateKnowledgeRequest;
use App\Models\Knowledge;
use App\Models\Module;
use App\Models\Step;
use Illuminate\Support\Facades\DB;

class KnowledgeController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        return Knowledge::query()->paginate(10);
    }

    /**
     * Show the form for creating a new resource.
     */
    public function create()
    {
        return "";
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(CreateKnowledgeRequest $request)
    {
        try {
            $step_id = $request->input('step_id');
            $title = $request->input('title');
            $description = $request->input('description');
            $content = $request->input('content');
            DB::beginTransaction();

            $knowledge = Knowledge::query()->create([
                'title' => $title,
                'description' => $description,
                'content' => $content,
            ]);

            if ($step_id) {
                $step = Step::query()->findOrFail($step_id);
                $step->entity()->associate($knowledge);
                $step->save();
            }

            DB::commit();
            return $knowledge;
        } catch (\Exception $exception) {
            return [
                "error" => true,
                "message" => $exception->getMessage()
            ];
        }
    }

    /**
     * Display the specified resource.
     */
    public function show(Knowledge $knowledge)
    {
        return $knowledge;
    }

    /**
     * Show the form for editing the specified resource.
     */
//    public function edit(Module $module)
//    {
//        return null;
//    }

    /**
     * Update the specified resource in storage.
     */
    public function update(CreateKnowledgeRequest $request, Knowledge $knowledge)
    {
        try {
            $step_id = $request->input('step_id');
            $title = $request->input('title');
            $description = $request->input('description');
            $content = $request->input('content');

            DB::beginTransaction();

            $knowledge->update([
                'title' => $title,
                'description' => $description,
                'content' => $content,
            ]);

            if ($step_id) {
                $step = Step::query()->findOrFail($step_id);
                $step->entity()->associate($knowledge);
                $step->save();
            }

            DB::commit();
            return $knowledge;
        } catch (\Exception $exception) {
            DB::rollBack();
            return response()->json([
                "error" => true,
                "message" => $exception->getMessage()
            ], 500);
        }
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(Module $module, Knowledge $knowledge)
    {
        try {
            DB::beginTransaction();
            $knowledge->delete();
            DB::commit();
            return response()->json([
                "error" => false
            ]);
        } catch (\Exception $exception) {
            DB::rollBack();
            return response()->json([
                "error" => true,
                "message" => $exception->getMessage()
            ], 500);
        }
    }
}
