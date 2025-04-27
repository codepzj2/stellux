export interface PostReq {
  title: string;
  content: string;
  author: string;
  description: string;
  category: string | undefined;
  tags: string[];
  is_top: boolean;
  is_publish: boolean;
  thumbnail: string;
}

export interface PostUpdateReq extends PostReq {
  id: string;
}

export interface PostVO {
  id?: string;
  created_at: string;
  updated_at: string;
  title: string;
  content: string;
  author: string;
  description: string;
  category: string;
  tags: string[];
  thumbnail: string;
  is_top: boolean;
  is_publish: boolean;
}

export interface PostUpdateStatusReq {
  id: string;
  is_publish: boolean;
}

export interface PostLabel {
  label: string;
  value: string;
}
