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

class CoursePublicController extends Controller
{
    public function __construct()
    {
//        $allowedEditCreate = Role::ROLE_OWNER . '|' . Role::ROLE_ADMINISTRATOR . '|' . Role::ROLE_MODERATOR . '|' . Role::ROLE_TEACHER;
//        $this->middleware("role:{$allowedEditCreate}")->only(['store', 'update', 'create', 'edit', 'delete']);
    }

    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        $courses = Course::query()->where('status', Course::COURSE_PUBLISHED)->paginate(10);
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

    public function show(string $id) {
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
}
