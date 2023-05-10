<script setup lang='ts'>
import { XMarkIcon } from '@heroicons/vue/24/solid'

const props = withDefaults(defineProps<{
    title: string
}>(), {
    title: "Внимание"
})

</script>

<template>
    <div
        @click.self='$emit("close")'
        class='fixed flex items-center justify-center backdrop-blur-[5px] backdrop-saturate-0 bg-gray-500 w-screen h-screen z-[1000] left-0 top-0 bg-opacity-5'>
        <div class='bg-white p-3 min-w-[300px] rounded shadow-2xl'>
            <div class='flex font-black text-xl text-gray-700 text-right items-center justify-between'>
                {{ title }}
                <a class='cursor-pointer' @click='$emit("close")'><XMarkIcon class='text-gray-700 w-[30px]'></XMarkIcon>
            </a></div>
            <slot></slot>
        </div>
    </div>
</template>

<style scoped>

</style>