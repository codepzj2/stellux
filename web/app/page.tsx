import { getPostDetailListAPI } from "@/api/post";
import { PostCard } from "@/components/PostCard";

export default async function Home() {
  const posts = await getPostDetailListAPI({
    page_no: 1,
    page_size: 10,
    post_type: "publish",
  });
  const postsList = posts?.data?.list || [];
  return (
    <div className="flex flex-col gap-4 max-w-2xl mx-auto">
      {postsList.map((item) => (
        <PostCard className="w-full" key={item.id} post={item} />
      ))}
    </div>
  );
}

