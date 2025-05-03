export const useSidebarStore = defineStore(
  "sidebar",
  () => {
    const selectedKeys = ref<string[]>([]);
    const openKeys = ref<string[]>([]);
    const collapse = ref(false);

    function setSelectedKeys(keys: string[]) {
      selectedKeys.value = keys;
    }

    function setOpenKeys(keys: string[]) {
      openKeys.value = keys;
    }

    function setCollapse(value: boolean) {
      collapse.value = value;
    }

    const ResetSidebarStore = () => {
      selectedKeys.value = [];
      openKeys.value = [];
      collapse.value = false;
    };

    return {
      selectedKeys,
      openKeys,
      setSelectedKeys,
      setOpenKeys,
      collapse,
      setCollapse,
      ResetSidebarStore,
    };
  },
  {
    persist: true,
  }
);
