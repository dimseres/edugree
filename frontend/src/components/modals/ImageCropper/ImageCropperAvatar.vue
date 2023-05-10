<script setup lang='ts'>

import { CircleStencil, Cropper, Preview } from 'vue-advanced-cropper'
import { ref } from 'vue'

const cropImage = ref()
const cropCoordinates = ref()
const cropCanvas = ref()
const cropData = ref()
const cropper = ref()

const emit = defineEmits(['saveImage'])
const props = defineProps<{
    image: null|string
}>()

const onChange = ({ coordinates, canvas, image }) => {
    cropCoordinates.value = coordinates
    console.log(coordinates, canvas, image)
    // You able to do different manipulations at a canvas
    // but there we just get a cropped image, that can be used
    // as src for <img/> to preview result
    cropCanvas.value = canvas
    cropImage.value = image
}

const saveImage = () => {

    let file = null

    emit("saveImage", {
        canvas: cropCanvas,
        image: cropper.value.getResult().canvas.toDataURL(),
    })
}

</script>

<template>
    <div>
        111
        <div class='flex'>
            <Cropper class='cropper'
                     ref='cropper'
                     :debounce='true'
                     :src='image'
                     :stencil-component='CircleStencil'
                     :stencil-props='{
                    aspectRatio: 12/3,
             }'
                     :canvas='{
		            maxHeight: 256,
		            maxWidth: 256
	            }'
                     @change='onChange'
            >

            </Cropper>
            <div class='ml-3'>
                <Preview
                    :width='128'
                    :height='128'
                    :image='cropImage'
                    :coordinates='cropCoordinates'
                />
                <Preview
                    class='mt-3'
                    :width='64'
                    :height='64'
                    :image='cropImage'
                    :coordinates='cropCoordinates'
                />

                <Preview
                    class='mt-3'
                    :width='32'
                    :height='32'
                    :image='cropImage'
                    :coordinates='cropCoordinates'
                />

            </div>
        </div>

        <div class='mt-3 flex items-center justify-end gap-x-6'>
            <!--                <button type="button" class="text-sm font-semibold leading-6 text-gray-900">Отмена</button>-->
            <button type='submit'
                    :disabled='!cropImage'
                    @click='saveImage'
                    class='rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600'>
                Сохранить
            </button>
        </div>
    </div>
</template>

<style lang='scss' scoped>
.cropper {
    width: 300px;
}
</style>