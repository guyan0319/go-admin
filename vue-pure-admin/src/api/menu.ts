import { http } from "../utils/http";

interface ResponseType extends Promise<any> {
  data?: object;
  code?: number;
  msg?: string;
}

// 获取菜单列表列表
export const getMenuList = (data?: object): ResponseType => {
  return http.request("get", "/routes", { data });
};
