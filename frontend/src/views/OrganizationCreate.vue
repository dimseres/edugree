<template>
    <div class='flex justify-center'>
        <div class='container w-1/3'>
            <div class='border-b border-gray-900/10 pb-12'>
                <h2 class='text-2xl text-gray-700 font-semibold leading-7'>Создание организации</h2>
                <p class='mt-1 text-sm leading-6 text-gray-600'>Данная информация полностью публичная</p>

                <div class='mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6'>
                    <div class='sm:col-span-4'>
                        <label for='username' class='block text-sm font-medium leading-6 text-gray-900'>Домен</label>
                        <div class='mt-2'>
                            <div
                                class='flex rounded-md shadow-sm ring-1 ring-inset ring-gray-300 focus-within:ring-2 focus-within:ring-inset focus-within:ring-indigo-600 sm:max-w-md'>
                                <input type='text' name='username' id='username' autocomplete='username'
                                       class='block flex-1 border-0 bg-transparent py-2 px-3 text-gray-900 placeholder:text-gray-400 focus:ring-0 sm:text-sm sm:leading-6'
                                       placeholder='janesmith' />
                            </div>
                            <p class='mt-3 text-sm leading-6 text-gray-600'>Домен должен быть уникальным</p>
                        </div>
                    </div>

                    <div class='col-span-full'>
                        <label for='about' class='block text-sm font-medium leading-6 text-gray-900'>Об организации</label>
                        <div class='mt-2'>
                        <textarea id='about' name='about' rows='3'
                                  class='block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6' />
                        </div>
                        <p class='mt-3 text-sm leading-6 text-gray-600'>Напишите пару слов о вашей организации</p>
                    </div>

                    <div class='col-span-full'>
                        <label for='photo' class='block text-sm font-medium leading-6 text-gray-900'>Аватар</label>
                        <div class='mt-2 flex items-center gap-x-3'>
                            <UserCircleIcon class='h-12 w-12 text-gray-300' aria-hidden='true' />
                            <input id='avatar-upload' name='avatar-upload' type='file' class='sr-only' />
                            <label for='avatar-upload'
                                    class='cursor-pointer rounded-md bg-white px-2.5 py-1.5 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50'>
                                Изменить
                            </label>
                        </div>
                    </div>

                    <div class='col-span-full'>
                        <label for='cover-photo' class='block text-sm font-medium leading-6 text-gray-900'>Обложка</label>
                        <div class='mt-2 flex justify-center rounded-lg border border-dashed border-gray-900/25 px-6 py-10'>
                            <div class='text-center'>
                                <PhotoIcon class='mx-auto h-12 w-12 text-gray-300' aria-hidden='true' />
                                <div class='mt-4 flex text-sm leading-6 text-gray-600'>
                                    <label for='file-upload'
                                           class='relative cursor-pointer rounded-md bg-white font-semibold text-indigo-600 focus-within:outline-none focus-within:ring-2 focus-within:ring-indigo-600 focus-within:ring-offset-2 hover:text-indigo-500'>
                                        <span>Загрузите файл</span>
                                        <input id='file-upload' name='file-upload' type='file' class='sr-only' />
                                    </label>
                                    <p class='pl-1'>или перетащите</p>
                                </div>
                                <p class='text-xs leading-5 text-gray-600'>PNG, JPG, GIF до 10MБ</p>
                            </div>
                        </div>
                    </div>
                </div>
            </div>


            <div class="mt-6 flex items-center justify-end gap-x-6">
<!--                <button type="button" class="text-sm font-semibold leading-6 text-gray-900">Отмена</button>-->
                <button type="submit" class="rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">Сохранить</button>
            </div>
        </div>
    </div>

</template>

<script lang='ts' setup>
import {PhotoIcon, UserCircleIcon} from '@heroicons/vue/24/solid'

</script>