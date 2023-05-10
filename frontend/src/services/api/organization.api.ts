import { POSITION, useToast } from 'vue-toastification'
import {axios} from './axios.config'

const toasted = useToast()

export interface IOrganizationCreateForm {
    title: string,
    domain: string,
    email: string,
    description: string|null
}

export async function createOrganization(form: IOrganizationCreateForm) {
    try {
        const {data} = await axios.post("/organization/create", form)
        if (data.error) {
            toasted.error(data.message)
            return null
        }
        return data.data
    } catch (e: any) {
        if (e.response && e.response.message) {
            return toasted.error(e.response.message, {
                position: POSITION.BOTTOM_RIGHT
            })
        }

        toasted(e.message)
    }
}