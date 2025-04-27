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

    return {
      selectedKeys,
      openKeys,
      setSelectedKeys,
      setOpenKeys,
      collapse,
      setCollapse,
    };
  },
  {
    persist: true,
  }
);
