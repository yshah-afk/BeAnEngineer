import { Link } from "react-router-dom";
import { motion } from "framer-motion";
import {
  Flame,
  Trophy,
  Clock,
  BookOpen,
  TrendingUp,
} from "lucide-react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Progress } from "@/components/ui/progress";
import { Skeleton } from "@/components/ui/skeleton";
import { Badge } from "@/components/ui/badge";
import { ProgressRing } from "@/components/shared/ProgressRing";
import { StreakBadge } from "@/components/shared/StreakBadge";
import { useDashboard } from "@/hooks/useProgress";
import { TRACK_ACCENT_TEXT } from "@/lib/constants";
import { formatTime, formatRelativeTime, cn } from "@/lib/utils";

export default function DashboardPage() {
  const { data, isLoading } = useDashboard();

  if (isLoading) {
    return (
      <div className="mx-auto max-w-5xl px-6 py-10 space-y-6">
        <Skeleton className="h-8 w-48" />
        <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-4">
          {Array.from({ length: 4 }, (_, i) => (
            <Skeleton key={i} className="h-28 rounded-xl" />
          ))}
        </div>
        <Skeleton className="h-80 rounded-xl" />
      </div>
    );
  }

  if (!data) {
    return (
      <div className="flex items-center justify-center py-20 text-muted-foreground">
        Unable to load dashboard data.
      </div>
    );
  }

  return (
    <div className="mx-auto max-w-5xl px-6 py-10">
      <motion.div
        initial={{ opacity: 0, y: 10 }}
        animate={{ opacity: 1, y: 0 }}
        className="mb-8"
      >
        <h1 className="text-3xl font-bold text-foreground">Dashboard</h1>
        <p className="mt-1 text-muted-foreground">Track your learning progress</p>
      </motion.div>

      {/* Stats grid */}
      <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-4 mb-8">
        {[
          {
            icon: BookOpen,
            label: "Completed",
            value: `${data.totalCompleted}/${data.totalLessons}`,
            sub: "lessons",
            color: "text-primary",
          },
          {
            icon: Flame,
            label: "Current Streak",
            value: data.currentStreak,
            sub: "days",
            color: "text-orange-400",
          },
          {
            icon: Trophy,
            label: "Longest Streak",
            value: data.longestStreak,
            sub: "days",
            color: "text-yellow-400",
          },
          {
            icon: Clock,
            label: "Time Spent",
            value: formatTime(data.totalTimeSpent),
            sub: "total",
            color: "text-cyan-400",
          },
        ].map(({ icon: Icon, label, value, sub, color }, i) => (
          <motion.div
            key={label}
            initial={{ opacity: 0, y: 10 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: i * 0.05 }}
          >
            <Card className="border-border/50 bg-card/80">
              <CardContent className="p-5">
                <div className="flex items-center gap-3">
                  <Icon className={cn("h-5 w-5", color)} />
                  <span className="text-sm text-muted-foreground">{label}</span>
                </div>
                <p className="mt-2 text-2xl font-bold text-foreground tabular-nums">
                  {value}
                </p>
                <p className="text-xs text-muted-foreground">{sub}</p>
              </CardContent>
            </Card>
          </motion.div>
        ))}
      </div>

      <div className="grid gap-6 lg:grid-cols-[1fr_300px]">
        {/* Left column */}
        <div className="space-y-6">
          {/* Overall progress */}
          <Card className="border-border/50 bg-card/80">
            <CardHeader>
              <CardTitle className="text-lg flex items-center gap-2">
                <TrendingUp className="h-5 w-5 text-primary" />
                Overall Progress
              </CardTitle>
            </CardHeader>
            <CardContent>
              <div className="flex items-center gap-8">
                <ProgressRing percentage={data.overallPercentage} size={100} />
                <div className="flex-1 space-y-4">
                  {data.trackProgress.map((tp) => (
                    <div key={tp.trackSlug}>
                      <div className="flex items-center justify-between text-sm mb-1">
                        <Link
                          to={`/tracks/${tp.trackSlug}`}
                          className={cn(
                            "font-medium hover:underline",
                            TRACK_ACCENT_TEXT[tp.trackSlug] ?? "text-foreground",
                          )}
                        >
                          {tp.trackSlug
                            .split("-")
                            .map((w) => w.charAt(0).toUpperCase() + w.slice(1))
                            .join(" ")}
                        </Link>
                        <span className="text-xs text-muted-foreground tabular-nums">
                          {tp.completedLessons}/{tp.totalLessons}
                        </span>
                      </div>
                      <Progress
                        value={tp.percentage}
                        className="h-1.5"
                      />
                    </div>
                  ))}
                </div>
              </div>
            </CardContent>
          </Card>

          {/* Streak calendar */}
          <Card className="border-border/50 bg-card/80">
            <CardHeader>
              <CardTitle className="text-lg flex items-center gap-2">
                <Flame className="h-5 w-5 text-orange-400" />
                Activity
              </CardTitle>
            </CardHeader>
            <CardContent>
              <div className="flex flex-wrap gap-1">
                {data.streakCalendar.map((day) => (
                  <div
                    key={day.date}
                    className={cn(
                      "h-3 w-3 rounded-sm transition-colors",
                      day.count === 0 && "bg-muted/50",
                      day.count === 1 && "bg-emerald-500/30",
                      day.count === 2 && "bg-emerald-500/50",
                      day.count >= 3 && "bg-emerald-500/80",
                    )}
                    title={`${day.date}: ${day.count} lesson${day.count !== 1 ? "s" : ""}`}
                  />
                ))}
              </div>
              <div className="flex items-center justify-end gap-2 mt-3 text-xs text-muted-foreground">
                <span>Less</span>
                <div className="flex gap-0.5">
                  <div className="h-2.5 w-2.5 rounded-sm bg-muted/50" />
                  <div className="h-2.5 w-2.5 rounded-sm bg-emerald-500/30" />
                  <div className="h-2.5 w-2.5 rounded-sm bg-emerald-500/50" />
                  <div className="h-2.5 w-2.5 rounded-sm bg-emerald-500/80" />
                </div>
                <span>More</span>
              </div>
            </CardContent>
          </Card>
        </div>

        {/* Right column */}
        <div className="space-y-6">
          <Card className="border-border/50 bg-card/80">
            <CardHeader>
              <CardTitle className="text-sm">Your Streak</CardTitle>
            </CardHeader>
            <CardContent className="text-center">
              <StreakBadge streak={data.currentStreak} className="text-base px-4 py-2" />
              <p className="mt-3 text-xs text-muted-foreground">
                Best: {data.longestStreak} days
              </p>
            </CardContent>
          </Card>

          <Card className="border-border/50 bg-card/80">
            <CardHeader>
              <CardTitle className="text-sm flex items-center justify-between">
                Recent Activity
                <Badge variant="secondary" className="text-[10px]">
                  {data.recentActivity.length}
                </Badge>
              </CardTitle>
            </CardHeader>
            <CardContent>
              {data.recentActivity.length === 0 ? (
                <p className="text-sm text-muted-foreground text-center py-4">
                  No recent activity
                </p>
              ) : (
                <div className="space-y-3">
                  {data.recentActivity.slice(0, 8).map((item) => (
                    <div key={item.id} className="flex items-start gap-2">
                      <div className="mt-1.5 h-1.5 w-1.5 rounded-full bg-primary shrink-0" />
                      <div className="min-w-0">
                        <p className="text-xs text-foreground line-clamp-2">
                          {item.description}
                        </p>
                        <p className="text-[10px] text-muted-foreground mt-0.5">
                          {formatRelativeTime(item.createdAt)}
                        </p>
                      </div>
                    </div>
                  ))}
                </div>
              )}
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  );
}
