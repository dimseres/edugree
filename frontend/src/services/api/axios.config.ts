import { usePreloaderStore } from '../../store/preloader.store'
import _axios, { AxiosInstance, AxiosResponse, CreateAxiosDefaults, InternalAxiosRequestConfig, request } from 'axios'
import {login, refresh} from "./auth.api.vue";


export class AxiosClient {
    private static axios: AxiosInstance

    static getInstance(config: any) {
        if (!AxiosClient.axios) {
            AxiosClient.axios = _axios.create(config)

            let lastRequest: InternalAxiosRequestConfig[] = [];

            AxiosClient.axios.interceptors.request.use((request: InternalAxiosRequestConfig): any => {
                lastRequest.push(request)
                return request
            })

            AxiosClient.axios.interceptors.request.use((config) => {
                return config
            }, (error) => {
                return Promise.reject(error)
            })

            AxiosClient.axios.interceptors.response.use((response: AxiosResponse<any, any>): AxiosResponse<any, any> => {
                return response
            }, async ({response}: any) => {
                const lastReq = lastRequest.pop()

                if (response.status === 401) {
                    if (lastReq && lastReq.url === lastReq.baseURL + "/auth/refresh") {
                        return response
                    }
                    const tokenTaken = await refresh()
                    if (!tokenTaken) {
                        return response
                    }
                }
                return response
            })
        }

        return AxiosClient.axios
    }

    static setXOrg(org: string) {
        AxiosClient.axios.defaults.headers.common["X-Org"] = org
    }
}

const config = {
    baseURL: import.meta.env.VITE_BASE_API ?? import.meta.env.BASE_URL,
    headers: {
        common: {
            "X-Org" : window.localStorage.getItem('tenant')
        }
    },
    withCredentials: true,
}

// export const axios = _axios.create()
export const axios = AxiosClient.getInstance(config)

// export function setXOrg(org: string) {
//     axios.defaults.headers.common["X-Org"] = org
// }
