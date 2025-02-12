// 对消息组件的封装
import { message } from "ant-design-vue";

message.config({
  top: "10px",
  duration: 2,
  maxCount: 1,
});

const success = (text: string) => {
  message.success({
    content: text,
    class: "message",
  });
};
const error = (text: string) => {
  message.error({
    content: text,
    class: "message",
  });
};
const warning = (text: string) => {
  message.warning({
    content: text,
    class: "message",
  });
};

export default {
  success,
  error,
  warning,
};
