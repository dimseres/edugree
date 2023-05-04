<template>
    <h2 class='mt-6 text-center text-2xl font-medium tracking-tight text-gray-700'>Введите логин и пароль</h2>
    <form class='mt-8 space-y-6' @submit.prevent='onSubmit'>
        <input type='hidden' name='remember' value='true' />
        <div class='-space-y-px rounded-md shadow-sm'>
            <div>
                <input id='email-address' name='email' type='email' autocomplete='email' required
                       v-model='form.email'
                       class='focus:outline-none p-2 border relative block w-full rounded-t-md border-0 py-1.5 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:z-10 focus:ring-2 focus:ring-inset focus:ring-gray-300 sm:text-sm sm:leading-6'
                       placeholder='Ваша почта или почта организации' />
            </div>
            <div>
                <input id='password' name='password' type='password' autocomplete='current-password' required
                       v-model='form.password'
                       class='focus:outline-none p-2 border relative block w-full rounded-b-md border-0 py-1.5 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:z-10 focus:ring-2 focus:ring-inset focus:ring-gray-300 sm:text-sm sm:leading-6'
                       placeholder='Пароль' />
            </div>
        </div>

        <div class='flex items-center justify-between'>
            <div class='flex items-center'>
                <input id='remember-me' name='remember-me' type='checkbox'
                       class='h-4 w-4 rounded bg-purple-500 text-indigo-600 focus:ring-indigo-600' />
                <label for='remember-me' class='ml-2 block text-sm text-gray-900'>Запомнить</label>
            </div>

            <div class='text-sm'>
                <router-link :to="{name: 'PasswordReset'}" class='font-medium text-gray-500 hover:text-indigo-500'>Забыли
                    пароль?
                </router-link>
            </div>
        </div>

        <div class='mt-auto'>
            <button type='submit'
                    class='border bg-purple-500 group relative flex w-full justify-center rounded-md text-white py-2 px-3 text-sm font-semibold text-white hover:text-purple hover:border hover:bg-white focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-purple-500'>
            <span class='absolute inset-y-0 left-0 flex items-center pl-3'>
              <LockClosedIcon class='h-5 w-5 text-white group-hover:text-purple-500' aria-hidden='true' />
            </span>
                Войти
            </button>
            <router-link :to="{name: 'Registration'}"
                         class='mt-3 group relative flex w-full justify-center rounded-md text-purple-500 py-2 px-3 text-sm font-semibold focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-purple-500'>
                Зарегистрироваться
            </router-link>
        </div>
    </form>
</template>

<script setup lang='ts'>
import { LockClosedIcon } from '@heroicons/vue/20/solid'
import { login } from '../../services/api/auth.api.vue'
import { ref } from 'vue'
import { useToast } from 'vue-toastification'
import { useRouter } from 'vue-router'
import { useUserStore } from '../../store/user.store'

const toast = useToast()

const form = ref({
    email: 'test1@example.net',
    password: 'admin',
})

const router = useRouter()
const state = useUserStore()

const onSubmit = async () => {
    const result = await login(form.value)
    if (result.error) {
        if (result.message) {
            toast.error(result.message)
        }
    }
    if (!state.tenant_role) {
        await router.push({ name: 'OrganizationChoose' })
    } else {
        await router.push('/')
    }
}


</script>
<style lang='ts' scoped>

</style>