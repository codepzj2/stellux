"use client";
import { useRouter } from "next/navigation";
import { useTheme } from "next-themes";
import Image from "next/image";
import { Skeleton } from "@/components/ui/skeleton";
import LightLogo from "@/assets/logo/logo-light.png";
import DarkLogo from "@/assets/logo/logo-dark.png";
import { useMounted } from "@/hooks/use-mounted";

export default function Logo() {
  const router = useRouter();
  const { resolvedTheme } = useTheme();
  const mounted = useMounted();

  return mounted ? (
    <Image
      src={resolvedTheme === "dark" ? DarkLogo : LightLogo}
      alt="logo"
      width={128}
      className="mr-4 flex items-center cursor-pointer lg:mr-6 hover:scale-95 transition-all duration-300"
      onClick={() => router.push("/")}
    />
  ) : (
    <div className="flex items-center space-x-4">
      <Skeleton className="h-12 w-12 rounded-full" />
      <div className="space-y-2">
        <Skeleton className="h-2 w-[84px]" />
        <Skeleton className="h-2 w-[48px]" />
      </div>
    </div>
  );
}
