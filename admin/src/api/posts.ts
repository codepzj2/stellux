import type { PostReq, PostUpdateStatusReq, PostVO } from "@/types/posts";
import type { PageReq, PageResponse, Response } from "@/types/response";
import request from "@/utils/request";

// 创建文章
export const createPostAPI: (
  data: PostReq
) => Promise<Response<any>> = data => {
  return request.post("/admin-api/post/create", data);
};

// 更新文章
export const updatePostAPI: (
  data: PostReq
) => Promise<Response<any>> = data => {
  return request.put("/admin-api/post/update", data);
};

// 根据id获取文章
export const getPostByIdAPI: (id: string) => Promise<Response<PostVO>> = id => {
  return request.get(`/admin-api/post/${id}`);
};

// 分页获取文章列表
export const getPostsByPageAPI: ({
  page_no,
  page_size,
}: PageReq) => Promise<PageResponse<PostVO>> = ({
  page_no = 1,
  page_size = 10,
  keyword,
  field,
  order,
}) => {
  return request.post("/admin-api/post/list", {
    page_no,
    page_size,
    keyword,
    field,
    order,
  });
};

// 更新文章状态
export const updatePostStatusAPI: (
  data: PostUpdateStatusReq
) => Promise<PageResponse<any>> = data => {
  return request.patch("/admin-api/post/status", data);
};

// 软删除文章
export const deletePostSoftAPI: (id: string) => Promise<Response<any>> = id => {
  return request.patch(`/admin-api/post/delete/${id}`);
};
