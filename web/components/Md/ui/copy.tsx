"use client";
import { Copy, Check } from "lucide-react";
import { useState } from "react";

const CopyButton = ({ id }: { id: string }) => {
  const [copied, setCopited] = useState(false);

  const onCopy = async () => {
    try {
      setCopited(true);
      const text = document.getElementById(id)!.innerText;
      await navigator.clipboard.writeText(text);
      setTimeout(() => {
        setCopited(false);
      }, 1500);
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <button
      onClick={onCopy}
      className={`inline-flex rounded-md p-2 bg-zinc-100 hover:bg-zinc-200 dark:bg-zinc-900 dark:hover:bg-zinc-800 absolute top-2 right-2 z-10`}
    >
      <Copy
        size={16}
        className={`transition-all
        ${copied ? "scale-0" : "scale-100"}
      `}
      />
      <Check
        size={16}
        className={`absolute transition-all text-green-600 font-bold ${
          copied ? "scale-100" : "scale-0"
        }`}
      />
    </button>
  );
};

export default CopyButton;
