<?php

namespace App\Http\Controllers;

use App\Http\Requests\Courses\CreateModuleRequest;
use App\Http\Requests\Courses\CreateUnitRequest;
use App\Models\Course;
use App\Models\Module;
use App\Models\Unit;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;
use mysql_xdevapi\Exception;
use PhpParser\Node\Expr\AssignOp\Mod;

class UnitController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index(Module $module)
    {
        return Unit::query()->where('module_id', $module->id)->with('steps')->get();
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
        $units = Unit::query()
            ->where("position", '>=', $insertedPos)
            ->where("module_id", $moduleId)->get();

        foreach ($units as $unit) {
            $unit->position += 1;
            $unit->save();
        }

        return $units;
    }

    private function decrementPositions(int $insertedPos, int $moduleId)
    {
        $units = Unit::query()
            ->where("position", '>=', $insertedPos)
            ->where("module_id", $moduleId)->get();

        foreach ($units as $unit) {
            $unit->position -= 1;
            $unit->save();
        }

        return $units;
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(CreateUnitRequest $request, Module $module)
    {
        try {
            $position = $request->input('position');
            $title = $request->input('title');
            $description = $request->input('description');
            DB::beginTransaction();

            $maxPos = Unit::query()->max('position');

            $position = $position <= 0 ? 1 : $position;

            if ($position > $maxPos + 1) {
                $position = $maxPos + 1;
            }

            $units = $this->incrementPositions($position, $module->id);

            $unit = Unit::query()->create([
                'title' => $title,
                'module_id' => $module->id,
                'description' => $description,
                'position' => count($units) == 0 && $position <= 1 ? 1 : $position
            ]);

            DB::commit();
            return $unit;
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
    public function show(Unit $unit)
    {
        return $unit->load(['steps']);
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
    public function update(CreateUnitRequest $request, Module $module, Unit $unit)
    {
        try {
            $position = $request->input('position');
            $title = $request->input('title');
            $description = $request->input('description');

            $position = $position <= 0 ? 1 : $position;
            $maxPos = Module::query()->count('id');

            if ($position > $maxPos) {
                $position = $maxPos;
            }

            DB::beginTransaction();

            if ($unit->position !== $position) {
                $replaced = Unit::query()->where('module_id', $module->id)->where('position', $position)->update([
                    'position' => $unit->position
                ]);
            }

            $unit->update([
                'title' => $title,
                'module_id' => $module->id,
                'description' => $description,
                'position' => $position,
            ]);

            DB::commit();
            return $unit;
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
    public function destroy(Module $module, Unit $unit)
    {
        try {
            DB::beginTransaction();

            $unit->delete();
            $units = Unit::query()->where('module_id', $module->id)->orderBy('position')->get();
            foreach ($units as $idx => $unit) {
                $unit->position = $idx + 1;
                $unit->save();
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
