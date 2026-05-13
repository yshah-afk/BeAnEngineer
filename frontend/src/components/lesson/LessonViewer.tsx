import { isValidElement } from "react";
import ReactMarkdown, { type Components } from "react-markdown";
import remarkGfm from "remark-gfm";
import rehypeHighlight from "rehype-highlight";
import rehypeRaw from "rehype-raw";
import { cn } from "@/lib/utils";
import { MermaidDiagram } from "@/components/lesson/MermaidDiagram";
import { slugifyHeading, stripLessonFrontmatter } from "@/lib/lesson-content";

interface LessonViewerProps {
  content: string;
  className?: string;
}

export function LessonViewer({ content, className }: LessonViewerProps) {
  const markdown = stripLessonFrontmatter(content);

  const components: Components = {
    h1: ({ children }) => (
      <h1 className="mt-0 mb-6 text-4xl font-bold tracking-tight text-foreground">{children}</h1>
    ),
    h2: ({ children }) => {
      const text = String(children);
      const id = slugifyHeading(text);
      return (
        <h2
          id={id}
          className="mt-12 scroll-mt-24 border-b border-border/50 pb-3 text-2xl font-semibold tracking-tight text-foreground"
        >
          <a href={`#${id}`} className="transition-colors hover:text-primary">
            {children}
          </a>
        </h2>
      );
    },
    h3: ({ children }) => {
      const text = String(children);
      const id = slugifyHeading(text);
      return (
        <h3
          id={id}
          className="mt-8 scroll-mt-24 text-xl font-semibold tracking-tight text-foreground"
        >
          <a href={`#${id}`} className="transition-colors hover:text-primary">
            {children}
          </a>
        </h3>
      );
    },
    p: ({ children }) => <p className="my-5 text-[1.02rem] leading-8 text-foreground/85">{children}</p>,
    a: ({ href, children }) => (
      <a
        href={href}
        target={href?.startsWith("http") ? "_blank" : undefined}
        rel={href?.startsWith("http") ? "noreferrer" : undefined}
        className="font-medium text-primary underline decoration-primary/35 underline-offset-4 transition-colors hover:decoration-primary"
      >
        {children}
      </a>
    ),
    ul: ({ children }) => <ul className="my-5 space-y-2 pl-6 text-foreground/85">{children}</ul>,
    ol: ({ children }) => <ol className="my-5 space-y-2 pl-6 text-foreground/85">{children}</ol>,
    li: ({ children }) => <li className="pl-1 leading-7 marker:text-primary/70">{children}</li>,
    strong: ({ children }) => <strong className="font-semibold text-foreground">{children}</strong>,
    blockquote: ({ children }) => (
      <blockquote className="my-6 rounded-r-2xl border-l-4 border-primary/45 bg-primary/6 px-5 py-4 text-foreground/75">
        {children}
      </blockquote>
    ),
    hr: () => <hr className="my-10 border-border/50" />,
    table: ({ children }) => (
      <div className="my-6 overflow-x-auto rounded-2xl border border-border/50">
        <table className="w-full border-collapse text-sm">{children}</table>
      </div>
    ),
    thead: ({ children }) => <thead className="bg-muted/40">{children}</thead>,
    th: ({ children }) => (
      <th className="border-b border-border/50 px-4 py-3 text-left font-semibold text-foreground">
        {children}
      </th>
    ),
    td: ({ children }) => <td className="border-t border-border/30 px-4 py-3 align-top text-foreground/80">{children}</td>,
    details: ({ children }) => (
      <details className="my-6 overflow-hidden rounded-2xl border border-border/50 bg-card/40">
        {children}
      </details>
    ),
    summary: ({ children }) => (
      <summary className="cursor-pointer list-none border-b border-border/40 bg-muted/25 px-4 py-3 font-medium text-foreground transition-colors hover:bg-muted/40">
        {children}
      </summary>
    ),
    pre: ({ children }) => {
      const child = Array.isArray(children) ? children[0] : children;
      const className =
        isValidElement<{ className?: string }>(child) ? child.props.className ?? "" : "";
      const language = className.replace("language-", "") || "code";

      if (className.includes("language-mermaid")) {
        return <>{children}</>;
      }

      return (
        <div className="my-7 overflow-hidden rounded-2xl border border-border/50 bg-[#0b1020] shadow-sm">
          <div className="border-b border-white/8 px-4 py-2 text-xs font-medium uppercase tracking-[0.18em] text-slate-400">
            {language}
          </div>
          <pre className="overflow-x-auto p-4 text-sm leading-7 text-slate-100">{children}</pre>
        </div>
      );
    },
    code: ({ className, children, ...props }) => {
      const value = String(children).replace(/\n$/, "");
      const language = className?.replace("language-", "");

      if (language === "mermaid") {
        return <MermaidDiagram chart={value} />;
      }

      if (className) {
        return (
          <code className={cn("font-mono text-sm", className)} {...props}>
            {children}
          </code>
        );
      }

      return (
        <code
          className="rounded-md border border-border/40 bg-muted/45 px-1.5 py-0.5 font-mono text-[0.95em] text-primary"
          {...props}
        >
          {children}
        </code>
      );
    },
  };

  return (
    <article
      className={cn(
        "lesson-markdown max-w-none text-foreground",
        "[&_img]:my-6 [&_img]:rounded-2xl [&_img]:border [&_img]:border-border/40 [&_img]:shadow-lg",
        className,
      )}
    >
      <ReactMarkdown
        remarkPlugins={[remarkGfm]}
        rehypePlugins={[rehypeHighlight, rehypeRaw]}
        components={components}
      >
        {markdown}
      </ReactMarkdown>
    </article>
  );
}
