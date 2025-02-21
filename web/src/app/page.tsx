"use client";

import { DatePicker } from "@heroui/react";
import { themeStore } from "@/store/theme";

export default  function App() {
  const { toggleTheme } = themeStore();
   
  return (
    <div>
      <div className="text-red-500">hello</div>
      <button onClick={toggleTheme}>按钮</button>
      <DatePicker className="max-w-[284px]" label="Birth date" />
    </div>
  );
}
