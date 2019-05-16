import store from '@/store'

/**
 * @param {Array} value
 * @returns {Boolean}
 * @example see @/views/permission/directive.vue
 */
export default function checkPermission(value) {
  if (value && value instanceof Array && value.length > 0) {
    const roles = store.getters && store.getters.roles
    const permissionRoles = value

    const hasPermission = roles.some(role => {
      return permissionRoles.includes(role)
    })

    if (!hasPermission) {
      return false
    }
    return true
  } else {
    console.error(`need roles! Like v-permission="['admin','editor']"`)
    return false
  }
}

export function checkTicketWorkPerm(value) {
  if (value && value instanceof Array && value.length > 0) {
    const perm = store.getters && store.getters.perms
    const permissionRoles = value

    const hasPermission = perm.some(role => {
      return permissionRoles.includes(role)
    })
    if (!hasPermission) {
      return false
    }
    return true
  } else {
    console.error(`need roles! Like v-permission="['admin','editor']"`)
    return false
  }
}

export function checkPermStage(value, stage, label) {
  if (value && value instanceof Array && value.length > 0) {
    const perm = store.getters && store.getters.perms
    const permissionRoles = value

    const hasPermission = perm.some(role => {
      return permissionRoles.includes(role)
    })

    if (!hasPermission) {
      return false
    }
  } else {
    console.error(`need roles! Like v-permission="['admin','editor']"`)
    return false
  }
  if (label && label instanceof Array && label.length > 0) {
    for (var i = 0; i <= label.length; i++) {
      if (stage === label[i]) {
        return true
      }
    }
    return false
  } else {
    console.error(`need label! Like "['新建工单','驳回']"`)
  }
}

export function checkPermNotStage(value, stage, label) {
  if (value && value instanceof Array && value.length > 0) {
    const perm = store.getters && store.getters.perms
    const permissionRoles = value

    const hasPermission = perm.some(role => {
      return permissionRoles.includes(role)
    })

    if (!hasPermission) {
      return false
    }
  } else {
    console.error(`need roles! Like v-permission="['admin','editor']"`)
    return false
  }
  if (label && label instanceof Array && label.length > 0) {
    for (var i = 0; i <= label.length; i++) {
      if (stage === label[i]) {
        return false
      }
    }
    return true
  } else {
    console.error(`need label! Like "['新建工单','驳回']"`)
  }
}

export function check_stage(stage, value) {
  if (stage === value) {
    return true
  } else {
    return false
  }
}

export function check_bool(status) {
  if (status === true) {
    return true
  }
  return false
}
