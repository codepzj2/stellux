import { useUserStore, useSystemStore, useSidebarStore } from "@/store";

export const clearStore = () => {
  useSystemStore().ResetSystemStore();
  useSidebarStore().ResetSidebarStore();
  useUserStore().ResetUserStore();
};
