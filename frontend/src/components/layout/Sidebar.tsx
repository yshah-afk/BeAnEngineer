import { NavLink } from "react-router-dom";
import { motion, AnimatePresence } from "framer-motion";
import {
  Home,
  BookOpen,
  BarChart3,
  Bookmark,
  Code2,
  Layers,
  GraduationCap,
  Shield,
  ChevronLeft,
  Sparkles,
} from "lucide-react";
import { cn } from "@/lib/utils";
import { useSidebarStore } from "@/stores/sidebarStore";
import { useAuthStore } from "@/stores/authStore";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { ScrollArea } from "@/components/ui/scroll-area";

const mainNav = [
  { to: "/", icon: Home, label: "Home" },
  { to: "/tracks", icon: BookOpen, label: "Tracks" },
  { to: "/playground", icon: Code2, label: "Playground" },
  { to: "/flashcards", icon: Layers, label: "Flashcards" },
];

const userNav = [
  { to: "/dashboard", icon: BarChart3, label: "Dashboard" },
  { to: "/bookmarks", icon: Bookmark, label: "Bookmarks" },
];

const adminNav = [
  { to: "/admin", icon: Shield, label: "Admin" },
];

export function Sidebar() {
  const { isCollapsed, setCollapsed } = useSidebarStore();
  const { isAuthenticated, user } = useAuthStore();

  return (
    <motion.aside
      initial={false}
      animate={{ width: isCollapsed ? 64 : 256 }}
      transition={{ duration: 0.2 }}
      className="fixed inset-y-0 left-0 z-30 flex flex-col border-r border-sidebar-border bg-sidebar"
    >
      <div className="flex h-16 items-center gap-2 px-4">
        <AnimatePresence mode="wait">
          {!isCollapsed && (
            <motion.div
              initial={{ opacity: 0 }}
              animate={{ opacity: 1 }}
              exit={{ opacity: 0 }}
              className="flex items-center gap-2 min-w-0"
            >
              <div className="flex h-8 w-8 shrink-0 items-center justify-center rounded-lg bg-primary">
                <GraduationCap className="h-4.5 w-4.5 text-primary-foreground" />
              </div>
              <span className="text-sm font-bold text-sidebar-foreground truncate">
                Mastery Hub
              </span>
            </motion.div>
          )}
        </AnimatePresence>
        {isCollapsed && (
          <div className="flex h-8 w-8 mx-auto items-center justify-center rounded-lg bg-primary">
            <GraduationCap className="h-4.5 w-4.5 text-primary-foreground" />
          </div>
        )}
      </div>

      <ScrollArea className="flex-1 px-3 py-2">
        <nav className="space-y-1">
          {mainNav.map(({ to, icon: Icon, label }) => (
            <NavLink
              key={to}
              to={to}
              end={to === "/"}
              className={({ isActive }) =>
                cn(
                  "flex items-center gap-3 rounded-lg px-3 py-2 text-sm transition-colors",
                  isActive
                    ? "bg-sidebar-accent text-sidebar-primary font-medium"
                    : "text-sidebar-foreground/70 hover:bg-sidebar-accent/50 hover:text-sidebar-foreground",
                  isCollapsed && "justify-center px-2",
                )
              }
            >
              <Icon className="h-4.5 w-4.5 shrink-0" />
              {!isCollapsed && <span>{label}</span>}
            </NavLink>
          ))}

          {isAuthenticated && (
            <>
              <Separator className="my-3 bg-sidebar-border" />
              {userNav.map(({ to, icon: Icon, label }) => (
                <NavLink
                  key={to}
                  to={to}
                  className={({ isActive }) =>
                    cn(
                      "flex items-center gap-3 rounded-lg px-3 py-2 text-sm transition-colors",
                      isActive
                        ? "bg-sidebar-accent text-sidebar-primary font-medium"
                        : "text-sidebar-foreground/70 hover:bg-sidebar-accent/50 hover:text-sidebar-foreground",
                      isCollapsed && "justify-center px-2",
                    )
                  }
                >
                  <Icon className="h-4.5 w-4.5 shrink-0" />
                  {!isCollapsed && <span>{label}</span>}
                </NavLink>
              ))}
            </>
          )}

          {user?.role === "admin" && (
            <>
              <Separator className="my-3 bg-sidebar-border" />
              {adminNav.map(({ to, icon: Icon, label }) => (
                <NavLink
                  key={to}
                  to={to}
                  className={({ isActive }) =>
                    cn(
                      "flex items-center gap-3 rounded-lg px-3 py-2 text-sm transition-colors",
                      isActive
                        ? "bg-sidebar-accent text-sidebar-primary font-medium"
                        : "text-sidebar-foreground/70 hover:bg-sidebar-accent/50 hover:text-sidebar-foreground",
                      isCollapsed && "justify-center px-2",
                    )
                  }
                >
                  <Icon className="h-4.5 w-4.5 shrink-0" />
                  {!isCollapsed && <span>{label}</span>}
                </NavLink>
              ))}
            </>
          )}
        </nav>

        {!isCollapsed && (
          <div className="mt-6 rounded-xl bg-primary/5 border border-primary/10 p-4">
            <div className="flex items-center gap-2 mb-2">
              <Sparkles className="h-4 w-4 text-primary" />
              <span className="text-xs font-semibold text-foreground">AI Tutor</span>
            </div>
            <p className="text-xs text-muted-foreground leading-relaxed">
              Get help from the AI tutor while studying any lesson.
            </p>
          </div>
        )}
      </ScrollArea>

      <div className="border-t border-sidebar-border p-2">
        <Button
          variant="ghost"
          size="sm"
          onClick={() => setCollapsed(!isCollapsed)}
          className={cn(
            "w-full text-sidebar-foreground/60 hover:text-sidebar-foreground",
            isCollapsed && "px-2",
          )}
        >
          <ChevronLeft
            className={cn(
              "h-4 w-4 transition-transform",
              isCollapsed && "rotate-180",
            )}
          />
          {!isCollapsed && <span className="ml-2 text-xs">Collapse</span>}
        </Button>
      </div>
    </motion.aside>
  );
}
