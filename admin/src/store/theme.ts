import { defineStore } from "pinia";
import { ref } from "vue";

export const useThemeStore = defineStore(
  "theme",
  () => {
    const tailwindTheme = ref("light");

    const toggleTheme = () => {
      tailwindTheme.value = tailwindTheme.value === "light" ? "dark" : "light";
      document.documentElement.classList.toggle(
        "dark",
        tailwindTheme.value === "dark"
      );
    };

    const initTheme = () => {
      document.documentElement.classList.toggle(
        "dark",
        tailwindTheme.value === "light"
      );
    };

    return { tailwindTheme, toggleTheme, initTheme };
  },
  {
    persist: true,
  }
);

export default useThemeStore;
