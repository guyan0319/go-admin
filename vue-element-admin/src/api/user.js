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
