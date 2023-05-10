<script setup lang='ts'>
import { IUserMembership, useUserStore } from '../../store/user.store'
import { Cog6ToothIcon } from '@heroicons/vue/24/outline'

// const props = defineProps<{ membership: IUserMembership }>()
const emits = defineEmits(['pickTenant'])

const { getOrganizationList, active_tenant, pickTenant: _pickTennant } = useUserStore()

const pickTenant = (membership: IUserMembership) => {
    emits('pickTenant', membership)
}

</script>

<template>
    <div class='flex w-full gap-3'>
        <div v-if='getOrganizationList' v-for='membership in getOrganizationList'
             class='w-1/5 shadow-md p-3 rounded-md hover:shadow-xl relative'>
            <div v-if='membership.organization.avatar'>
                <img :src='membership.organization.avatar' alt=''
                     class='w-full h-full absolute left-0 top-0 object-cover rounded-md'>
                <div class='w-full h-full absolute left-0 top-0 bg-purple opacity-30 rounded-md'></div>
                <div class='z-50 text-white relative'>
                    <h2 class='font-bold z-20 mb-0'>{{ membership.organization.title }}</h2>
                    <span class='block text-sm z-20'>Вы: {{ membership.role.title }}</span>
                    <div class='flex mt-3 gap-2'>
                        <button
                            :disabled='active_tenant === membership.organization.domain'
                            @click='pickTenant(membership)'
                            class='z-50 text-gray-500 bg-white font-bold px-5 py-2 rounded-xl w-4/5 text-sm lowercase'>
                            Войти
                        </button>
                        <button title='настройки'
                                class='z-50 bg-white text-gray-700 rounded-xl text-sm lowercase w-1/5 flex justify-center items-center border'>
                            <Cog6ToothIcon class='h-[20px]' />
                        </button>
                    </div>
                </div>
            </div>
            <div v-else>
                <h2 class='font-bold'>{{ membership.organization.title }}</h2>
                <span class='block text-sm'>Вы: {{ membership.role.title }}</span>
                <div class='flex mt-3 gap-2'>
                    <button
                        @click='pickTenant(membership)'
                        :disabled='active_tenant === membership.organization.domain'
                        :class='{"bg-gray-400": active_tenant === membership.organization.domain}'
                        class='z-50 text-white bg-purple px-5 py-2 rounded-xl font-bold w-4/5 text-sm lowercase'>
                        Войти
                    </button>
                    <button title='настройки'
                            class='z-50 text-gray-700 rounded-xl text-sm lowercase w-1/5 flex justify-center items-center border'>
                        <Cog6ToothIcon class='h-[20px]' />
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>

</style>