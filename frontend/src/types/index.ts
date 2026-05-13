export interface User {
  id: string;
  email: string;
  name: string;
  avatarUrl?: string;
  provider: "local" | "github" | "google";
  role: "user" | "admin";
  streak: number;
  longestStreak: number;
  lastActiveDate?: string;
  createdAt: string;
}

export interface Track {
  id: string;
  slug: string;
  title: string;
  description: string;
  icon: string;
  color: string;
  moduleCount: number;
  lessonCount: number;
  estimatedHours: number;
  modules: Module[];
}

export interface Module {
  id: string;
  trackId: string;
  slug: string;
  title: string;
  description?: string;
  order: number;
  lessonCount: number;
  lessons: LessonSummary[];
}

export interface LessonSummary {
  id: string;
  moduleId: string;
  slug: string;
  title: string;
  order: number;
  difficulty: "Beginner" | "Intermediate" | "Advanced";
  estimatedMinutes: number;
  isCompleted?: boolean;
}

export interface Lesson extends LessonSummary {
  trackSlug: string;
  trackTitle: string;
  moduleTitle: string;
  content: string;
  quiz?: QuizQuestion[];
  hasPlayground: boolean;
  playgroundLanguage?: string;
  starterCode?: string;
  prevLesson?: { slug: string; title: string } | null;
  nextLesson?: { slug: string; title: string } | null;
}

export interface LessonContent {
  markdown: string;
  quiz?: QuizQuestion[];
}

export interface Progress {
  userId: string;
  lessonId: string;
  trackSlug: string;
  completed: boolean;
  completedAt?: string;
  quizScore?: number;
  timeSpent: number;
}

export interface TrackProgress {
  trackSlug: string;
  totalLessons: number;
  completedLessons: number;
  percentage: number;
  lastAccessedAt?: string;
  lastLessonSlug?: string;
  lastLessonTitle?: string;
}

export interface Bookmark {
  id: string;
  userId: string;
  lessonId: string;
  lessonSlug: string;
  lessonTitle: string;
  trackSlug: string;
  trackTitle: string;
  createdAt: string;
}

export interface Note {
  id: string;
  userId: string;
  lessonId: string;
  content: string;
  updatedAt: string;
}

export interface Flashcard {
  id: string;
  moduleId: string;
  front: string;
  back: string;
}

export interface FlashcardProgress {
  cardId: string;
  userId: string;
  easeFactor: number;
  interval: number;
  repetitions: number;
  nextReviewAt: string;
  lastReviewedAt?: string;
}

export interface FlashcardDeckStats {
  moduleId: string;
  moduleTitle: string;
  trackSlug: string;
  totalCards: number;
  dueCards: number;
  reviewedToday: number;
}

export interface QuizQuestion {
  id: string;
  question: string;
  options: string[];
  correctIndex: number;
  explanation: string;
}

export interface QuizResult {
  lessonId: string;
  score: number;
  totalQuestions: number;
  answers: { questionId: string; selectedIndex: number; correct: boolean }[];
}

export interface SearchResult {
  type: "lesson" | "module" | "track";
  id: string;
  title: string;
  description: string;
  trackSlug: string;
  trackTitle: string;
  slug?: string;
  highlight?: string;
}

export interface ChatMessage {
  id: string;
  role: "user" | "assistant";
  content: string;
  timestamp: Date;
  isStreaming?: boolean;
}

export interface ApiError {
  message: string;
  code?: string;
  status: number;
}

export interface PaginatedResponse<T> {
  data: T[];
  total: number;
  page: number;
  pageSize: number;
  hasMore: boolean;
}

export interface ActivityItem {
  id: string;
  type: "lesson_completed" | "quiz_passed" | "streak_milestone" | "track_started";
  description: string;
  trackSlug?: string;
  lessonSlug?: string;
  createdAt: string;
}

export interface StreakDay {
  date: string;
  count: number;
}

export interface DashboardStats {
  totalCompleted: number;
  totalLessons: number;
  overallPercentage: number;
  currentStreak: number;
  longestStreak: number;
  totalTimeSpent: number;
  trackProgress: TrackProgress[];
  recentActivity: ActivityItem[];
  streakCalendar: StreakDay[];
}

export interface AuthResponse {
  user: User;
  accessToken: string;
  expiresIn: number;
}

export interface LoginPayload {
  email: string;
  password: string;
}

export interface RegisterPayload {
  name: string;
  email: string;
  password: string;
}
