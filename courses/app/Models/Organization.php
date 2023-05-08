<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class Organization extends Model
{
    public $guarded = [];

    public function owner() {
        return $this->hasOne(User::class, "owner_id");
    }
}
