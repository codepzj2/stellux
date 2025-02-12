const errorCode = ["401", "403", "404", "500"];
export const errorView = errorCode.map(code => {
  return {
    path: code,
    name: code,
    component: () => import(`@/views/auth/error/${code}.vue`),
  };
});
