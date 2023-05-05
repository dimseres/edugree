<script setup lang='ts'>
import InviteUserModalList from './InviteUserModalList.vue'
import { reactive, ref } from 'vue'

const props = defineProps<{
    accept: Function
}>()

const userList = reactive([])

const roleList = [
    {
        label: "owner",
        value: 1
    },
    {
        label: "admin",
        value: 2
    }
]

</script>

<template>
    <InviteUserModalList v-model='userList' :role-list='roleList'/>
    <div class='flex justify-center mt-6'>
        <button class='bg-purple text-white py-2 px-3 rounded' @click='accept'>пригласить</button>
    </div>
</template>

<style scoped>

</style>