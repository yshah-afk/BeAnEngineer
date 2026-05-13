import { motion } from "framer-motion";
import { RotateCcw, Trophy, Star } from "lucide-react";
import { Button } from "@/components/ui/button";
import { FlashcardCard } from "./FlashcardCard";
import { cn } from "@/lib/utils";

interface FlashcardDeckProps {
  front: string;
  back: string;
  isFlipped: boolean;
  currentIndex: number;
  totalCards: number;
  reviewedCount: number;
  isComplete: boolean;
  isRating: boolean;
  onFlip: () => void;
  onRate: (quality: number) => void;
  onReset: () => void;
}

const QUALITY_LABELS = [
  { value: 0, label: "Again", color: "bg-red-500/15 text-red-400 hover:bg-red-500/25" },
  { value: 1, label: "Hard", color: "bg-orange-500/15 text-orange-400 hover:bg-orange-500/25" },
  { value: 2, label: "Okay", color: "bg-yellow-500/15 text-yellow-400 hover:bg-yellow-500/25" },
  { value: 3, label: "Good", color: "bg-emerald-500/15 text-emerald-400 hover:bg-emerald-500/25" },
  { value: 4, label: "Easy", color: "bg-cyan-500/15 text-cyan-400 hover:bg-cyan-500/25" },
  { value: 5, label: "Perfect", color: "bg-violet-500/15 text-violet-400 hover:bg-violet-500/25" },
];

export function FlashcardDeck({
  front,
  back,
  isFlipped,
  currentIndex,
  totalCards,
  reviewedCount,
  isComplete,
  isRating,
  onFlip,
  onRate,
  onReset,
}: FlashcardDeckProps) {
  if (isComplete) {
    return (
      <motion.div
        initial={{ opacity: 0, scale: 0.9 }}
        animate={{ opacity: 1, scale: 1 }}
        className="flex flex-col items-center justify-center py-16 text-center"
      >
        <Trophy className="h-20 w-20 text-yellow-400 mb-6" />
        <h2 className="text-2xl font-bold mb-2">Deck Complete!</h2>
        <p className="text-muted-foreground mb-1">
          You reviewed{" "}
          <span className="font-bold text-foreground">{reviewedCount}</span> cards
        </p>
        <div className="flex items-center gap-1 text-sm text-muted-foreground mb-8">
          <Star className="h-4 w-4 text-yellow-400" />
          Great job keeping up with your reviews!
        </div>
        <Button onClick={onReset} variant="outline" className="gap-2">
          <RotateCcw className="h-4 w-4" />
          Review Again
        </Button>
      </motion.div>
    );
  }

  return (
    <div className="space-y-8">
      <div className="flex items-center justify-between">
        <span className="text-sm text-muted-foreground">
          Card{" "}
          <span className="font-medium text-foreground tabular-nums">
            {currentIndex + 1}
          </span>{" "}
          of <span className="tabular-nums">{totalCards}</span>
        </span>
        <span className="text-sm text-muted-foreground">
          Reviewed: <span className="tabular-nums font-medium text-foreground">{reviewedCount}</span>
        </span>
      </div>

      <div className="flex gap-1">
        {Array.from({ length: totalCards }, (_, i) => (
          <div
            key={i}
            className={cn(
              "h-1 flex-1 rounded-full transition-colors",
              i < currentIndex ? "bg-primary" : i === currentIndex ? "bg-primary/60" : "bg-muted",
            )}
          />
        ))}
      </div>

      <FlashcardCard front={front} back={back} isFlipped={isFlipped} onFlip={onFlip} />

      {isFlipped && (
        <motion.div
          initial={{ opacity: 0, y: 10 }}
          animate={{ opacity: 1, y: 0 }}
          className="space-y-3"
        >
          <p className="text-center text-sm text-muted-foreground">
            How well did you know this?
          </p>
          <div className="flex flex-wrap justify-center gap-2">
            {QUALITY_LABELS.map(({ value, label, color }) => (
              <Button
                key={value}
                variant="ghost"
                size="sm"
                disabled={isRating}
                onClick={() => onRate(value)}
                className={cn("min-w-[72px]", color)}
              >
                {label}
              </Button>
            ))}
          </div>
        </motion.div>
      )}
    </div>
  );
}
