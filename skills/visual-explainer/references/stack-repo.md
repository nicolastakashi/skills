# Stack: inside a repo

Use when the page or component lives in a real project. Follow the project's existing design system first. This file only fills gaps.

## Packages

```bash
npm i gsap @gsap/react lucide-react
npx shadcn@latest init      # only if the project has no component system yet
npx shadcn@latest add button card badge tabs slider
```

Tailwind v4 through the Vite plugin (`@tailwindcss/vite`) or `@tailwindcss/postcss`. Fonts through `next/font` (Geist ships as `geist/font`) or self-hosted `@font-face` with `font-display: swap`.

Code blocks: [Shiki](https://shiki.style) at build time (or `rehype-pretty-code` in MDX), with a custom theme whose colors come from the same `--code-*` tokens as the artifact stack.

## Rules

- Use the shadcn components as they are. Theme through CSS variables in the project's globals, not by editing component internals.
- GSAP lives in leaf client components. In React, use `useGSAP()` from `@gsap/react` with a scoped ref, so cleanup is automatic.
- Register plugins once: `gsap.registerPlugin(useGSAP, DrawSVGPlugin, MotionPathPlugin, SplitText, Flip, ScrollTrigger)`.
- Keep the data (nodes, edges, steps) in a typed constant next to the component, and render from it.
- Same tokens, motion rules and critique pass as the artifact stack.
