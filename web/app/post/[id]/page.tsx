import { getPostDetailAPI } from "@/api/post";
import Md from "@/components/Md";
import { Toc } from "@/components/Toc";
import { getTableOfContents } from "@/lib/toc";
import { ScrollShadow } from "@heroui/scroll-shadow";
import { Metadata } from "next";
import { title } from "process";

export default async function PostPage({ params }: { params: Promise<{ id: string }> }) {
  const { id } = await params;
  const post = await getPostDetailAPI(id);
  const toc = await getTableOfContents(post.data.content);
  return (
    <div className="relative max-w-7xl mx-auto px-4">
      <div className="flex flex-col justify-end lg:flex-row gap-12">
        {/* Markdown 内容 */}
        <div className="w-full lg:w-2/3">
          <h1 className="text-4xl font-bold mb-10">{post.data.title}</h1>
          <Md content={post.data.content} />
        </div>
        {/* 目录 */}
        <div className="hidden lg:block lg:w-1/5">
          <div className="sticky top-32">
            <ScrollShadow className="w-full max-h-[600px]" size={20} hideScrollBar  >
              <Toc toc={toc} />
            </ScrollShadow>
          </div>
        </div>
      </div>
    </div>
  );
}

export const metadata: Metadata = {
  title: "文章详情",
  description: "文章详情",
};
