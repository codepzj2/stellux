import StelluxMarkdown from "./components/markdown/md";

import React from "react";

export default async function PostPage() {
  const content = await fetch(
    "https://gitee.com/go-admin/go-admin/raw/master/README.md"
  ).then((res) => res.text());

  return (
    <div className="w-[800px] m-auto">
      <StelluxMarkdown content={content} />
    </div>
  );
}
