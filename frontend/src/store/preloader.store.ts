import { defineStore } from 'pinia'

interface IPreloadState {
    activeRequests: number,
}

export const usePreloaderStore = defineStore('preloader', {
    state: (): IPreloadState => {
        return {
            activeRequests: 0,
        }
    },
    getters: {
        isActive(): boolean {
            return this.activeRequests > 0
        },
    },
    actions: {
        incrementRequest() {
            this.activeRequests++
        },
        decrementRequest() {
            this.activeRequests--
        }
    },
})