import ReactMarkdown from "react-markdown";
import rehypeHighlight from "rehype-highlight";
import { Terminal } from "lucide-react";

import "highlight.js/styles/atom-one-dark.css";
import "@/style/markdown.scss"

import remarkGfm from 'remark-gfm'
import CopyButton from "@/components/markdown/copy";
const Markdown = ({ content }: { content: string }) => {
  return (
    <ReactMarkdown
      remarkPlugins={[remarkGfm]}
      rehypePlugins={[rehypeHighlight]}
      components={{
        pre: ({ children }) => <pre className="not-prose">{children}</pre>,
        code: ({ node, className, children, ...props }) => {
          const match = /language-(\w+)/.exec(className || "");
          if (match?.length) {
            const id = Math.random().toString(36).substr(2, 9);
            return (
              <div>
                <div className="flex h-12 items-center justify-between bg-zinc-100 px-4">
                  <div className="flex items-center gap-2">
                    <Terminal size={18} />
                    <p className="text-sm text-zinc-600 dark:text-zinc-400">
                      {node?.data?.meta}
                    </p>
                  </div>
                  <CopyButton id={id} />
                </div>
                <div className="overflow-x-auto">
                  <div id={id}>
                    {children}
                  </div>
                </div>
              </div>
            );
          } else {
            return (
              <code
                {...props}
                className="not-prose rounded bg-gray-100 px-1"
              >
                {children}
              </code>
            );
          }
        },
      }}
    >
      {content}
    </ReactMarkdown>
  );
};

export default Markdown;
