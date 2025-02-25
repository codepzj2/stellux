export interface PostsReq {
  title: string;
  content: string;
  author: string;
  category: string;
  tags: string[];
  cover: string;
  isTop: boolean;
}
