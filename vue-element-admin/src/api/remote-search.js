import request from '@/utils/request'
import { httphost } from '@/utils/global'

export function searchUser(name) {
  // return request({
  //   url: '/search/user',
  //   method: 'get',
  //   params: { name }
  // })
  return request({
    url: '',
    method: 'get',
    params: { name },
    baseURL: httphost + '/user/search'
  })
}

export function transactionList(query) {
  return request({
    url: '/transaction/list',
    method: 'get',
    params: query
  })
}
