import { defineStore } from 'pinia'

export interface INotification {
    type: string,
    data: string,
}

export interface IInvite {
    id: number,
    link: string,
    created_at: string,
    organization: {
        id: number,
        title: string,
        domain: string,
        email: string,
        description: string,
    },
    role: {
        id: number,
        title: string,
        slug: string,
        description: string
    }
}

interface INotificationState {
    notifications: null|any[],
    invites: null|IInvite[]
}

export const useNotificationStore = defineStore('notification', {
    state: (): INotificationState => {
        return {
            notifications: null,
            invites: null,
        }
    },
    getters: {
        getNotifications(): INotification[] | null {
            if (this.notifications) {
                return this.notifications
            }
            return null
        },
        getInvites(): INotification[] | null {
            if (this.invites) {
                return this.invites
            }
            return null
        },
    },
    actions: {
        setInvites(invites: any[]) {
            this.invites = invites
        },
    },
})