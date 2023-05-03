import { defineStore } from 'pinia'
import { state } from 'vue-tsc/out/shared'

export interface IOrganization {
    id: number,
    title: string,
    domain: string,
    avatar: string|null
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
    user: IUser|null,
    active_tenant: string|null,
    tenant_role: string|null
}

export const useUserStore = defineStore("user", {
    state: (): IUserState => {
        return {
            user: null,
            active_tenant: window.localStorage.getItem("tenant"),
            tenant_role: null
        }
    },
    getters: {
        getOrganizationList(): IUserMembership[]|null {
            if (this.user) {
                return this.user.membership
            }
            return null;
        }
    },
    actions: {
        setUser(user: IUser) {
            this.user = user
        },

        pickTenant(membership: IUserMembership) {
            window.localStorage.setItem("tenant", membership.organization.domain)
            window.localStorage.setItem("tenant_role", membership.role.slug)
            this.active_tenant = membership.organization.domain
            this.tenant_role = membership.role.slug
        }
    }
})