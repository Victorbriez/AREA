import React, { useState, useEffect } from "react";
import {
  Navbar as NextUINavbar,
  NavbarContent,
  NavbarBrand,
  NavbarItem,
  NavbarMenu,
  NavbarMenuToggle,
  NavbarMenuItem,
} from "@nextui-org/navbar";
import { Link } from "@nextui-org/link";
import { link as linkStyles } from "@nextui-org/theme";
import NextLink from "next/link";
import clsx from "clsx";
import { useRouter } from "next/router";
import Cookies from "js-cookie";
import { Button } from "@nextui-org/react";

import { siteConfig } from "@/config/site";
import { ThemeSwitch } from "@/components/theme-switch";

export const Navbar = () => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const router = useRouter();

  useEffect(() => {
    const token = Cookies.get("token");

    setIsAuthenticated(!!token);
  }, []);

  const handleLogout = () => {
    Cookies.remove("token");
    setIsAuthenticated(false);
    router.push("/");
  };

  const renderAuthButton = () => {
    if (isAuthenticated) {
      return (
        <Button color="secondary" variant="flat" onClick={handleLogout}>
          Déconnexion
        </Button>
      );
    }

    return (
      <Button
        color="secondary"
        variant="flat"
        onClick={() => router.push("/auth/login")}
      >
        Connexion
      </Button>
    );
  };

  const renderAuthButtonMobile = () => {
    if (isAuthenticated) {
      return (
        <Link as="button" color="secondary" size="lg" onClick={handleLogout}>
          Déconnexion
        </Link>
      );
    }

    return (
      <Link
        as="button"
        color="secondary"
        size="lg"
        onClick={() => router.push("/auth/login")}
      >
        Connexion
      </Link>
    );
  };

  return (
    <NextUINavbar isBordered maxWidth="xl" position="sticky">
      <NavbarContent className="basis-1/5 sm:basis-1/3" justify="start">
        <NavbarBrand as="li" className="gap-3 max-w-fit">
          <NextLink className="flex justify-start items-center gap-1" href="/">
            <p className="font-bold text-inherit">Area</p>
          </NextLink>
        </NavbarBrand>
      </NavbarContent>

      <NavbarContent className="hidden sm:flex basis-1/3" justify="center">
        <ul className="flex gap-8 justify-center items-center h-full">
          {siteConfig.navItems.map((item) => (
            <NavbarItem key={item.href}>
              <NextLink
                className={clsx(
                  linkStyles({ color: "foreground" }),
                  "relative group data-[active=true]:text-secondary data-[active=true]:font-medium",
                )}
                color="foreground"
                href={item.href}
              >
                {item.label}
                <span className="absolute left-0 bottom-0 w-0 h-[4px] bg-secondary transition-all duration-100 group-hover:w-full" />
              </NextLink>
            </NavbarItem>
          ))}
        </ul>
      </NavbarContent>

      <NavbarContent className="hidden sm:flex basis-1/3" justify="end">
        <NavbarItem>
          <ThemeSwitch />
        </NavbarItem>
        <NavbarItem>{renderAuthButton()}</NavbarItem>
      </NavbarContent>

      <NavbarContent className="sm:hidden basis-1 pl-4" justify="end">
        <ThemeSwitch />
        <NavbarMenuToggle />
      </NavbarContent>

      <NavbarMenu>
        <div className="mx-4 mt-2 flex flex-col gap-2">
          {siteConfig.navMenuItems.map((item, index) => (
            <NavbarMenuItem key={`${item}-${index}`}>
              <Link color={"foreground"} href={item.href} size="lg">
                {item.label}
              </Link>
            </NavbarMenuItem>
          ))}
          <NavbarMenuItem key="auth">{renderAuthButtonMobile()}</NavbarMenuItem>
        </div>
      </NavbarMenu>
    </NextUINavbar>
  );
};
