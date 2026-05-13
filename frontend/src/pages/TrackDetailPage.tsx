import { useParams, Link } from "react-router-dom";
import { useQuery } from "@tanstack/react-query";
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
  ChevronRight,
  BarChart3,
} from "lucide-react";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Progress } from "@/components/ui/progress";
import { Skeleton } from "@/components/ui/skeleton";
import { ModuleAccordion } from "@/components/shared/ModuleAccordion";
import { tracksApi } from "@/lib/api";
import { useTrackProgress } from "@/hooks/useProgress";
import { TRACK_COLORS } from "@/lib/constants";
import { cn } from "@/lib/utils";
import { formatHours } from "@/lib/utils";

const iconMap: Record<string, React.ElementType> = {
  Brain, Server, Layout, Cloud, Code2, Network,
};

export default function TrackDetailPage() {
  const { slug } = useParams<{ slug: string }>();

  const { data: track, isLoading } = useQuery({
    queryKey: ["tracks", slug],
    queryFn: () => tracksApi.get(slug!),
    enabled: !!slug,
  });

  const { data: allProgress } = useTrackProgress();
  const progress = allProgress?.find((p) => p.trackSlug === slug);

  if (isLoading) {
    return (
      <div className="mx-auto max-w-4xl px-6 py-10 space-y-6">
        <Skeleton className="h-48 rounded-xl" />
        <Skeleton className="h-12 rounded-xl" />
        <div className="space-y-4">
          {Array.from({ length: 5 }, (_, i) => (
            <Skeleton key={i} className="h-16 rounded-xl" />
          ))}
        </div>
      </div>
    );
  }

  if (!track) {
    return (
      <div className="flex flex-col items-center justify-center py-20 text-center">
        <h2 className="text-2xl font-bold mb-2">Track not found</h2>
        <p className="text-muted-foreground mb-4">
          The track you're looking for doesn't exist.
        </p>
        <Button render={<Link to="/tracks" />}>Browse Tracks</Button>
      </div>
    );
  }

  const Icon = iconMap[track.icon] ?? BookOpen;
  const gradient = TRACK_COLORS[track.slug] ?? "from-indigo-500 to-purple-600";

  return (
    <div className="mx-auto max-w-4xl px-6 py-10">
      {/* Breadcrumb */}
      <nav className="flex items-center gap-1.5 text-sm text-muted-foreground mb-6">
        <Link to="/tracks" className="hover:text-foreground transition-colors">
          Tracks
        </Link>
        <ChevronRight className="h-3.5 w-3.5" />
        <span className="text-foreground font-medium">{track.title}</span>
      </nav>

      {/* Track Header */}
      <motion.div
        initial={{ opacity: 0, y: 20 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.4 }}
      >
        <Card className="border-border/50 bg-card/80 overflow-hidden">
          <div className={cn("h-2 bg-gradient-to-r", gradient)} />
          <CardContent className="p-6 sm:p-8">
            <div className="flex flex-col sm:flex-row gap-6">
              <div
                className={cn(
                  "flex h-16 w-16 shrink-0 items-center justify-center rounded-2xl bg-gradient-to-br text-white shadow-lg",
                  gradient,
                )}
              >
                <Icon className="h-8 w-8" />
              </div>
              <div className="flex-1">
                <h1 className="text-2xl sm:text-3xl font-bold text-foreground">
                  {track.title}
                </h1>
                <p className="mt-2 text-muted-foreground leading-relaxed">
                  {track.description}
                </p>
                <div className="mt-4 flex flex-wrap items-center gap-4 text-sm text-muted-foreground">
                  <span className="flex items-center gap-1.5">
                    <BookOpen className="h-4 w-4" />
                    {track.lessonCount} lessons
                  </span>
                  <span className="flex items-center gap-1.5">
                    <BarChart3 className="h-4 w-4" />
                    {track.moduleCount} modules
                  </span>
                  <span className="flex items-center gap-1.5">
                    <Clock className="h-4 w-4" />
                    ~{formatHours(track.estimatedHours)}
                  </span>
                </div>

                {progress && progress.percentage > 0 && (
                  <div className="mt-5 space-y-2">
                    <div className="flex justify-between text-sm">
                      <span className="text-muted-foreground">
                        {progress.completedLessons} of {progress.totalLessons} lessons complete
                      </span>
                      <span className="font-medium text-primary tabular-nums">
                        {progress.percentage}%
                      </span>
                    </div>
                    <Progress value={progress.percentage} className="h-2" />
                  </div>
                )}
              </div>
            </div>
          </CardContent>
        </Card>
      </motion.div>

      {/* Modules */}
      <div className="mt-8">
        <h2 className="text-xl font-bold text-foreground mb-4">Modules</h2>
        <ModuleAccordion modules={track.modules} trackSlug={track.slug} />
      </div>
    </div>
  );
}
