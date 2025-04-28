import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";
import rehypeRaw from "rehype-raw";
import rehypeHighlight from "rehype-highlight";
import "./md.scss";
import CopyButton from "./copy";
import React from "react";

export default function Md({ content }: { content: string }) {
  let index = 1;

  return (
    <div className="markdown-body">
      <ReactMarkdown
        rehypePlugins={[rehypeRaw, rehypeHighlight]}
        remarkPlugins={[remarkGfm]}
        components={{
          h1: ({ children }) => (
            <h1 className="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl">
              {children}
            </h1>
          ),
          h2: ({ children }) => (
            <h2
              id={`header-${index++}`}
              className="mt-10 scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0"
            >
              {children}
            </h2>
          ),
          h3: ({ children }) => (
            <h3
              id={`header-${index++}`}
              className="mt-8 scroll-m-20 text-2xl font-semibold tracking-tight"
            >
              {children}
            </h3>
          ),
          h4: ({ children }) => (
            <h4 className="scroll-m-20 text-xl font-semibold tracking-tight">
              {children}
            </h4>
          ),
          p: ({ children }) => (
            <p className="leading-7 [&:not(:first-child)]:mt-6">{children}</p>
          ),
          ul: ({ children }) => (
            <ul className="my-6 ml-6 list-disc [&>li]:mt-2">{children}</ul>
          ),
          blockquote: ({ children }) => (
            <div className="mt-6 border-l-2 pl-6 italic">{children}</div>
          ),
          a: ({ children, ...props }) => (
            <a
              className="font-medium text-primary underline underline-offset-4"
              {...props}
            >
              {children}
            </a>
          ),
          table: ({ children }) => <table className="w-full">{children}</table>,
          th: ({ children }) => (
            <th className="border px-4 py-2 text-left font-bold [&[align=center]]:text-center [&[align=right]]:text-right">
              {children}
            </th>
          ),
          tr: ({ children }) => (
            <tr className="m-0 border-t p-0 even:bg-muted">{children}</tr>
          ),
          td: ({ children }) => (
            <td className="border px-4 py-2 text-left [&[align=center]]:text-center [&[align=right]]:text-right">
              {children}
            </td>
          ),
          pre: ({ children }) => (
            <pre className="rounded-md bg-muted !p-0">{children}</pre>
          ),
          code: ({ className, children }) => {
            const match = /language-(\w+)/.exec(className || "");
            const count =
              React.Children.toArray(children).length === 1
                ? (
                    React.Children.toArray(children)[0]
                      .toString()
                      .match(/\n/g) || []
                  ).length
                : 0;

            if (match?.length || count > 0) {
              const id = Math.random().toString(36).slice(2, 11);
              // 多行代码块
              return (
                <div className="not-prose relative rounded-md">
                  <CopyButton id={id} />
                  <div className="overflow-x-auto p-4" id={id}>
                    {children}
                  </div>
                </div>
              );
            }
            // 单行代码块
            return (
              <code className="relative rounded bg-muted px-[0.3rem] py-[0.2rem] font-mono text-sm font-semibold">
                {children}
              </code>
            );
          },
        }}
      >
        {content}
      </ReactMarkdown>
    </div>
  );
}
