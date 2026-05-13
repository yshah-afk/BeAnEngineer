import { GraduationCap, GitBranch } from "lucide-react";

export function Footer() {
  return (
    <footer className="border-t border-border/50 bg-card/30 py-8 mt-auto">
      <div className="mx-auto max-w-6xl px-6">
        <div className="flex flex-col items-center gap-4 sm:flex-row sm:justify-between">
          <div className="flex items-center gap-2 text-muted-foreground">
            <GraduationCap className="h-5 w-5 text-primary" />
            <span className="text-sm font-medium">AI & Full-Stack Mastery Hub</span>
          </div>
          <div className="flex items-center gap-4 text-xs text-muted-foreground">
            <span>170 lessons · 6 tracks · 12 projects</span>
            <a
              href="https://github.com"
              target="_blank"
              rel="noreferrer"
              className="hover:text-foreground transition-colors"
            >
              <GitBranch className="h-4 w-4" />
            </a>
          </div>
        </div>
      </div>
    </footer>
  );
}
