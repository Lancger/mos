import request from '@/utils/request'

export function login(data) {
  return request({
    // url: '/account/account_login',
    url: '/UserLogin',
    method: 'post',
    data
  })
}

export function getInfo(token) {
  return request({
    url: '/sys/AccountInfo',
    method: 'get',
    params: { token }
  })
}

export function logout() {
  return request({
    url: '/sys/UserLogout',
    method: 'post'
  })
}


export function getAccountList(query) {
  return request({
    url: '/sys/UserList',
    method: 'get',
    params: query
  })
}

export function deleteAccount(id) {
  return request({
    url: '/sys/UserDelete',
    method: 'post',
    data: {
      id: id
    }
  })
}

export function addAccount(account_data) {
  return request({
    url: '/sys/UserAdd',
    method: 'post',
    data: {
      username: account_data.username,
      password: account_data.password,
      nick_name: account_data.nick_name,
      email: account_data.email,
      phone: account_data.phone
    }
  })
}

export function getAccountMsg(query) {
  return request({
    url: '/sys/UserMsg',
    method: 'get',
    params: query
  })
}

export function updateAccount(account_data) {
  return request({
    url: '/sys/UserUpdate',
    method: 'post',
    data: {
      id: account_data.id,
      username: account_data.username,
      password: account_data.password,
      nick_name: account_data.nick_name,
      email: account_data.email,
      phone: account_data.phone
    }
  })
}