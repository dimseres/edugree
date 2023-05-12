<?php

namespace App\Http\Controllers;

use App\Http\Requests\Courses\CreateModuleRequest;
use App\Http\Requests\Courses\CreateUnitRequest;
use App\Models\Course;
use App\Models\Step;
use App\Models\Unit;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;
use mysql_xdevapi\Exception;
use PhpParser\Node\Expr\AssignOp\Mod;

class StepController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index(Unit $unit)
    {
        return Step::query()->where('unit_id', $unit->id)->with('steps')->get();
    }

    /**
     * Show the form for creating a new resource.
     */
    public function create()
    {
        return "";
    }


    private function incrementPositions(int $insertedPos, int $moduleId)
    {
        $steps = Step::query()
            ->where("position", '>=', $insertedPos)
            ->where("unit_id", $moduleId)->get();

        foreach ($steps as $step) {
            $step->position += 1;
            $step->save();
        }

        return $steps;
    }

    private function decrementPositions(int $insertedPos, int $moduleId)
    {
        $steps = Step::query()
            ->where("position", '>=', $insertedPos)
            ->where("unit_id", $moduleId)->get();

        foreach ($steps as $step) {
            $step->position -= 1;
            $step->save();
        }

        return $steps;
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(CreateUnitRequest $request, Unit $unit)
    {
        try {
            $position = $request->input('position');
            $title = $request->input('title');
            $description = $request->input('description');
            DB::beginTransaction();

            $maxPos = Step::query()->max('position');

            $position = $position <= 0 ? 1 : $position;

            if ($position > $maxPos + 1) {
                $position = $maxPos + 1;
            }

            $units = $this->incrementPositions($position, $unit->id);

            $step = Step::query()->create([
                'title' => $title,
                'unit_id' => $unit->id,
                'description' => $description,
                'position' => count($units) == 0 && $position <= 1 ? 1 : $position
            ]);

            DB::commit();
            return $step;
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
    public function show(int $unit, Step $step)
    {
        return $step->load(['entity']);
    }

    /**
     * Show the form for editing the specified resource.
     */
    public function edit(Course $course, Unit $module)
    {
        return null;
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(CreateUnitRequest $request, Unit $unit, Step $step)
    {
        try {
            $position = $request->input('position');
            $title = $request->input('title');
            $description = $request->input('description');

            $position = $position <= 0 ? 1 : $position;
            $maxPos = Step::query()->count('id');

            if ($position > $maxPos) {
                $position = $maxPos;
            }

            DB::beginTransaction();

            if ($step->position !== $position) {
                $replaced = Step::query()->where('unit_id', $unit->id)->where('position', $position)->update([
                    'position' => $step->position
                ]);
            }

            $step->update([
                'title' => $title,
                'unit_id' => $unit->id,
                'description' => $description,
                'position' => $position,
            ]);

            DB::commit();
            return $step;
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
    public function destroy(Unit $unit, Step $step)
    {
        try {
            DB::beginTransaction();

            $step->delete();
            $steps = Step::query()->where('unit_id', $unit->id)->orderBy('position')->get();
            foreach ($steps as $idx => $step) {
                $step->position = $idx + 1;
                $step->save();
            }
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
