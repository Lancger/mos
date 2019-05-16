import request from '@/utils/request'
// 获取项目列表
export function getProjectList(query) {
    return request({
      url: '/project/ProjectList',
      method: 'get',
      params: query
    })
  }
export function getProjectMsg() {
  return request({
    url: '/project/ProjectMsg',
    method: 'get'
  })
}  

// 增加项目
export function addProject(data) {
  return request({
    url: '/project/ProjectAdd',
    method: 'post',
    data: data
  })
}

// 更新项目
export function updateProject(data) {
    return request({
      url: '/project/ProjectUpdate',
      method: 'post',
      data: {
          id: data.id,
          name: data.name
      }
    })
}

// 删除项目
export function deleteProject(id) {
    return request({
      url: '/project/ProjectDelete',
      method: 'post',
      data: {
        id: id
      }
    })
}