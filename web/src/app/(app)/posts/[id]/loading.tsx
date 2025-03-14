import { Skeleton } from "@/components/ui/skeleton";
import { Separator } from "@/components/ui/separator";
import { Book, Clock } from "lucide-react";

export default function PostSkeleton() {
  return (
    <div className="max-w-[850px] m-auto p-6 my-4">
      {/* 文章标题骨架 */}
      <Skeleton className="h-10 w-3/4 my-4" />

      {/* 元信息骨架 */}
      <div className="flex justify-start items-center gap-3 text-sm text-gray-500">
        <Skeleton className="h-4 w-24" />
        <Skeleton className="h-4 w-20" />
        <span className="flex items-center gap-1">
          <Book size={14} />
          <Skeleton className="h-4 w-10" />
        </span>
        <span className="flex items-center gap-1">
          <Clock size={14} />
          <Skeleton className="h-4 w-10" />
        </span>
      </div>

      <Separator className="mt-4" />

      {/* 文章内容骨架 */}
      <div className="space-y-3 mt-6">
        <Skeleton className="h-5 w-full" />
        <Skeleton className="h-5 w-5/6" />
        <Skeleton className="h-5 w-4/6" />
        <Skeleton className="h-5 w-3/6" />
      </div>
    </div>
  );
}
