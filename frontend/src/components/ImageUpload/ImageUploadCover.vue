<script setup lang='ts'>

import { PhotoIcon } from '@heroicons/vue/24/solid'
import { reactive, ref } from 'vue'

const props = defineProps<{
    image: string | null,
    label: string,
    accept: string
}>()

const emit = defineEmits(['fileChanged'])

const output = reactive({
    src: null,
    file: null,
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

    emit('fileChanged', output)


    // console.log('ss')
    // debugger
}

const clear = (event) => {
    handleChange(event)
}

</script>

<template>
    <label for='cover-photo'
           class='block text-sm font-medium leading-6 text-gray-900'>{{ props.label }}</label>
    <div
        class='mt-2 flex justify-center rounded-lg border border-dashed border-gray-900/25 px-6 py-10 relative overflow-hidden'>
        <div class='text-center'>
            <img v-if='props.image' :src='image' alt='' class='object-cover absolute left-0 top-0 w-full h-full'>
            <PhotoIcon class='mx-auto h-12 w-12 text-gray-300' aria-hidden='true' :class='{"invisible": props.image}' />
            <div class='z-50' :class='{"invisible": props.image}'>
                <div class='mt-4 flex text-sm leading-6 text-gray-600'>
                    <label for='file-upload'
                           class='relative cursor-pointer rounded-md bg-white font-semibold text-indigo-600 focus-within:outline-none focus-within:ring-2 focus-within:ring-indigo-600 focus-within:ring-offset-2 hover:text-indigo-500'>
                        <span>Загрузите файл</span>
                        <input id='file-upload' name='file-upload' :accept='props.accept' type='file' class='sr-only'
                               @change='handleChange' />
                    </label>
                    <p class='pl-1'>или перетащите</p>
                </div>
                <p class='text-xs leading-5 text-gray-600'>PNG, JPG, GIF до 10MБ</p>
            </div>
        </div>
    </div>
    <div class='text-right mt-2'>
        <button v-if='props.image' @click='clear' class='rounded bg-red-400 px-3 py-0.5 text-white text-sm'>Очистить</button>
    </div>
</template>

<style scoped>

</style>