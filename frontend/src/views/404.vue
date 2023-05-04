<template>
  <auth-template>
    <div class="flex min-h-full items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
      <div class="w-full">
        <h1 class="text-4xl inline-block px-10 bg-white">404 Такой страницы нет :(</h1>
        <p class="mt-0 text-md px-10 bg-white mt-3">скорее всего опечатка в адресе</p>
        <router-link :to="link" class="mt-0 cursor-pointer text-sm ml-10 px-10 py-2 inline-block mt-3 text-white bg-gray-500" >Назад</router-link>
      </div>
    </div>
  </auth-template>
</template>

<script setup lang="ts">
import {LockClosedIcon} from '@heroicons/vue/20/solid'
import AuthTemplate from "../components/templates/AuthTemplate.vue";
import {computed} from "vue";
import {useRouter} from "vue-router";

const router = useRouter()

const link = computed((): string => {
  return ""+router.options.history.state.back ?? '/'
})

</script>

<style lang="scss">
.form-container {
  min-height: 250px;
}
</style>