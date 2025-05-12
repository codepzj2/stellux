import request from "@/utils/request";
import type { PageResponse, Response } from "@/types/dto";
import type { UploadFile } from "ant-design-vue/lib/upload/interface";
import type { FileVO } from "@/types/file";
import { ALLOWED_IMAGE_TYPES } from "@/constant";

const isValidImageType = (type: string): boolean => {
  return ALLOWED_IMAGE_TYPES.includes(type as any);
};

export const uploadFilesAPI = async ({
  files,
}: {
  files: UploadFile[];
}): Promise<Response<any>> => {
  const formData = new FormData();

  for (const file of files) {
    try {
      if (!file.type || !isValidImageType(file.type)) {
        throw new Error(
          `不支持的文件类型: ${file.type || "未知类型"}。仅支持: ${ALLOWED_IMAGE_TYPES.join(", ")}`
        );
      }
      formData.append("files", file as unknown as File);
    } catch (error) {
      console.error("文件处理失败:", file.name, error);
      throw error;
    }
  }
  return request.post("/admin-api/file/upload", formData);
};

export const queryFileList = ({
  page_no,
  page_size,
}: {
  page_no?: number;
  page_size?: number;
}): Promise<PageResponse<FileVO>> => {
  return request.get("/file/list", { params: { page_no, page_size } });
};

export const deleteFilesByIdListAPI = ({
  id_list,
}: {
  id_list: string[];
}): Promise<Response<any>> => {
  return request.delete("/admin-api/file/delete", { data: { id_list } });
};
