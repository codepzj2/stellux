"use client";
import {
  Card,
  CardContent,
  CardDescription,
  CardTitle,
} from "@/components/ui/card";
import Image from "next/image";
import { Separator } from "@/components/ui/separator";
import { Badge } from "@/components/ui/badge";
import { FolderOpen, Tag } from "lucide-react";
import { useRouter } from "next/navigation";

// 假设这些工具函数和类型已定义
import { timeAgo } from "@/utils/time";
import { IPostCard } from "@/types/posts";
import NotFoundPng from "@/assets/png/not-found.png";

export default function PostsCard({ post }: { post: IPostCard }) {
  const router = useRouter();

  return (
    <Card
      className="w-full py-6 cursor-pointer transition-shadow hover:shadow-sm hover:bg-zinc-50 dark:bg-zinc-900 dark:hover:bg-zinc-800"
      onClick={() => {
        router.push(`/post/${post.id}`);
      }}
    >
      <CardContent className="flex flex-row-reverse items-center gap-4">
        <div className="w-[108px] h-[84px] relative hidden md:block">
          <Image
            src={post.thumbnail || NotFoundPng}
            alt={post.title}
            fill
            className="border-1 border-gray-200 rounded-sm object-cover"
          />
        </div>
        <div className="flex-1 flex flex-col justify-between">
          <CardTitle className="text-xl font-bold">{post.title}</CardTitle>
          <CardDescription className="line-clamp-1 my-1">
            {post.description || post.content}
          </CardDescription>
          <div className="w-full flex justify-between items-center text-xs text-gray-500">
            <div className="flex h-5 items-center space-x-2.5">
              <span>{post.author}</span>
              <Separator orientation="vertical" />
              <span>{timeAgo(post.created_at)}</span>
            </div>
            <div className="flex items-center space-x-2 py-1">
              <Badge variant="secondary">
                <FolderOpen className="w-4 h-4 mr-1" />
                {post.category}
              </Badge>
              {post.tags?.map(tag => (
                <Badge variant="secondary" key={tag}>
                  <Tag className="w-4 h-4 mr-1" />
                  {tag}
                </Badge>
              ))}
            </div>
          </div>
        </div>
      </CardContent>
    </Card>
  );
}
