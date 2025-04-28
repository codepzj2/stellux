"use client";

import {
  Sheet,
  SheetContent,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet";
import { Button } from "@/components/ui/button";
import { ScrollArea } from "@/components/ui/scroll-area";
import Toc from "@/components/toc";
import { NotebookText } from "lucide-react";

export default function TocButton({
  toc,
  className,
}: {
  toc: any;
  className?: string;
}) {
  return (
    <Sheet>
      <SheetTrigger asChild>
        <Button variant="outline" size="icon" className={className}>
          <NotebookText />
        </Button>
      </SheetTrigger>
      <SheetContent
        className="px-4 pb-4"
        onCloseAutoFocus={event => {
          event.preventDefault();
        }}
      >
        <SheetHeader>
          <SheetTitle>文章目录</SheetTitle>
        </SheetHeader>
        {toc && (
          <ScrollArea className="max-h-[90vh]">
            <Toc toc={toc} className="max-h-[90vh]" />
          </ScrollArea>
        )}
      </SheetContent>
    </Sheet>
  );
}
