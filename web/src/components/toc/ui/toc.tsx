// @ts-nocheck
"use client";

import * as React from "react";
import { TableOfContents } from "@/lib/toc";
import { cn } from "@/lib/utils";

interface TocProps {
  toc: TableOfContents;
  className?: string;
}

export default function Toc({ toc, className }: TocProps) {
  // 提取所有 URL（去除 @ 前缀），并过滤空值
  const itemIds = React.useMemo(
    () =>
      toc
        ?.flatMap(item => [
          item.url?.replace(/^@/, ""),
          ...(item.items?.map(subItem => subItem.url?.replace(/^@/, "")) || []),
        ])
        .filter(Boolean) ?? [],
    [toc]
  );

  const activeHeading = useActiveItem(itemIds);

  // 如果没有目录数据则不渲染
  if (!toc?.length) return null;

  return (
    <nav className={cn("space-y-2 p-2", className)}>
      <Tree tree={toc} activeItem={activeHeading} level={2} />
    </nav>
  );
}

// 用于检测当前页面活跃的目录项
function useActiveItem(itemIds: string[]) {
  const [activeId, setActiveId] = React.useState<string | null>(null);

  React.useEffect(() => {
    const observer = new IntersectionObserver(
      entries => {
        entries.forEach(entry => {
          if (entry.isIntersecting) {
            setActiveId(entry.target.id);
          }
        });
      },
      { rootMargin: "0% 0% -80% 0%" }
    );

    itemIds.forEach(id => {
      const el = document.getElementById(id);
      if (el) observer.observe(el);
    });

    return () => {
      itemIds.forEach(id => {
        const el = document.getElementById(id);
        if (el) observer.unobserve(el);
      });
    };
  }, [itemIds]);

  return activeId;
}

interface TreeProps {
  tree: TableOfContents;
  level: number; // 当前级别
  activeItem?: string;
}

function Tree({ tree, level, activeItem }: TreeProps) {
  return tree?.length && level <= 3 ? (
    <ul
      className={cn("m-0 list-none", {
        "pl-3 border-l border-muted": level === 2,
      })}
    >
      {tree.map((item, idx) => {
        const itemId = item.url?.replace(/^@/, "");
        const isActive = itemId === activeItem;

        return (
          // 只渲染二级和三级目录项
          level === 2 || level === 3 ? (
            <li key={idx} className="relative text-sm">
              <a
                href={item.url}
                className={cn(
                  "block px-2 py-1 transition-colors duration-200 relative",
                  "hover:text-primary",
                  isActive
                    ? "font-semibold text-primary before:absolute before:left-[-8px] before:top-1/2 before:-translate-y-1/2 before:w-1 before:h-4 before:bg-primary before:rounded-full"
                    : "text-muted-foreground"
                )}
              >
                {item.title}
              </a>
              {/* 递归渲染子目录项 */}
              {item.items?.length && level < 3 ? (
                <ul className="pl-3 transition-all duration-300 ease-in-out text-xs block">
                  <Tree
                    tree={item.items}
                    level={level + 1}
                    activeItem={activeItem}
                  />
                </ul>
              ) : null}
            </li>
          ) : null
        );
      })}
    </ul>
  ) : null;
}
