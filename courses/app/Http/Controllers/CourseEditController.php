<?php

namespace App\Http\Controllers;

use App\Http\Requests\Courses\CreateCourseRequest;
use App\Models\Course;
use App\Models\Role;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Facades\Storage;
use Illuminate\Support\Facades\URL;

class CourseEditController extends Controller
{
    public function __construct()
    {
        $allowedEditCreate = Role::ROLE_OWNER . '|' . Role::ROLE_ADMINISTRATOR . '|' . Role::ROLE_MODERATOR . '|' . Role::ROLE_TEACHER;
        $this->middleware("role:{$allowedEditCreate}")->only(['store', 'update', 'create', 'edit', 'delete']);
    }

    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        $courses = Course::query()->where('status', Course::COURSE_EDITED)->paginate(10);
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

    public function joinCourse(Course $course, Request $request)
    {
        $user = Auth::user();
        if ($course->status == Course::COURSE_EDITED) {
            return response()->json([
                'error' => true,
                'message' => 'course not created',
            ], 422);
        }

        $course->userCourses()->attach($user->getAuthIdentifier());

        return $course;
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
            $cover = $request->hasFile('cover');
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

            if ($cover) {
                $path = Storage::disk('public')->put('images/covers', $request->file('cover'));
                $course->cover = URL::asset('storage/' . $path);
            }

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
        $user = Auth::user();
        $course = Course::query()
            ->whereHas('userCourses', function ($q) use ($user) {
                $q->where('user_courses.id', $user->getAuthIdentifier());
            })->orWhereHas('courseAuthors', function ($q) use ($user) {
                $q->whereIn('course_authors.id', [$user->getAuthIdentifier()]);
            })
            ->where('id', $id)
            ->with(['modules', 'modules.units'])
            ->firstOrFail();
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
        try {
            $user = Auth::user();
            $isAdmin = $user->hasRole([Role::ROLE_MODERATOR, Role::ROLE_ADMINISTRATOR, Role::ROLE_OWNER]);
            if (!$isAdmin) {
                $course = Course::query()->where('id', $id)->with(['modules', 'modules.units'])->firstOrFail();
            } else {
                $course = Course::query()->where('id', $id)->with(['modules', 'modules.units'])
                    ->whereHas('courseAuthors', function ($q) use ($user) {
                    $q->whereIn('course_authors.id', [$user->getAuthIdentifier()]);
                })->firstOrFail();
            }
            if (!$course) {
                return response()->json([
                    'error' => true,
                    'message' => 'course not found'
                ], 422);
            }
            $cover = $request->hasFile('cover');
            $courseName = $request->input('name');
            $courseDescription = $request->input('description');
            $course = Course::query()->where('id', $id)->with(['modules', 'modules.units'])->firstOrFail();

            if ($cover) {
                $path = Storage::disk('public')->put('images/covers', $request->file('cover'));
                $course->cover = URL::asset('storage/' . $path);
            }

            $course->title = $courseName ?? $course->title;
            $course->description = $courseDescription ?? $course->description;

            $course->save();

            return $course;
        } catch (\Exception $exception) {
            return response()->json([
                'error' => true,
                'message' => $exception->getMessage()
            ]);
        }
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
