<template>
    <div class='dashboard-wrapper'>
        <Navbar />
        <div class='p-6'>
            <router-view></router-view>
        </div>
    </div>
</template>

<script lang='ts'>
import Navbar from '../navbar/Navbar.vue'

export default {
    name: 'DashboardTemplate',
    components: { Navbar },
}
</script>

<style scoped>

</style>