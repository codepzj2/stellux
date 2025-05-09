import { Page, PageVO } from "@/types/page";
import { PostVO } from "@/types/post";
import request from "@/utils/request";

// 获取首页文章列表
export const getPostDetailListAPI = (page: Page) => {
  return request.get<PageVO<PostVO>>("/post/detail/list", page);
};

// 获取文章详情
export const getPostDetailAPI = (id: string) => {
  return request.get<PostVO>(`/post/detail/${id}`);
};
