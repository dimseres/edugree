import { defineStore } from 'pinia'

export interface IOrganization {
    id: number,
    title: string,
    domain: string,
    avatar: string | null
}

export interface IRole {
    id: number,
    name: string,
    slug: string
}

export interface IUserMembership {
    organization: IOrganization
    role: IRole
}

export interface IUser {
    id: number,
    full_name: string,
    phone: string,
    avatar: string,
    bio: string,
    membership: IUserMembership[]
}

interface IUserState {
    user: IUser | null,
    active_tenant: string | null,
    tenant_role: string | null
}

export const useUserStore = defineStore('user', {
    state: (): IUserState => {
        return {
            user: null,
            active_tenant: window.localStorage.getItem('tenant'),
            tenant_role: null,
        }
    },
    getters: {
        getOrganizationList(): IUserMembership[] | null {
            if (this.user) {
                return this.user.membership
            }
            return null
        },
    },
    actions: {
        setUser(user: IUser) {
            this.user = user
        },

        setDefaultTenant() {
            const tenant = window.localStorage.getItem('tenant');
            const tenantRole = window.localStorage.getItem('tenant_role');
            let validTenant = false
            if (this.user && tenant && tenantRole) {
                for (const member of this.user.membership) {
                    if (member.organization.domain === tenant) {
                        return this.pickTenant(member)
                    }
                }
            }
            return this.clearLocalStorage()
        },

        clearLocalStorage() {
            window.localStorage.clear()
        },

        pickTenant(membership: IUserMembership) {
            window.localStorage.setItem('tenant', membership.organization.domain)
            window.localStorage.setItem('tenant_role', membership.role.slug)
            this.active_tenant = membership.organization.domain
            this.tenant_role = membership.role.slug
        },
    },
})