<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;

class Poll extends Model
{
    use SoftDeletes;
    protected $table = 'polls';
    protected $guarded = [];
    public function editors() {
        return $this->belongsToMany(User::class, 'poll_editors');
    }

    public function questions() {
        return $this->hasMany(Question::class);
    }
}
