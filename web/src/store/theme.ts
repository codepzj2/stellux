import { create } from "zustand";
import { persist } from "zustand/middleware";

type ThemeStore = {
  theme: "light" | "dark";
  toggleTheme: () => void;
};

export const themeStore = create<ThemeStore>()(
  persist(
    (set, get) => ({
      theme: "light",
      toggleTheme() {
        set((state: ThemeStore) => ({
          theme: state.theme === "light" ? "dark" : "light",
        }));
        document.documentElement.classList.toggle(
          "dark",
          get().theme === "dark"
        );
      },
    }),
    {
      name: "stellux-theme",
    }
  )
);
