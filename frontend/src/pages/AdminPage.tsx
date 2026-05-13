import { useState } from "react";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { motion } from "framer-motion";
import {
  Shield,
  Plus,
  Pencil,
  Trash2,
  Database,
  Loader2,
  ChevronLeft,
  ChevronRight,
} from "lucide-react";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { adminApi } from "@/lib/api";
import { DIFFICULTY_COLORS } from "@/lib/constants";
import { cn } from "@/lib/utils";

export default function AdminPage() {
  const queryClient = useQueryClient();
  const [page, setPage] = useState(1);
  const [editDialog, setEditDialog] = useState(false);
  const [editForm, setEditForm] = useState({
    id: "",
    title: "",
    slug: "",
    content: "",
    difficulty: "Beginner" as "Beginner" | "Intermediate" | "Advanced",
    estimatedMinutes: 30,
  });

  const { data, isLoading } = useQuery({
    queryKey: ["admin", "lessons", page],
    queryFn: () => adminApi.lessons(page),
  });

  const deleteMutation = useMutation({
    mutationFn: (id: string) => adminApi.deleteLesson(id),
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ["admin"] }),
  });

  const seedMutation = useMutation({
    mutationFn: () => adminApi.seed(),
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ["admin"] }),
  });

  const saveMutation = useMutation({
    mutationFn: () =>
      editForm.id
        ? adminApi.updateLesson(editForm.id, editForm)
        : adminApi.createLesson(editForm),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["admin"] });
      setEditDialog(false);
    },
  });

  const handleEdit = (lesson: typeof editForm) => {
    setEditForm(lesson);
    setEditDialog(true);
  };

  const handleNew = () => {
    setEditForm({
      id: "",
      title: "",
      slug: "",
      content: "",
      difficulty: "Beginner" as const,
      estimatedMinutes: 30,
    });
    setEditDialog(true);
  };

  return (
    <div className="mx-auto max-w-5xl px-6 py-10">
      <motion.div
        initial={{ opacity: 0, y: 10 }}
        animate={{ opacity: 1, y: 0 }}
        className="flex items-center justify-between mb-8"
      >
        <div>
          <h1 className="text-3xl font-bold text-foreground flex items-center gap-3">
            <Shield className="h-8 w-8 text-primary" />
            Admin
          </h1>
          <p className="mt-1 text-muted-foreground">Manage lessons and content</p>
        </div>
        <div className="flex items-center gap-2">
          <Button
            variant="outline"
            size="sm"
            onClick={() => seedMutation.mutate()}
            disabled={seedMutation.isPending}
            className="gap-1.5"
          >
            {seedMutation.isPending ? (
              <Loader2 className="h-4 w-4 animate-spin" />
            ) : (
              <Database className="h-4 w-4" />
            )}
            Seed DB
          </Button>
          <Dialog open={editDialog} onOpenChange={setEditDialog}>
            <DialogTrigger render={<Button size="sm" onClick={handleNew} className="gap-1.5" />}>
              <Plus className="h-4 w-4" />
              New Lesson
            </DialogTrigger>
            <DialogContent className="max-w-lg">
              <DialogHeader>
                <DialogTitle>
                  {editForm.id ? "Edit Lesson" : "Create Lesson"}
                </DialogTitle>
              </DialogHeader>
              <form
                onSubmit={(e) => {
                  e.preventDefault();
                  saveMutation.mutate();
                }}
                className="space-y-4"
              >
                <div className="space-y-2">
                  <Label>Title</Label>
                  <Input
                    value={editForm.title}
                    onChange={(e) =>
                      setEditForm((f) => ({ ...f, title: e.target.value }))
                    }
                    placeholder="Lesson title"
                  />
                </div>
                <div className="space-y-2">
                  <Label>Slug</Label>
                  <Input
                    value={editForm.slug}
                    onChange={(e) =>
                      setEditForm((f) => ({ ...f, slug: e.target.value }))
                    }
                    placeholder="lesson-slug"
                  />
                </div>
                <div className="grid grid-cols-2 gap-4">
                  <div className="space-y-2">
                    <Label>Difficulty</Label>
                    <select
                      value={editForm.difficulty}
                      onChange={(e) =>
                        setEditForm((f) => ({ ...f, difficulty: e.target.value as "Beginner" | "Intermediate" | "Advanced" }))
                      }
                      className="w-full rounded-md border border-input bg-background px-3 py-2 text-sm"
                    >
                      <option>Beginner</option>
                      <option>Intermediate</option>
                      <option>Advanced</option>
                    </select>
                  </div>
                  <div className="space-y-2">
                    <Label>Est. Minutes</Label>
                    <Input
                      type="number"
                      value={editForm.estimatedMinutes}
                      onChange={(e) =>
                        setEditForm((f) => ({
                          ...f,
                          estimatedMinutes: Number(e.target.value),
                        }))
                      }
                    />
                  </div>
                </div>
                <div className="space-y-2">
                  <Label>Content (Markdown)</Label>
                  <Textarea
                    value={editForm.content}
                    onChange={(e) =>
                      setEditForm((f) => ({ ...f, content: e.target.value }))
                    }
                    rows={8}
                    className="font-mono text-xs"
                    placeholder="# Lesson content..."
                  />
                </div>
                <div className="flex justify-end gap-2">
                  <Button
                    type="button"
                    variant="ghost"
                    onClick={() => setEditDialog(false)}
                  >
                    Cancel
                  </Button>
                  <Button type="submit" disabled={saveMutation.isPending}>
                    {saveMutation.isPending && (
                      <Loader2 className="h-4 w-4 animate-spin mr-1.5" />
                    )}
                    Save
                  </Button>
                </div>
              </form>
            </DialogContent>
          </Dialog>
        </div>
      </motion.div>

      <Card className="border-border/50">
        <CardHeader>
          <CardTitle className="text-base">
            Lessons{" "}
            {data && (
              <span className="text-muted-foreground font-normal">
                ({data.total} total)
              </span>
            )}
          </CardTitle>
        </CardHeader>
        <CardContent>
          {isLoading ? (
            <div className="space-y-3">
              {Array.from({ length: 8 }, (_, i) => (
                <Skeleton key={i} className="h-12 rounded-lg" />
              ))}
            </div>
          ) : (
            <>
              <div className="divide-y divide-border/30">
                {data?.data.map((lesson) => (
                  <div
                    key={lesson.id}
                    className="flex items-center gap-3 py-3 first:pt-0 last:pb-0"
                  >
                    <div className="min-w-0 flex-1">
                      <p className="text-sm font-medium text-foreground truncate">
                        {lesson.title}
                      </p>
                      <p className="text-xs text-muted-foreground truncate">
                        {lesson.slug}
                      </p>
                    </div>
                    <Badge
                      variant="outline"
                      className={cn(
                        "text-[10px] shrink-0",
                        DIFFICULTY_COLORS[lesson.difficulty],
                      )}
                    >
                      {lesson.difficulty}
                    </Badge>
                    <span className="text-xs text-muted-foreground tabular-nums shrink-0">
                      {lesson.estimatedMinutes}m
                    </span>
                    <div className="flex items-center gap-1 shrink-0">
                      <Button
                        variant="ghost"
                        size="icon"
                        className="h-7 w-7"
                        onClick={() =>
                          handleEdit({
                            id: lesson.id,
                            title: lesson.title,
                            slug: lesson.slug,
                            content: lesson.content,
                            difficulty: lesson.difficulty,
                            estimatedMinutes: lesson.estimatedMinutes,
                          })
                        }
                      >
                        <Pencil className="h-3.5 w-3.5" />
                      </Button>
                      <Button
                        variant="ghost"
                        size="icon"
                        className="h-7 w-7 text-destructive hover:text-destructive"
                        onClick={() => deleteMutation.mutate(lesson.id)}
                        disabled={deleteMutation.isPending}
                      >
                        <Trash2 className="h-3.5 w-3.5" />
                      </Button>
                    </div>
                  </div>
                ))}
              </div>

              {data && data.total > 20 && (
                <div className="flex items-center justify-between mt-4 pt-4 border-t border-border/30">
                  <Button
                    variant="outline"
                    size="sm"
                    disabled={page <= 1}
                    onClick={() => setPage((p) => p - 1)}
                    className="gap-1"
                  >
                    <ChevronLeft className="h-4 w-4" />
                    Prev
                  </Button>
                  <span className="text-xs text-muted-foreground">Page {page}</span>
                  <Button
                    variant="outline"
                    size="sm"
                    disabled={!data.hasMore}
                    onClick={() => setPage((p) => p + 1)}
                    className="gap-1"
                  >
                    Next
                    <ChevronRight className="h-4 w-4" />
                  </Button>
                </div>
              )}
            </>
          )}
        </CardContent>
      </Card>
    </div>
  );
}
