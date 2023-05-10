<script lang='ts' setup>
import { PhotoIcon } from '@heroicons/vue/24/solid'
import ImageUpload from '../components/ImageUpload/ImageUpload.vue'
import ImageCropperAvatar from '../components/modals/ImageCropper/ImageCropperAvatar.vue'
import ImageCropperCover from '../components/modals/ImageCropper/ImageCropperCover.vue'
import ModalWrapper from '../components/modals/ModalWrapper.vue'
import { ref, watch } from 'vue'
import ImageUploadCover from '../components/ImageUpload/ImageUploadCover.vue'

const image = ref()

const showModal = ref(false)

const toggleModal = () => {
    showModal.value = !showModal.value
}

const avatar = ref()
const avatarOut = ref()
const croppedAvatar = ref()
const avatarImage = ref()
const modalComponent = ref()
const modalSaveFunc = ref()

const cover = ref()
const coverOut = ref()
const croppedCover = ref()
const coverImage = ref()

watch(avatar, (newValue) => {
    if (newValue.src) {
        image.value = newValue.src
        modalComponent.value = ImageCropperAvatar
        modalSaveFunc.value = saveCroppedImage
        showModal.value = true
    }
}, { deep: true })

watch(cover, (newValue) => {
    if (newValue.src) {
        image.value = newValue.src
        modalComponent.value = ImageCropperCover
        modalSaveFunc.value = saveCroppedCover
        showModal.value = true
    }
}, { deep: true })

const saveCroppedImage = (image) => {
    croppedAvatar.value = image
    if (avatar.value) {
        debugger
        avatarImage.value = croppedAvatar.value.image
        avatarOut.value = croppedAvatar
    }
    showModal.value = false
}

const saveCroppedCover = (image) => {
    croppedCover.value = image
    if (avatar.value) {
        debugger
        coverImage.value = croppedCover.value.image
        coverOut.value = croppedCover
    }
    showModal.value = false
}

const avatarUploaded = (file) => {
    avatar.value = file
    avatarImage.value = file.src
}

const coverUploaded = (file) => {
    cover.value = file
    coverImage.value = file.src
    debugger
}

</script>
<template>
    <ModalWrapper title='Выберите секцию' v-if='showModal' @close='toggleModal'>
        <component :is='modalComponent' :image='image' @saveImage='modalSaveFunc'/>
<!--        <ImageCropperAvatar :image='image' @saveImage='saveCroppedImage'></ImageCropperAvatar>-->
    </ModalWrapper>
    <div class='flex justify-center'>
        <div class='container w-1/3'>
            <div class='border-b border-gray-900/10 pb-12'>
                <h2 class='text-2xl text-gray-700 font-semibold leading-7'>Создание организации</h2>
                <p class='mt-1 text-sm leading-6 text-gray-600'>Данная информация полностью публичная</p>

                <div class='mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6'>
                    <div class='sm:col-span-4'>
                        <label for='username' class='block text-sm font-medium leading-6 text-gray-900'>Домен</label>
                        <div class='mt-2'>
                            <div
                                class='flex rounded-md shadow-sm ring-1 ring-inset ring-gray-300 focus-within:ring-2 focus-within:ring-inset focus-within:ring-indigo-600 sm:max-w-md'>
                                <input type='text' name='username' id='username' autocomplete='username'
                                       class='block flex-1 border-0 bg-transparent py-2 px-3 text-gray-900 placeholder:text-gray-400 focus:ring-0 sm:text-sm sm:leading-6'
                                       placeholder='janesmith' />
                            </div>
                            <p class='mt-3 text-sm leading-6 text-gray-600'>Домен должен быть уникальным</p>
                        </div>
                    </div>

                    <div class='col-span-full'>
                        <label for='about' class='block text-sm font-medium leading-6 text-gray-900'>Об
                            организации</label>
                        <div class='mt-2'>
                        <textarea id='about' name='about' rows='3'
                                  class='block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6' />
                        </div>
                        <p class='mt-3 text-sm leading-6 text-gray-600'>Напишите пару слов о вашей организации</p>
                    </div>

                    <div class='col-span-full'>
                        <ImageUpload label='Аватар'
                                     accept='.jpg,.png,.jpeg,.gif'
                                     :image='avatarImage'
                                     @fileChanged='avatarUploaded'
                        ></ImageUpload>
                    </div>

                    <div class='col-span-full'>
                        <ImageUploadCover :image='coverImage' label='Обложка' accept='.jpg,.png,.jpeg,.gif' @fileChanged='coverUploaded'></ImageUploadCover>
                    </div>
                </div>
            </div>


            <div class='mt-6 flex items-center justify-end gap-x-6'>
                <!--                <button type="button" class="text-sm font-semibold leading-6 text-gray-900">Отмена</button>-->
                <button type='submit'
                        class='rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600'>
                    Сохранить
                </button>
            </div>
        </div>
    </div>

</template>