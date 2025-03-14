import { ThemeProvider } from "@/components/theme-provider";

import NavBar from "@/components/navbar";

// 引入全局样式
import "@/styles/global.scss";

import "@/styles/var.scss";

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html suppressHydrationWarning lang="zh-CN">
      <body>
        {/* 亮暗模式容器 */}
        <ThemeProvider
          attribute="class"
          defaultTheme="system"
          enableSystem
          disableTransitionOnChange
        >
          <div className="flex flex-col h-screen">
            <NavBar />
            <div className="flex-1">{children}</div>
          </div>
        </ThemeProvider>
      </body>
    </html>
  );
}
