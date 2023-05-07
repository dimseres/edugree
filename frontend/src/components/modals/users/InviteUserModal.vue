<script setup lang='ts'>
import InviteUserModalList from './InviteUserModalList.vue'
import { onMounted, reactive, ref } from 'vue'
import { getInviteConstants } from '../../../services/api/membership.api'

const props = defineProps<{
    accept: Function
}>()

const userList = reactive([])

const roleList = reactive([
    // {
    //     label: "owner",
    //     value: 1
    // },
    // {
    //     label: "admin",
    //     value: 2
    // }
])

onMounted(async () => {
    const data = await getInviteConstants()
    if (data.data) {
        roleList.push(...data.data.roles.map(item => {
            return {
                label: item.title,
                value: item.slug
            }
        }))
    }
})

const invite = () => {
    const payload = userList.map(item => {
        return {
            role: item.role.value,
            email: item.email
        }
    }).filter(item => item.email)
    props.accept(payload)
}

</script>

<template>
    <InviteUserModalList v-model='userList' :role-list='roleList'/>
    <div class='flex justify-center mt-6'>
        <button class='bg-purple text-white py-2 px-3 rounded disabled:bg-gray-300' @click='invite' :disabled="userList.length == 0">пригласить</button>
    </div>
</template>

<style scoped>

</style>