"use client";
import "@/style/main.scss";
import { HeroUIProvider } from "@heroui/react";
import { themeStore } from "@/store/theme";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const { theme } = themeStore();
  return (
    <html lang="zh-CN" className={theme}>
      <body>
        <HeroUIProvider>{children}</HeroUIProvider>
      </body>
    </html>
  );
}
