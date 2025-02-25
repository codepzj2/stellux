<template>
  <div class="flex justify-center">
    <a-upload-dragger
      v-model:file-list="fileList"
      :multiple="true"
      :beforeUpload="handleBeforeUpload"
      :directory="true"
      list-type="picture"
      class="w-3/5 h-[250px] inline-block"
    >
      <div class="flex flex-col items-center justify-center h-full">
        <p class="ant-upload-drag-icon">
          <inbox-outlined></inbox-outlined>
        </p>
        <p class="ant-upload-text">点击或拖拽文件到此区域上传</p>
        <p class="ant-upload-hint">支持单个或批量上传</p>
      </div>
    </a-upload-dragger>
  </div>
</template>

<script lang="ts" setup>
import { InboxOutlined } from "@ant-design/icons-vue";
import { ref } from "vue";
import type { UploadProps } from "ant-design-vue/lib/upload/interface";
import { uploadPicturesLocal } from "@/api/modules/upload";
import { message } from "ant-design-vue";


const fileList = ref<UploadProps['fileList']>([]);

const key = "upload-pictures-local";
const handleBeforeUpload: UploadProps['beforeUpload'] = (file,FileList) =>{
  fileList.value = [...(fileList.value || []), file];
  if (fileList.value.length === FileList.length) {
    try {
      message.loading({
        content: "图片上传中...",
        key,
      });
      uploadPicturesLocal({
        files: fileList.value,
      }).then((res) => {
        message.success({
          content: res.msg,
          key,
          duration: 1.5,
        });
        setTimeout(() => {
          fileList.value = [];
        }, 1500);
      });
    } catch (error) {
      console.error(error);
      message.error({
        content: "图片上传失败",
        key,
      });
    }
  }

  return false;
};

</script>
<style scoped>
.ant-upload-list-item {
  transition: opacity 0.5s ease-in-out;
}
</style>
