"use client";

import { FC } from "react";
import { useTheme } from "next-themes";
import { useIsSSR } from "@react-aria/ssr";
import { Skeleton } from "@heroui/skeleton";

export const Logo: FC = () => {
    const { theme } = useTheme();
    const isSSR = useIsSSR();
    const logoSrc = theme === "light" ? "/logo-light.png" : "/logo-dark.png";
    return isSSR ? <div className="w-36 h-10 flex flex-col gap-2">
        <Skeleton className="h-3 w-3/5 rounded-lg" />
        <Skeleton className="h-3 w-4/5 rounded-lg" />
    </div> : <img className="w-36 h-10 hover:scale-95 transition-all duration-300" src={logoSrc} alt="logo" />;
};




