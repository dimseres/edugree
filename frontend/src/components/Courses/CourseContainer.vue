<script setup lang='ts'>

import CourseCard from './CourseCard.vue'
import { ref } from 'vue'

const fakeData = ref([
    {
        id: 1,
        title: 'Разработка программного кода',
        courseCover: 'https://images.unsplash.com/photo-1515879218367-8466d910aaa4?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8OHx8Y29kZXxlbnwwfHwwfHw%3D&w=1000&q=80',
        lessonCount: 12,
        taskCount: 12,
        active: true,
        progress: 6,
        author: {
            name: 'Надточий П.Н',
            avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTvFWTVF8t77qNUOW-M0MJt-97KXpoW6mRp7Gx22Jlr-X1wRtW63QggWyUVhnXhjvXjcSA&usqp=CAU',
        },
    },
    {
        id: 2,
        title: 'Проектирование надежных и отказоустойчивых вычислительных систем',
        courseCover: 'https://thumbs.dreamstime.com/b/internet-information-technology-businessman-hand-showing-concept-75784736.jpg',
        lessonCount: 14,
        taskCount: 12,
        progress: 7,
        active: true,
        author: {
            name: 'Потапов В.И',
            avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTvFWTVF8t77qNUOW-M0MJt-97KXpoW6mRp7Gx22Jlr-X1wRtW63QggWyUVhnXhjvXjcSA&usqp=CAU',
        },
    },
    {
        id: 3,
        title: 'Системы обнаружения вторжений',
        courseCover: 'https://picsum.photos/id/3/480/640',
        lessonCount: 12,
        taskCount: 12,
        progress: 1,
        active: false,
        author: {
            name: 'Надточий П.Н',
            avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTvFWTVF8t77qNUOW-M0MJt-97KXpoW6mRp7Gx22Jlr-X1wRtW63QggWyUVhnXhjvXjcSA&usqp=CAU',
        },
    },
    {
        id: 4,
        title: 'Разработка операционных систем',
        courseCover: 'https://picsum.photos/id/9/480/640',
        lessonCount: 12,
        taskCount: 12,
        progress: 3,
        active: true,
        author: {
            name: 'Надточий П.Н',
            // avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTvFWTVF8t77qNUOW-M0MJt-97KXpoW6mRp7Gx22Jlr-X1wRtW63QggWyUVhnXhjvXjcSA&usqp=CAU',
        },
    },
    {
        id: 5,
        title: 'Технологии разработки программных систем',
        courseCover: 'https://picsum.photos/id/7/480/640',
        lessonCount: 12,
        taskCount: 12,
        progress: 4,
        active: false,
        author: {
            name: 'Надточий П.Н',
            avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTvFWTVF8t77qNUOW-M0MJt-97KXpoW6mRp7Gx22Jlr-X1wRtW63QggWyUVhnXhjvXjcSA&usqp=CAU',
        },
    },
])

</script>

<template>
    <div class='flex flex-wrap flex-grow-0'>
        <CourseCard class='course-container p-2 w-1/4' v-for='item in fakeData' :key='item.id' :cardData='item'></CourseCard>
    </div>
</template>

<style scoped>

</style>