---
name: design-direction
description: 'Use when an HTML page, explainer or UI needs a look, or when output feels generic, templated or "AI-made". Picks a distinctive visual direction (palette, typography, layout, one bold move) and enforces a banned list of AI design tells.'
---

# Design direction

A clean shadcn page is the floor, not the goal. Clean and predictable reads as "fine". Every page gets one deliberate direction and one bold move.

Precedence: the user's words, then the project's existing design system, then this skill.

## 1. Pick a direction

Pick from the subject, not by habit. Do not reuse the direction of the previous page in the same conversation unless the user asked for it.

| Direction | Feels like | Palette | Type | Motion character | Best for |
|---|---|---|---|---|---|
| **Instrument** | Linear, Vercel dashboard | default palette, light first | Geist + Geist Mono | precise, short, expo.out | product UI, reference sheets |
| **Cinematic** | a film title sequence over a system | default palette dark, deeper blue-black ground, heavy vignette | Geist or Manrope, large and tight; mono for data | camera moves, depth of field, slow pull-backs | traces, request flows, "follow the X" stories |
| **Night ops** | a control room at 3 am | default palette dark, near-black ground, blue as light, glow only on active things | IBM Plex Sans + IBM Plex Mono | light trails, particles on paths, counters ticking | live systems, pipelines, incidents |
| **Editorial** | an interactive long-read | default palette light, off-white paper, ink text | Bricolage Grotesque or Instrument Sans display + Hanken Grotesk body | scroll-driven, pinned diagram, text reveals | concepts, decisions, postmortems |
| **Blueprint** | an engineering drawing | default palette with a blue ground, thin line work, dimension marks | JetBrains Mono for labels + Onest body | lines drawing themselves, measured reveals | architecture, protocols, data models |

If none fits, define a new one in the same shape (feel, palette, type, motion, why) before writing code.

## 2. Make one bold move

Choose exactly one and push it. Keep everything else quiet.

- A camera that travels the diagram (see motion-choreography).
- An oversized number or headline that carries the insight ("236 ms").
- A strong asymmetric layout: diagram bleeding off one edge, text in a narrow column.
- A dark stage inside an otherwise light page (at container width, aligned with the other blocks). Scope its tokens with a class that redefines them (`.c-stage{--card:#000;--fg:#fafafa;...;color-scheme:dark}`).
- Type as the visual: kinetic headline that resolves into the diagram.

## 3. Color: ask, then assign by meaning

Color is a UX decision, not decoration. Read `references/color.md` in full before choosing any color. It holds the palette, the default values and the mapping rules, and it is the only copy.

## 4. Typography

Write the scale as classes before markup and use only these roles: seven roles, and `h3` may share the `body` size. Code blocks use the `small` size. Prefix the classes (`t-display`, `t-body`, `t-small` ...) so they never clash with component names.

| Role | Size / line-height | Weight | Use |
|---|---|---|---|
| `display` | 40 to 48px / 1.05, tracking -0.02em | 600 | the thesis headline, once |
| `stat` | 64 to 88px / 1, tabular | 600 | the one hero number, once |
| `h2` | 22px / 1.25 | 600 | section headings |
| `h3` | 16px / 1.35 | 600 | card titles |
| `body` | 15 to 16px / 1.6 | 400 | paragraphs, full width of their container (same edges as the stage); use columns, not a narrow measure, to shorten lines |
| `small` | 13px / 1.45 | 400 or 500 | captions, labels next to values, helper text |
| `label` | 11px uppercase, tracking 0.12em | 500 | eyebrows, at most one per block |

Rules:
- Two families at most: one sans for everything readable, one mono. Never default to Inter, Roboto, Arial or the system stack as the display face.
- Mono is only for code, ids, commands, field names and numbers in data rows. Words and units stay in the sans: `3` in mono, "réplicas" in sans.
- Siblings share a role. Values in the same list have the same size and weight; labels in the same list have the same size. Never mix a 15px mono value with a 13px sans label in one row and a 28px number in the next.
- Inline code inside body text (`maxSurge`) uses mono at 0.92em with a subtle background, so it does not shout louder than the sentence.
- Headlines: `text-wrap: balance`, two lines at most.

## 5. Icons

- One icon set (Lucide in the stack). No emoji, ever.
- All-or-none per group: if one sibling heading has an icon, all siblings do, at the same size (16px), stroke and position. A group where some cards have icons and others do not is a bug.
- Icons in headings are neutral (`--muted-fg`). Semantic color (ok, error) is only for real state in the content: a pod that is ready, a check that failed. A heading that talks about failure is not a failure.
- Icons support a label, they never replace one.

## 6. Banned tells (unless the user asks for them)

- Purple-to-blue gradient heroes, neon outer glows on everything, glassmorphism by default.
- Cream background with terracotta accent and a serif display.
- Three equal feature cards. Every block in the same rounded card with the same shadow.
- Emoji as icons or section markers. Hand-drawn decorative SVG.
- Numbered eyebrows (01 / 02) when the content is not a sequence. A step counter on a real sequence ("Passo 3 de 8") is fine.
- Bounce, elastic or back easing.
- Fake-precise numbers, "Acme", "John Doe", filler verbs (elevate, seamless, unleash).
- Em dashes in visible copy.
- Default Mermaid look in a page that is meant to be beautiful.

## 7. Diagrams in the direction's language

**Depict the mechanism, not its name.** The concept the page is about must be a place or a thing on the stage: a buffer is a box that fills, a queue is a line that grows, a decision is a gate things pass or do not pass. A table of rows with marks that light up is not a mechanism. Label the arrows with what flows (`OTLP`, `trace inteiro`, `polls every 30s`). Things that move travel visibly from one place to another; they never fade out in transit.


Nodes are real UI (cards, chips, icons from the direction's kit), placed in a grid or absolute layout. Connections are an SVG overlay with paths computed from node positions. This keeps the diagram in the same visual language as the page. Mermaid is only for quick drafts where looks do not matter.
