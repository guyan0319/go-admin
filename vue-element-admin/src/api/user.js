import request from '@/utils/request'
import { httphost } from '@/utils/global'
export function login(data) {
  return request({
    url: '',
    method: 'post',
    baseURL: httphost + '/login',
    data
  })
}

export function getInfo(token) {
  return request({
    url: '',
    method: 'get',
    params: { token },
    baseURL: httphost + '/info'
  })
}

export function logout() {
  return request({
    url: '',
    method: 'post',
    baseURL: httphost + '/logout'
  })
}

export function getAuthMenu(token) {
  return request({
    url: '/dashboard',
    method: 'get',
    baseURL: httphost + '',
    params: { token }

  })
}
export function fetchList(data) {
  return request({
    url: '',
    method: 'get',
    params: data,
    baseURL: httphost + '/user'
  })
}
export function createUser(data) {
  return request({
    url: '',
    method: 'post',
    data,
    baseURL: httphost + '/user/create'
  })
}
export function editUser(data) {
  return request({
    url: '',
    method: 'post',
    data,
    baseURL: httphost + '/user/edit'
  })
}
export function repasswdUser(data) {
  return request({
    url: '',
    method: 'post',
    data,
    baseURL: httphost + '/user/repasswd'
  })
}
export function fetchUser(id) {
  return request({
    url: '',
    method: 'get',
    params: { id },
    baseURL: httphost + '/user/detail'
  })
}
export function deleteUser(id) {
  return request({
    url: '',
    method: 'get',
    params: { id },
    baseURL: httphost + '/user/delete'
  })
}
