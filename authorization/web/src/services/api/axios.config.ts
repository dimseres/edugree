import _axios, {AxiosResponse, InternalAxiosRequestConfig} from "axios";
import {login, refresh} from "./auth.api.vue";

export const axios = _axios.create({
    baseURL: import.meta.env.VITE_BASE_API ?? import.meta.env.BASE_URL,
    headers: {
        "x-org": ""
    },
    withCredentials: true,
})

let lastRequest: InternalAxiosRequestConfig[] = [];

axios.interceptors.request.use((request: InternalAxiosRequestConfig): any => {
    lastRequest.push(request)
    return request
})

axios.interceptors.response.use((response: AxiosResponse<any, any>): AxiosResponse<any, any> => {
    return response
}, async (response: any) => {
    const lastReq = lastRequest.pop()

    if (response.status === 401) {
        if (lastReq && lastReq.url === lastReq.baseURL + "/auth/refresh") {
            return response
        }

        return await refresh()
    }

    return response
})