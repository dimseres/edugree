<template>
    <Disclosure as='nav' class='shadow' v-slot='{ open }'>
        <div class='px-2 sm:px-6 lg:px-8'>
            <div class='relative flex h-16 items-center justify-between'>
                <div class='absolute inset-y-0 left-0 flex items-center sm:hidden'>
                    <!-- Mobile menu button-->
                    <DisclosureButton
                        class='inline-flex items-center justify-center rounded-md p-2 text-gray-400 hover:bg-gray-100 hover:text-white focus:outline-none focus:ring-2 focus:ring-inset focus:ring-white'>
                        <span class='sr-only'>Open main menu</span>
                        <Bars3Icon v-if='!open' class='block h-6 w-6' aria-hidden='true' />
                        <XMarkIcon v-else class='block h-6 w-6' aria-hidden='true' />
                    </DisclosureButton>
                </div>
                <div class='flex flex-1 items-center justify-center sm:items-stretch sm:justify-start'>
                    <div class='flex flex-shrink-0 items-center'>
                        <img class='block h-8 w-auto lg:hidden' src='/src/assets/logo.svg' alt='Your Company' />
                        <img class='hidden h-8 w-auto lg:block' src='/src/assets/logo.svg' alt='Your Company' />
                    </div>
                    <div class='hidden sm:ml-6 sm:block'>
                        <div class='flex space-x-4'>
                            <router-link v-for='item in navigation' class='page-link' :key='item.name' :to='item.href'
                                         :class="['text-gray-700 hover:text-purple-500', 'rounded-md px-3 py-2 text-sm font-medium tracking-wide']"
                                         :aria-current="item.current ? 'page' : undefined">
                                <div class='flex'>
                                    <component :is='item.icon' class='w-4 mr-1' />
                                    {{ item.name }}
                                </div>
                            </router-link>
                        </div>
                    </div>
                </div>
                <div class='absolute inset-y-0 right-0 flex items-center pr-2 sm:static sm:inset-auto sm:ml-6 sm:pr-0'>
                    <Notification/>

                    <!-- Profile dropdown -->
                    <Menu as='div' class='relative ml-3'>
                        <div>
                            <MenuButton
                                class='flex rounded-full bg-gray-800 text-sm focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800'>
                                <span class='sr-only'>Open user menu</span>
                                <img class='h-8 w-8 rounded-full'
                                     src='https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80'
                                     alt='' />
                            </MenuButton>
                        </div>
                        <transition enter-active-class='transition ease-out duration-100'
                                    enter-from-class='transform opacity-0 scale-95'
                                    enter-to-class='transform opacity-100 scale-100'
                                    leave-active-class='transition ease-in duration-75'
                                    leave-from-class='transform opacity-100 scale-100'
                                    leave-to-class='transform opacity-0 scale-95'>
                            <MenuItems
                                class='absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-md bg-white py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none'>
                                <MenuItem v-slot='{ active }'>
                                    <a href='#'
                                       :class="[active ? 'bg-gray-100' : '', 'block px-4 py-2 text-sm text-gray-700']">Профиль</a>
                                </MenuItem>
                                <MenuItem v-slot='{ active }'>
                                    <a href='#'
                                       :class="[active ? 'bg-gray-100' : '', 'block px-4 py-2 text-sm text-gray-700']">Выйти</a>
                                </MenuItem>
                                <MenuItem v-slot='{ active }'>
                                    <a href='#'
                                       :class="[active ? 'bg-gray-100' : '', 'block px-4 py-2 text-sm text-gray-700']">Сменить организцию</a>
                                </MenuItem>
                            </MenuItems>
                        </transition>
                    </Menu>
                </div>
            </div>
        </div>

        <DisclosurePanel class='sm:hidden'>
            <div class='space-y-1 px-2 pt-2 pb-3'>
                <DisclosureButton v-for='item in navigation' :key='item.name' as='a' :href='item.href'
                                  :class="[item.current ? 'bg-gray-500 text-white' : 'text-gray-300 hover:bg-gray-700 hover:text-white', 'block rounded-md px-3 py-2 text-base font-medium']"
                                  :aria-current="item.current ? 'page' : undefined">{{ item.name }}
                </DisclosureButton>
            </div>
        </DisclosurePanel>
    </Disclosure>
</template>

<script lang='ts' setup>
import { Disclosure, DisclosureButton, DisclosurePanel, Menu, MenuButton, MenuItem, MenuItems } from '@headlessui/vue'
import { Bars3Icon, BellIcon, BriefcaseIcon, PlusIcon, CircleStackIcon, UserIcon, XMarkIcon } from '@heroicons/vue/24/outline'
import { useUserStore } from '../../store/user.store'
import { computed } from 'vue'
import Notification from './notification/Notification.vue'

const {user, tenant_role} = useUserStore()

const userWithoutMembership = !user.membership

const _navigation = [
    {
        name: 'Пользователи',
        href: '/users',
        current: false,
        icon: UserIcon,
        roles: ['owner', 'administrator', 'manager'],
    },
    {
        name: 'Организация',
        href: '/organizations',
        current: false,
        icon: BriefcaseIcon,
        roles: ['owner', 'administrator', 'manager'],
    },
    {
        name: 'Сервисы',
        href: '/services',
        current: false,
        icon: CircleStackIcon,
        roles: ['owner', 'administrator', 'manager'],
    },
]

const emptyMemberNavigation = [
    {
        name: 'Создать организацию',
        href: '/organizations/create',
        current: false,
        icon: PlusIcon,
    },
]

const navigation = computed(() => {
    if (userWithoutMembership) {
        return emptyMemberNavigation
    }
    return _navigation.filter((item) => {
        return item.roles && item.roles.includes(tenant_role)
    })
})

</script>

<style lang='scss'>
.page-link {

&
.router-link-exact-active {
    color: theme('colors.purple.500');
}

}
</style>