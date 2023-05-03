<template>
    <div class='users'>
        <h1 class='text-2xl font-bold mb-2'>Пользователи</h1>
        <div class='control-bar flex justify-between'>
            <div class='w-1/4 flex'>
                <input id='email' name='email' type='email' autocomplete='email' required=''
                       class='px-3 block w-full rounded-l-md border-0 py-1.5 text-gray shadow-sm ring-1 ring-inset ring-gray-light placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6' />
                <button class='bg-gray-light px-6 rounded-r-md'>поиск</button>
            </div>
            <div>
                <button class='h-full bg-blue text-sm text-white px-6 py-1 rounded-md mr-2 flex items-center'>
                    <PlusIcon class='h-4' />
                    добавить
                </button>
            </div>
        </div>
        <div class='search mt-6'>

        </div>
        <div class='p-3 mt-3 shadow-2xl rounded-md'>
            <table class='w-full' v-if='users'>
                <thead class='text-left'>
                <tr class='border-b border-gray-light uppercase text-sm'>
                    <th class='w-[100px]'>id</th>
                    <th>телефон</th>
                    <th>фио</th>
                    <th>почта</th>
                    <th>роль</th>
                </tr>
                </thead>
                <tbody>
                <tr v-for='user in users' :key='user.id' class='border-b border-gray-light hover:bg-gray-light transition last:border-none'>
                    <td class='p-2'>{{ user.id }}</td>
                    <td class='p-2'>{{ user.phone }}</td>
                    <td class='p-2'>{{ user.full_name }}</td>
                    <td class='p-2'>{{ user.email }}</td>
                    <td class='p-2'></td>
                </tr>
                </tbody>
            </table>

            <div class='mt-6'>
                <nav class='isolate inline-flex -space-x-px rounded-md shadow-sm' aria-label='Pagination'>
                    <a href='#'
                       class='relative inline-flex items-center rounded-l-md px-2 py-2 text-gray-400 ring-1 ring-inset ring-gray-light hover:bg-gray-50 focus:z-20 focus:outline-offset-0'>
                        <span class='sr-only'>Previous</span>
                        <ChevronLeftIcon class='h-5 w-5' aria-hidden='true' />
                    </a>
                    <!-- Current: "z-10 bg-indigo-600 text-white focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600", Default: "text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:outline-offset-0" -->
                    <a href='#' aria-current='page'
                       class='relative z-10 inline-flex items-center bg-gray px-4 py-2 text-sm font-semibold text-white focus:z-20 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600'>1</a>
                    <a href='#'
                       class='relative inline-flex items-center px-4 py-2 text-sm font-semibold text-gray-900 ring-1 ring-inset ring-gray-light hover:bg-gray-50 focus:z-20 focus:outline-offset-0'>2</a>
                    <a href='#'
                       class='relative hidden items-center px-4 py-2 text-sm font-semibold text-gray-900 ring-1 ring-inset ring-gray-light hover:bg-gray-50 focus:z-20 focus:outline-offset-0 md:inline-flex'>3</a>
                    <span
                        class='relative inline-flex items-center px-4 py-2 text-sm font-semibold text-gray-700 ring-1 ring-inset ring-gray-light focus:outline-offset-0'>...</span>
                    <a href='#'
                       class='relative hidden items-center px-4 py-2 text-sm font-semibold text-gray-900 ring-1 ring-inset ring-gray-light hover:bg-gray-50 focus:z-20 focus:outline-offset-0 md:inline-flex'>8</a>
                    <a href='#'
                       class='relative inline-flex items-center px-4 py-2 text-sm font-semibold text-gray-900 ring-1 ring-inset ring-gray-light hover:bg-gray-50 focus:z-20 focus:outline-offset-0'>9</a>
                    <a href='#'
                       class='relative inline-flex items-center px-4 py-2 text-sm font-semibold text-gray-900 ring-1 ring-inset ring-gray-light hover:bg-gray-50 focus:z-20 focus:outline-offset-0'>10</a>
                    <a href='#'
                       class='relative inline-flex items-center rounded-r-md px-2 py-2 text-gray-400 ring-1 ring-inset ring-gray-light hover:bg-gray-50 focus:z-20 focus:outline-offset-0'>
                        <span class='sr-only'>Next</span>
                        <ChevronRightIcon class='h-5 w-5' aria-hidden='true' />
                    </a>
                </nav>
            </div>
        </div>
    </div>
</template>

<script setup lang='ts'>
import { ChevronLeftIcon, ChevronRightIcon, PlusIcon } from '@heroicons/vue/24/solid'
import { onMounted, ref } from 'vue'
import { fetchOrganizationMembers } from '../services/api/membership.api'


const users = ref([])
const pagination = ref({})

onMounted(async () => {
    const data = await fetchOrganizationMembers({})
    if (!data.error) {
        users.value = data.data
        pagination.value = {
            total: data.total,
            pages: data.pages,
        }
    }
})

</script>

<style scoped>

</style>