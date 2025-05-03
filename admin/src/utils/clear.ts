import {
  useUserStore,
  useSystemStore,
  useSidebarStore,
  useTabStore,
} from "@/store";

export const clearStore = () => {
  useSystemStore().ResetSystemStore();
  useSidebarStore().ResetSidebarStore();
  useTabStore().ResetTabStore();
  useUserStore().ResetUserStore();
};
