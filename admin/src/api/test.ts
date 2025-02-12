import request from "@/utils/axios";

export const testApi = async () => {
  const res = await request({
    url: "/test",
    method: "get",
  });
  console.log(res);
};
export const testApi2 = async () => {
  const res = await request({
    url: "/test2",
    method: "get",
  });
  console.log(res);
};
