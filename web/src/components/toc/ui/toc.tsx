// @ts-nocheck
"use client";

import * as React from "react";
import { TableOfContents } from "@/lib/toc";
import { cn } from "@/lib/utils";
import { useMounted } from "@/hooks/use-mounted";

interface TocProps {
  toc: TableOfContents;
}

export default function TocSimple({ toc }: TocProps) {
  const itemIds = React.useMemo(
    () =>
      toc
        ? toc
            .flatMap(item => [
              item.url?.replace(/^#/, ""),
              ...(item.items?.map(subItem => subItem.url?.replace(/^#/, "")) ||
                []),
            ])
            .filter(Boolean)
        : [],
    [toc]
  );

  const activeHeading = useActiveItem(itemIds);

  if (!toc?.length) {
    return null;
  }

  return (
    <nav className="space-y-2 rounded-lg border p-4">
      <p className="text-md font-bold">目录</p>
      <Tree tree={toc} activeItem={activeHeading} />
    </nav>
  );
}

function useActiveItem(itemIds: string[]) {
  const [activeId, setActiveId] = React.useState(null);

  React.useEffect(() => {
    const observer = new IntersectionObserver(
      entries => {
        entries.forEach(entry => {
          if (entry.isIntersecting) {
            setActiveId(entry.target.id);
          }
        });
      },
      { rootMargin: `0% 0% -80% 0%` }
    );

    itemIds?.forEach(id => {
      const element = document.getElementById(id);
      if (element) {
        observer.observe(element);
      }
    });

    return () => {
      itemIds?.forEach(id => {
        const element = document.getElementById(id);
        if (element) {
          observer.unobserve(element);
        }
      });
    };
  }, [itemIds]);

  return activeId;
}

interface TreeProps {
  tree: TableOfContents;
  level?: number;
  activeItem?: string;
}

function Tree({ tree, level = 1, activeItem }: TreeProps) {
  function hasActiveChild(
    items: TableOfContents[],
    activeItem: string
  ): boolean {
    return items?.some(
      item =>
        item.url?.replace(/^#/, "") === activeItem ||
        (item.items && hasActiveChild(item.items, activeItem))
    );
  }

  return tree?.length && level < 3 ? (
    <ul
      className={cn("m-0 list-none", {
        "pl-3 border-l border-muted": level === 1,
      })}
    >
      {tree.map((item, index) => {
        const itemId = item.url?.replace(/^#/, "");
        const isActive = itemId === activeItem;
        const isActiveParent = hasActiveChild(item.items || [], activeItem);

        return (
          <li key={index} className="relative group text-sm">
            <a
              href={item.url}
              className={cn(
                "block transition-colors duration-200 px-2 py-1 relative",
                "hover:text-primary",
                isActive
                  ? "font-semibold text-primary before:absolute before:left-[-8px] before:top-1/2 before:-translate-y-1/2 before:w-1 before:h-4 before:bg-primary before:rounded-full"
                  : "text-muted-foreground"
              )}
            >
              {item.title}
            </a>
            {item.items?.length ? (
              <ul
                className={cn("pl-3 transition-all duration-300 ease-in-out text-xs", {
                  block: isActiveParent || isActive,
                  "hidden group-hover:block": !isActiveParent && !isActive,
                })}
              >
                <Tree
                  tree={item.items}
                  level={level + 1}
                  activeItem={activeItem}
                />
              </ul>
            ) : null}
          </li>
        );
      })}
    </ul>
  ) : null;
}