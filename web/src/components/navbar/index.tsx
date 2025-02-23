"use client";
import { Navbar, NavbarBrand, NavbarContent, NavbarItem } from "@heroui/react";
import Logo from "./logo";
import ThemeSwitcher from "./switcher";
import Link from "next/link";


export default function NavBar() {
  return (
    <Navbar
      className="hover:bg-zinc-50 dark:hover:bg-zinc-900"
      isBordered
      maxWidth="2xl"
      shouldHideOnScroll
    >
      <NavbarBrand>
        <Logo />
      </NavbarBrand>
      <NavbarContent className="hidden sm:flex gap-4" justify="start">
        <NavbarItem>
          <Link href="/">首页</Link>
        </NavbarItem>
        <NavbarItem>
          <Link href={"/posts/67b876b0e2f2b89ab22d3ca0"}>文章</Link>
        </NavbarItem>
        <NavbarItem>
          <Link href={"/posts/67b876b0e2f2b89ab22d3ca0"}>文章</Link>
        </NavbarItem>
      </NavbarContent>
      <NavbarContent justify="end">
        <NavbarItem>
          <ThemeSwitcher />
        </NavbarItem>
      </NavbarContent>
    </Navbar>
  );
}
