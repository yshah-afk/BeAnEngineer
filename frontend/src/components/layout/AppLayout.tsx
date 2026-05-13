import { Outlet } from "react-router-dom";
import { motion } from "framer-motion";
import { Sidebar } from "./Sidebar";
import { TopBar } from "./TopBar";
import { useSidebarStore } from "@/stores/sidebarStore";
import { cn } from "@/lib/utils";

export function AppLayout() {
  const isCollapsed = useSidebarStore((s) => s.isCollapsed);

  return (
    <div className="min-h-screen bg-background">
      <Sidebar />
      <motion.div
        initial={false}
        animate={{ marginLeft: isCollapsed ? 64 : 256 }}
        transition={{ duration: 0.2 }}
        className={cn("flex min-h-screen flex-col transition-[margin] max-lg:!ml-0")}
      >
        <TopBar />
        <main className="flex-1">
          <Outlet />
        </main>
      </motion.div>
    </div>
  );
}
