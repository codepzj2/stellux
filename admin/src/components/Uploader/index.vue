<template>
  <div>
    <a-button type="primary" @click="showModal"> 上传附件 </a-button>
  </div>
  <a-modal
    v-model:open="open"
    title="附件上传"
    width="100%"
    wrap-class-name="full-modal"
    :showUploadList="false"
    :footer="null"
    :destroyOnClose="true"
    @drop="handleDrop"
    @cancel="handleCancel"
  >
    <div class="flex flex-col items-center">
      <a-tabs v-model:activeKey="activeKey">
        <a-tab-pane key="1" tab="文件上传"></a-tab-pane>
        <a-tab-pane key="2" tab="文件夹上传"></a-tab-pane>
      </a-tabs>
      <a-upload-dragger
        :file-list="fileList"
        :multiple="true"
        :beforeUpload="handleBeforeUpload"
        list-type="picture"
        :directory="activeKey === '2'"
        class="w-3/5 h-[280px] inline-block"
      >
        <div class="flex flex-col items-center justify-center h-full">
          <p class="ant-upload-drag-icon">
            <inbox-outlined></inbox-outlined>
          </p>
          <p class="ant-upload-text hidden md:block">
            点击或拖拽{{ activeKey === "1" ? "文件" : "文件夹" }}到此区域上传
          </p>
          <p class="ant-upload-hint hidden md:block">支持单个或批量上传</p>
        </div>
      </a-upload-dragger>
    </div>
  </a-modal>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { uploadFilesAPI } from "@/api/file";
import { message } from "ant-design-vue";
import { InboxOutlined } from "@ant-design/icons-vue";
import type { UploadProps } from "ant-design-vue/lib/upload/interface";

const open = ref<boolean>(false);
const fileList = ref<UploadProps["fileList"]>([]);
const activeKey = ref("1");

const handleBeforeUpload: UploadProps["beforeUpload"] = async (
  file,
  FileList
) => {
  fileList.value = [...(fileList.value || []), file];
  if (fileList.value.length === FileList.length) {
    message.loading({
      content: "图片上传中...",
      duration: 0,
    });
    try {
      const res = await uploadFilesAPI({
        files: fileList.value,
      });
      message.success({
        content: res.msg,
      });
      fileList.value = [];
    } catch (error) {
      fileList.value = [];
    }
  }
  return false;
};

const showModal = () => {
  open.value = true;
  fileList.value = [];
};

const emit = defineEmits<{
  (e: "close"): void;
}>();

const handleCancel = () => {
  open.value = false;
  emit("close");
  fileList.value = [];
};

const handleDrop = (e: any) => {
  console.log(e);
  fileList.value = [...(fileList.value || []), e.file];
};
</script>

<style lang="scss">
.full-modal {
  .ant-modal {
    max-width: 100%;
    top: 0;
    padding-bottom: 0;
    margin: 0;
  }
  .ant-modal-content {
    display: flex;
    flex-direction: column;
    height: calc(100vh);
  }
  .ant-modal-body {
    flex: 1;
    overflow-y: auto;
  }
}

::-webkit-scrollbar {
  width: 8px;
}
</style>
