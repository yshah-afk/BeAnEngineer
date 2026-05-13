import axios, { type AxiosError, type InternalAxiosRequestConfig } from "axios";
import type {
  AuthResponse,
  Bookmark,
  DashboardStats,
  Flashcard,
  FlashcardDeckStats,
  FlashcardProgress,
  Lesson,
  LoginPayload,
  Note,
  PaginatedResponse,
  RegisterPayload,
  SearchResult,
  Track,
  TrackProgress,
} from "@/types";
import { API_URL } from "./constants";

const api = axios.create({
  baseURL: API_URL,
  headers: { "Content-Type": "application/json" },
});

let isRefreshing = false;
let failedQueue: Array<{
  resolve: (token: string) => void;
  reject: (err: unknown) => void;
}> = [];

function processQueue(error: unknown, token: string | null) {
  failedQueue.forEach((prom) => {
    if (token) prom.resolve(token);
    else prom.reject(error);
  });
  failedQueue = [];
}

api.interceptors.request.use((config: InternalAxiosRequestConfig) => {
  const token = localStorage.getItem("access_token");
  if (token && config.headers) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

api.interceptors.response.use(
  (response) => response,
  async (error: AxiosError) => {
    const original = error.config;
    if (!original || error.response?.status !== 401) {
      return Promise.reject(error);
    }

    if (isRefreshing) {
      return new Promise((resolve, reject) => {
        failedQueue.push({
          resolve: (token: string) => {
            if (original.headers) original.headers.Authorization = `Bearer ${token}`;
            resolve(api(original));
          },
          reject,
        });
      });
    }

    isRefreshing = true;
    try {
      const refreshToken = localStorage.getItem("refresh_token");
      if (!refreshToken) throw new Error("No refresh token");

      const { data } = await axios.post<AuthResponse>(`${API_URL}/auth/refresh`, {
        refreshToken,
      });

      localStorage.setItem("access_token", data.accessToken);
      processQueue(null, data.accessToken);

      if (original.headers) {
        original.headers.Authorization = `Bearer ${data.accessToken}`;
      }
      return api(original);
    } catch (refreshError) {
      processQueue(refreshError, null);
      localStorage.removeItem("access_token");
      localStorage.removeItem("refresh_token");
      window.location.href = "/login";
      return Promise.reject(refreshError);
    } finally {
      isRefreshing = false;
    }
  },
);

export const authApi = {
  login: (payload: LoginPayload) =>
    api.post<AuthResponse>("/auth/login", payload).then((r) => r.data),
  register: (payload: RegisterPayload) =>
    api.post<AuthResponse>("/auth/register", payload).then((r) => r.data),
  oauthCallback: (provider: string, code: string) =>
    api.post<AuthResponse>(`/auth/${provider}/callback`, { code }).then((r) => r.data),
  refresh: (refreshToken: string) =>
    api.post<AuthResponse>("/auth/refresh", { refreshToken }).then((r) => r.data),
  me: () => api.get<import("@/types").User>("/auth/me").then((r) => r.data),
};

export const tracksApi = {
  list: () => api.get<Track[]>("/tracks").then((r) => r.data),
  get: (slug: string) => api.get<Track>(`/tracks/${slug}`).then((r) => r.data),
};

export const lessonsApi = {
  get: (slug: string) => api.get<Lesson>(`/lessons/${slug}`).then((r) => r.data),
};

export const progressApi = {
  dashboard: () => api.get<DashboardStats>("/progress/dashboard").then((r) => r.data),
  trackProgress: () => api.get<TrackProgress[]>("/progress/tracks").then((r) => r.data),
  complete: (lessonId: string, timeSpent: number, quizScore?: number) =>
    api.post("/progress/complete", { lessonId, timeSpent, quizScore }).then((r) => r.data),
};

export const bookmarksApi = {
  list: () => api.get<Bookmark[]>("/bookmarks").then((r) => r.data),
  toggle: (lessonId: string) =>
    api.post<{ bookmarked: boolean }>("/bookmarks/toggle", { lessonId }).then((r) => r.data),
  check: (lessonId: string) =>
    api.get<{ bookmarked: boolean }>(`/bookmarks/check/${lessonId}`).then((r) => r.data),
};

export const notesApi = {
  get: (lessonId: string) => api.get<Note>(`/notes/${lessonId}`).then((r) => r.data),
  save: (lessonId: string, content: string) =>
    api.put<Note>(`/notes/${lessonId}`, { content }).then((r) => r.data),
};

export const flashcardsApi = {
  decks: () => api.get<FlashcardDeckStats[]>("/flashcards/decks").then((r) => r.data),
  cards: (moduleId: string) =>
    api.get<Flashcard[]>(`/flashcards/module/${moduleId}`).then((r) => r.data),
  dueCards: (moduleId: string) =>
    api.get<(Flashcard & { progress?: FlashcardProgress })[]>(
      `/flashcards/module/${moduleId}/due`,
    ).then((r) => r.data),
  review: (cardId: string, quality: number) =>
    api.post<FlashcardProgress>("/flashcards/review", { cardId, quality }).then((r) => r.data),
};

export const searchApi = {
  search: (query: string, page = 1) =>
    api
      .get<PaginatedResponse<SearchResult>>("/search", { params: { q: query, page } })
      .then((r) => r.data),
};

export const playgroundApi = {
  run: (language: string, code: string) =>
    api.post<{ output: string; error?: string }>("/playground/run", { language, code }).then((r) => r.data),
};

export const adminApi = {
  lessons: (page = 1) =>
    api.get<PaginatedResponse<Lesson>>("/admin/lessons", { params: { page } }).then((r) => r.data),
  createLesson: (data: Partial<Lesson>) =>
    api.post<Lesson>("/admin/lessons", data).then((r) => r.data),
  updateLesson: (id: string, data: Partial<Lesson>) =>
    api.put<Lesson>(`/admin/lessons/${id}`, data).then((r) => r.data),
  deleteLesson: (id: string) => api.delete(`/admin/lessons/${id}`).then((r) => r.data),
  seed: () => api.post("/admin/seed").then((r) => r.data),
};

export default api;
