export interface LessonHeading {
  id: string;
  text: string;
  level: 2 | 3;
}

export interface LessonFrontmatter {
  title?: string;
  track?: string;
  module?: string;
  difficulty?: string;
  estimatedMinutes?: number;
}

const FRONTMATTER_REGEX = /^---\r?\n[\s\S]*?\r?\n---\r?\n*/;

export function stripLessonFrontmatter(content: string): string {
  return content.replace(FRONTMATTER_REGEX, "").trim();
}

export function parseLessonFrontmatter(content: string): LessonFrontmatter {
  const match = content.match(/^---\r?\n([\s\S]*?)\r?\n---/);
  if (!match) return {};

  const frontmatter: LessonFrontmatter = {};
  for (const rawLine of match[1].split(/\r?\n/)) {
    const line = rawLine.trim();
    if (!line || !line.includes(":")) continue;

    const [key, ...rest] = line.split(":");
    const value = rest.join(":").trim().replace(/^["']|["']$/g, "");

    if (key === "title") frontmatter.title = value;
    if (key === "track") frontmatter.track = value;
    if (key === "module") frontmatter.module = value;
    if (key === "difficulty") frontmatter.difficulty = value;
    if (key === "estimatedMinutes") {
      const minutes = Number(value);
      if (!Number.isNaN(minutes)) frontmatter.estimatedMinutes = minutes;
    }
  }

  return frontmatter;
}

export function slugifyHeading(text: string): string {
  return text
    .toLowerCase()
    .trim()
    .replace(/[`*_~()[\]{}:!?'",./\\]/g, "")
    .replace(/\s+/g, "-")
    .replace(/-+/g, "-");
}

export function extractLessonHeadings(content: string): LessonHeading[] {
  const stripped = stripLessonFrontmatter(content);
  const headings: LessonHeading[] = [];
  const lines = stripped.split(/\r?\n/);
  let insideFence = false;

  for (const line of lines) {
    if (line.trim().startsWith("```")) {
      insideFence = !insideFence;
      continue;
    }

    if (insideFence) continue;

    const match = /^(##|###)\s+(.+)$/.exec(line.trim());
    if (!match) continue;

    const level = match[1].length as 2 | 3;
    const text = match[2].trim();
    headings.push({
      id: slugifyHeading(text),
      text,
      level,
    });
  }

  return headings;
}
