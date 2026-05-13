import { lazy, Suspense } from "react";
import { BrowserRouter, Routes, Route, Navigate, Outlet } from "react-router-dom";
import { Loader2 } from "lucide-react";
import { AppLayout } from "@/components/layout/AppLayout";
import { useAuthStore } from "@/stores/authStore";

const HomePage = lazy(() => import("@/pages/HomePage"));
const TracksPage = lazy(() => import("@/pages/TracksPage"));
const TrackDetailPage = lazy(() => import("@/pages/TrackDetailPage"));
const LessonPage = lazy(() => import("@/pages/LessonPage"));
const PlaygroundPage = lazy(() => import("@/pages/PlaygroundPage"));
const FlashcardsPage = lazy(() => import("@/pages/FlashcardsPage"));
const DashboardPage = lazy(() => import("@/pages/DashboardPage"));
const BookmarksPage = lazy(() => import("@/pages/BookmarksPage"));
const SearchPage = lazy(() => import("@/pages/SearchPage"));
const LoginPage = lazy(() => import("@/pages/LoginPage"));
const RegisterPage = lazy(() => import("@/pages/RegisterPage"));
const AuthCallbackPage = lazy(() => import("@/pages/AuthCallbackPage"));
const AdminPage = lazy(() => import("@/pages/AdminPage"));
const NotFoundPage = lazy(() => import("@/pages/NotFoundPage"));

function PageLoader() {
  return (
    <div className="flex h-[50vh] items-center justify-center">
      <Loader2 className="h-8 w-8 animate-spin text-primary" />
    </div>
  );
}

function ProtectedRoute() {
  const isAuthenticated = useAuthStore((s) => s.isAuthenticated);
  if (!isAuthenticated) return <Navigate to="/login" replace />;
  return <Outlet />;
}

export default function App() {
  return (
    <BrowserRouter>
      <Suspense fallback={<PageLoader />}>
        <Routes>
          <Route element={<AppLayout />}>
            {/* Public routes */}
            <Route path="/" element={<HomePage />} />
            <Route path="/tracks" element={<TracksPage />} />
            <Route path="/tracks/:slug" element={<TrackDetailPage />} />
            <Route path="/tracks/:slug/lessons/:lessonSlug" element={<LessonPage />} />
            <Route path="/login" element={<LoginPage />} />
            <Route path="/register" element={<RegisterPage />} />
            <Route path="/auth/callback/:provider" element={<AuthCallbackPage />} />
            <Route path="/search" element={<SearchPage />} />

            {/* Protected routes */}
            <Route element={<ProtectedRoute />}>
              <Route path="/dashboard" element={<DashboardPage />} />
              <Route path="/bookmarks" element={<BookmarksPage />} />
              <Route path="/playground" element={<PlaygroundPage />} />
              <Route path="/flashcards" element={<FlashcardsPage />} />
              <Route path="/admin" element={<AdminPage />} />
            </Route>

            {/* 404 */}
            <Route path="*" element={<NotFoundPage />} />
          </Route>
        </Routes>
      </Suspense>
    </BrowserRouter>
  );
}
