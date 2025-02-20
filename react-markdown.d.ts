declare module 'react-markdown' {
  import { ComponentType } from 'react';

  export interface MarkdownProps {
    className?: string;
    children: React.ReactNode;
    remarkPlugins?: any[];
  }

  const Markdown: ComponentType<MarkdownProps>;
  export default Markdown;
} 