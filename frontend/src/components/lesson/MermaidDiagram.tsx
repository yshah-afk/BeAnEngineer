"use client";

import { useEffect, useId, useState } from "react";
import mermaid from "mermaid";
import { AlertTriangle, Network } from "lucide-react";

let initializedTheme: "dark" | "neutral" | null = null;

function getMermaidTheme(): "dark" | "neutral" {
  if (typeof document === "undefined") return "dark";
  return document.documentElement.classList.contains("light") ? "neutral" : "dark";
}

function ensureMermaidInitialized(theme: "dark" | "neutral") {
  if (initializedTheme === theme) return;

  mermaid.initialize({
    startOnLoad: false,
    securityLevel: "loose",
    theme,
    fontFamily: "Inter, ui-sans-serif, system-ui, sans-serif",
    themeVariables: {
      darkMode: theme === "dark",
      primaryColor: theme === "dark" ? "#312454" : "#ede9fe",
      primaryTextColor: theme === "dark" ? "#f8fafc" : "#1f2937",
      lineColor: theme === "dark" ? "#94a3b8" : "#64748b",
      tertiaryColor: theme === "dark" ? "#161b27" : "#f8fafc",
      noteBkgColor: theme === "dark" ? "#1f2937" : "#f8fafc",
      noteTextColor: theme === "dark" ? "#e5e7eb" : "#1f2937",
    },
  });

  initializedTheme = theme;
}

interface MermaidDiagramProps {
  chart: string;
}

export function MermaidDiagram({ chart }: MermaidDiagramProps) {
  const [svg, setSvg] = useState("");
  const [error, setError] = useState<string | null>(null);
  const diagramId = useId().replace(/:/g, "-");

  useEffect(() => {
    let cancelled = false;

    const renderDiagram = async () => {
      try {
        const theme = getMermaidTheme();
        ensureMermaidInitialized(theme);
        const { svg: renderedSvg } = await mermaid.render(`lesson-diagram-${diagramId}`, chart);
        if (!cancelled) {
          setSvg(renderedSvg);
          setError(null);
        }
      } catch (err) {
        if (!cancelled) {
          setError(err instanceof Error ? err.message : "Unable to render diagram.");
          setSvg("");
        }
      }
    };

    renderDiagram();

    return () => {
      cancelled = true;
    };
  }, [chart, diagramId]);

  if (error) {
    return (
      <div className="my-8 overflow-hidden rounded-2xl border border-destructive/30 bg-destructive/5">
        <div className="flex items-center gap-2 border-b border-destructive/20 px-4 py-3 text-sm font-medium text-destructive">
          <AlertTriangle className="h-4 w-4" />
          Diagram couldn’t be rendered
        </div>
        <pre className="overflow-x-auto p-4 text-xs leading-6 text-muted-foreground">{chart}</pre>
      </div>
    );
  }

  if (!svg) {
    return (
      <div className="my-8 flex min-h-44 items-center justify-center rounded-2xl border border-border/50 bg-muted/20 text-sm text-muted-foreground">
        <span className="inline-flex items-center gap-2">
          <Network className="h-4 w-4" />
          Rendering diagram...
        </span>
      </div>
    );
  }

  return (
    <div className="lesson-mermaid my-8 overflow-hidden rounded-2xl border border-border/50 bg-card/70 shadow-sm">
      <div className="flex items-center gap-2 border-b border-border/40 bg-muted/30 px-4 py-3 text-sm font-medium text-foreground">
        <Network className="h-4 w-4 text-primary" />
        Diagram
      </div>
      <div
        className="overflow-x-auto p-4 [&_svg]:mx-auto [&_svg]:h-auto [&_svg]:max-w-full"
        dangerouslySetInnerHTML={{ __html: svg }}
      />
    </div>
  );
}
