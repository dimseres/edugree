<template>
<!--    <ModalWrapper v-if='activeComponent' @close='activeComponent = null' :title='modalTitle'>-->
<!--        <component :is='activeComponent' v-bind='modalComponentProps'></component>-->
<!--    </ModalWrapper>-->
    <div class='flex justify-center'>
        <div class='container'>
            <div class='head'>
                <h1 class='text-2xl text-gray-700 font-bold'>Организации</h1>
                <p class='text-gray-500 font-light text-sm'>Здесь находятся все ваши организации, а также вы можете их
                    создать</p>
            </div>
            <div class='mt-3'>
                <div v-if='getOrganizationList' class='flex gap-3 w-full'>
                    <OrganizationCardExtended @pickTenant='pickTenant'></OrganizationCardExtended>
<!--                    <div v-for='membership in getOrganizationList' class='w-1/5 shadow-md p-3 rounded-md hover:shadow-xl'>-->
<!--                    </div>-->
                </div>
                <div v-else>
                    <h2 class='text-gray-500'>
                        <ExclamationCircleIcon />
                        Вы пока не состоите ни в одной организации
                    </h2>
                </div>
            </div>
            <div class='mt-6'>
                <div class='flex'>
                    <router-link :to='{name: "OrganizationCreate"}' class='p-3 text-gray-700 shadow rounded-md bg-white hover:bg-purple hover:text-white'><PlusIcon class='h-[20px] inline'/>создать организацию</router-link>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang='ts' setup>
import { IUserMembership, useUserStore } from '../store/user.store'
import { useRouter } from 'vue-router'
import { setTenant } from '../services/api/auth.api.vue'
import { ExclamationCircleIcon, PlusIcon } from '@heroicons/vue/24/solid'
import OrganizationCardExtended from '../components/organization/OrganizationCardExtended.vue'
import ModalWrapper from '../components/modals/ModalWrapper.vue'
import { ref } from 'vue'
import CreateOrganizationModal from '../components/modals/organization/CreateOrganizationModal.vue'

const { getOrganizationList, pickTenant: _pickTennant } = useUserStore()

const router = useRouter()

const pickTenant = async (tenant: IUserMembership) => {
    const data = await setTenant(tenant.organization.id)
    _pickTennant(tenant)
    await router.go(0)
}

const activeComponent = ref()
const modalTitle = ref()
const modalComponentProps = ref()


const collectForm = () => {
    alert()
}


const openCreateOrganizationModal = () => {
    activeComponent.value = CreateOrganizationModal
    modalTitle.value = 'Создание организации'
    modalComponentProps.value = {
        accept: collectForm,
    }
}

</script>