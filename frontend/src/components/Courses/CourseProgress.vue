<script setup lang='ts'>
const props = defineProps<{
    progress: any,
    rounded?: any
}>()
</script>

<template>
    <div class='course-progress'>
        <div class='px-4 text-white'>
            <div class='absolute flex h-full bg-purple left-0' :class='{"rounded-md": rounded}' :style='{width: progress + "%"}'></div>
            <div class='left-4 top-0.5 absolute w-full text-xs'>Завершено на {{ progress }}%</div>
        </div>
    </div>
</template>

<style scoped>

</style>