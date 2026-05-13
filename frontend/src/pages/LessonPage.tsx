import { useState, useEffect, useMemo } from "react";
import { useParams, Link } from "react-router-dom";
import { useQuery } from "@tanstack/react-query";
import { motion } from "framer-motion";
import {
  ChevronRight,
  CheckCircle2,
  Bookmark,
  BookmarkCheck,
  StickyNote,
  Clock,
  Loader2,
  ListTree,
} from "lucide-react";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { Textarea } from "@/components/ui/textarea";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { LessonViewer } from "@/components/lesson/LessonViewer";
import { LessonNav } from "@/components/lesson/LessonNav";
import { Quiz } from "@/components/lesson/Quiz";
import { ChatPanel } from "@/components/tutor/ChatPanel";
import { lessonsApi } from "@/lib/api";
import {
  useCompleteLesson,
  useToggleBookmark,
  useBookmarkCheck,
  useNote,
  useSaveNote,
} from "@/hooks/useProgress";
import { useAuthStore } from "@/stores/authStore";
import { DIFFICULTY_COLORS } from "@/lib/constants";
import { formatTime, cn } from "@/lib/utils";
import { extractLessonHeadings, parseLessonFrontmatter } from "@/lib/lesson-content";

function titleFromSlug(value?: string) {
  if (!value) return "";
  return value
    .split("-")
    .map((part) => part.charAt(0).toUpperCase() + part.slice(1))
    .join(" ");
}

