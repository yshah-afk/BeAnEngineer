import { useState, useCallback } from "react";
import Editor from "@monaco-editor/react";
import { useMutation } from "@tanstack/react-query";
import { Play, Loader2, RotateCcw, Copy, Check } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { playgroundApi } from "@/lib/api";
import { useThemeStore } from "@/stores/themeStore";
import { PLAYGROUND_LANGUAGES } from "@/lib/constants";
import { cn } from "@/lib/utils";

interface CodePlaygroundProps {
  initialCode?: string;
  initialLanguage?: string;
  className?: string;
}

const DEFAULT_CODE: Record<string, string> = {
  go: `package main

import "fmt"

func main() {
    fmt.Println("Hello from Go!")
}`,
  python: `def main():
    print("Hello from Python!")

if __name__ == "__main__":
    main()`,
  javascript: `function main() {
  console.log("Hello from JavaScript!");
}

main();`,
  typescript: `function greet(name: string): string {
  return \`Hello, \${name}!\`;
}

console.log(greet("TypeScript"));`,
};

export function CodePlayground({
  initialCode,
  initialLanguage = "go",
  className,
}: CodePlaygroundProps) {
  const [language, setLanguage] = useState(initialLanguage);
  const [code, setCode] = useState(initialCode ?? DEFAULT_CODE[initialLanguage] ?? "");
  const [output, setOutput] = useState("");
  const [copied, setCopied] = useState(false);
  const theme = useThemeStore((s) => s.theme);

  const runMutation = useMutation({
    mutationFn: () => playgroundApi.run(language, code),
    onSuccess: (data) => {
      setOutput(data.error ? `Error:\n${data.error}` : data.output);
    },
    onError: () => {
      setOutput("Error: Failed to execute code. Please try again.");
    },
  });

  const handleRun = useCallback(() => {
    setOutput("");
    runMutation.mutate();
  }, [runMutation]);

  const handleCopy = useCallback(async () => {
    await navigator.clipboard.writeText(code);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  }, [code]);

  const handleReset = useCallback(() => {
    setCode(DEFAULT_CODE[language] ?? "");
    setOutput("");
  }, [language]);

  const handleLanguageChange = (lang: string) => {
    setLanguage(lang);
    setCode(DEFAULT_CODE[lang] ?? "");
    setOutput("");
  };

  const monacoLang =
    PLAYGROUND_LANGUAGES.find((l) => l.value === language)?.monacoId ?? language;

  return (
    <div className={cn("flex flex-col h-full", className)}>
      <div className="flex items-center gap-2 p-3 border-b border-border/50 bg-card/50">
        <div className="flex gap-1">
          {PLAYGROUND_LANGUAGES.map((lang) => (
            <Button
              key={lang.value}
              variant={language === lang.value ? "secondary" : "ghost"}
              size="sm"
              onClick={() => handleLanguageChange(lang.value)}
              className="text-xs"
            >
              {lang.label}
            </Button>
          ))}
        </div>
        <div className="ml-auto flex items-center gap-2">
          <Button variant="ghost" size="sm" onClick={handleCopy} className="gap-1.5">
            {copied ? <Check className="h-3.5 w-3.5" /> : <Copy className="h-3.5 w-3.5" />}
            <span className="text-xs">{copied ? "Copied" : "Copy"}</span>
          </Button>
          <Button variant="ghost" size="sm" onClick={handleReset} className="gap-1.5">
            <RotateCcw className="h-3.5 w-3.5" />
            <span className="text-xs">Reset</span>
          </Button>
          <Button
            size="sm"
            onClick={handleRun}
            disabled={runMutation.isPending}
            className="gap-1.5"
          >
            {runMutation.isPending ? (
              <Loader2 className="h-3.5 w-3.5 animate-spin" />
            ) : (
              <Play className="h-3.5 w-3.5" />
            )}
            Run
          </Button>
        </div>
      </div>

      <div className="flex flex-1 flex-col lg:flex-row min-h-0">
        <div className="flex-1 min-h-[300px]">
          <Editor
            height="100%"
            language={monacoLang}
            value={code}
            onChange={(v) => setCode(v ?? "")}
            theme={theme === "dark" ? "vs-dark" : "light"}
            options={{
              minimap: { enabled: false },
              fontSize: 14,
              lineHeight: 22,
              padding: { top: 16 },
              scrollBeyondLastLine: false,
              smoothScrolling: true,
              cursorBlinking: "smooth",
              renderLineHighlight: "gutter",
              bracketPairColorization: { enabled: true },
              fontFamily: "var(--font-mono)",
            }}
          />
        </div>

        <div className="lg:w-[40%] border-t lg:border-t-0 lg:border-l border-border/50 bg-muted/30">
          <div className="flex items-center justify-between px-4 py-2 border-b border-border/50">
            <span className="text-xs font-medium text-muted-foreground uppercase tracking-wider">
              Output
            </span>
            {runMutation.isPending && (
              <Loader2 className="h-3.5 w-3.5 animate-spin text-muted-foreground" />
            )}
          </div>
          <Card className="m-0 border-0 rounded-none bg-transparent shadow-none">
            <pre className="p-4 text-sm font-mono whitespace-pre-wrap min-h-[200px] text-foreground/80">
              {output || (
                <span className="text-muted-foreground italic">
                  Click "Run" to execute your code...
                </span>
              )}
            </pre>
          </Card>
        </div>
      </div>
    </div>
  );
}
