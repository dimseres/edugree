<script setup lang='ts'>
import { IUserMembership } from '../../store/user.store'

const props = defineProps<{membership: IUserMembership}>()
const emits = defineEmits(['pickTenant'])

const pickTenant = (membership: IUserMembership) => {
    emits("pickTenant", membership)
}

</script>

<template>
    <div>
        <div v-if='membership.organization.avatar'>
            <img :src='membership.organization.avatar' alt=''
                 class='w-full h-full absolute left-0 top-0 object-cover'>
            <div class='w-full h-full absolute left-0 top-0 bg-purple opacity-30'></div>
            <div class='z-50 text-white relative'>
                <h2 class='font-bold z-20 mb-0'>{{ membership.organization.title }}</h2>
                <span class='block text-sm z-20'>Вы: {{ membership.role.title }}</span>
                <button
                    @click='pickTenant(membership)'
                    class='mt-3 z-50 text-gray-500 bg-white font-bold px-5 py-2 rounded-xl w-full text-sm lowercase'>
                    Войти
                </button>
            </div>
        </div>
        <div v-else>
            <h2 class='font-bold'>{{ membership.organization.title }}</h2>
            <span class='block text-sm'>Вы: {{ membership.role.title }}</span>
            <button
                @click='pickTenant(membership)'
                class='mt-3 z-50 text-white bg-purple px-5 py-2 rounded-xl font-bold w-full text-sm lowercase'>
                Войти
            </button>
        </div>
    </div>
</template>

<style scoped>

</style>