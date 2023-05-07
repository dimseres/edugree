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

export async function getInviteConstants() {
    try {
        const {data} = await axios.get("/membership/invites/create")
        return data
    } catch (e) {
        if (e.response) {
            toasted.error(e.response.message, {position: POSITION.BOTTOM_RIGHT})
        }
        console.error(e)
    }
}

export interface IInviteUserList {
    role: string,
    email: string
}

export async function inviteUsers(userList: IInviteUserList[]) {
    try {
        const {data} = await axios.post("/membership/invites", {members: userList})
        return data
    } catch (e) {
        if (e.response) {
            toasted.error(e.response.message, {position: POSITION.BOTTOM_RIGHT})
        }
        console.error(e)
    }
}