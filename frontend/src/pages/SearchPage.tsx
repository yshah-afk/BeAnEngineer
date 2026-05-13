import { useSearchParams, Link } from "react-router-dom";
import { useQuery } from "@tanstack/react-query";
import { motion } from "framer-motion";
import { Search, ArrowRight, Loader2 } from "lucide-react";
import { Card, CardContent } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { SearchBar } from "@/components/shared/SearchBar";
import { searchApi } from "@/lib/api";

export default function SearchPage() {
  const [searchParams] = useSearchParams();
  const query = searchParams.get("q") ?? "";

  const { data, isLoading } = useQuery({
    queryKey: ["search", "page", query],
    queryFn: () => searchApi.search(query),
    enabled: query.length >= 2,
  });

  return (
    <div className="mx-auto max-w-3xl px-6 py-10">
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-foreground mb-4">Search</h1>
        <SearchBar className="max-w-full" autoFocus />
      </div>

      {query && (
        <p className="text-sm text-muted-foreground mb-6">
          {isLoading
            ? "Searching..."
            : data
              ? `${data.total} result${data.total !== 1 ? "s" : ""} for "${query}"`
              : `No results for "${query}"`}
        </p>
      )}

      {isLoading ? (
        <div className="flex justify-center py-12">
          <Loader2 className="h-8 w-8 animate-spin text-muted-foreground" />
        </div>
      ) : data?.data && data.data.length > 0 ? (
        <div className="space-y-3">
          {data.data.map((result, i) => (
            <motion.div
              key={`${result.type}-${result.id}`}
              initial={{ opacity: 0, y: 10 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ delay: i * 0.03 }}
            >
              <Link
                to={
                  result.type === "track"
                    ? `/tracks/${result.trackSlug}`
                    : result.slug
                      ? `/tracks/${result.trackSlug}/lessons/${result.slug}`
                      : `/tracks/${result.trackSlug}`
                }
              >
                <Card className="border-border/50 hover:border-primary/30 transition-all hover:-translate-y-0.5 group">
                  <CardContent className="p-4 flex items-start gap-4">
                    <Badge variant="secondary" className="text-[10px] mt-0.5 shrink-0 uppercase">
                      {result.type}
                    </Badge>
                    <div className="min-w-0 flex-1">
                      <p className="font-medium text-foreground text-sm group-hover:text-primary transition-colors">
                        {result.title}
                      </p>
                      <p className="text-xs text-muted-foreground mt-0.5 line-clamp-2">
                        {result.description || result.trackTitle}
                      </p>
                      {result.highlight && (
                        <p
                          className="text-xs text-muted-foreground mt-1 line-clamp-1"
                          dangerouslySetInnerHTML={{ __html: result.highlight }}
                        />
                      )}
                    </div>
                    <ArrowRight className="h-4 w-4 text-muted-foreground group-hover:text-primary transition-colors shrink-0 mt-1" />
                  </CardContent>
                </Card>
              </Link>
            </motion.div>
          ))}
        </div>
      ) : query.length >= 2 ? (
        <div className="text-center py-16">
          <Search className="h-12 w-12 mx-auto text-muted-foreground/30 mb-4" />
          <h3 className="font-medium text-foreground mb-1">No results found</h3>
          <p className="text-sm text-muted-foreground">
            Try different keywords or browse our tracks
          </p>
        </div>
      ) : null}
    </div>
  );
}
