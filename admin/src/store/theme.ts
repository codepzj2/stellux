import { defineStore } from "pinia";
import { ref } from "vue";

export const useThemeStore = defineStore("theme", () => {
  const tailwindTheme = ref("light");

  const toggleTheme = () => {
    tailwindTheme.value = tailwindTheme.value === "light" ? "dark" : "light";
    document.documentElement.classList.toggle("dark", tailwindTheme.value === "dark");
  };

  const initTheme = () => {
    const savedTheme = JSON.parse(localStorage.getItem("theme") || '{"tailwindTheme":"light"}');
    tailwindTheme.value = savedTheme.tailwindTheme;
    localStorage.setItem("theme", JSON.stringify({ tailwindTheme: tailwindTheme.value }));
    document.documentElement.classList.toggle("dark", savedTheme.tailwindTheme === "dark");
  };

  return { tailwindTheme, toggleTheme, initTheme };
}, {
  persist: true,
});

export default useThemeStore;
