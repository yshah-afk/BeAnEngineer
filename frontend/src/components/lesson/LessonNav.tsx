import { Link } from "react-router-dom";
import { ChevronLeft, ChevronRight } from "lucide-react";
import { Button } from "@/components/ui/button";

interface NavTarget {
  slug: string;
  title: string;
}

interface LessonNavProps {
  trackSlug: string;
  prev?: NavTarget | null;
  next?: NavTarget | null;
}

export function LessonNav({ trackSlug, prev, next }: LessonNavProps) {
  return (
    <div className="flex items-center justify-between gap-4 pt-8 border-t border-border/30">
      {prev ? (
        <Button variant="outline" size="sm" render={<Link to={`/tracks/${trackSlug}/lessons/${prev.slug}`} />} className="gap-2">
          <ChevronLeft className="h-4 w-4" />
          <span className="hidden sm:inline max-w-[200px] truncate">{prev.title}</span>
          <span className="sm:hidden">Previous</span>
        </Button>
      ) : (
        <div />
      )}

      {next ? (
        <Button variant="outline" size="sm" render={<Link to={`/tracks/${trackSlug}/lessons/${next.slug}`} />} className="gap-2">
          <span className="hidden sm:inline max-w-[200px] truncate">{next.title}</span>
          <span className="sm:hidden">Next</span>
          <ChevronRight className="h-4 w-4" />
        </Button>
      ) : (
        <div />
      )}
    </div>
  );
}
