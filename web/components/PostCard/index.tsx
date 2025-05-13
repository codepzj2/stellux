"use client";

import { IPostCard } from "@/types/post";
import { Card, CardBody, CardHeader, CardFooter } from "@heroui/card";
import { useRouter } from "next/navigation";
import { cn } from "@/lib/utils";
import { Bookmark, Tag, Calendar } from "lucide-react";
import dayjs from "dayjs";
import { Image } from "@heroui/image";
import { TopIcon } from "@/components/SvgIcon";
export function PostCard({
  post,
  className,
}: {
  post: IPostCard;
  className?: string;
}) {
  const router = useRouter();
  const tags = post.tags?.join(", ");

  return (
    <Card
      key={post.id}
      className={cn("p-6 shadow transition-all hover:shadow-md", className)}
      onPress={() => router.push(`/post/${post.id}`)}
      isHoverable
      isPressable
    >
      <div className="flex flex-wrap md:flex-nowrap">
        {/* Left side: Title, Description, Footer */}
        <div className="flex-1">
          <CardHeader className="p-0 mb-4">
            <span
              className="text-lg font-bold tracking-tight truncate px-2"
              title={post.title}
            >
              {post.title}
            </span>
            {post.is_top && (
              <TopIcon className="text-red-500" size={24} />
            )}
          </CardHeader>

          <CardBody className="p-0 mb-4">
            <p
              className="text-sm text-muted-foreground leading-relaxed line-clamp-2 px-2"
              title={post.description}
            >
              {post.description}
            </p>
          </CardBody>

          <CardFooter className="hidden md:flex flex-wrap gap-2 p-0 text-sm text-muted-foreground">
            {post.category && (
              <span
                className="flex items-center gap-1.5 px-2 py-0.5 rounded-full bg-muted/60 hover:bg-muted transition"
                title={`分类：${post.category}`}
              >
                <Bookmark size={14} />
                {post.category}
              </span>
            )}
            {tags && (
              <span
                className="flex items-center gap-1.5 px-2 py-0.5 rounded-full bg-muted/60 hover:bg-muted transition"
                title={`标签：${tags}`}
              >
                <Tag size={14} />
                {tags}
              </span>
            )}
            <span
              className="flex items-center gap-1.5 px-2 py-0.5 rounded-full bg-muted/60 hover:bg-muted transition"
              title={post.created_at}
            >
              <Calendar size={14} />
              {dayjs(post.created_at).format("YYYY-MM-DD")}
            </span>
          </CardFooter>
        </div>

        {/* Right side: Image */}
        <div className="hidden md:flex justify-center items-center md:w-48 h-auto">
          {post.thumbnail && (
            <Image
              src={post.thumbnail}
              alt={post.title}
              width={180}
              height={120}
              shadow="none"
              loading="lazy"
              isZoomed
            />
          )}
        </div>
      </div>
    </Card>
  );
}
