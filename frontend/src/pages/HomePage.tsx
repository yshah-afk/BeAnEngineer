import { useQuery } from "@tanstack/react-query";
import { motion } from "framer-motion";
import { Link } from "react-router-dom";
import {
  BookOpen,
  Clock,
  Code2,
  Layers,
  GraduationCap,
  ArrowRight,
  Sparkles,
} from "lucide-react";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";
import { TrackCard } from "@/components/shared/TrackCard";
import { Footer } from "@/components/layout/Footer";
import { tracksApi } from "@/lib/api";
import { useTrackProgress } from "@/hooks/useProgress";
import { useAuthStore } from "@/stores/authStore";

const STATS = [
  { icon: BookOpen, label: "Lessons", value: "170" },
  { icon: Layers, label: "Tracks", value: "6" },
  { icon: Code2, label: "Projects", value: "12" },
  { icon: Clock, label: "Hours", value: "~170" },
];

export default function HomePage() {
  const { isAuthenticated } = useAuthStore();

  const { data: tracks, isLoading } = useQuery({
    queryKey: ["tracks"],
    queryFn: tracksApi.list,
  });

  const { data: trackProgress } = useTrackProgress();

  const continueTrack = trackProgress?.find((tp) => tp.percentage > 0 && tp.percentage < 100);

  return (
    <div className="flex flex-col">
      {/* Hero */}
      <section className="relative overflow-hidden px-6 py-16 sm:py-24">
        <div className="absolute inset-0 bg-gradient-to-b from-primary/5 via-transparent to-transparent" />
        <div className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[600px] h-[600px] bg-primary/5 rounded-full blur-3xl" />

        <div className="relative mx-auto max-w-3xl text-center">
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.5 }}
          >
            <div className="inline-flex items-center gap-2 rounded-full bg-primary/10 px-4 py-1.5 text-sm text-primary mb-6">
              <Sparkles className="h-4 w-4" />
              AI-Powered Learning Platform
            </div>
            <h1 className="text-4xl sm:text-5xl lg:text-6xl font-bold tracking-tight text-foreground leading-[1.1]">
              AI & Full-Stack{" "}
              <span className="bg-gradient-to-r from-primary to-purple-400 bg-clip-text text-transparent">
                Mastery Hub
              </span>
            </h1>
            <p className="mt-6 text-lg text-muted-foreground max-w-xl mx-auto leading-relaxed">
              Master LLMs, Go, React, DevOps, DSA, and System Design through
              170+ interactive lessons, hands-on projects, and AI-powered tutoring.
            </p>
          </motion.div>

          <motion.div
            initial={{ opacity: 0, y: 10 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.5, delay: 0.2 }}
            className="mt-8 flex flex-wrap items-center justify-center gap-4"
          >
            <Button size="lg" render={<Link to="/tracks" />} className="gap-2">
              <GraduationCap className="h-5 w-5" />
              Start Learning
            </Button>
            {!isAuthenticated && (
              <Button size="lg" variant="outline" render={<Link to="/register" />}>
                Create Account
              </Button>
            )}
          </motion.div>
        </div>
      </section>

      {/* Stats */}
      <section className="px-6 pb-12">
        <div className="mx-auto max-w-4xl">
          <div className="grid grid-cols-2 sm:grid-cols-4 gap-4">
            {STATS.map(({ icon: Icon, label, value }, i) => (
              <motion.div
                key={label}
                initial={{ opacity: 0, y: 10 }}
                animate={{ opacity: 1, y: 0 }}
                transition={{ duration: 0.3, delay: 0.3 + i * 0.05 }}
              >
                <Card className="border-border/30 bg-card/60 backdrop-blur-sm text-center">
                  <CardContent className="py-5">
                    <Icon className="h-5 w-5 mx-auto text-primary mb-2" />
                    <p className="text-2xl font-bold text-foreground">{value}</p>
                    <p className="text-xs text-muted-foreground">{label}</p>
                  </CardContent>
                </Card>
              </motion.div>
            ))}
          </div>
        </div>
      </section>

      {/* Continue learning */}
      {isAuthenticated && continueTrack && (
        <section className="px-6 pb-8">
          <div className="mx-auto max-w-4xl">
            <Card className="border-primary/20 bg-primary/5">
              <CardContent className="flex items-center gap-4 p-5">
                <div className="flex-1">
                  <p className="text-xs text-primary font-medium mb-1">
                    Continue where you left off
                  </p>
                  <p className="text-sm font-medium text-foreground">
                    {continueTrack.lastLessonTitle ?? "Continue learning"}
                  </p>
                  <p className="text-xs text-muted-foreground mt-0.5">
                    {continueTrack.completedLessons}/{continueTrack.totalLessons} lessons ·{" "}
                    {continueTrack.percentage}% complete
                  </p>
                </div>
                <Button
                  size="sm"
                  className="gap-1.5 shrink-0"
                  render={
                    <Link
                      to={
                        continueTrack.lastLessonSlug
                          ? `/tracks/${continueTrack.trackSlug}/lessons/${continueTrack.lastLessonSlug}`
                          : `/tracks/${continueTrack.trackSlug}`
                      }
                    />
                  }
                >
                  Continue
                  <ArrowRight className="h-4 w-4" />
                </Button>
              </CardContent>
            </Card>
          </div>
        </section>
      )}

      {/* Tracks */}
      <section className="px-6 pb-16">
        <div className="mx-auto max-w-4xl">
          <div className="flex items-center justify-between mb-6">
            <h2 className="text-2xl font-bold text-foreground">Learning Tracks</h2>
            <Button variant="ghost" size="sm" render={<Link to="/tracks" />} className="gap-1 text-muted-foreground">
              View all <ArrowRight className="h-4 w-4" />
            </Button>
          </div>

          {isLoading ? (
            <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
              {Array.from({ length: 6 }, (_, i) => (
                <Skeleton key={i} className="h-48 rounded-xl" />
              ))}
            </div>
          ) : (
            <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
              {tracks?.map((track, i) => (
                <TrackCard
                  key={track.id}
                  track={track}
                  progress={trackProgress?.find((p) => p.trackSlug === track.slug)}
                  index={i}
                />
              ))}
            </div>
          )}
        </div>
      </section>

      <Footer />
    </div>
  );
}
