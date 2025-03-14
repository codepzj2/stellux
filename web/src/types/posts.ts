export interface PostVO {
  id: string;
  title: string;
  content: string;
  created_at: string;
  updated_at: string;
  author: string;
  description: string;
  category: string;
  tags: string[];
  cover: string;
  is_top: boolean;
  is_publish: boolean;
}

export interface IPostCard {
  id: string;
  title: string;
  content: string;
  description: string;
  created_at: string;
  author: string;
  cover?: string;
  category?: string;
  tags?: string[];
}
