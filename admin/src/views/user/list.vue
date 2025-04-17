<template>
  <div class="px-4 sm:px-6 lg:px-8">
    <div class="sm:flex sm:items-center">
      <div class="sm:flex-auto">
        <h1 class="text-base font-semibold text-gray-900">Users</h1>
        <p class="mt-2 text-sm text-gray-700">
          A list of all the users in your account including their name, title,
          email and role.
        </p>
      </div>
      <div class="mt-4 sm:mt-0 sm:ml-16 sm:flex-none">
        <button
          type="button"
          class="block rounded-md bg-indigo-600 px-3 py-2 text-center text-sm font-semibold text-white shadow-xs hover:bg-indigo-500 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
        >
          Add user
        </button>
      </div>
    </div>
    <div class="mt-8 flow-root">
      <div class="-mx-4 -my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
        <div class="inline-block min-w-full py-2 align-middle sm:px-6 lg:px-8">
          <table class="min-w-full divide-y divide-gray-300">
            <tbody class="divide-y divide-gray-200 bg-white">
              <tr v-for="user in userList" :key="user.id">
                <td class="w-64 py-5 pr-3 pl-4 text-sm whitespace-nowrap sm:pl-0">
                  <div class="flex items-center">
                    <div class="size-11 shrink-0">
                      <img
                        class="size-11 rounded-full"
                        :src="user.avatar"
                        alt=""
                      />
                    </div>
                    <div class="ml-4">
                      <div class="font-medium text-gray-900">
                        {{ user.username }}
                      </div>
                      <div class="mt-1 text-gray-500">{{ user.email }}</div>
                    </div>
                  </div>
                </td>
                <td class="px-3 py-5 text-sm whitespace-nowrap text-gray-500">
                  <span
                    :class="[
                      'inline-flex items-center rounded-md px-2 py-1 text-xs font-medium ring-1 ring-inset',
                      user.role_id === 0
                        ? 'bg-blue-50 text-blue-700 ring-blue-600/20'
                        : '',
                      user.role_id === 1
                        ? 'bg-green-50 text-green-700 ring-green-600/20'
                        : '',
                      user.role_id === 2
                        ? 'bg-yellow-50 text-yellow-700 ring-yellow-600/20'
                        : '',
                    ]"
                  >
                    {{
                      user.role_id === 0
                        ? "管理员"
                        : user.role_id === 1
                          ? "普通用户"
                          : "游客"
                    }}
                  </span>
                </td>
                <td
                  class="relative py-5 pr-4 pl-3 text-right text-sm font-medium whitespace-nowrap space-x-2 sm:pr-0"
                >
                  <a href="#" class="text-indigo-600 hover:text-indigo-900"
                    >编辑<span class="sr-only">, {{ user.username }}</span></a
                  >
                  <a href="#" class="text-red-600 hover:text-red-900"
                    >删除<span class="sr-only">, {{ user.username }}</span></a
                  >
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from "vue";
import type { UserVO } from "@/types/user";
import { getUserList } from "@/api/user";
const userList = ref<UserVO[]>([]);

const GetUserList = async ({
  page_no,
  page_size,
}: {
  page_no: number;
  page_size: number;
}) => {
  const res = await getUserList({ page_no, page_size });
  userList.value = res.data.list;
};

onMounted(() => {
  GetUserList({ page_no: 1, page_size: 10 });
});
</script>
