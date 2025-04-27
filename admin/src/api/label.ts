import request from "@/utils/request";
import type {
  CreateLabelReq,
  EditLabelReq,
  LabelPageReq,
  LabelVO,
} from "@/types/label";
import type { PageResponse, Response } from "@/types/response";

export const getLabelListAPI: (
  params: LabelPageReq
) => Promise<PageResponse<LabelVO>> = params => {
  return request.get("/label/list", { params });
};

export const getLabelTypeByIdAPI: (
  id: string
) => Promise<Response<LabelVO>> = id => {
  return request.get(`/label/type/${id}`);
};

export const createLabelAPI: (
  data: CreateLabelReq
) => Promise<Response<any>> = data => {
  return request.post("/admin-api/label/create", data);
};

export const editLabelAPI: (
  data: EditLabelReq
) => Promise<Response<any>> = data => {
  return request.put("/admin-api/label/edit", data);
};

export const deleteLabelAPI: (id: string) => Promise<Response<any>> = id => {
  return request.delete(`/admin-api/label/delete/${id}`);
};
