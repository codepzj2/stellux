declare module "*.vue" {
  import { DefineComponent } from "vue";
  const component: DefineComponent<{}, {}, any>;
  export default component;
}

declare module "@/lib/editor" {
  import type { Component } from "vue";
  const Editor: Component;
  export default Editor;
}
