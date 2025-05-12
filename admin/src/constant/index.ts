export const API_BASE_URL = import.meta.env.VITE_API_URL as string; // 后端服务地址
export const FILE_BASE_URL = API_BASE_URL + "/images"; // 文件服务地址
export const ALLOWED_IMAGE_TYPES = [
  "image/png",
  "image/jpeg",
  "image/jpg",
  "image/gif",
  "image/svg+xml",
  "image/x-icon",
  "image/avif",
  "image/webp",
] as const;
