export type SiteConfig = typeof siteConfig;

export const siteConfig = {
  name: "stellux",
  description: "stellux博客,记录生活,记录工作,记录学习",
  navItems: [
    {
      label: "博客",
      href: "/",
    },
    {
      label: "文档",
      href: "/docs",
    },
    {
      label: "关于",
      href: "/about",
    },
  ],
  navMenuItems: [
    {
      label: "博客",
      href: "/",
    },
    {
      label: "文档",
      href: "/docs",
    },
    {
      label: "关于",
      href: "/about",
    },
  ],
  links: {
    github: "https://github.com/codepzj/stellux",
  },
};
