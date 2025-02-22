import { useTheme } from "next-themes";
import { useEffect, useState } from "react";

const useThemeLoaded: () => boolean | undefined = () => {
  const { theme } = useTheme();
  const [isDark, setDark] = useState<boolean | undefined>(undefined);
  useEffect(() => {
    setDark(theme === "dark");
  }, [theme]);
  return isDark;
};

export default useThemeLoaded;
