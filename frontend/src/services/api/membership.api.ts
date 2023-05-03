import { axios } from './axios.config'
import { POSITION, useToast } from 'vue-toastification'

export interface IOrganizationFetchParams {
    page?: number
}

const toasted = useToast()

export async function fetchOrganizationMembers(params: IOrganizationFetchParams) {
    try {
        const {data} = await axios.get("/membership/users", {params})
        if (data.error) {
            toasted.error(data.message, {position: POSITION.BOTTOM_RIGHT})
        }
        return data
    } catch (e) {

    }
}