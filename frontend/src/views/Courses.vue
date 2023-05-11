<template>
    <div class='flex justify-center'>
        <div class='container'>
            <div class='head'>
                <h1 class='text-2xl text-gray-700 font-bold'>Ваши курсы</h1>
                <p class='text-gray-500 font-light text-sm'>Здесь находятся все ваши курсы</p>
            </div>

            <div class='flex'>
                <router-link :to='{name: "CourseCreate"}' class='p-3 text-gray-700 shadow rounded-md bg-white hover:bg-purple hover:text-white'><PlusIcon class='h-[20px] inline'/>создать курс</router-link>
            </div>

            <div v-if='courses'>
                <CourseContainer :courses='courses'></CourseContainer>
            </div>
            <div v-else>
                <h2 class='text-xl my-6 text-gray-500'>
                    Пока у вас нет активных курсов
                </h2>
            </div>

            <hr class='my-3'>

            <div class='head'>
                <h1 class='text-2xl text-gray-700 font-bold'>Доступные</h1>
                <p class='text-gray-500 font-light text-sm'>Здесь находятся доступные для прохождения курсы</p>
            </div>

            <div class='mt-3'>
                <CourseContainer :courses='courses'></CourseContainer>
            </div>
        </div>
    </div>
</template>

<script lang='ts' setup>
import { useUserStore } from '../store/user.store'
import { ref } from 'vue'
import CourseContainer from '../components/Courses/CourseContainer.vue'
import { PlusIcon } from '@heroicons/vue/24/outline'

const { getOrganizationList, pickTenant: _pickTennant } = useUserStore()

const courses = ref()

</script>