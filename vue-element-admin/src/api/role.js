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
    baseURL: httphost + '/role/list'
  })
}

export function addRole(data) {
  return request({
    url: '',
    method: 'post',
    data,
    baseURL: httphost + '/role/add'
  })
}

export function updateRole(id, data) {
  return request({
    url: ``,
    method: 'post',
    data,
    baseURL: httphost + '/role/update'
  })
}

export function deleteRole(data) {
  return request({
    url: ``,
    method: 'post',
    data,
    baseURL: httphost + `/role/delete`
  })
}

export function getAllRole() {
  return request({
    url: ``,
    method: 'get',
    baseURL: httphost + `/role/index`
  })
}
