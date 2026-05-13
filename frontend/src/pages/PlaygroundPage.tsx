import { CodePlayground } from "@/components/playground/CodePlayground";

export default function PlaygroundPage() {
  return (
    <div className="flex flex-col h-[calc(100vh-4rem)]">
      <div className="flex items-center justify-between px-6 py-4 border-b border-border/50">
        <div>
          <h1 className="text-xl font-bold text-foreground">Code Playground</h1>
          <p className="text-sm text-muted-foreground">
            Write and run code in Go, Python, JavaScript, or TypeScript
          </p>
        </div>
      </div>
      <div className="flex-1 min-h-0">
        <CodePlayground />
      </div>
    </div>
  );
}
