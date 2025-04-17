import { defineStore } from "pinia";

import { ref } from "vue";

export const useUserStore = defineStore(
  "user",
  () => {
    const access_token = ref<string | null>(null);
    const refresh_token = ref<string | null>(null);

    const setAccessToken = (accessToken: string) => {
      access_token.value = accessToken;
    };

    const setRefreshToken = (refreshToken: string) => {
      refresh_token.value = refreshToken;
    };

    const clearAccessToken = () => {
      access_token.value = null;
    };

    const clearRefreshToken = () => {
      refresh_token.value = null;
    };

    return {
      access_token,
      refresh_token,
      setAccessToken,
      setRefreshToken,
      clearAccessToken,
      clearRefreshToken,
    };
  },
  {
    persist: true,
  }
);
