<?php

namespace App\Http\Controllers;

use App\Http\Requests\Courses\CreateCourseRequest;
use App\Models\Course;
use Illuminate\Http\Request;

class CourseController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        $courses = Course::query()->paginate(10);
        return $courses;
    }

    /**
     * Show the form for creating a new resource.
     */
    public function create()
    {

    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(CreateCourseRequest $request)
    {
        $courseName = $request->input('name');
        $courseDescription = $request->input('description');
        $course = Course::query()->where('title', $courseName)->first();
        if ($course) {
            return [
                'error' => true,
                'message' => 'курс с таким названием уже существует'
            ];
        }

        $course = Course::query()->create([
            'title' => $courseName,
            'description' => $courseDescription
        ]);

        return [
            "error" => false,
            "data" => $course
        ];
    }

    /**
     * Display the specified resource.
     */
    public function show(string $id)
    {
        $course = Course::query()->where('id', $id)->with(['modules', 'modules.units'])->firstOrFail();
        return $course;
    }

    /**
     * Show the form for editing the specified resource.
     */
    public function edit(string $id)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(CreateCourseRequest $request, string $id)
    {
        $courseName = $request->input('name');
        $courseDescription = $request->input('description');
        $course = Course::query()->where('id', $id)->with(['modules', 'modules.units'])->firstOrFail();

        $course->title = $courseName;
        $course->description = $courseDescription;

        return $course;
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(string $id)
    {
        $destroyed = Course::query()->find($id)->delete();
        if (!$destroyed) {
            return [
                'error' => true,
                'message' => 'something went wrong'
            ];
        }

        return [
            'error' => false
        ];
    }
}
