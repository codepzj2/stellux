import Md from "@/components/markdown/ui/md";

export default function Markdown({ content }: { content: string }) {
  return <Md content={content} />;
}
