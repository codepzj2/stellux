import { createApp } from "vue";
import App from "./App.vue";
import router from "@/router";
import pinia from "@/store";
import "virtual:svg-icons-register";
import VueLazyLoad from "vue3-lazyload";
import { message } from "ant-design-vue";

const app = createApp(App);
app.use(router).use(pinia).use(VueLazyLoad, {});

message.config({
  duration: 2,
  maxCount: 1,
  rtl: true,
});

app.mount("#app");