export default function LessonPage() {
  const { slug, lessonSlug } = useParams<{ slug: string; lessonSlug: string }>();
  const { isAuthenticated } = useAuthStore();
  const [chatOpen, setChatOpen] = useState(false);
  const [notesOpen, setNotesOpen] = useState(false);
  const [noteContent, setNoteContent] = useState("");
  const [startTime] = useState(Date.now());

  const { data: lesson, isLoading } = useQuery({
    queryKey: ["lesson", lessonSlug],
    queryFn: () => lessonsApi.get(lessonSlug!),
    enabled: !!lessonSlug,
  });

  const { data: bookmarkData } = useBookmarkCheck(lesson?.id);
  const { data: noteData } = useNote(lesson?.id);
  const completeMutation = useCompleteLesson();
  const bookmarkMutation = useToggleBookmark();
  const saveNoteMutation = useSaveNote();
  const lessonContent = lesson?.content ?? "";
  const frontmatter = useMemo(() => parseLessonFrontmatter(lessonContent), [lessonContent]);
  const headings = useMemo(() => extractLessonHeadings(lessonContent), [lessonContent]);

  useEffect(() => {
    if (noteData?.content) setNoteContent(noteData.content);
  }, [noteData]);

  const handleComplete = () => {
    if (!lesson) return;
    const timeSpent = Math.round((Date.now() - startTime) / 60000);
    completeMutation.mutate({ lessonId: lesson.id, timeSpent });
  };

  const handleBookmark = () => {
    if (!lesson) return;
    bookmarkMutation.mutate(lesson.id);
  };

  const handleSaveNote = () => {
    if (!lesson) return;
    saveNoteMutation.mutate({ lessonId: lesson.id, content: noteContent });
  };

  if (isLoading) {
    return (
      <div className="mx-auto max-w-4xl px-6 py-10 space-y-6">
        <Skeleton className="h-8 w-64" />
        <Skeleton className="h-6 w-96" />
        <div className="space-y-4">
          {Array.from({ length: 8 }, (_, i) => (
            <Skeleton key={i} className="h-4 w-full" />
          ))}
        </div>
      </div>
    );
  }

  if (!lesson) {
    return (
      <div className="flex flex-col items-center justify-center py-20 text-center">
        <h2 className="text-2xl font-bold mb-2">Lesson not found</h2>
        <p className="text-muted-foreground mb-4">
          The lesson you're looking for doesn't exist.
        </p>
        <Button render={<Link to="/tracks" />}>Browse Tracks</Button>
      </div>
    );
  }

  const displayTitle = frontmatter.title || lesson.title;
  const displayTrackTitle = lesson.trackTitle || titleFromSlug(frontmatter.track || slug);
  const displayModuleTitle = lesson.moduleTitle || titleFromSlug(frontmatter.module);
  const displayDifficulty = lesson.difficulty || frontmatter.difficulty;
  const displayDuration = lesson.estimatedMinutes || frontmatter.estimatedMinutes || 0;

  return (
    <div className="min-h-[calc(100vh-4rem)] bg-background">
      <div className="mx-auto max-w-7xl px-6 py-8">
        <nav className="mb-6 flex flex-wrap items-center gap-1.5 text-sm text-muted-foreground">
          <Link to={`/tracks/${slug}`} className="transition-colors hover:text-foreground">
            {displayTrackTitle}
          </Link>
          <ChevronRight className="h-3.5 w-3.5 shrink-0" />
          <span className="truncate">{displayModuleTitle}</span>
          <ChevronRight className="h-3.5 w-3.5 shrink-0" />
          <span className="font-medium text-foreground">{displayTitle}</span>
        </nav>

        <div className="grid gap-8 xl:grid-cols-[minmax(0,1fr)_280px]">
          <div className="min-w-0">
            <motion.div initial={{ opacity: 0, y: 10 }} animate={{ opacity: 1, y: 0 }} className="mb-8">
              <Card className="overflow-hidden border-border/60 bg-card/80 shadow-sm">
                <div className="h-1.5 bg-gradient-to-r from-primary/80 via-primary to-chart-2/80" />
                <CardContent className="space-y-6 p-6 sm:p-8">
                  <div className="space-y-4">
                    <div className="flex flex-wrap items-center gap-2">
                      {displayDifficulty && (
                        <Badge
                          variant="outline"
                          className={cn("text-xs", DIFFICULTY_COLORS[displayDifficulty])}
                        >
                          {displayDifficulty}
                        </Badge>
                      )}
                      {displayDuration > 0 && (
                        <span className="inline-flex items-center gap-1.5 rounded-full bg-muted/40 px-3 py-1 text-xs text-muted-foreground">
                          <Clock className="h-3 w-3" />
                          {formatTime(displayDuration)}
                        </span>
                      )}
                    </div>

                    <div className="space-y-2">
                      <h1 className="text-3xl font-bold tracking-tight text-foreground sm:text-4xl">
                        {displayTitle}
                      </h1>
                      <p className="max-w-3xl text-sm leading-6 text-muted-foreground sm:text-base">
                        A focused lesson in the {displayTrackTitle} track. Use the navigation on the
                        right to jump between sections and keep the page easy to scan.
                      </p>
                    </div>
                  </div>

                  {isAuthenticated && (
                    <div className="flex flex-wrap items-center gap-2">
                      <Button
                        variant={lesson.isCompleted ? "secondary" : "default"}
                        size="sm"
                        onClick={handleComplete}
                        disabled={completeMutation.isPending}
                        className="gap-1.5"
                      >
                        {completeMutation.isPending ? (
                          <Loader2 className="h-4 w-4 animate-spin" />
                        ) : (
                          <CheckCircle2 className="h-4 w-4" />
                        )}
                        {lesson.isCompleted ? "Completed" : "Mark Complete"}
                      </Button>
                      <Button
                        variant="outline"
                        size="sm"
                        onClick={handleBookmark}
                        disabled={bookmarkMutation.isPending}
                        className="gap-1.5"
                      >
                        {bookmarkData?.bookmarked ? (
                          <BookmarkCheck className="h-4 w-4 text-primary" />
                        ) : (
                          <Bookmark className="h-4 w-4" />
                        )}
                        {bookmarkData?.bookmarked ? "Bookmarked" : "Bookmark"}
                      </Button>
                      <Button
                        variant="outline"
                        size="sm"
                        onClick={() => setNotesOpen((current) => !current)}
                        className="gap-1.5"
                      >
                        <StickyNote className="h-4 w-4" />
                        {notesOpen ? "Hide Notes" : "Open Notes"}
                      </Button>
                    </div>
                  )}
                </CardContent>
              </Card>
            </motion.div>

            {notesOpen && isAuthenticated && (
              <motion.div
                initial={{ opacity: 0, height: 0 }}
                animate={{ opacity: 1, height: "auto" }}
                className="mb-8"
              >
                <Card className="border-border/50 bg-card/60">
                  <CardHeader className="pb-3">
                    <CardTitle className="flex items-center gap-2 text-sm">
                      <StickyNote className="h-4 w-4" />
                      Personal Notes
                    </CardTitle>
                  </CardHeader>
                  <CardContent className="space-y-3">
                    <Textarea
                      value={noteContent}
                      onChange={(e) => setNoteContent(e.target.value)}
                      placeholder="Capture the mental models, gotchas, and examples you want to remember."
                      rows={5}
                      className="resize-y bg-muted/20"
                    />
                    <Button size="sm" onClick={handleSaveNote} disabled={saveNoteMutation.isPending}>
                      {saveNoteMutation.isPending ? (
                        <Loader2 className="mr-1.5 h-4 w-4 animate-spin" />
                      ) : null}
                      Save Notes
                    </Button>
                  </CardContent>
                </Card>
              </motion.div>
            )}

            <Card className="border-border/50 bg-card/40 shadow-sm">
              <CardContent className="p-6 sm:p-8 lg:p-10">
                <LessonViewer content={lesson.content} />
              </CardContent>
            </Card>

            {lesson.quiz && lesson.quiz.length > 0 && (
              <div className="mt-12">
                <Quiz
                  questions={lesson.quiz}
                  onComplete={(score, total) => {
                    if (lesson && isAuthenticated) {
                      const timeSpent = Math.round((Date.now() - startTime) / 60000);
                      completeMutation.mutate({
                        lessonId: lesson.id,
                        timeSpent,
                        quizScore: Math.round((score / total) * 100),
                      });
                    }
                  }}
                />
              </div>
            )}

            <div className="mt-12 mb-8">
              <LessonNav trackSlug={slug!} prev={lesson.prevLesson} next={lesson.nextLesson} />
            </div>
          </div>

          <aside className="hidden xl:block">
            <div className="sticky top-24 space-y-4">
              <Card className="border-border/50 bg-card/50">
                <CardHeader className="pb-3">
                  <CardTitle className="text-sm">Lesson Overview</CardTitle>
                </CardHeader>
                <CardContent className="space-y-3 text-sm">
                  <div className="rounded-xl bg-muted/25 p-3">
                    <div className="text-xs uppercase tracking-[0.16em] text-muted-foreground">
                      Track
                    </div>
                    <div className="mt-1 font-medium text-foreground">{displayTrackTitle}</div>
                  </div>
                  <div className="rounded-xl bg-muted/25 p-3">
                    <div className="text-xs uppercase tracking-[0.16em] text-muted-foreground">
                      Module
                    </div>
                    <div className="mt-1 font-medium text-foreground">{displayModuleTitle}</div>
                  </div>
                  {displayDuration > 0 && (
                    <div className="rounded-xl bg-muted/25 p-3">
                      <div className="text-xs uppercase tracking-[0.16em] text-muted-foreground">
                        Reading Time
                      </div>
                      <div className="mt-1 font-medium text-foreground">
                        {formatTime(displayDuration)}
                      </div>
                    </div>
                  )}
                </CardContent>
              </Card>

              {headings.length > 0 && (
                <Card className="border-border/50 bg-card/50">
                  <CardHeader className="pb-3">
                    <CardTitle className="flex items-center gap-2 text-sm">
                      <ListTree className="h-4 w-4 text-primary" />
                      On This Page
                    </CardTitle>
                  </CardHeader>
                  <CardContent>
                    <nav className="space-y-1.5">
                      {headings.map((heading) => (
                        <a
                          key={heading.id}
                          href={`#${heading.id}`}
                          className={cn(
                            "block rounded-lg px-3 py-2 text-sm text-muted-foreground transition-colors hover:bg-muted/35 hover:text-foreground",
                            heading.level === 3 && "ml-4 text-[13px]",
                          )}
                        >
                          {heading.text}
                        </a>
                      ))}
                    </nav>
                  </CardContent>
                </Card>
              )}
            </div>
          </aside>
        </div>
      </div>

      {isAuthenticated && (
        <>
          <div className="hidden lg:block">
            {chatOpen ? (
              <div className="fixed inset-y-24 right-6 z-40 w-[380px] overflow-hidden rounded-2xl border border-border/60 bg-card/95 shadow-2xl backdrop-blur">
                <ChatPanel
                  lessonSlug={lessonSlug}
                  isOpen={chatOpen}
                  onToggle={() => setChatOpen(false)}
                />
              </div>
            ) : (
              <ChatPanel
                lessonSlug={lessonSlug}
                isOpen={false}
                onToggle={() => setChatOpen(true)}
              />
            )}
          </div>

          <div className="lg:hidden">
            {chatOpen ? (
              <div className="fixed inset-x-4 bottom-4 top-20 z-40 overflow-hidden rounded-2xl border border-border/60 bg-card/95 shadow-2xl backdrop-blur">
                <ChatPanel
                  lessonSlug={lessonSlug}
                  isOpen={chatOpen}
                  onToggle={() => setChatOpen(false)}
                />
              </div>
            ) : (
              <ChatPanel
                lessonSlug={lessonSlug}
                isOpen={false}
                onToggle={() => setChatOpen(true)}
              />
            )}
          </div>
        </>
      )}
    </div>
  );
}
