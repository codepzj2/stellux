import type { PostsReq } from "@/api/interfaces/posts";
import type { Response } from "@/api/interfaces/resp";
import request from "@/utils/axios";
export const createPost: (data: PostsReq) => Promise<Response<any>> = (
  data
) => {
  console.log("data", data);
  return request.post("/posts/create", data);
};
