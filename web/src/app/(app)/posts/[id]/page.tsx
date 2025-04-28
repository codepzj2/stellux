import { Separator } from "@/components/ui/separator";
import { Book, Clock } from "lucide-react";

import request from "@/utils/request";
import Md from "@/components/md";
import { getTableOfContents } from "@/lib/toc";

import { timeAgo, readTime } from "@/utils/time";
import { PostVO } from "@/types/posts";
import TocButton from "@/components/button/toc-button";

export default async function PostPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = await params;
  const post = await request.get<PostVO>(`/post/${id}`);
  const { data } = post;
  const { words, minutes } = readTime(data.content);
  const toc = await getTableOfContents(data.content);

  return (
    <div className="relative p-6 my-4 max-w-screen-xl mx-auto">
      <TocButton className="fixed right-4 top-1/2 -translate-y-1/2" toc={toc} />
      {/* 主内容区 */}
      <div className="max-w-[700px] w-full mx-auto">
        <h1 className="text-4xl font-bold my-4">{data.title}</h1>
        <div className="h-4 flex justify-start items-center gap-3 text-sm text-gray-500">
          <span className="text-gray-700 dark:text-slate-400">
            {data.author}
          </span>
          <span>{timeAgo(data.created_at)}</span>
          <span className="flex items-center gap-1">
            <Book size={14} /> {words}字
          </span>
          <span className="flex items-center gap-1">
            <Clock size={14} /> {minutes}分钟
          </span>
        </div>
        <Separator className="my-4" />
        <Md content={data.content} />
      </div>

      {/* 目录（固定在右上角） */}
    </div>
  );
}
