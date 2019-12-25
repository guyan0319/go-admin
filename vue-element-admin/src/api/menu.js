import request from '@/utils/request'
import { httphost } from '@/utils/global'
export function getMenus() {
  return request({
    url: '',
    method: 'get',
    baseURL: httphost + '/menu'
  })
}
export function addMenu(data) {
  return request({
    url: '',
    method: 'post',
    data,
    baseURL: httphost + '/menu'
  })
}
export function updateMenu(id, data) {
  return request({
    url: '',
    method: 'put',
    data,
    baseURL: httphost + '/menu'
  })
}
export function deleteMenu(data) {
  return request({
    url: '',
    method: 'delete',
    data,
    baseURL: httphost + '/menu'
  })
}
