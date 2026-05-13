import { useQuery } from "@tanstack/react-query";
import { Skeleton } from "@/components/ui/skeleton";
import { TrackCard } from "@/components/shared/TrackCard";
import { tracksApi } from "@/lib/api";
import { useTrackProgress } from "@/hooks/useProgress";

export default function TracksPage() {
  const { data: tracks, isLoading } = useQuery({
    queryKey: ["tracks"],
    queryFn: tracksApi.list,
  });

  const { data: trackProgress } = useTrackProgress();

  return (
    <div className="mx-auto max-w-5xl px-6 py-10">
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-foreground">Learning Tracks</h1>
        <p className="mt-2 text-muted-foreground">
          Choose a track to start your learning journey. Each track has modules with
          structured lessons and hands-on projects.
        </p>
      </div>

      {isLoading ? (
        <div className="grid gap-5 sm:grid-cols-2 lg:grid-cols-3">
          {Array.from({ length: 6 }, (_, i) => (
            <Skeleton key={i} className="h-52 rounded-xl" />
          ))}
        </div>
      ) : (
        <div className="grid gap-5 sm:grid-cols-2 lg:grid-cols-3">
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
  );
}
