import { Separator } from "@/components/ui/separator";
import { Book, Clock } from "lucide-react";

import request from "@/utils/request";
import StelluxMarkdown from "@/components/markdown/md";

import { timeAgo, readTime } from "@/utils/time";
import { PostVO } from "@/types/posts";

export default async function PostPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = await params;
  const post = await request.get<PostVO>(`/posts/${id}`);
  const { data } = post;
  const { words, minutes } = readTime(data.content);
  return (
    <div className="max-w-[850px] w-full m-auto p-6 my-4">
      <h1 className="text-4xl font-bold my-4">{data.title}</h1>
      <div className="h-4 flex justify-start items-center gap-3 text-sm text-gray-500">
        <span className="text-gray-700 dark:text-slate-400">{data.author}</span>
        <span>{timeAgo(data.created_at)}</span>
        <span className="flex items-center gap-1">
          <Book size={14} /> {words}字
        </span>
        <span className="flex items-center gap-1">
          <Clock size={14} /> {minutes}分钟
        </span>
      </div>
      <Separator className="mt-4" />
      <StelluxMarkdown content={data.content} />
    </div>
  );
}
