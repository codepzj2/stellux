import { theme } from "ant-design-vue";

export const useSystemStore = defineStore(
  "system",
  () => {
    const themeColor = ref("#1677FF");
    const themeMode = ref("light");

    const appTheme = computed(() => {
      return {
        algorithm: themeMode.value === "dark" ? theme.darkAlgorithm : theme.defaultAlgorithm,
        token: {
          colorPrimary: themeColor.value,
          fontSize: 15,
          fontFamily: "var(--font-sans)",
        },
      };
    });

    // 设置暗黑模式
    const setThemeMode = (value: string) => {
      themeMode.value = value;
    };

    const setThemeColor = (value: string) => {
      themeColor.value = value;
    };

    const ResetSystemStore = () => {
      themeMode.value = "light";
      themeColor.value = "#1677FF";
    };

    return {
      themeMode,
      setThemeMode,
      themeColor,
      setThemeColor,
      appTheme,
      ResetSystemStore,
    };
  },
  {
    persist: true,
  }
);
