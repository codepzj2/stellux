"use client";

import TocSimple from "@/components/toc/ui/toc";
import { TableOfContents } from "@/lib/toc";

export default function Toc({ toc }: { toc: TableOfContents }) {
  return <TocSimple toc={toc} />;
}
