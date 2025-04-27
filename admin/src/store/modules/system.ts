export const useSystemStore = defineStore(
  "system",
  () => {
    const darkMode = ref(false);
    const spinning = ref(false);
    const loading = ref(false);

    // 设置暗黑模式
    const setDarkMode = (value: boolean) => {
      darkMode.value = value;
    };

    // 设置页面加载状态
    const setSpinning = (value: boolean) => {
      spinning.value = value;
    };

    // 设置骨架屏加载状态
    const setSkeleton = (value: boolean) => {
      loading.value = value;
    };

    return {
      darkMode,
      setDarkMode,
      spinning,
      setSpinning,
      loading,
      setSkeleton,
    };
  },
  {
    persist: [
      {
        pick: ["darkMode"],
        storage: localStorage,
      },
    ],
  }
);
