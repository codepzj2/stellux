"use client";

import { Button } from "@/components/ui/button";

import { useRouter } from "next/navigation";
import Image from "next/image";
import NotFoundSvg from "@/assets/svg/404.svg";

export default function NotFound() {
  const router = useRouter();

  return (
    <>
      <div className="absolute w-screen h-[calc(100vh-4em)] bg-white dark:bg-black">
        <div className="w-full h-[73vh] mt-20">
          <div className="w-full h-full flex justify-center items-center flex-col md:flex-row">
            <Image
              src={NotFoundSvg}
              alt="404"
              className="w-full xl:w-[32rem] lg:w-[30rem] md:w-[24rem]"
            />

            <div className="xl:w-[32rem] lg:w-[30rem] md:w-[24rem] sm:text-start mx-4 text-center">
              <h1 className="text-5xl sm:text-8xl font-bold">404</h1>
              <h2 className="text-3xl sm:text-3xl font-bold my-4">
                页面未找到
              </h2>
              <p>您要查找的页面不存在或已被删除。</p>
              <Button className="mt-6" onClick={() => router.push("/")}>
                返回首页
              </Button>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
