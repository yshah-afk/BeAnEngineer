import { Link } from "react-router-dom";
import { motion } from "framer-motion";
import { BookmarkCheck, ArrowRight, BookOpen } from "lucide-react";
import { Card, CardContent } from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";
import { useBookmarks } from "@/hooks/useProgress";
import { formatDate } from "@/lib/utils";

export default function BookmarksPage() {
  const { data: bookmarks, isLoading } = useBookmarks();

  return (
    <div className="mx-auto max-w-3xl px-6 py-10">
      <motion.div
        initial={{ opacity: 0, y: 10 }}
        animate={{ opacity: 1, y: 0 }}
        className="mb-8"
      >
        <h1 className="text-3xl font-bold text-foreground flex items-center gap-3">
          <BookmarkCheck className="h-8 w-8 text-primary" />
          Bookmarks
        </h1>
        <p className="mt-2 text-muted-foreground">
          Lessons you've saved for later review
        </p>
      </motion.div>

      {isLoading ? (
        <div className="space-y-3">
          {Array.from({ length: 5 }, (_, i) => (
            <Skeleton key={i} className="h-20 rounded-xl" />
          ))}
        </div>
      ) : !bookmarks || bookmarks.length === 0 ? (
        <Card className="border-border/50">
          <CardContent className="py-16 text-center">
            <BookOpen className="h-12 w-12 mx-auto text-muted-foreground/30 mb-4" />
            <h3 className="font-medium text-foreground mb-1">No bookmarks yet</h3>
            <p className="text-sm text-muted-foreground">
              Bookmark lessons while studying to find them here later.
            </p>
          </CardContent>
        </Card>
      ) : (
        <div className="space-y-3">
          {bookmarks.map((bm, i) => (
            <motion.div
              key={bm.id}
              initial={{ opacity: 0, y: 10 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ delay: i * 0.03 }}
            >
              <Link to={`/tracks/${bm.trackSlug}/lessons/${bm.lessonSlug}`}>
                <Card className="border-border/50 hover:border-primary/30 transition-all hover:-translate-y-0.5 group">
                  <CardContent className="p-4 flex items-center gap-4">
                    <BookmarkCheck className="h-5 w-5 text-primary shrink-0" />
                    <div className="min-w-0 flex-1">
                      <p className="font-medium text-foreground text-sm group-hover:text-primary transition-colors truncate">
                        {bm.lessonTitle}
                      </p>
                      <p className="text-xs text-muted-foreground mt-0.5">
                        {bm.trackTitle} · Saved {formatDate(bm.createdAt)}
                      </p>
                    </div>
                    <ArrowRight className="h-4 w-4 text-muted-foreground group-hover:text-primary transition-colors shrink-0" />
                  </CardContent>
                </Card>
              </Link>
            </motion.div>
          ))}
        </div>
      )}
    </div>
  );
}
