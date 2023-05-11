<script setup lang='ts'>
import CourseProgress from './CourseProgress.vue'
import {UserCircleIcon} from '@heroicons/vue/24/outline'

const props = defineProps<{
    progress: any,
    cover: any,
    author: any
}>()
</script>

<template>
    <div class='course-head relative'>
        <CourseProgress :progress='progress'/>
        <div class='image overflow-hidden h-[134px]'>
            <div class='image-overlay'>
            </div>
            <img :src='cover' alt='#' class='w-full'>
        </div>
        <div class='author mx-4'>
            <div class='author-image shadow-md' v-if='author.avatar'>
                <img :src='author.avatar' alt=''>
            </div>
            <div class='author-image shadow-md' v-else>
                <UserCircleIcon/>
            </div>
            <p class='ml-4 mt-2 font-medium text-white'>{{ author.name }}</p>
        </div>
    </div>
</template>

<style lang='scss' scoped>

.image {
    max-height: 134px;

    &-overlay {
        width: 100%;
        height: 100%;
        position: absolute;
        background: rgba(#000, .2);
    }
}
.author {
    position: absolute;
    bottom: -12px;
    display: flex;
    &-image {
        background: #fff;
        width: 56px;
        height: 56px;
        border-radius: 50%;
        overflow: hidden;
    }
}
</style>