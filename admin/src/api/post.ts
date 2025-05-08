import type { PostReq, PostDetailVO, PostVO } from "@/types/post";
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
  return request.put(`/admin-api/post/update`, data);
};

// 根据id获取文章
export const getPostByIdAPI: (id: string) => Promise<Response<PostVO>> = id => {
  return request.get(`/post/${id}`);
};

// 分页获取文章详情列表
export const getPostDetailListAPI: (
  params: PageReq
) => Promise<PageResponse<PostDetailVO>> = params => {
  return request.get("/post/detail/list", { params });
};

// 根据id获取文章详情
export const getPostDetailByIdAPI: (
  id: string
) => Promise<Response<PostDetailVO>> = id => {
  return request.get(`/post/detail/${id}`);
};

// 软删除文章
export const softDeletePostAPI: (id: string) => Promise<Response<any>> = id => {
  return request.delete(`/admin-api/post/soft-delete/${id}`);
};
