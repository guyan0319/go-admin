import { http } from "../utils/http";

export const getMenuList = (params?: object) => {
  return http.request("get", "/routes", { params });
};
