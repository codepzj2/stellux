import { useWindowSize } from "@vueuse/core";

export const useMobile = () => {
  const { width } = useWindowSize();
  return width.value < 768;
};
