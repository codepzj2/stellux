import { getPostDetailAPI } from "@/api/post";
import Md from "@/components/Md";
import { Toc } from "@/components/Toc";
import { getTableOfContents } from "@/lib/toc";
import { ScrollShadow } from "@heroui/scroll-shadow";
import { Metadata } from "next";
import { headers } from "next/headers";

type Props = {
  params: Promise<{ id: string }>
}

export default async function PostPage({ params }: Props) {
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

export async function generateMetadata({ params }: Props): Promise<Metadata> {
  const { id } = await params;
  const post = await getPostDetailAPI(id);
  const data = post.data;

  const title = data.title || "默认标题";
  const description = data.description || "默认描述";
  const image = data.thumbnail || "/default-thumbnail.jpg";
  const keywords = [data.category, ...(data.tags || [])].filter(Boolean);

  const headerList = await headers();
  const host = headerList.get("host") || "localhost:3000";
  const protocol = host.startsWith("localhost") ? "http" : "https";
  const url = `${protocol}://${host}/post/${id}`;

  return {
    title,
    description,
    keywords,
    openGraph: {
      title,
      description,
      url,
      type: "article",
      images: [{ url: image }],
    },
    twitter: {
      card: "summary_large_image",
      title,
      description,
      images: [image],
    },
    authors: [{ name: data.author }],
    metadataBase: new URL(`${protocol}://${host}`),
  };
}
