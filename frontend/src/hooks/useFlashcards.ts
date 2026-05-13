import { useState, useCallback } from "react";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { flashcardsApi } from "@/lib/api";
import { useAuthStore } from "@/stores/authStore";
import type { Flashcard, FlashcardProgress } from "@/types";

export function useFlashcardDecks() {
  const isAuthenticated = useAuthStore((s) => s.isAuthenticated);
  return useQuery({
    queryKey: ["flashcards", "decks"],
    queryFn: flashcardsApi.decks,
    enabled: isAuthenticated,
  });
}

export function useFlashcardReview(moduleId: string | undefined) {
  const isAuthenticated = useAuthStore((s) => s.isAuthenticated);
  const queryClient = useQueryClient();

  const [currentIndex, setCurrentIndex] = useState(0);
  const [isFlipped, setIsFlipped] = useState(false);
  const [reviewedCount, setReviewedCount] = useState(0);

  const cardsQuery = useQuery({
    queryKey: ["flashcards", "due", moduleId],
    queryFn: () => flashcardsApi.dueCards(moduleId!),
    enabled: isAuthenticated && !!moduleId,
  });

  const reviewMutation = useMutation({
    mutationFn: ({ cardId, quality }: { cardId: string; quality: number }) =>
      flashcardsApi.review(cardId, quality),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["flashcards"] });
    },
  });

  const currentCard: (Flashcard & { progress?: FlashcardProgress }) | undefined =
    cardsQuery.data?.[currentIndex];

  const totalCards = cardsQuery.data?.length ?? 0;

  const flipCard = useCallback(() => {
    setIsFlipped((prev) => !prev);
  }, []);

  const rateCard = useCallback(
    async (quality: number) => {
      if (!currentCard) return;
      await reviewMutation.mutateAsync({ cardId: currentCard.id, quality });
      setReviewedCount((prev) => prev + 1);
      setIsFlipped(false);
      setCurrentIndex((prev) => Math.min(prev + 1, totalCards - 1));
    },
    [currentCard, reviewMutation, totalCards],
  );

  const resetDeck = useCallback(() => {
    setCurrentIndex(0);
    setIsFlipped(false);
    setReviewedCount(0);
    cardsQuery.refetch();
  }, [cardsQuery]);

  const isComplete = currentIndex >= totalCards - 1 && reviewedCount > 0;

  return {
    currentCard,
    currentIndex,
    totalCards,
    isFlipped,
    reviewedCount,
    isComplete,
    isLoading: cardsQuery.isLoading,
    flipCard,
    rateCard,
    resetDeck,
    isRating: reviewMutation.isPending,
  };
}
