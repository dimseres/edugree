<template>
    <h2 class='mt-6 text-center text-2xl font-medium tracking-tight text-gray-500'>Введите вашу почту</h2>
    <!--          <p class="mt-2 text-center text-sm text-gray-600">-->
    <!--            Или-->
    <!--            {{ ' ' }}-->
    <!--            <a href="#" class="font-medium text-indigo-600 hover:text-indigo-500">зарегистрируйтесь</a>-->
    <!--          </p>-->
    <form class='flex flex-col content-between mt-6 space-y-6' action='#' method='POST'>
        <input type='hidden' name='remember' value='true' />
        <div class='-space-y-px rounded-md shadow-sm'>
            <input id='email' name='email' type='email' autocomplete='email' required
                   class='p-2 border relative block w-full rounded-md border-0 py-1.5 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:z-10 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6'
                   placeholder='Ваша почта' />
        </div>
        <div class='mt-auto'>
            <button type='submit'
                    class='border bg-purple-500 group relative flex w-full justify-center rounded-md text-white py-2 px-3 text-sm font-semibold text-white hover:bg-white hover:border hover:text-purple-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-purple-500'>
            <span class='absolute inset-y-0 left-0 flex items-center pl-3'>
              <LockClosedIcon class='h-5 w-5 text-white group-hover:text-purple-500' aria-hidden='true' />
            </span>
                Войти
            </button>
            <button @click='$router.go(-1)'
                    class='mt-3 border group relative flex w-full justify-center rounded-md text-purple-500 py-2 px-3 text-sm font-semibold focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-purple-500'>
                назад
            </button>
        </div>
    </form>
</template>

<script setup lang='ts'>
import { LockClosedIcon } from '@heroicons/vue/20/solid'
</script>

<style lang='ts' scoped>

</style>