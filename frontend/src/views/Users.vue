<template>
    <div class='users flex justify-center'>
        <ModalWrapper v-if='activeComponent' @close='activeComponent = null' :title='modalTitle'>
            <component :is='activeComponent' v-bind='modalComponentProps'></component>
        </ModalWrapper>
        <div class='container'>
            <h1 class='text-2xl text-gray-700 font-bold'>Пользователи</h1>
            <p class='text-gray-500 font-light text-sm'>Здесь находятся все члены вашей организации, а также вы можете
                их
                пригласить</p>
            <div class='control-bar flex justify-between mt-3'>
                <div class='w-1/4 flex'>
                    <input id='email' name='email' type='email' autocomplete='email' required=''
                           class='px-3 block w-full rounded-l-md border-0 py-1.5 text-gray-200 shadow-sm ring-1 ring-inset ring-gray-50 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6' />
                    <button class='bg-gray-50 px-6 rounded-r-md'>поиск</button>
                </div>
                <div>
                    <button
                        @click='openInviteUsersModal'
                        class='h-full bg-purple-500 hover:bg-purple-700 text-sm text-white px-6 py-1 rounded-md mr-2 flex items-center'>
                        <PlusIcon class='h-4 text-white' />
                        добавить
                    </button>
                </div>
            </div>
            <div class='search mt-6'>

            </div>
            <div class='p-3 mt-3 border border-gray-100 rounded-md'>
                <table class='w-full' v-if='users'>
                    <thead class='text-left text-gray-700 bg-gray-100'>
                    <tr class='border-b border-gray-50 uppercase text-sm'>
                        <th class='p-2 w-[80px]'>id</th>
                        <th class='p-2'>телефон</th>
                        <th class='p-2'>фио</th>
                        <th class='p-2'>почта</th>
                        <th class='p-2'>роль</th>
                        <th class='p-2'></th>
                    </tr>
                    </thead>
                    <tbody>
                    <tr v-for='user in users' :key='user.id'
                        class='border-b text-gray-800 border-gray-50 hover:bg-gray-50 transition last:border-none font-light text-sm'>
                        <td class='p-2'>{{ user.id }}</td>
                        <td class='p-2'>{{ user.phone }}</td>
                        <td class='p-2'>{{ user.full_name }}</td>
                        <td class='p-2'>{{ user.email }}</td>
                        <td class='p-2'>
                            <div class='flex'>
                        <span
                            class='bg-gray-100 text-gray-800 text-xs font-medium mr-2 px-2.5 py-0.5 rounded-full dark:bg-gray-700 dark:text-gray-300'>
                        <StarIcon v-if='user.domain_role.slug === "owner"'
                                  class='inline w-[16px] text-yellow' /> {{ user.domain_role.title }}</span>
                            </div>
                        </td>
                        <td class='w-[30px]'>
                            <Menu as='div' class='relative ml-3'>
                                <div class='w-[30px] h-[30px]'>
                                    <MenuButton
                                        class='flex rounded-full bg-gray-50 text-sm focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800'>
                                        <span class='sr-only'>Open user menu</span>
                                        <EllipsisVerticalIcon class='w-full'></EllipsisVerticalIcon>
                                        <!--                                    <img class='h-8 w-8 rounded-full'-->
                                        <!--                                         src='https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80'-->
                                        <!--                                         alt='' />-->
                                    </MenuButton>
                                </div>
                                <transition enter-active-class='transition ease-out duration-100'
                                            enter-from-class='transform opacity-0 scale-95'
                                            enter-to-class='transform opacity-100 scale-100'
                                            leave-active-class='transition ease-in duration-75'
                                            leave-from-class='transform opacity-100 scale-100'
                                            leave-to-class='transform opacity-0 scale-95'>
                                    <MenuItems
                                        v-if='user.domain_role.slug !== "owner"'
                                        class='absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-md bg-white py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none'>
                                        <MenuItem v-slot='{ active }'>
                                            <a href='#'
                                               :class="[active ? 'bg-gray-100' : '', 'block px-4 py-2 text-sm text-gray-700']">Сменить
                                                роль</a>
                                        </MenuItem>
                                        <MenuItem v-slot='{ active }'>
                                            <a href='#'
                                               :class="[active ? 'bg-gray-100' : '', 'block px-4 py-2 text-sm text-gray-700']">Профиль</a>
                                        </MenuItem>
                                        <MenuItem
                                            v-slot='{ active }'>
                                            <a href='#'
                                               class='text-red'
                                               :class="[active ? 'bg-gray-100' : '', 'block px-4 py-2 text-sm text-gray-700']">Удалить</a>
                                        </MenuItem>
                                    </MenuItems>
                                </transition>
                            </Menu>
                        </td>
                    </tr>
                    </tbody>
                </table>
                <div class='flex justify-center text-2xl font-bold' v-else>
                    <h3>Нет данных</h3>
                </div>

                <div class='mt-2' v-if='Object.keys(pagination) && pagination.pages > 1'>
                    <Pagination v-model='currentPage' :pagination='pagination' />
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang='ts'>
import { EllipsisVerticalIcon, PlusIcon, StarIcon } from '@heroicons/vue/24/solid'
import { defineComponent, onMounted, ref, watch } from 'vue'
import { fetchOrganizationMembers, inviteUsers } from '../services/api/membership.api'
import { Menu, MenuButton, MenuItem, MenuItems } from '@headlessui/vue'
import Pagination from '../components/pagination/Pagination.vue'
import { useRoute } from 'vue-router'
import { router } from '../routes/routes'
import ModalWrapper from '../components/modals/ModalWrapper.vue'
import InviteUserModal from '../components/modals/users/InviteUserModal.vue'
import { POSITION, useToast } from 'vue-toastification'

const route = useRoute()

const users = ref([])
const pagination = ref({})
const currentPage = ref(route.query.page ?? 1)
const activeComponent = ref()
const modalTitle = ref()
const modalComponentProps = ref()

defineComponent({ Pagination })

watch(currentPage, async (newVal, oldVal) => {
    await router.push({
        replace: true,
        query: { page: newVal },
    })
    await fetchUsers()
})

const fetchUsers = async () => {
    const data = await fetchOrganizationMembers({ page: currentPage.value as number })
    if (!data.error) {
        users.value = data.data

        pagination.value = {
            total: data.total,
            pages: data.pages,
        }
    }
}

onMounted(async () => {
    await fetchUsers()
})

const toast = useToast()

const sendInviteUser = async (users) => {
    const result = await inviteUsers(users)
    if (result.error) {
        toast.error(result.message, { position: POSITION.BOTTOM_RIGHT })
    } else {
        toast.success('Успешно', { position: POSITION.BOTTOM_RIGHT })
    }
}


const openInviteUsersModal = () => {
    activeComponent.value = InviteUserModal
    modalTitle.value = 'Добавление участников'
    modalComponentProps.value = {
        accept: sendInviteUser,
    }
}

</script>

<style scoped>

</style>