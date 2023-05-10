import { axios, setXOrg } from './axios.config'
import { useUserStore } from '../../store/user.store'
import { POSITION, useToast } from 'vue-toastification'


export interface IApiResponse {
    error: boolean;
    message?: string | null;
    payload?: any | null;
}

interface LoginFormDTO {
    email: string,
    password: string
}

interface ErrorDataDTO {
    error: boolean,
    message: string,
}

const toasted = useToast()

export async function login(form: LoginFormDTO): Promise<IApiResponse> {
    try {
        const { data } = await axios.post('/auth/login', form)
        if (data.error) {
            return {
                error: true,
                message: data.message,
            }
        }

        const { setUser, setDefaultTenant } = useUserStore()
        setUser(data.user)
        setDefaultTenant()

        if (data.membership) {

        }

        return {
            error: false,
            message: null,
        }
    } catch (e: any) {
        return formatErrorRequest(e)
    }
}

export async function refresh() {
    try {
        const {data} = await axios.post('/auth/refresh')
        if (data.error) {
            return false
        }
        return true
    } catch (e) {
        return false
    }
}

export function logout() {

}

export interface ISignForm {
    email: string
    phone: string
    full_name: string
    password: string
    repeat_password: string
}

export async function signIn(form: ISignForm) {
    try {
        const {data} = await axios.post("/auth/register", form)
        if (data.error) {
            return data
        }

        const { setUser } = useUserStore()
        debugger
        setUser(data.user)

        return data.user
    } catch (e) {
        if (e.response) {
            return {
                error: e.response.error,
                message: e.response.message
            }
        }
        return e
    }
}

export async function getProfile(): Promise<IApiResponse> {
    try {
        const { data } = await axios.get('/users/profile')
        if (data.error) {
            return {
                error: true,
                message: data.message,
            }
        }
        return {
            error: false,
            payload: {
                user: data.user,
                invites: data.invites
            },
        }
    } catch (e: any) {
        return formatErrorRequest(e)
    }
}

function formatErrorRequest(e: any): IApiResponse {
    if (e.response && e.response.data) {
        return {
            error: true,
            message: e.response.data.message,
        }
    }
    return {
        error: true,
        message: e.message,
    }
}

export async function setTenant(tenantId: number) {
    try {
        const {data} = await axios.post("/auth/setTenant", {
            tenant_id: tenantId
        })
        if (data.error) {
            toasted.error(data.message, {position: POSITION.BOTTOM_RIGHT})
            return null
        }
        setXOrg(data.data.organization.domain)

        return data
    } catch (e: any) {
        toasted.error(e, {position: POSITION.BOTTOM_RIGHT})
    }
}