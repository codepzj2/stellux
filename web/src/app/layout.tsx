import ProgressProvider from "@/components/progress-provider";
import ThemeProvider from "@/components/theme-provider";
import NavBar from "@/components/navbar";

// 引入全局样式
import "@/styles/global.scss";

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
          // 系统跟随
          enableSystem
          disableTransitionOnChange
        >
          <ProgressProvider>
            <div className="h-screen flex flex-col">
              <div className="fixed top-0 left-0 right-0 z-50">
                <NavBar />
            </div>
            <div className="mt-14">{children}</div>
            </div>
          </ProgressProvider>
        </ThemeProvider>
      </body>
    </html>
  );
}
