import {
  NavigationMenu,
  NavigationMenuItem,
  NavigationMenuLink,
  navigationMenuTriggerStyle,
} from "@/components/ui/navigation-menu";

function NavBar() {
  return (
    <NavigationMenu>
      <NavigationMenuItem>
        <NavigationMenuLink className={navigationMenuTriggerStyle()}>
          Home
        </NavigationMenuLink>
      </NavigationMenuItem>
      <NavigationMenuItem>
        <NavigationMenuLink className={navigationMenuTriggerStyle()}>
          Recommender
        </NavigationMenuLink>
      </NavigationMenuItem>
      <NavigationMenuItem>
        <NavigationMenuLink
          href="https://github.com/ttamre/go.tv"
          target="_blank"
          rel="noopener noreferrer"
          className={navigationMenuTriggerStyle()}
        >
          Github
        </NavigationMenuLink>
      </NavigationMenuItem>
    </NavigationMenu>
  );
}

export default NavBar;
