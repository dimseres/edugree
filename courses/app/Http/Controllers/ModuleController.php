<?php

namespace App\Http\Controllers;

use App\Http\Requests\Courses\CreateModuleRequest;
use App\Models\Course;
use App\Models\Module;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;
use mysql_xdevapi\Exception;
use PhpParser\Node\Expr\AssignOp\Mod;

class ModuleController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index(Course $course)
    {
        return Module::query()->where('course_id', $course->id)->with('units')->get();
    }

    /**
     * Show the form for creating a new resource.
     */
    public function create()
    {
        return "";
    }


    private function incrementPositions(int $insertedPos, int $courseId)
    {
        $modules = Module::query()
            ->where("position", '>=', $insertedPos)
            ->where("course_id", $courseId)->get();

        foreach ($modules as $module) {
            $module->position += 1;
            $module->save();
        }

        return $modules;
    }

    private function decrementPositions(int $insertedPos, int $courseId)
    {
        $modules = Module::query()
            ->where("position", '>=', $insertedPos)
            ->where("course_id", $courseId)->get();

        foreach ($modules as $module) {
            $module->position -= 1;
            $module->save();
        }

        return $modules;
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(CreateModuleRequest $request, Course $course)
    {
        try {
            $position = $request->input('position');
            $title = $request->input('title');
            $description = $request->input('description');
            DB::beginTransaction();

            $maxPos = Module::query()->max('position');

            if ($position > $maxPos + 1) {
                $position = $maxPos + 1;
            }

            $modules = $this->incrementPositions($position, $course->id);

            $module = Module::query()->create([
                'title' => $title,
                'course_id' => $course->id,
                'description' => $description,
                'position' => count($modules) == 0 && $position <= 1 ? 1 : $position
            ]);

            DB::commit();
            return $module;
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
    public function show(Course $course, Module $module)
    {
        return $module->load('units');
    }

    /**
     * Show the form for editing the specified resource.
     */
    public function edit(Course $course, Module $module)
    {
        return null;
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(CreateModuleRequest $request, Course $course, Module $module)
    {
        try {
            $position = $request->input('position');
            $title = $request->input('title');
            $description = $request->input('description');

            $maxPos = Module::query()->count('id');

            if ($position > $maxPos) {
                $position = $maxPos;
            }

            DB::beginTransaction();

            if ($module->position !== $position) {
                $replaced = Module::query()->where('course_id', $course->id)->where('position', $position)->update([
                    'position' => $module->position
                ]);
            }

            $module->update([
                'title' => $title,
                'course_id' => $course->id,
                'description' => $description,
                'position' => $position,
            ]);

            DB::commit();
            return $module;
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
    public function destroy(Course $course, Module $module)
    {
        $pos = $module->position;
        $units = $module->units;
        if ($units->count()) {
            return response()->json([
                'error' => true,
                'message' => 'сначала удалите юниты'
            ], 422);
        }

        $module->delete();
        $this->decrementPositions($pos, $course->id);
        return response()->json([
            "error" => false
        ]);
    }
}
