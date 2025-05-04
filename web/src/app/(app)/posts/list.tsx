import { getPostsList } from "@/api/posts";
import PostsCard from "@/components/card/posts";
import EmptySvg from "@/assets/svg/empty.svg";
import Image from "next/image";
import type { PostVO } from "@/types/posts";
export default async function PostsListPage() {
  const posts = await getPostsList({ page_no: 1, page_size: 10 });
  const postsList = posts?.data?.list || [];
  return (
    <div className="w-full md:w-4/5 max-w-screen-md mx-auto px-4 py-8 space-y-4">
      {postsList.length > 0 ? (
        postsList.map((post: PostVO) => <PostsCard key={post.id} post={post} />)
      ) : (
        <div className="w-full h-[calc(100vh-10em)] flex flex-col justify-center items-center">
          <Image src={EmptySvg} alt="404" className="w-3/5" />
          <div className="text-xl md:text-2xl font-bold">暂未发布文章</div>
        </div>
      )}
    </div>
  );
}
