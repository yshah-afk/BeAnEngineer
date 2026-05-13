import { motion } from "framer-motion";
import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";
import { cn } from "@/lib/utils";

interface FlashcardCardProps {
  front: string;
  back: string;
  isFlipped: boolean;
  onFlip: () => void;
}

export function FlashcardCard({ front, back, isFlipped, onFlip }: FlashcardCardProps) {
  return (
    <div
      className="relative w-full max-w-xl mx-auto cursor-pointer perspective-1000"
      style={{ perspective: "1000px" }}
      onClick={onFlip}
    >
      <motion.div
        className="relative w-full min-h-[280px]"
        animate={{ rotateY: isFlipped ? 180 : 0 }}
        transition={{ duration: 0.5, type: "spring", stiffness: 200, damping: 25 }}
        style={{ transformStyle: "preserve-3d" }}
      >
        {/* Front */}
        <div
          className={cn(
            "absolute inset-0 rounded-2xl border border-border/50 bg-card p-8 flex flex-col items-center justify-center text-center shadow-lg",
            "backface-hidden",
          )}
          style={{ backfaceVisibility: "hidden" }}
        >
          <span className="text-xs font-medium text-primary mb-4 uppercase tracking-wider">
            Question
          </span>
          <div className="prose prose-invert prose-sm max-w-none">
            <ReactMarkdown remarkPlugins={[remarkGfm]}>{front}</ReactMarkdown>
          </div>
          <span className="mt-6 text-xs text-muted-foreground">Click to reveal answer</span>
        </div>

        {/* Back */}
        <div
          className={cn(
            "absolute inset-0 rounded-2xl border border-primary/20 bg-card p-8 flex flex-col items-center justify-center text-center shadow-lg",
            "backface-hidden",
          )}
          style={{ backfaceVisibility: "hidden", transform: "rotateY(180deg)" }}
        >
          <span className="text-xs font-medium text-success mb-4 uppercase tracking-wider">
            Answer
          </span>
          <div className="prose prose-invert prose-sm max-w-none">
            <ReactMarkdown remarkPlugins={[remarkGfm]}>{back}</ReactMarkdown>
          </div>
        </div>
      </motion.div>
    </div>
  );
}
