import request from '@/utils/request'

// 获取用户组列表
export function getGroupList(query) {
    return request({
      url: '/sys/GroupList',
      method: 'get',
      params: query
    })
  }

// 增加用户组
export function addGroup(group_data) {
  return request({
    url: '/sys/GroupAdd',
    method: 'post',
    data: {
      group_name: group_data.group_name,
      nick_name: group_data.nick_name,
      user_selected: group_data.user_selected
    }
  })
}
// 删除用户组
export function deleteGroup(id) {
  return request({
    url: '/sys/GroupDelete',
    method: 'post',
    data: {
      id: id
    }
  })
}

// 用户组下用户列表
export function getGroupMsg(id) {
  return request({
    url: '/sys/GroupMsg',
    method: 'get',
    params: {
      id: id
    }
  })
}

export function updateGroup(group_data) {
  return request({
    url: '/sys/GroupUpdate',
    method: 'post',
    data: {
      id: group_data.id,
      nick_name: group_data.nick_name,
      user_selected: group_data.user_selected
    }
  })
}


