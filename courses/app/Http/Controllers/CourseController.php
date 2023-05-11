<?php

namespace App\Http\Controllers;

use App\Http\Requests\Courses\CreateCourseRequest;
use App\Models\Course;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\DB;

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

    public function userCourses()
    {
        $user = Auth::user();
        $courses = Course::query()
            ->whereHas('userCourses', function ($q) use ($user) {
                $q->where('user_courses.id', $user->getAuthIdentifier());
            })->orWhereHas('courseAuthors', function ($q) use ($user) {
                $q->whereIn('course_authors.id', [$user->getAuthIdentifier()]);
            })
            ->withCount('modules')
            ->with('userCourses');

        return $courses->paginate(25);
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
        try {
            $courseName = $request->input('name');
            $courseDescription = $request->input('description');
            $course = Course::query()->where('title', $courseName)->first();
            if ($course) {
                return [
                    'error' => true,
                    'message' => 'курс с таким названием уже существует'
                ];
            }

            DB::beginTransaction();
            $course = Course::query()->create([
                'title' => $courseName,
                'description' => $courseDescription
            ]);

            $course->userCourses()->attach(Auth::user());
            $course->courseAuthors()->attach(Auth::user(), ['owner_id' => Auth::user()->getAuthIdentifier()]);
            DB::commit();
            return [
                "error" => false,
                "data" => $course
            ];
        } catch (\Exception $exception) {
            DB::rollBack();
            return [
                "error" => true,
                "message" => $exception->getMessage()
            ];
        }



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
    public function destroy(Course $course)
    {
        $modules = $course->modules;
        if ($modules) {
            return [
                'error' => true,
                'message' => 'сначала удалите модули'
            ];
        }
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
