import type { UserInfoVO } from "@/types/user";

export const useUserStore = defineStore(
  "user",
  () => {
    const access_token = ref<string | null>(null);
    const refresh_token = ref<string | null>(null);
    const userInfo = ref<UserInfoVO | null>(null);

    const setAccessToken = (accessToken: string) => {
      access_token.value = accessToken;
    };

    const setRefreshToken = (refreshToken: string) => {
      refresh_token.value = refreshToken;
    };

    const setUserInfo = (newUserInfo: UserInfoVO) => {
      userInfo.value = newUserInfo;
    };

    const clearAccessToken = () => {
      access_token.value = null;
    };

    const clearRefreshToken = () => {
      refresh_token.value = null;
    };

    const ResetUserStore = () => {
      userInfo.value = null;
      access_token.value = null;
      refresh_token.value = null;
    };

    return {
      access_token,
      refresh_token,
      userInfo,
      setAccessToken,
      setRefreshToken,
      clearAccessToken,
      clearRefreshToken,
      setUserInfo,
      ResetUserStore,
    };
  },
  {
    persist: true,
  }
);
