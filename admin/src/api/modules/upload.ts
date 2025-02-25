import request from "@/utils/axios";
import type { Response } from "@/api/interfaces/resp";
import type { UploadFile } from "ant-design-vue/lib/upload/interface";
export const uploadPicturesLocal: (data: {
  files: UploadFile[];
}) => Promise<Response<any>> = ({ files }: { files: UploadFile[] }) => {
  const formData = new FormData();
  files.forEach((file: any) => {
    formData.append("files", file);
  });
  return request.post("/picture/upload", formData);
};
