export interface Response<T> {
  code: number;
  msg: string;
  data: T;
}

export interface PageResponse<T> {
  code: number;
  msg: string;
  data: {
    page_no: number;
    page_size: number;
    total_count: number;
    total_page: number;
    list: T[];
  };
}

export interface PageReq {
  page_no: number;
  page_size: number;
  keyword?: string;
  field?: string;
  order?: "ASC" | "DESC";
}

export interface PageData<T> {
  page_no: number;
  page_size: number;
  total_count: number;
  total_page: number;
  list: T[];
}
