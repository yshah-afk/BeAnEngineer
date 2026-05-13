import { useState } from "react";
import { motion } from "framer-motion";
import { Layers, Clock, BarChart3 } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";
import { FlashcardDeck } from "@/components/flashcard/FlashcardDeck";
import { useFlashcardDecks, useFlashcardReview } from "@/hooks/useFlashcards";
import { cn } from "@/lib/utils";

export default function FlashcardsPage() {
  const [selectedModule, setSelectedModule] = useState<string | undefined>();
  const { data: decks, isLoading: decksLoading } = useFlashcardDecks();
  const review = useFlashcardReview(selectedModule);

  if (decksLoading) {
    return (
      <div className="mx-auto max-w-3xl px-6 py-10 space-y-4">
        <Skeleton className="h-8 w-48" />
        <div className="grid gap-3 sm:grid-cols-2">
          {Array.from({ length: 4 }, (_, i) => (
            <Skeleton key={i} className="h-24 rounded-xl" />
          ))}
        </div>
      </div>
    );
  }

  return (
    <div className="mx-auto max-w-3xl px-6 py-10">
      <motion.div
        initial={{ opacity: 0, y: 10 }}
        animate={{ opacity: 1, y: 0 }}
        className="mb-8"
      >
        <h1 className="text-3xl font-bold text-foreground">Flashcards</h1>
        <p className="mt-2 text-muted-foreground">
          Review your knowledge with spaced repetition flashcards
        </p>
      </motion.div>

      {!selectedModule ? (
        <div>
          <h2 className="text-lg font-semibold text-foreground mb-4">Select a Deck</h2>
          {!decks || decks.length === 0 ? (
            <Card className="border-border/50">
              <CardContent className="py-12 text-center">
                <Layers className="h-10 w-10 text-muted-foreground/30 mx-auto mb-4" />
                <p className="text-muted-foreground">
                  No flashcard decks available yet. Complete some lessons to unlock flashcards.
                </p>
              </CardContent>
            </Card>
          ) : (
            <div className="grid gap-3 sm:grid-cols-2">
              {decks.map((deck, i) => (
                <motion.div
                  key={deck.moduleId}
                  initial={{ opacity: 0, y: 10 }}
                  animate={{ opacity: 1, y: 0 }}
                  transition={{ delay: i * 0.05 }}
                >
                  <Card
                    className={cn(
                      "cursor-pointer border-border/50 hover:border-primary/30 transition-all hover:-translate-y-0.5",
                      deck.dueCards === 0 && "opacity-60",
                    )}
                    onClick={() => deck.dueCards > 0 && setSelectedModule(deck.moduleId)}
                  >
                    <CardContent className="p-4">
                      <h3 className="font-medium text-foreground text-sm mb-2">
                        {deck.moduleTitle}
                      </h3>
                      <div className="flex items-center gap-3 text-xs text-muted-foreground">
                        <span className="flex items-center gap-1">
                          <Layers className="h-3 w-3" />
                          {deck.totalCards} cards
                        </span>
                        <span
                          className={cn(
                            "flex items-center gap-1 font-medium",
                            deck.dueCards > 0 ? "text-primary" : "text-muted-foreground",
                          )}
                        >
                          <Clock className="h-3 w-3" />
                          {deck.dueCards} due
                        </span>
                        <span className="flex items-center gap-1">
                          <BarChart3 className="h-3 w-3" />
                          {deck.reviewedToday} today
                        </span>
                      </div>
                    </CardContent>
                  </Card>
                </motion.div>
              ))}
            </div>
          )}
        </div>
      ) : (
        <div>
          <div className="flex items-center gap-3 mb-6">
            <Button
              variant="ghost"
              size="sm"
              onClick={() => setSelectedModule(undefined)}
            >
              ← Back to decks
            </Button>
          </div>

          {review.isLoading ? (
            <div className="space-y-4">
              <Skeleton className="h-4 w-32" />
              <Skeleton className="h-64 rounded-2xl" />
            </div>
          ) : review.totalCards === 0 ? (
            <Card className="border-border/50">
              <CardContent className="py-12 text-center">
                <p className="text-muted-foreground">No cards due for review in this deck.</p>
                <Button
                  variant="outline"
                  size="sm"
                  onClick={() => setSelectedModule(undefined)}
                  className="mt-4"
                >
                  Choose another deck
                </Button>
              </CardContent>
            </Card>
          ) : review.currentCard ? (
            <FlashcardDeck
              front={review.currentCard.front}
              back={review.currentCard.back}
              isFlipped={review.isFlipped}
              currentIndex={review.currentIndex}
              totalCards={review.totalCards}
              reviewedCount={review.reviewedCount}
              isComplete={review.isComplete}
              isRating={review.isRating}
              onFlip={review.flipCard}
              onRate={review.rateCard}
              onReset={review.resetDeck}
            />
          ) : null}
        </div>
      )}
    </div>
  );
}
