<script setup lang='ts'>
import { BellIcon } from '@heroicons/vue/24/outline'
import { useNotificationStore } from '../../../store/notification.store'
import { Menu, MenuButton, MenuItem, MenuItems } from '@headlessui/vue'

const { getInvites } = useNotificationStore()

</script>

<template>
    <Menu as='div' class='relative ml-3'>
        <div class='relative'>
            <MenuButton
                class='p-1 w-8 flex rounded-full bg-gray-100 text-sm text-gray-700 focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800'>
                <span class='sr-only'>Open user menu</span>
                <BellIcon class='w-full' aria-hidden='true' />
            </MenuButton>
            <span class='absolute w-4 h-4 bg-red-300 rounded-full text-center text-xs -right-2 -top-2' v-if='getInvites?.length'>{{ getInvites?.length }}</span>
        </div>
        <transition enter-active-class='transition ease-out duration-100'
                    enter-from-class='transform opacity-0 scale-95'
                    enter-to-class='transform opacity-100 scale-100'
                    leave-active-class='transition ease-in duration-75'
                    leave-from-class='transform opacity-100 scale-100'
                    leave-to-class='transform opacity-0 scale-95'>
            <MenuItems
                class='absolute right-0 z-10 mt-2 w-64 origin-top-right rounded-md bg-white py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none'>
                <div v-if='getInvites'>
                    <MenuItem v-for='invite in getInvites' v-slot='{ active }'>
                        <div class='p-2' :class="{'bg-gray-50': active}">
                            <div class='text-sm text-gray-500'>Приглашение от:
                                <div class='text-gray-800'>{{ invite?.organization.title }}</div>
                            </div>
                            <div class='text-sm text-gray-500'>Роль:
                                <div class='text-gray-800'>{{ invite?.role.title }}</div>
                            </div>
                            <div class='flex text-center text-sm mt-2'>
                                <a href='' class='w-1/2 bg-purple p-1 rounded-md text-white hover:shadow-xl'>Принять</a>
                                <a href='' class='w-1/2 ml-2 p-1 bg-gray-300 rounded-md'>Отклонить</a>
                            </div>
                        </div>
                    </MenuItem>
                    <!--                    <MenuItem v-slot='{ active }'>-->
                    <!--                        <a href='#'-->
                    <!--                           :class="[active ? 'bg-gray-100' : '', 'block px-4 py-2 text-sm text-gray-700']">Выйти</a>-->
                    <!--                    </MenuItem>-->
                    <!--                    <MenuItem v-slot='{ active }'>-->
                    <!--                        <a href='#'-->
                    <!--                           :class="[active ? 'bg-gray-100' : '', 'block px-4 py-2 text-sm text-gray-700']">Сменить-->
                    <!--                            организцию</a>-->
                    <!--                    </MenuItem>-->
                </div>
            </MenuItems>
        </transition>
    </Menu>
</template>

<style scoped>

</style>