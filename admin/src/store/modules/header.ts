import type { JSX } from "vue/jsx-runtime";

interface HeaderAction {
  label: string;
  type?: "primary" | "default" | "dashed" | "text" | "link";
  size?: "small" | "middle" | "large";
  icon?: JSX.Element;
  onClick: () => void;
}

export const useHeaderStore = defineStore("header", () => {
  const leftHeaderActions = ref<HeaderAction[]>([]);
  const rightHeaderActions = ref<HeaderAction[]>([]);

  const setLeftHeaderActions = (actions: HeaderAction[]) => {
    leftHeaderActions.value = actions;
  };

  const setRightHeaderActions = (actions: HeaderAction[]) => {
    rightHeaderActions.value = actions;
  };

  const clearHeaderActions = () => {
    leftHeaderActions.value = [];
    rightHeaderActions.value = [];
  };

  return {
    leftHeaderActions,
    rightHeaderActions,
    setLeftHeaderActions,
    setRightHeaderActions,
    clearHeaderActions,
  };
});
