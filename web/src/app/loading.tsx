"use client";
export default function Loading() {
  return (
    <div className="flex justify-center items-center w-full h-[calc(100vh-4em)]">
      <div className="w-24 h-24 animate-spin rounded-full border-y-2 border-black dark:border-zinc-100"></div>
    </div>
  );
}
