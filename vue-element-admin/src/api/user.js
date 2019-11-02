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
    url: '',
    method: 'get',
    params: { token },
    baseURL: 'http://localhost:8090/info'
  })
}

export function logout() {
  return request({
    url: '',
    method: 'post',
    baseURL: 'http://localhost:8090/logout'
  })
}
