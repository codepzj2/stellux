import Link from "next/link";

export default function Nav() {
  return (
    <>
      <Link
        href="/"
        className="transition-colors hover:text-foreground/80 text-foreground/80"
      >
        🏠首页
      </Link>
      <Link
        href="/docs"
        className="transition-colors hover:text-foreground/80 text-foreground/80"
      >
        📝文档
      </Link>
      <Link
        href="/notes"
        className="transition-colors hover:text-foreground/80 text-foreground/80"
      >
        📒便签
      </Link>
      <Link
        href="/about"
        className="transition-colors hover:text-foreground/80 text-foreground/80"
      >
        👋关于
      </Link>
    </>
  );
}
