import request from '@/utils/request'
import { httphost } from '@/utils/global'
export function getMenus() {
  return request({
    url: '',
    method: 'get',
    baseURL: httphost + '/menu'
  })
}
export function addRole(data) {
  return request({
    url: '',
    method: 'post',
    data,
    baseURL: httphost + '/menu'
  })
}
