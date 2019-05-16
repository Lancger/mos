import request from '@/utils/request'
// 获取工单类型列表
export function getTicketTypeList(query) {
  return request({
    url: '/ticket/TicketTypeList',
    method: 'get',
    params: query
  })
}

export function getTicketTypeMsg() {
  return request({
    url: '/ticket/TicketTypeMsg',
    method: 'get'
  })
}
// 增加工单类型
export function addTicketType(data) {
  return request({
    url: '/ticket/TicketTypeAdd',
    method: 'post',
    data: data
  })
}

// 增加工单类型
export function updateTicketType(data) {
  return request({
    url: '/ticket/TicketTypeUpdate',
    method: 'post',
    data: {
      id: data.id,
      name: data.name
    }
  })
}

// 删除工单类型
export function deleteTicketType(id) {
  return request({
    url: '/ticket/TicketTypeDelete',
    method: 'post',
    data: {
      id: id
    }
  })
}

// 获取工单类型列表
export function getTicketLevelList() {
  return request({
    url: '/ticket/TicketLevelList',
    method: 'get'
  })
}

// 增加工单等级
export function addTicketLevel(data) {
  return request({
    url: '/ticket/TicketLevelAdd',
    method: 'post',
    data: data
  })
}

// 更新工单等级
export function updateTicketLevel(data) {
  return request({
    url: '/ticket/TicketLevelUpdate',
    method: 'post',
    data: {
      id: data.id,
      name: data.name,
      level: data.level,
      time: data.time
    }
  })
}

// 删除工单等级
export function deleteTicketLevel(id) {
  return request({
    url: '/ticket/TicketLevelDelete',
    method: 'post',
    data: {
      id: id
    }
  })
}

// 获取工单来源列表
export function getTicketSourceList() {
  return request({
    url: '/ticket/TicketSourceList',
    method: 'get'
  })
}

// 增加工单来源
export function addTicketSource(data) {
  return request({
    url: '/ticket/TicketSourceAdd',
    method: 'post',
    data: data
  })
}

// 更新工单等级
export function updateTicketSource(data) {
  return request({
    url: '/ticket/TicketSourceUpdate',
    method: 'post',
    data: {
        id: data.id,
        name: data.name
    }
  })
}

// 删除工单来源
export function deleteTicketSource(id) {
  return request({
    url: '/ticket/TicketSourceDelete',
    method: 'post',
    data: {
      id: id
    }
  })
}

// 获取工单来源列表
export function getTicketList(query_data) {
  return request({
    url: '/ticket/TicketList',
    method: 'post',
    data: query_data
  })
}

// 添加工单
export function addTicket(data) {
  return request({
    url: '/ticket/TicketAdd',
    method: 'post',
    data: data
  })
}

// 删除工单
export function deleteTicket(id) {
  return request({
    url: '/ticket/TicketDelete',
    method: 'post',
    data: {
      id: id
    }
  })
}

// 更新工单
export function updateTicket(data) {
  return request({
    url: '/ticket/TicketUpdate',
    method: 'post',
    data: data
  })
}


// 工单信息
export function fetchTicket(id) {
  return request({
    url: '/ticket/TicketMsg',
    method: 'get',
    params: {
      id: id
    }
  })
}

export function sendTicket(data, status_type) {
  return request({
    url: '/ticket/TicketSend',
    method: 'post',
    data: {
      id: data.id,
      // number: data.ticket_number,
      content: data.content,
      from: data.from,
      to: data.to,
      action: status_type
    }
  })
}

export function ctlTicket(id, ctl, user, content) {
  return request({
    url: '/ticket/TicketCtl',
    method: 'post',
    data: {
      id: id,
      ctl: ctl,
      user: user,
      content: content
    }
  })
}

