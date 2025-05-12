export const useSidebarStore = defineStore(
  "sidebar",
  () => {
    const collapse = ref(false);

    function setCollapse(value: boolean) {
      collapse.value = value;
    }

    const ResetSidebarStore = () => {
      collapse.value = false;
    };

    return {
      collapse,
      setCollapse,
      ResetSidebarStore,
    };
  },
  {
    persist: true,
  }
);
