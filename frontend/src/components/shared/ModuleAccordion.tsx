import { Link } from "react-router-dom";
import { CheckCircle2, Circle, Clock, ChevronDown } from "lucide-react";
import { Badge } from "@/components/ui/badge";
import { cn } from "@/lib/utils";
import { DIFFICULTY_COLORS } from "@/lib/constants";
import { formatTime } from "@/lib/utils";
import type { Module } from "@/types";
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion";

interface ModuleAccordionProps {
  modules: Module[];
  trackSlug: string;
}

export function ModuleAccordion({ modules, trackSlug }: ModuleAccordionProps) {
  return (
    <Accordion multiple className="space-y-3">
      {modules.map((mod) => {
        const completed = mod.lessons?.filter((l) => l.isCompleted).length ?? 0;
        const total = mod.lessons?.length ?? 0;
        const allDone = completed === total && total > 0;

        return (
          <AccordionItem
            key={mod.slug}
            value={mod.slug}
            className="rounded-xl border border-border/50 bg-card/60 backdrop-blur-sm overflow-hidden"
          >
            <AccordionTrigger className="px-5 py-4 hover:no-underline hover:bg-accent/30 transition-colors [&[data-state=open]>svg]:rotate-180">
              <div className="flex items-center gap-3 text-left">
                <div
                  className={cn(
                    "flex h-8 w-8 shrink-0 items-center justify-center rounded-lg text-sm font-bold",
                    allDone
                      ? "bg-success/15 text-success"
                      : "bg-muted text-muted-foreground",
                  )}
                >
                  {mod.order}
                </div>
                <div>
                  <h3 className="font-medium text-foreground">{mod.title}</h3>
                  <p className="text-xs text-muted-foreground mt-0.5">
                    {completed}/{total} lessons completed
                  </p>
                </div>
              </div>
              <ChevronDown className="h-4 w-4 text-muted-foreground transition-transform duration-200" />
            </AccordionTrigger>
            <AccordionContent>
              <div className="divide-y divide-border/30">
                {(mod.lessons ?? []).map((lesson) => (
                  <Link
                    key={lesson.slug}
                    to={`/tracks/${trackSlug}/lessons/${lesson.slug}`}
                    className="flex items-center gap-3 px-5 py-3 transition-colors hover:bg-accent/20"
                  >
                    {lesson.isCompleted ? (
                      <CheckCircle2 className="h-4.5 w-4.5 shrink-0 text-success" />
                    ) : (
                      <Circle className="h-4.5 w-4.5 shrink-0 text-muted-foreground/50" />
                    )}
                    <span
                      className={cn(
                        "flex-1 text-sm",
                        lesson.isCompleted
                          ? "text-muted-foreground"
                          : "text-foreground",
                      )}
                    >
                      {lesson.title}
                    </span>
                    <div className="flex items-center gap-2">
                      {lesson.difficulty && (
                        <Badge
                          variant="outline"
                          className={cn(
                            "text-[10px] px-1.5 py-0",
                            DIFFICULTY_COLORS[lesson.difficulty],
                          )}
                        >
                          {lesson.difficulty}
                        </Badge>
                      )}
                      {lesson.estimatedMinutes > 0 && (
                        <span className="flex items-center gap-1 text-xs text-muted-foreground tabular-nums">
                          <Clock className="h-3 w-3" />
                          {formatTime(lesson.estimatedMinutes)}
                        </span>
                      )}
                    </div>
                  </Link>
                ))}
              </div>
            </AccordionContent>
          </AccordionItem>
        );
      })}
    </Accordion>
  );
}
