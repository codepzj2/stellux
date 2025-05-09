import Md from "./ui/md";

export default function Markdown({ content, className }: { content: string, className?: string }) {
  return <Md content={content} className={className} />;
}
