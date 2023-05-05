<script setup lang='ts'>

import { CheckIcon, ChevronUpDownIcon } from '@heroicons/vue/24/solid'
import { Listbox, ListboxButton, ListboxOption, ListboxOptions } from '@headlessui/vue'
import { ref, watch } from 'vue'

const selected = ref()

const emit = defineEmits(["update:modelValue"])

watch(selected, () => {
    emit("update:modelValue", selected)
})

const props = withDefaults(defineProps<{
    modelValue: any,
    placeholder: string,
    options: Array<{
        value: string | number | boolean,
        label: string
    } | string>
}>(), {
    placeholder: 'выберите роль',
})
</script>

<template>
    <Listbox as='div' class='absolute w-[150px] right-0' v-model='selected'>
        <div class='relative'>
            <ListboxButton
                class='relative w-full cursor-default rounded-md bg-white py-1.5 pl-3 pr-10 text-left text-gray-900 focus:outline-none focus:ring-2 focus:ring-indigo-200 sm:text-sm sm:leading-6'>
                <span class='block truncate'>{{ modelValue ? modelValue.label ?? modelValue : placeholder }}</span>
                <span class='pointer-events-none absolute inset-y-0 right-0 flex items-center pr-2'>
                        <ChevronUpDownIcon class='h-5 w-5 text-gray-400' aria-hidden='true' />
                    </span>
            </ListboxButton>

            <transition leave-active-class='transition ease-in duration-100' leave-from-class='opacity-100'
                        leave-to-class='opacity-0'>
                <ListboxOptions
                    class='absolute z-10 mt-1 max-h-56 w-full overflow-auto rounded-md bg-white py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm'>
                    <ListboxOption as='template' v-for='(option, idx) in options' :key='idx' :value='option'
                                   v-slot='{ active, selected }'>
                        <li :class="[active ? 'bg-indigo-600 text-white' : 'text-gray-900', 'relative cursor-default select-none py-2 pl-3 pr-9']">
                            <div class='flex items-center'>
                                        <span
                                            :class="[selected ? 'font-semibold' : 'font-normal', 'ml-3 block truncate']">
                                            {{ option.label ?? option }}</span>
                            </div>

                            <span v-if='selected'
                                  :class="[active ? 'text-white' : 'text-indigo-600', 'absolute inset-y-0 right-0 flex items-center pr-4']">
                <CheckIcon class='h-5 w-5' aria-hidden='true' />
              </span>
                        </li>
                    </ListboxOption>
                </ListboxOptions>
            </transition>
        </div>
    </Listbox>
</template>

<style scoped>

</style>