import { POSITION, useToast } from 'vue-toastification'
import axios from 'axios/index'

const toasted = useToast()

export interface IOrganizationCreateForm {
    title: string,
    domain: string,
    email: string,
    description: string
}

export async function createOrganization(form: IOrganizationCreateForm) {
    try {
        const {data} = await axios.post("/organization/create", form)
        if (data.error) {
            toasted.error(data.error)
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