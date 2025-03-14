import { LoadingSpinner } from "@/components/ui/spin";

export default function Loading() {
  return (
    <div className="flex flex-col justify-center items-center gap-2 h-[calc(100vh-6em)]">
      <LoadingSpinner size={36} />
      加载中...
    </div>
  );
}
