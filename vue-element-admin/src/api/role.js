import request from '@/utils/request'
import { httphost } from '@/utils/global'
export function getRoutes() {
  return request({
    url: '',
    method: 'get',
    baseURL: httphost + '/routes'
  })
}
export function getRoles() {
  return request({
    url: '',
    method: 'get',
    baseURL: httphost + '/roles'
  })
}

export function addRole(data) {
  return request({
    url: '',
    method: 'post',
    data,
    baseURL: httphost + '/addrole'
  })
}

export function updateRole(id, data) {
  return request({
    url: `/role/${id}`,
    method: 'put',
    data
  })
}
export function deleteRole(id) {
  return request({
    url: `/role/${id}`,
    method: 'delete'
  })
}
