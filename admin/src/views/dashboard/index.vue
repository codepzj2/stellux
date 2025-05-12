<template>
  <div>
    <a-card class="mx-auto my-4 p-6 rounded-2xl" :bordered="false">
      <a-typography-title :level="2">你好 👋</a-typography-title>
      <a-typography-paragraph>
        ✨ <strong>欢迎来到我的 GitHub 主页！</strong> ✨
      </a-typography-paragraph>
      <a-typography-paragraph>
        我是 <strong>codepzj</strong>，我对 <strong>Golang</strong> 和
        <strong>Python</strong> 很感兴趣。
      </a-typography-paragraph>

      <a-divider />

      <a-typography-title :level="4">联系方式</a-typography-title>
      <ul class="list-disc pl-5 space-y-1 text-gray-700">
        <li>QQ: 2363435714@qq.com</li>
        <li>
          Email:
          <a
            href="mailto:email@codepzj.cn"
            class="text-blue-600 hover:underline"
            >email@codepzj.cn</a
          >
        </li>
        <li>
          Blog:
          <a
            href="https://haohanxinghe.com"
            target="_blank"
            class="text-blue-600 hover:underline"
            >haohanxinghe.com</a
          >
        </li>
        <li>
          Docs:
          <a
            href="https://docs.codepzj.cn"
            target="_blank"
            class="text-blue-600 hover:underline"
            >docs.codepzj.cn</a
          >
        </li>
      </ul>

      <a-divider />

      <div class="flex justify-center my-4">
        <a href="https://github.com/codepzj" target="_blank">
          <img
            class="rounded-lg"
            src="https://github-readme-stats.vercel.app/api?username=codepzj&show_icons=true&count_private=true"
            alt="GitHub stats"
          />
        </a>
      </div>

      <a-divider />

      <a-typography-title :level="4">技术栈</a-typography-title>
      <div class="flex justify-center mb-6">
        <img
          src="https://skillicons.dev/icons?i=vue,nuxt,react,cpp,python,go,docker,git,github"
          alt="My Skills"
          class="rounded-md"
        />
      </div>

      <a-divider />

      <a-typography-title :level="4">座右铭</a-typography-title>
      <blockquote class="border-l-4 border-gray-300 pl-4 italic text-gray-600">
        低级的欲望通过放纵就可获得；高级的欲望通过自律方可获得；顶级的欲望通过煎熬才可获得。
        “所谓自由，不是随心所欲，而是自我主宰。”
      </blockquote>
    </a-card>
  </div>
</template>
