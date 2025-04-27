import { useSidebarStore } from "./modules/sidebar";
import { useTabStore } from "./modules/tab";
import { useSystemStore } from "./modules/system";
import { useUserStore } from "./modules/user";
import { useHeaderStore } from "./modules/header";
import { createPinia } from "pinia";
import piniaPluginPersistedstate from "pinia-plugin-persistedstate";

// 创建pinia实例
const pinia = createPinia();
pinia.use(piniaPluginPersistedstate);

// 创建store实例
export { useSidebarStore, useTabStore, useSystemStore, useUserStore, useHeaderStore };

export default pinia;
