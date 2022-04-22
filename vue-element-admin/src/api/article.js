import request from '@/utils/request'
import { httphost } from '@/utils/global'
export function fetchList(query) {
  return request({
    url: '',
    method: 'get',
    params: query,
    baseURL: httphost + '/articles/list'
  })
}

export function fetchArticle(id) {
  return request({
    url: '',
    method: 'get',
    params: { id },
    baseURL: httphost + '/articles/detail'
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
    baseURL: httphost + '/articles/create'
  })
}

export function updateArticle(data) {
  return request({
    url: '',
    method: 'post',
    data,
    baseURL: httphost + '/articles/edit'
  })
}

export function delImage(url) {
  return request({
    url: '',
    method: 'get',
    params: { url },
    baseURL: httphost + '/del/image'
  })
}
