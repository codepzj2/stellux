'use client'
import Markdown from 'react-markdown'
import remarkGfm from 'remark-gfm'
import rehypeHighlight from "rehype-highlight";
import { useState } from 'react';

import "highlight.js/styles/atom-one-light.css";
import "@/style/markdown.scss"
import { Terminal } from 'lucide-react'
import CopyButton from '@/components/markdown/copy'

export default function PostPage() {
  try {
    const [content, setContent] = useState("");

  const handleInput = (e: React.FormEvent<HTMLTextAreaElement>) => {
    setContent(e.currentTarget.value);
  };

    return (
      <div style={{ width: '800px', margin: '0 auto' }}>
        <textarea name="" id="" style={{ width: '100%', height: '100px' }} value={content} onChange={handleInput}></textarea>
        <div className="markdown-body">
          <Markdown rehypePlugins={[rehypeHighlight]} remarkPlugins={[remarkGfm]} components={{
        pre: ({ children }) => <pre style={{ padding: '0' }}>{children}</pre>,
        code: ({ node, className, children, ...props }) => {
          const match = /language-(\w+)/.exec(className || "");
          if (match?.length) {
            const id = Math.random().toString(36).substr(2, 9);
            return (
              <div className="not-prose rounded-md border">
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
                  <div id={id} className="p-4">
                    {children}
                  </div>
                </div>
              </div>
            );
          } else {
            return (
              <code
                {...props}
                className="not-prose rounded bg-gray-100 px-1 dark:bg-zinc-900"
              >
                {children}
              </code>
            );
          }
        },
      }}
      >
        {content}
      </Markdown>
        </div>
      </div>
      
    );
  } catch (error) {
    console.error('Fetch error:', error);
    return <div>Error loading post</div>;
  }
}
