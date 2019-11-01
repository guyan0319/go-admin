import request from '@/utils/request'

export function login(data) {
  return request({
    url: '',
    method: 'post',
    baseURL: 'http://localhost:8090/login',
    data
  })
}

export function getInfo(token) {
  return request({
    url: '/user/info',
    method: 'get',
    baseURL: 'http://localhost:8090/login',
    params: { token }
  })
}

export function logout() {
  return request({
    url: '/user/logout',
    method: 'post'
  })
}
