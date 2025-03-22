"use client";

import { useState, useEffect } from "react";
import { Button } from "@/components/ui/button";
import Logo from "./logo";
import Nav from "./nav";
import ThemeSwitcher from "./theme-switcher";
import DropDownMenu from "./drop-menu";
import Github from "@/components/icons/github";
import clsx from "clsx";

export default function NavBar() {
  const [hidden, setHidden] = useState(false);
  const [lastScrollY, setLastScrollY] = useState(0);

  useEffect(() => {
    const handleScroll = () => {
      const currentScrollY = window.scrollY;
      setHidden(currentScrollY > lastScrollY && currentScrollY > 50);
      setLastScrollY(currentScrollY);
    };

    window.addEventListener("scroll", handleScroll);
    return () => window.removeEventListener("scroll", handleScroll);
  }, [lastScrollY]);

  return (
    <header
      className={clsx(
        "fixed top-0 z-50 w-full border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60 transition-transform duration-300",
        hidden ? "-translate-y-full" : "translate-y-0"
      )}
    >
      <div className="container-wrapper">
        <div className="flex h-14 items-center gap-2 md:gap-4">
          <div className="flex h-full w-full items-center justify-between px-4">
            <div className="mr-4 hidden md:flex">
              <div className="flex items-center gap-10">
                <Logo />
                <Nav />
              </div>
            </div>
            <div className="md:hidden">
              <DropDownMenu />
            </div>
            <div className="ml-auto flex items-center gap-2 md:flex-1 md:justify-end">
              <Button
                variant="secondary"
                size="icon"
                className="hover:bg-gray-100 dark:hover:bg-gray-800 cursor-pointer"
                onClick={() =>
                  window.open("https://github.com/codepzj/Stellux", "_blank")
                }
              >
                <Github />
              </Button>
              <ThemeSwitcher />
            </div>
          </div>
        </div>
      </div>
    </header>
  );
}
