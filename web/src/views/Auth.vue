<template>
  <auth-template>
    <div class="flex min-h-full items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
      <div class="w-full max-w-md space-y-8 bg-white p-10 rounded-2xl shadow-2xl">
        <img class="mx-auto h-20 w-auto" src="/src/assets/logo.png"
             alt="Your Company"/>
        <div class="form-container">
          <router-view />
        </div>
      </div>
    </div>
  </auth-template>
</template>

<script setup lang="ts">
import {LockClosedIcon} from '@heroicons/vue/20/solid'
import AuthTemplate from "../components/templates/AuthTemplate.vue";</script>

<style lang="scss">
.form-container {
  min-height: 250px;
}
</style>