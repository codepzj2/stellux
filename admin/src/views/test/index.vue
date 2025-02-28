


<template>
  <div class="grid grid-cols-4 gap-4">
    <div v-for="item in photos" :key="item.id">
      <img :src="item.url" alt="photo" />
    </div>
  </div>
</template>
<script lang="ts" setup>
import { ref } from "vue";

const photos = ref<Photo[]>([]);

const getPhotos = async () => {
  const res = await getPhotos();
  photos.value = res.data;
};

getPhotos();

</script>
