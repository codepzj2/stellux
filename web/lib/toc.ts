// @ts-nocheck

import { toc } from "mdast-util-toc";
import { remark } from "remark";
import { visit } from "unist-util-visit";

const textTypes = ["text", "emphasis", "strong", "inlineCode"];

function flattenNode(node) {
  const p = [];
  visit(node, node => {
    if (!textTypes.includes(node.type)) return;
    p.push(node.value);
  });
  return p.join(``);
}

interface Item {
  title: string;
  url: string;
  items?: Item[];
}

interface Items {
  items?: Item[];
}

function getItems(node, current, level = 1): Items {
  if (!node) {
    return {};
  }

  if (node.type === "paragraph") {
    visit(node, item => {
      if (item.type === "link") {
        current.url = item.url;
        current.title = flattenNode(node);
      }

      if (item.type === "text") {
        current.title = flattenNode(node);
      }
    });

    return current;
  }

  if (node.type === "list") {
    // 如果当前是列表，递归获取列表项
    if (level >= 3) {
      return current; // 超过三级，停止递归
    }

    current.items = node.children.map(i => getItems(i, {}, level + 1));

    return current;
  } else if (node.type === "listItem") {
    const heading = getItems(node.children[0], {}, level);

    if (node.children.length > 1) {
      getItems(node.children[1], heading, level);
    }

    return heading;
  }

  return {};
}

const updateUrlsWithIncrement = (items, counter = { count: 1 }) => {
  return items?.map(item => {
    const newItem = { ...item, url: `#header-${counter.count++}` };

    if (item.items) {
      newItem.items = updateUrlsWithIncrement(item.items, counter);
    }

    return newItem;
  });
};

const getToc = () => (node, file) => {
  const table = toc(node);
  const items = getItems(table.map, {}, 1); // 从一级开始递归

  file.data = updateUrlsWithIncrement(items.items);
};

export type TableOfContents = Items;

export async function getTableOfContents(
  content: string
): Promise<TableOfContents> {
  const result = await remark().use(getToc).process(content);

  return result.data;
}
