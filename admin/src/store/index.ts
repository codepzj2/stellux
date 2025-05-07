import { useSidebarStore } from "./modules/sidebar";
import { useSystemStore } from "./modules/system";
import { useUserStore } from "./modules/user";
import { createPinia } from "pinia";
import piniaPluginPersistedstate from "pinia-plugin-persistedstate";

// 创建pinia实例
const pinia = createPinia();
pinia.use(piniaPluginPersistedstate);

// 创建store实例
export { useSidebarStore, useSystemStore, useUserStore };

export default pinia;
