import type { IRouteGuard } from '../../routes/routes'
import { IUser, useUserStore } from '../../store/user.store'
import { getProfile } from '../api/auth.api.vue'

export class AuthGuard implements IRouteGuard {
  protected roles: string[] | null

  public constructor(roles: string[] | null = null) {
    this.roles = roles
  }

  async routeAllowed(route: any) {
    let message = 'вы не авторизованы'
    const { user, tenant_role, setUser } = useUserStore()
    if (user) {
      if (tenant_role && this.isRoleGranted(tenant_role)) {
        return {
          success: true,
        }
      } else {
        return {
          role_failed: true,
          success: false,
          message: 'нет прав',
        }
      }
    }

    const profile = await getProfile()
    if (profile.error) {
      return {
        success: false,
        message: profile.message as string,
      }
    }

    setUser(profile.payload)

    return {
      success: false
    }
  }

  private isRoleGranted(role: string) {
    if (this.roles && this.roles.length > 0) {
      if (!this.roles.includes(role)) {
        return false
      }
    }
    return true
  }
}
