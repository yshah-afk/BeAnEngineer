import { useState } from "react";
import { motion, AnimatePresence } from "framer-motion";
import { CheckCircle2, XCircle, RotateCcw, Trophy } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { cn } from "@/lib/utils";
import type { QuizQuestion } from "@/types";

interface QuizProps {
  questions: QuizQuestion[];
  onComplete?: (score: number, total: number) => void;
}

export function Quiz({ questions, onComplete }: QuizProps) {
  const [currentIndex, setCurrentIndex] = useState(0);
  const [selectedIndex, setSelectedIndex] = useState<number | null>(null);
  const [isRevealed, setIsRevealed] = useState(false);
  const [score, setScore] = useState(0);
  const [isComplete, setIsComplete] = useState(false);
  const [answers, setAnswers] = useState<boolean[]>([]);

  const question = questions[currentIndex];

  const handleSelect = (index: number) => {
    if (isRevealed) return;
    setSelectedIndex(index);
  };

  const handleCheck = () => {
    if (selectedIndex === null) return;
    setIsRevealed(true);
    const correct = selectedIndex === question.correctIndex;
    if (correct) setScore((s) => s + 1);
    setAnswers((a) => [...a, correct]);
  };

  const handleNext = () => {
    if (currentIndex < questions.length - 1) {
      setCurrentIndex((i) => i + 1);
      setSelectedIndex(null);
      setIsRevealed(false);
    } else {
      setIsComplete(true);
      const finalScore = score + (selectedIndex === question.correctIndex ? 0 : 0);
      onComplete?.(finalScore, questions.length);
    }
  };

  const handleReset = () => {
    setCurrentIndex(0);
    setSelectedIndex(null);
    setIsRevealed(false);
    setScore(0);
    setIsComplete(false);
    setAnswers([]);
  };

  if (isComplete) {
    const percentage = Math.round((score / questions.length) * 100);
    return (
      <Card className="border-border/50 bg-card/80">
        <CardContent className="py-10 text-center">
          <motion.div
            initial={{ scale: 0 }}
            animate={{ scale: 1 }}
            transition={{ type: "spring", duration: 0.5 }}
          >
            <Trophy
              className={cn(
                "mx-auto h-16 w-16 mb-4",
                percentage >= 70 ? "text-yellow-400" : "text-muted-foreground",
              )}
            />
          </motion.div>
          <h3 className="text-2xl font-bold mb-2">Quiz Complete!</h3>
          <p className="text-lg text-muted-foreground mb-1">
            You scored{" "}
            <span className="font-bold text-foreground">
              {score}/{questions.length}
            </span>
          </p>
          <p className="text-sm text-muted-foreground mb-6">{percentage}% correct</p>
          <Button onClick={handleReset} variant="outline" className="gap-2">
            <RotateCcw className="h-4 w-4" />
            Try Again
          </Button>
        </CardContent>
      </Card>
    );
  }

  return (
    <Card className="border-border/50 bg-card/80">
      <CardHeader className="pb-4">
        <div className="flex items-center justify-between">
          <CardTitle className="text-lg">Quiz</CardTitle>
          <span className="text-sm text-muted-foreground tabular-nums">
            {currentIndex + 1} / {questions.length}
          </span>
        </div>
        <div className="flex gap-1 mt-2">
          {questions.map((_, i) => (
            <div
              key={i}
              className={cn(
                "h-1 flex-1 rounded-full transition-colors",
                i < currentIndex
                  ? answers[i]
                    ? "bg-success"
                    : "bg-destructive"
                  : i === currentIndex
                    ? "bg-primary"
                    : "bg-muted",
              )}
            />
          ))}
        </div>
      </CardHeader>
      <CardContent className="space-y-4">
        <AnimatePresence mode="wait">
          <motion.div
            key={currentIndex}
            initial={{ opacity: 0, x: 20 }}
            animate={{ opacity: 1, x: 0 }}
            exit={{ opacity: 0, x: -20 }}
          >
            <p className="text-base font-medium mb-4">{question.question}</p>
            <div className="space-y-2">
              {question.options.map((option, i) => (
                <button
                  key={i}
                  onClick={() => handleSelect(i)}
                  disabled={isRevealed}
                  className={cn(
                    "w-full text-left rounded-lg border p-3 text-sm transition-all",
                    selectedIndex === i && !isRevealed && "border-primary bg-primary/10",
                    isRevealed && i === question.correctIndex && "border-success bg-success/10",
                    isRevealed &&
                      selectedIndex === i &&
                      i !== question.correctIndex &&
                      "border-destructive bg-destructive/10",
                    !isRevealed &&
                      selectedIndex !== i &&
                      "border-border/50 hover:border-border hover:bg-accent/30",
                    isRevealed && i !== question.correctIndex && selectedIndex !== i && "opacity-50",
                  )}
                >
                  <div className="flex items-center gap-3">
                    <span
                      className={cn(
                        "flex h-6 w-6 shrink-0 items-center justify-center rounded-full text-xs font-medium",
                        selectedIndex === i && !isRevealed && "bg-primary text-primary-foreground",
                        isRevealed && i === question.correctIndex && "bg-success text-white",
                        isRevealed && selectedIndex === i && i !== question.correctIndex && "bg-destructive text-white",
                        !(selectedIndex === i || (isRevealed && i === question.correctIndex)) &&
                          "bg-muted text-muted-foreground",
                      )}
                    >
                      {isRevealed && i === question.correctIndex ? (
                        <CheckCircle2 className="h-4 w-4" />
                      ) : isRevealed && selectedIndex === i ? (
                        <XCircle className="h-4 w-4" />
                      ) : (
                        String.fromCharCode(65 + i)
                      )}
                    </span>
                    <span>{option}</span>
                  </div>
                </button>
              ))}
            </div>

            {isRevealed && question.explanation && (
              <motion.div
                initial={{ opacity: 0, height: 0 }}
                animate={{ opacity: 1, height: "auto" }}
                className="mt-4 rounded-lg bg-primary/5 border border-primary/10 p-3"
              >
                <p className="text-sm text-muted-foreground">
                  <strong className="text-foreground">Explanation:</strong>{" "}
                  {question.explanation}
                </p>
              </motion.div>
            )}
          </motion.div>
        </AnimatePresence>

        <div className="flex justify-end gap-2 pt-2">
          {!isRevealed ? (
            <Button onClick={handleCheck} disabled={selectedIndex === null}>
              Check Answer
            </Button>
          ) : (
            <Button onClick={handleNext}>
              {currentIndex < questions.length - 1 ? "Next Question" : "See Results"}
            </Button>
          )}
        </div>
      </CardContent>
    </Card>
  );
}
