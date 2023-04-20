import { axios } from './axios.config'
import { useUserStore } from '../../store/user.store'


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

export async function login(form: LoginFormDTO): Promise<IApiResponse> {
    try {
        const { data } = await axios.post('/auth/login', form)
        if (data.error) {
            return {
                error: true,
                message: data.message,
            }
        }

        const { setUser } = useUserStore()
        debugger
        setUser(data.user)

        return {
            error: false,
            message: null,
        }
    } catch (e: any) {
        return formatErrorRequest(e)
    }
}

export async function refresh() {
    await axios.post('/auth/refresh')
}

export function logout() {

}

export function signIn() {

}

export async function setTenant() {
    try {
        const { data } = await axios.get('/auth/setTenant')
        if (data.error) {
            return {
                error: true,
                message: data.message
            }
        }
        return {
            error: false,
            payload: data.data
        }
    } catch (e) {

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
            payload: data.user,
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