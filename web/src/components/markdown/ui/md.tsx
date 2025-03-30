import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";
import rehypeRaw from "rehype-raw";
import rehypeHighlight from "rehype-highlight";
import "@/styles/md.scss";
import { Link as LinkIcon, Terminal } from "lucide-react";
import CopyButton from "@/components/markdown/ui/copy";
import React, { JSX, ReactNode } from "react";

// 定义CustomHeading的props接口
interface CustomHeadingProps {
  level: 2 | 3;
  children: ReactNode;
  index: number;
}

const CustomHeading: React.FC<CustomHeadingProps> = ({ level, children, index }) => {
  const Tag = `h${level}` as keyof JSX.IntrinsicElements;
  
  const iconSize = 24 - (level - 2) * 2;

  return (
    <Tag
      id={`header-${index}`}
      className="flex items-center gap-2"
      style={{ scrollMarginTop: "4em" }}
    >
      <LinkIcon className="text-zinc-500" size={iconSize} />
      {children}
    </Tag>
  );
};

// 定义MdSimple的props接口
interface MdSimpleProps {
  content: string;
}

export default function Markdown({ content }: MdSimpleProps) {
  let index = 1;

  return (
    <div className="markdown-body">
      <ReactMarkdown
        rehypePlugins={[rehypeRaw, rehypeHighlight]}
        remarkPlugins={[remarkGfm]}
        components={{
          h2: ({ children }) => <CustomHeading level={2} children={children} index={index++} />,
          h3: ({ children }) => <CustomHeading level={3} children={children} index={index++} />,
          h4: ({ children }) => <h4 id={`header-${index++}`}>{children}</h4>,
          h5: ({ children }) => <h5 id={`header-${index++}`}>{children}</h5>,
          h6: ({ children }) => <h6 id={`header-${index++}`}>{children}</h6>,
          ul: ({ children }) => <ul style={{ listStyleType: "disc", listStylePosition: "outside" }}>{children}</ul>,
          li: ({ children }) => <li style={{ listStyleType: "disc", listStylePosition: "outside" }}>{children}</li>,
          a: ({ children, ...props }) => <a style={{ textDecoration: "none" }} {...props}>{children}</a>,
        pre: ({ children }) => <pre style={{ padding: "0" }}>{children}</pre>,
        blockquote: ({ children }) => <div className="my-4 p-3 rounded-md border bg-sky-100 border-sky-200 dark:bg-transparent dark:border-gray-600 blockquote">{children}</div>,
        code: ({ node, className, children, ...props }) => {
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
              return (
                <div className="not-prose rounded-md">
                  <div className="flex h-10 items-center justify-between px-4 bg-zinc-100 dark:bg-zinc-800">
                    <div className="flex items-center gap-2">
                      <Terminal size={16} />
                      <span className="h-full flex items-center">
                        {node?.data?.meta || match?.[1] || "plaintext"}
                      </span>
                    </div>
                    <CopyButton id={id} />
                  </div>
                  <div className="overflow-x-auto">
                    <div id={id} className="p-4">
                      {children}
                    </div>
                  </div>
                </div>
              );
            }
            return (
              <code {...props} className="not-prose rounded-md">
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