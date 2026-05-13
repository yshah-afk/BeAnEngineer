import { Flame } from "lucide-react";
import { cn } from "@/lib/utils";

interface StreakBadgeProps {
  streak: number;
  className?: string;
}

export function StreakBadge({ streak, className }: StreakBadgeProps) {
  const isActive = streak > 0;

  return (
    <div
      className={cn(
        "inline-flex items-center gap-1.5 rounded-full px-3 py-1 text-sm font-medium",
        isActive
          ? "bg-orange-500/15 text-orange-400"
          : "bg-muted text-muted-foreground",
        className,
      )}
    >
      <Flame className={cn("h-4 w-4", isActive && "text-orange-400")} />
      <span className="tabular-nums">{streak}</span>
      <span className="text-xs opacity-70">day{streak !== 1 ? "s" : ""}</span>
    </div>
  );
}
