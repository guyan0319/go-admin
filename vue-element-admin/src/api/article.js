import request from '@/utils/request'
import { httphost } from '@/utils/global'
export function fetchList(query) {
  return request({
    url: '',
    method: 'get',
    params: query,
    baseURL: httphost + '/article/list'
  })
}

export function fetchArticle(id) {
  return request({
    url: '/article/detail',
    method: 'get',
    params: { id }
  })
}

export function fetchPv(pv) {
  return request({
    url: '/article/pv',
    method: 'get',
    params: { pv }
  })
}

export function createArticle(data) {
  return request({
    // url: '/article/create',
    // method: 'post',
    // data
    url: '',
    method: 'post',
    data,
    baseURL: httphost + '/article/create'
  })
}

export function updateArticle(data) {
  return request({
    url: '/article/update',
    method: 'post',
    data
  })
}
