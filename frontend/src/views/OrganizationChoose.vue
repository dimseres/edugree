<template>
    <auth-template>
        <div class='flex items-center justify-center w-screen h-screen bg-white'>
            <div class='p-3'>
                <h1 class='font-black text-center text-2xl mb-4'>Выберите организацию</h1>
                <div class='flex max-w-[750px]'>
                    <div
                        class='organization w-[250px] relative m-2 p-3 shadow-xl hover:shadow-2xl transition duration-150 rounded-xl overflow-hidden'
                        v-for='organization in getOrganizationList'
                    >
                        <div v-if='organization.organization.avatar'>
                            <img :src='organization.organization.avatar' alt=''
                                 class='w-full h-full absolute left-0 top-0 object-cover'>
                            <div class='w-full h-full absolute left-0 top-0 bg-purple opacity-30'></div>
                            <div class='z-50 text-white relative'>
                                <h2 class='font-bold z-20 mb-0'>{{ organization.organization.title }}</h2>
                                <span class='block text-sm z-20'>Вы: {{ organization.role.title }}</span>
                                <button
                                    @click='pickTenant(organization)'
                                    class='mt-3 z-50 text-gray-500 bg-white font-bold px-5 py-2 rounded-xl w-full text-sm lowercase'>
                                    Войти
                                </button>
                            </div>
                        </div>
                        <div v-else>
                            <h2 class='font-bold'>{{ organization.organization.title }}</h2>
                            <span class='block text-sm'>Вы: {{ organization.role.title }}</span>
                            <button
                                @click='pickTenant(organization)'
                                class='mt-3 z-50 text-white bg-purple px-5 py-2 rounded-xl font-bold w-full text-sm lowercase'>
                                Войти
                            </button>
                        </div>
                    </div>
                </div>
                <div class='flex justify-center mt-6'>
                    <img class='w-14' src='/src/assets/logo.svg' alt='Your Company' />
                </div>
            </div>
        </div>
    </auth-template>
</template>

<script lang='ts' setup>
import AuthTemplate from '../components/templates/AuthTemplate.vue'
import { IUserMembership, useUserStore } from '../store/user.store'
import { useRouter } from 'vue-router'
import {setTenant} from "../services/api/auth.api.vue"

const { getOrganizationList, pickTenant: _pickTennant } = useUserStore()

const router = useRouter();

const pickTenant = async (tenant: IUserMembership) => {
    const data = await setTenant(tenant.organization.id)
    _pickTennant(tenant)
    await router.push("/")
}

</script>