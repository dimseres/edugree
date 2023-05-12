import { POSITION, useToast } from 'vue-toastification'
import { AxiosError } from 'axios'
import { axios } from './axios.config'


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

export async function getMyCourses() {
    try {
        const {data} = await axios.get("courses/courses/my")
        if (data.error) {
            toasted.error(data.message)
            return null
        }
        return data;
    } catch (e: any) {
        if (e.response && e.response.data?.error) {
            toasted.error(e.response.data.error)
            return null
        }
        toasted.error(e.message)
    }
}