export interface PostVO {
  id: string;
  created_at: string;
  title: string;
  content: string;
  description: string;
  author: string;
  category: string;
  tags: string[];
  thumbnail: string;
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
  thumbnail?: string;
  category?: string;
  tags?: string[];
}
