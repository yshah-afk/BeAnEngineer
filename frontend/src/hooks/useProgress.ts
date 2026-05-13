import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { progressApi, bookmarksApi, notesApi } from "@/lib/api";
import { useAuthStore } from "@/stores/authStore";

export function useDashboard() {
  const isAuthenticated = useAuthStore((s) => s.isAuthenticated);
  return useQuery({
    queryKey: ["progress", "dashboard"],
    queryFn: progressApi.dashboard,
    enabled: isAuthenticated,
  });
}

export function useTrackProgress() {
  const isAuthenticated = useAuthStore((s) => s.isAuthenticated);
  return useQuery({
    queryKey: ["progress", "tracks"],
    queryFn: progressApi.trackProgress,
    enabled: isAuthenticated,
  });
}

export function useCompleteLesson() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: ({
      lessonId,
      timeSpent,
      quizScore,
    }: {
      lessonId: string;
      timeSpent: number;
      quizScore?: number;
    }) => progressApi.complete(lessonId, timeSpent, quizScore),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["progress"] });
      queryClient.invalidateQueries({ queryKey: ["tracks"] });
      queryClient.invalidateQueries({ queryKey: ["lesson"] });
    },
  });
}

export function useBookmarks() {
  const isAuthenticated = useAuthStore((s) => s.isAuthenticated);
  return useQuery({
    queryKey: ["bookmarks"],
    queryFn: bookmarksApi.list,
    enabled: isAuthenticated,
  });
}

export function useToggleBookmark() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (lessonId: string) => bookmarksApi.toggle(lessonId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["bookmarks"] });
    },
  });
}

export function useBookmarkCheck(lessonId: string | undefined) {
  const isAuthenticated = useAuthStore((s) => s.isAuthenticated);
  return useQuery({
    queryKey: ["bookmarks", "check", lessonId],
    queryFn: () => bookmarksApi.check(lessonId!),
    enabled: isAuthenticated && !!lessonId,
  });
}

export function useNote(lessonId: string | undefined) {
  const isAuthenticated = useAuthStore((s) => s.isAuthenticated);
  return useQuery({
    queryKey: ["notes", lessonId],
    queryFn: () => notesApi.get(lessonId!),
    enabled: isAuthenticated && !!lessonId,
  });
}

export function useSaveNote() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: ({ lessonId, content }: { lessonId: string; content: string }) =>
      notesApi.save(lessonId, content),
    onSuccess: (_data, variables) => {
      queryClient.invalidateQueries({ queryKey: ["notes", variables.lessonId] });
    },
  });
}
