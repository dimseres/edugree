<script setup lang='ts'>

import { UserCircleIcon } from '@heroicons/vue/24/solid'
import { reactive, ref } from 'vue'

const props = defineProps<{
    image: string|null,
    label: string,
    accept: string
}>()

const emit = defineEmits(['fileChanged'])

const output = reactive({
    src: null,
    file: null
})

const objUrl = ref()

const handleChange = (event) => {
    let file = null
    if (event.target.files) {
        URL.revokeObjectURL(objUrl.value)
        file = event.target.files[0]
        objUrl.value = URL.createObjectURL(file)

        output.src = objUrl.value
        output.file = file
    } else {
        output.src = null
        output.file = null
    }

    emit("fileChanged", output)


    // console.log('ss')
    // debugger
}

const clear = (event) => {
    handleChange(event)
}

</script>

<template>
    <label for='photo' class='block text-sm font-medium leading-6 text-gray-900'>{{ label }}</label>
    <div class='mt-2 flex items-center gap-x-3'>
        <img class='h-12 w-12 rounded-full' v-if='props.image' :src='props.image' alt=''>
        <UserCircleIcon class='h-12 w-12 text-gray-300' aria-hidden='true' v-else/>
        <input id='avatar-upload' name='avatar-upload' type='file' class='sr-only' @change='handleChange' :accept='accept'/>
        <label for='avatar-upload'
               class='cursor-pointer rounded-md bg-white px-2.5 py-1.5 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50'>
            Изменить
        </label>

    </div>
    <div class='text-right mt-2'>
        <button v-if='props.image' @click='clear' class='rounded bg-red-400 px-3 py-0.5 text-white text-sm'>Очистить</button>
    </div>
</template>

<style scoped>

</style>