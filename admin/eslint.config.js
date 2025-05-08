import globals from "globals";
import pluginJs from "@eslint/js";
import tseslint from "typescript-eslint";
import pluginVue from "eslint-plugin-vue";


/** @type {import('eslint').Linter.Config[]} */
export default [
  { files: ["**/*.{js,mjs,cjs,ts,vue}"] },
  { languageOptions: { globals: globals.browser } },
  pluginJs.configs.recommended,
  ...tseslint.configs.recommended,
  ...pluginVue.configs["flat/essential"],
  {
    files: ["**/*.vue"],
    languageOptions: { parserOptions: { parser: tseslint.parser } },
  },
  {
    rules: {
      "vue/multi-word-component-names": "off", // 禁用 文件大驼峰 规则
      "@typescript-eslint/no-explicit-any": ["off"], // 允许 any 类型
      "no-undef": "off", // 允许未定义的变量
      "@typescript-eslint/no-unused-vars": "off", // 允许未使用的变量
      "@typescript-eslint/no-empty-object-type": "off", // 允许空对象类型
      "@typescript-eslint/ban-ts-comment": "off", // 允许 ts-comment
    },
  },
];