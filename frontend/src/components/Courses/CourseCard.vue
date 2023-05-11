<script setup lang='ts'>
import CourseHead from './CourseHead.vue'
import {AcademicCapIcon, PencilSquareIcon} from '@heroicons/vue/24/outline'

const props = defineProps<{
    cardData: {
        author: string,
        progress: string,
        cover: string,
        lessonCount: string,
        title: string,
        id: string|number,
        courseCover: string,
        active: boolean,
    },
    progress?: any
}>()
</script>

<template>
    <div class='course-card'>
        <div
            class='card flex flex-col justify-between relative h-full rounded-xl shadow-sm overflow-hidden transform duration-200 hover:-translate-y-1.5 hover:shadow-xl'>
            <div>
                <CourseHead
                    :author='cardData.author'
                    :progress='progress'
                    :cover='cardData.courseCover' />
                <div class='card-body p-4 mt-1 text-gray-700 bg-white'>
                    <h4 class='font-montserrat card-title font-bold text-lg'>{{ cardData.title }}</h4>
                    <div class='flex mt-3 text-sm'>
                        <div class='flex w-1/3 items-center mr-2'>
<!--                            <img src='@/assets/icons/student.svg' alt='#'>-->
                            <AcademicCapIcon class='w-[20px]'/>
                            <span class='block'>{{ cardData.lessonCount }} тем</span>
                        </div>
                        <div class='elem w-1/3 flex'>
<!--                            <img src='@/assets/icons/pen.svg' alt='#'>-->
                            <PencilSquareIcon class='w-[20px]'/>
                            <span>{{ cardData.lessonCount }} задач</span>
                        </div>
                    </div>
                </div>
            </div>
            <div>
                <div class='p-4 flex card-control mt-auto'>
                    <a v-if='cardData.active'
                                 class='w-1/3 text-sm cursor-pointer block text-center text-purple rounded-md p-2 border border-purple hover:bg-purple hover:text-white transform duration-100'>
                        Продолжить
                    </a>
                    <button
                        :class='{"bg-purple ml-0 text-white": !cardData.active}'
                        class='ml-4 text-purple text-sm rounded-md border border-white p-2 w-1/3 hover:underline transform duration-100'>
                        Подробнее
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>

</style>