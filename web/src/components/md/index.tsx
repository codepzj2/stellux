import Md from "./ui/md";

export default function Markdown({ content }: { content: string }) {
  return <Md content={content} />;
}
