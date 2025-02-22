import useThemeLoaded from "@/hooks/theme";
import { MoonIcon } from "@/components/icons/moon";
import { SunLinearIcon } from "@/components/icons/sun";
import { useTheme } from "next-themes";

export default function ThemeSwitcher() {
    const isDark = useThemeLoaded();
    const { setTheme } = useTheme();
    const toggleTheme = () => setTheme(isDark ? "light" : "dark");

    return isDark === undefined ? (
        <MoonIcon onClick={toggleTheme} />
    ) : isDark ? (
        <SunLinearIcon onClick={toggleTheme} />
    ) : (
        <MoonIcon onClick={toggleTheme} />
    );
}
