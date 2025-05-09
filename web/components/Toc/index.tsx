// @ts-nocheck
"use client";

import * as React from "react";
import { useEffect, useState, useMemo } from "react";
import type { TableOfContents } from "@/lib/toc";
import { cn } from "@/lib/utils";

interface TocProps {
  toc: TableOfContents;
  className?: string;
}

export function Toc({ toc, className }: TocProps) {
  const itemIds = useMemo(
    () =>
      toc
        ? toc
          .flatMap((item) => [
            item.url?.replace(/^#/, ""),
            ...(item.items?.map((subItem) => subItem.url?.replace(/^#/, "")) || []),
          ])
          .filter(Boolean)
        : [],
    [toc]
  );

  const activeId = useActiveItem(itemIds);

  if (!toc?.length) {
    return null;
  }

  return (
    <nav className={cn("text-sm", className)}>
      <Tree tree={toc} activeItem={activeId} />
    </nav>
  );
}

function useActiveItem(itemIds: string[]) {
  const [activeId, setActiveId] = useState<string | null>(null);

  useEffect(() => {
    if (!itemIds.length) return;

    const handleScroll = () => {
      let current: string | null = null;

      for (const id of itemIds) {
        const el = document.getElementById(id);
        if (el) {
          const rect = el.getBoundingClientRect();
          if (rect.top <= 100) {
            current = id;
          }
        }
      }

      setActiveId(current);
    };

    window.addEventListener("scroll", handleScroll, { passive: true });
    handleScroll(); // 初始调用一次

    return () => {
      window.removeEventListener("scroll", handleScroll);
    };
  }, [itemIds]);

  return activeId;
}

interface TreeProps {
  tree: TableOfContents;
  level?: number;
  activeItem?: string | null;
}

function Tree({ tree, level = 1, activeItem }: TreeProps) {
  return tree?.length && level < 3 ? (
    <ul
      className={cn("m-0 list-none", {
        "pl-3 border-l border-zinc-300 dark:border-zinc-700": level === 1,
      })}
    >
      {tree.map((item, index) => {
        const itemId = item.url?.replace(/^#/, "");
        const isActive = itemId === activeItem;

        return (
          <li key={index} className="relative">
            <a
              href={item.url}
              className={cn(
                "block transition-colors duration-200 px-2 py-1",
                "hover:text-primary",
                isActive
                  ? "font-semibold text-primary"
                  : "text-muted-foreground"
              )}
            >
              {item.title}
            </a>
            {item.items?.length ? (
              <ul className="pl-3">
                <Tree tree={item.items} level={level + 1} activeItem={activeItem} />
              </ul>
            ) : null}
          </li>
        );
      })}
    </ul>
  ) : null;
}
