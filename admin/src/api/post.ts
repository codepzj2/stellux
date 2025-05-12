import type { PostReq, PostDetailVO, PostVO } from "@/types/post";
import type { PageReq, PageResponse, Response } from "@/types/dto";
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

// 更新文章发布状态
export const updatePostPublishStatusAPI: (
  id: string,
  is_publish: boolean
) => Promise<Response<any>> = (id, is_publish) => {
  return request.put(`/admin-api/post/update/publish-status`, {
    id,
    is_publish,
  });
};

// 恢复文章
export const restorePostAPI: (id: string) => Promise<Response<any>> = id => {
  return request.put(`/admin-api/post/restore/${id}`);
};

// 批量恢复文章
export const restorePostBatchAPI: (
  id_list: string[]
) => Promise<Response<any>> = id_list => {
  return request.put(`/admin-api/post/restore/batch`, { id_list });
};

// 根据id获取文章
export const getPostByIdAPI: (id: string) => Promise<Response<PostVO>> = id => {
  return request.get(`/post/${id}`);
};

// 分页获取发布文章详情列表
export const getPublishPostDetailListAPI: (
  params: PageReq
) => Promise<PageResponse<PostDetailVO>> = params => {
  return request.get("/post/detail/list", { params });
};

// 分页获取草稿箱文章详情列表
export const getDraftPostDetailListAPI: (
  params: PageReq
) => Promise<PageResponse<PostDetailVO>> = params => {
  return request.get("/admin-api/post/draft/list", { params });
};

// 分页获取回收站文章详情列表
export const getBinPostDetailListAPI: (
  params: PageReq
) => Promise<PageResponse<PostDetailVO>> = params => {
  return request.get("/admin-api/post/bin/list", { params });
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

// 批量软删除文章
export const softDeletePostBatchAPI: (
  id_list: string[]
) => Promise<Response<any>> = id_list => {
  return request.delete(`/admin-api/post/soft-delete/batch`, {
    data: { id_list },
  });
};

// 硬删除文章
export const deletePostAPI: (id: string) => Promise<Response<any>> = id => {
  return request.delete(`/admin-api/post/delete/${id}`);
};

// 批量硬删除文章
export const deletePostBatchAPI: (
  id_list: string[]
) => Promise<Response<any>> = id_list => {
  return request.delete(`/admin-api/post/delete/batch`, {
    data: { id_list },
  });
};
