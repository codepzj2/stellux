"use client";

import Toc from "./ui/toc";
import { TableOfContents } from "@/lib/toc";

export default ({
  toc,
  className,
}: {
  toc: TableOfContents;
  className?: string;
}) => {
  return <Toc toc={toc} className={className} />;
};
