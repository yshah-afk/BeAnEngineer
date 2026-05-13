export const API_URL = import.meta.env.VITE_API_URL ?? "/api";

export const TRACK_ICONS: Record<string, string> = {
  "llm-ai": "Brain",
  golang: "Code2",
  react: "Layout",
  devops: "Server",
  dsa: "Network",
  "system-design": "Network",
};

export const TRACK_COLORS: Record<string, string> = {
  "llm-ai": "from-violet-500 to-purple-600",
  golang: "from-cyan-500 to-blue-600",
  react: "from-orange-400 to-rose-500",
  devops: "from-emerald-500 to-teal-600",
  dsa: "from-amber-400 to-orange-500",
  "system-design": "from-pink-500 to-fuchsia-600",
};

export const TRACK_ACCENT_BG: Record<string, string> = {
  "llm-ai": "bg-violet-500/10",
  golang: "bg-cyan-500/10",
  react: "bg-orange-400/10",
  devops: "bg-emerald-500/10",
  dsa: "bg-amber-400/10",
  "system-design": "bg-pink-500/10",
};

export const TRACK_ACCENT_TEXT: Record<string, string> = {
  "llm-ai": "text-violet-400",
  golang: "text-cyan-400",
  react: "text-orange-400",
  devops: "text-emerald-400",
  dsa: "text-amber-400",
  "system-design": "text-pink-400",
};

export const DIFFICULTY_COLORS: Record<string, string> = {
  Beginner: "bg-emerald-500/15 text-emerald-400 border-emerald-500/20",
  Intermediate: "bg-amber-500/15 text-amber-400 border-amber-500/20",
  Advanced: "bg-rose-500/15 text-rose-400 border-rose-500/20",
};

export const PLAYGROUND_LANGUAGES = [
  { value: "go", label: "Go", monacoId: "go" },
  { value: "python", label: "Python", monacoId: "python" },
  { value: "javascript", label: "JavaScript", monacoId: "javascript" },
  { value: "typescript", label: "TypeScript", monacoId: "typescript" },
];
