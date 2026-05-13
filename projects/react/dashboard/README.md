# Project: Build a Dashboard with Zustand + TanStack Query + shadcn/ui

## Description

Create a data-rich analytics dashboard for the "AI & Full-Stack Mastery Hub" learning platform. The dashboard displays student progress, course analytics, leaderboards, and activity trends with real-time updates, interactive charts, responsive layout, and dark mode support. Built with the recommended React state management stack.

## Learning Objectives

By completing this project, you will:

- Scaffold and configure a React + Vite + TypeScript project with Tailwind CSS v4
- Install and customize shadcn/ui components with theme variables
- Manage client/UI state with Zustand (sidebar, theme, filters)
- Manage server state with TanStack Query (data fetching, caching, polling)
- Build responsive layouts using CSS Grid and Flexbox
- Implement dark mode with CSS custom properties and system preference detection
- Create interactive charts using a charting library (Recharts or Chart.js)

## Prerequisites

- React project setup with Vite and TypeScript (completed)
- Components, JSX, and Hooks knowledge (completed)
- Node.js 18+, npm or pnpm

## Architecture Overview

```
┌──────────────────────────────────────────────────────┐
│                     App Shell                         │
│  ┌─────────┐  ┌──────────────────────────────────┐   │
│  │ Sidebar  │  │         Main Content              │   │
│  │          │  │  ┌──────────────────────────────┐ │   │
│  │ • Home   │  │  │  Stats Cards (4 KPIs)        │ │   │
│  │ • Courses│  │  └──────────────────────────────┘ │   │
│  │ • Users  │  │  ┌──────────┐  ┌───────────────┐ │   │
│  │ • Scores │  │  │ Line     │  │  Bar Chart    │ │   │
│  │          │  │  │ Chart    │  │  (by track)   │ │   │
│  │ ──────── │  │  └──────────┘  └───────────────┘ │   │
│  │ Settings │  │  ┌──────────────────────────────┐ │   │
│  │ Theme 🌙 │  │  │  Data Table + Pagination     │ │   │
│  └─────────┘  │  └──────────────────────────────┘ │   │
│               └──────────────────────────────────┘   │
└──────────────────────────────────────────────────────┘

State Management:
  Zustand  → sidebar open/closed, theme, active filters
  TanStack → dashboard stats, user list, course data, leaderboard
```

## Acceptance Criteria

### Layout & Navigation

- [ ] Responsive sidebar that collapses to icons on mobile
- [ ] Top header with breadcrumbs, search, and user avatar
- [ ] Main content area with grid layout
- [ ] Smooth transitions when sidebar toggles

### Dashboard Page

- [ ] **KPI Cards** — 4 stat cards showing: Total Students, Active Courses, Completion Rate, Average Score
- [ ] **Activity Chart** — Line chart showing daily active users over the past 30 days
- [ ] **Track Distribution** — Bar or pie chart showing student distribution across tracks
- [ ] **Leaderboard** — Top 10 students by XP/points with rank, avatar, name, and score
- [ ] **Recent Activity** — Feed of recent completions, quiz results, and enrollments

### Data Management

- [ ] TanStack Query for all API data with stale-while-revalidate caching
- [ ] Loading skeletons while data is being fetched
- [ ] Error states with retry buttons
- [ ] Automatic background refetching every 60 seconds
- [ ] Optimistic updates where applicable

### UI State (Zustand)

- [ ] Sidebar collapsed/expanded state (persisted to localStorage)
- [ ] Theme mode (light/dark/system) persisted
- [ ] Active date range filter for charts
- [ ] Selected track filter for filtering dashboard data

### Dark Mode

- [ ] Toggle between light, dark, and system preference
- [ ] All components properly themed
- [ ] Charts adapt colors to current theme
- [ ] Smooth transition without flash

### Responsive Design

- [ ] Desktop (1280px+): Full sidebar + 2-column charts
- [ ] Tablet (768-1279px): Collapsed sidebar + 1-column charts
- [ ] Mobile (< 768px): Hidden sidebar (hamburger menu) + stacked layout

### Data Table

- [ ] Sortable columns (click header to sort)
- [ ] Pagination (10/20/50 per page)
- [ ] Search/filter input
- [ ] Row click to view details

## Getting Started

### Step 1: Scaffold the Project

```bash
npm create vite@latest learning-dashboard -- --template react-ts
cd learning-dashboard
npm install
```

### Step 2: Install Dependencies

```bash
npm install @tanstack/react-query zustand recharts
npx shadcn@latest init
npx shadcn@latest add button card table badge input select
npm install -D tailwindcss @tailwindcss/vite
```

### Step 3: Create a Mock API

Create `src/lib/mock-api.ts` that returns realistic dashboard data with simulated network delays:

```typescript
export async function fetchDashboardStats() {
  await new Promise(r => setTimeout(r, 800))
  return {
    totalStudents: 12847,
    activeCourses: 6,
    completionRate: 68.5,
    averageScore: 82.3,
  }
}
```

### Step 4: Build Incrementally

1. App shell with sidebar and routing
2. Zustand store for UI state
3. Mock API + TanStack Query setup
4. KPI cards with loading skeletons
5. Charts with responsive containers
6. Data table with sorting and pagination
7. Dark mode toggle
8. Responsive breakpoints

## Hints and Tips

- **Mock data first** — Build the entire UI with mock data before connecting to a real API. This lets you iterate on the design quickly.
- **Zustand is minimal** — A single store with 5-10 properties is fine. Don't over-engineer.
- **TanStack Query keys** — Use structured query keys like `['dashboard', 'stats', { dateRange }]` for proper cache invalidation.
- **Chart responsiveness** — Wrap charts in a `ResponsiveContainer` component. Set height with CSS, not props.
- **shadcn/ui customization** — Modify the CSS variables in `globals.css` to match your brand colors.

## Bonus Challenges

1. **Real API Integration** — Connect to the actual backend API instead of mocks
2. **Data Export** — Export dashboard data as CSV or PDF
3. **Custom Date Range Picker** — Allow selecting custom date ranges for all charts
4. **Drag-and-Drop Layout** — Let users rearrange dashboard widgets
5. **Accessibility Audit** — Run axe-core and fix all a11y issues to WCAG 2.1 AA

## Resources

- [TanStack Query Documentation](https://tanstack.com/query/latest)
- [Zustand Documentation](https://zustand-demo.pmnd.rs/)
- [shadcn/ui Components](https://ui.shadcn.com/)
- [Recharts](https://recharts.org/)
- [Tailwind CSS v4](https://tailwindcss.com/docs)
