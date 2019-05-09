import request from '@/utils/request'
// 权限列表请求
export function getPermList(id) {
    return request({
        url: '/sys/PermissionList',
        method: 'get',
    })
  }

export function getPermMsg(query) {
    return request({
        url: '/sys/PermissionMsg',
        method: 'get',
        params: query
    })
}  


// 增加权限
export function addPerm(data) {
    return request({
        url: '/sys/PermissionAdd',
        method: 'post',
        data: {
            name: data.name,
            nick_name: data.nick_name,
            type: data.type
        }
    })
}

// 删除权限
export function deletePerm(id) {
    return request({
        url: '/sys/PermissionDelete',
        method: 'post',
        data: {
        id: id
        }
    })
}

// 增加权限
export function updatePerm(data) {
    return request({
        url: '/sys/PermissionUpdate',
        method: 'post',
        data: {
            id: data.id,
            name: data.name,
            nick_name: data.nick_name,
            type: data.type
        }
    })
}
  