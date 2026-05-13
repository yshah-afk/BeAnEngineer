import { Link } from "react-router-dom";
import { motion } from "framer-motion";
import {
  Brain,
  Server,
  Layout,
  Cloud,
  Code2,
  Network,
  BookOpen,
  Clock,
} from "lucide-react";
import { Card, CardContent } from "@/components/ui/card";
import { Progress } from "@/components/ui/progress";
import { cn } from "@/lib/utils";
import { TRACK_COLORS } from "@/lib/constants";
import type { Track, TrackProgress } from "@/types";

const iconMap: Record<string, React.ElementType> = {
  Brain, Server, Layout, Cloud, Code2, Network,
};

interface TrackCardProps {
  track: Track;
  progress?: TrackProgress;
  index?: number;
}

export function TrackCard({ track, progress, index = 0 }: TrackCardProps) {
  const Icon = iconMap[track.icon] ?? BookOpen;
  const gradient = TRACK_COLORS[track.slug] ?? "from-indigo-500 to-purple-600";

  return (
    <motion.div
      initial={{ opacity: 0, y: 20 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.4, delay: index * 0.08 }}
    >
      <Link to={`/tracks/${track.slug}`}>
        <Card className="group relative overflow-hidden border-border/50 bg-card/80 backdrop-blur-sm transition-all duration-300 hover:border-primary/30 hover:shadow-lg hover:shadow-primary/5 hover:-translate-y-1">
          <div
            className={cn(
              "absolute inset-x-0 top-0 h-1 bg-gradient-to-r opacity-60 transition-opacity group-hover:opacity-100",
              gradient,
            )}
          />
          <CardContent className="p-6">
            <div className="flex items-start gap-4">
              <div
                className={cn(
                  "flex h-12 w-12 shrink-0 items-center justify-center rounded-xl bg-gradient-to-br text-white shadow-lg",
                  gradient,
                )}
              >
                <Icon className="h-6 w-6" />
              </div>
              <div className="min-w-0 flex-1">
                <h3 className="font-semibold text-foreground truncate text-base">
                  {track.title}
                </h3>
                <p className="mt-1 text-sm text-muted-foreground line-clamp-2 leading-relaxed">
                  {track.description}
                </p>
              </div>
            </div>

            <div className="mt-4 flex items-center gap-4 text-xs text-muted-foreground">
              <span className="flex items-center gap-1">
                <BookOpen className="h-3.5 w-3.5" />
                {track.lessonCount} lessons
              </span>
              <span className="flex items-center gap-1">
                <Clock className="h-3.5 w-3.5" />
                ~{track.estimatedHours}h
              </span>
            </div>

            {progress && progress.percentage > 0 && (
              <div className="mt-4 space-y-1.5">
                <div className="flex justify-between text-xs">
                  <span className="text-muted-foreground">
                    {progress.completedLessons}/{progress.totalLessons} lessons
                  </span>
                  <span className="font-medium text-primary tabular-nums">
                    {progress.percentage}%
                  </span>
                </div>
                <Progress value={progress.percentage} className="h-1.5" />
              </div>
            )}
          </CardContent>
        </Card>
      </Link>
    </motion.div>
  );
}
