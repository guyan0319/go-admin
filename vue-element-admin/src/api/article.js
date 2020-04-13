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
    url: '',
    method: 'get',
    params: { id },
    baseURL: httphost + '/article/detail'
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
    url: '',
    method: 'post',
    data,
    baseURL: httphost + '/article/create'
  })
}

export function updateArticle(data) {
  return request({
    url: '',
    method: 'post',
    data,
    baseURL: httphost + '/article/edit'
  })
}
