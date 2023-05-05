<script setup lang='ts'>
import { reactive } from 'vue'
import { MinusIcon, PlusIcon } from '@heroicons/vue/24/solid'
import Dropdown from '../../dropdown/Dropdown.vue'

const props = defineProps<{
    roleList: Array<{
        label: string,
        value: number
    }>
    modelValue: Array<{
        email: string,
        role: string
    }>
}>()

const clearAddedUser = () => {
    return { email: '', role: '' }
}

const addedUser = reactive({ email: '', role: '' })

const addUser = () => {
    props.modelValue.push({ ...addedUser })
    addedUser.email = ''
    addedUser.role = ''
}

const remove = (idx) => {
    props.modelValue.splice(idx, 1)
}

</script>

<template>
    <div class='my-3'>
        <div v-for='(invitor, idx) in modelValue' class='flex items-center'>
            <div class='relative w-[450px] border py-1 flex items-center'>
                <input placeholder='email участника' class='px-3 py-1 w-[250px] focus:outline-none' type='email'
                       v-model='invitor.email'>
                <Dropdown :options='roleList' v-model='invitor.role'></Dropdown>
            </div>
            <div>
                <button class='ml-3 w-[30px] h-[30px] block text-sm px-2 py-1 bg-red-100 rounded' @click='remove(idx)'>
                    <MinusIcon />
                </button>
            </div>
            <!--            <input class='border px-3 py-1' type='email' v-model='invitor.email'>-->
            <!--            <select v-model='invitor.role' class='border px-3 py-1'>-->
            <!--                <option disabled>роль</option>-->
            <!--                <option v-for='role in roleList'>{{ role.name }}</option>-->
            <!--            </select>-->

        </div>

        <div class='flex items-center'>
            <div class='relative w-[450px] border py-1 flex items-center'>
                <input placeholder='email участника' class='px-3 py-1 w-[250px] focus:outline-none' type='email'
                       v-model='addedUser.email'>
                <Dropdown :options='roleList' v-model='addedUser.role'></Dropdown>
            </div>
            <!--            <select v-model='addedUser.role' class='border px-3 py-1'>-->
            <!--            <option v-for='role in roleList'>{{ role.name }}</option>-->
            <!--        </select>-->
            <button class='ml-3 w-[30px] h-[30px] block text-sm px-2 py-1 bg-gray-100 rounded' @click='addUser'>
                <PlusIcon />
            </button>
        </div>
    </div>
</template>

<style scoped>

</style>