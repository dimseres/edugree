<template>
  <h2 class="mt-6 text-center text-2xl font-black tracking-tight text-gray-500">Регистрация</h2>
  <form class="mt-8 space-y-6" @submit.prevent='register'>
    <input type="hidden" name="remember" value="true"/>
    <div class="-space-y-px rounded-md shadow-sm">
      <div>
        <input id="email-address" v-model='email' name="email" type="email" autocomplete="email" required
               class="focus:outline-none p-3 border relative block w-full rounded-t-md border-0 py-2 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:z-10 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
               placeholder="Ваша почта"/>
      </div>
      <div>
        <input id="phone" v-model='phone' name="phone" type="tel" autocomplete="email" required
               class="focus:outline-none p-3 border relative block w-full border-0 py-2 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:z-10 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
               placeholder="Телефон"/>
      </div>
      <div>
        <input id="name" v-model='fullName' name="name" type="text" autocomplete="email" required
               class="focus:outline-none p-3 border relative block w-full border-0 py-2 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:z-10 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
               placeholder="Фамилия Имя Отчество"/>
      </div>
      <div>
        <input id="password" v-model='password' name="password" type="password" autocomplete="email" required
               class="focus:outline-none p-3 border relative block w-full border-0 py-2 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:z-10 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
               placeholder="Пароль"/>
      </div>
      <div>
        <input id="password_confirm" v-model='repeatPassword' name="password_confirm" type="password" autocomplete="current-password" required
               class="focus:outline-none p-3 border relative block w-full rounded-b-md border-0 py-2 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:z-10 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
               placeholder="Пароль еще раз"/>
      </div>
    </div>

    <div class="mt-8">
      <button type="submit"
              class="border bg-purple-500 group relative flex w-full justify-center rounded-md text-white py-2 px-3 text-sm font-semibold text-white hover:bg-white hover:border hover:text-purple-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-purple-500">
        Зарегистрироваться
      </button>
      <router-link :to="{name: 'Login'}" class="mt-3 border group relative flex w-full justify-center rounded-md text-purple-500 py-2 px-3 text-sm font-semibold focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-purple-500">
        У меня уже есть аккаунт
      </router-link>
    </div>
    <div class='text-red text-xs' v-if='errorString'>{{ errorString }}</div>
  </form>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { signIn } from '../../services/api/auth.api.vue'
import { useRouter } from 'vue-router'
import { POSITION, useToast } from 'vue-toastification'

const baseApi = import.meta.env.VITE_BASE_API

const email = ref("")
const phone = ref("")
const fullName = ref("")
const password = ref("")
const repeatPassword = ref("")

const errorString = ref("")

const router = useRouter()
const toasted = useToast()

const register = async () => {
  const result = await signIn({
    email: email.value,
    phone: phone.value,
    full_name: fullName.value,
    password: password.value,
    repeat_password: repeatPassword.value
  })
  if (result.error) {
    errorString.value = result.message
  }
  await router.push('/')
  toasted.success("Добро пожаловать", {position: POSITION.BOTTOM_RIGHT})
}

</script>
<style lang="ts" scoped>

</style>