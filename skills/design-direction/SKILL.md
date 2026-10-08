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

Color is a UX decision, not decoration. Never pick it silently from the direction table.

**a. Find or ask for a palette.** Look for one first: the project's tokens or theme file, a brand guide, CLAUDE.md, earlier messages. If there is none and the user is present, ask once before building (a multiple-choice question when the tool exists):
- "Use my brand palette" (they paste hex values or a link),
- two or three proposed palettes, each named and with a one-line rationale tied to the subject (for example "Signal: neutral slate, violet for the new version, green and amber for health").
If the user has no palette, or the session is unattended, use the default palette below and say so in one line.

**Default palette (no user palette).** Blue and grey carry the page; green and red are reserved for good and bad. Values were checked for WCAG contrast and for separation under protanopia and deuteranopia (green is lighter and slightly teal, red darker, so the pair differs in lightness, not only hue).

| Role | Light | Dark | Use |
|---|---|---|---|
| `--bg` | `oklch(0.985 0.003 247)` cool off-white | `oklch(0.17 0.015 255)` | page ground |
| `--card` | `oklch(1 0 0)` | `oklch(0.21 0.017 255)` | surfaces |
| `--fg` | `oklch(0.22 0.02 255)` | `oklch(0.95 0.006 255)` | text, 16:1 |
| `--muted-fg` | `oklch(0.5 0.02 255)` | `oklch(0.7 0.018 255)` | secondary text, 5.7:1 / 7:1 |
| `--border` | `oklch(0.9 0.008 255)` | `oklch(0.31 0.02 255)` | lines |
| `--neutral` | `oklch(0.62 0.02 255)` slate | same | old, inactive, baseline, "the other one" |
| `--accent` | `oklch(0.55 0.18 258)` blue | `oklch(0.72 0.15 252)` | new, active, the thing to follow |
| `--ok` | `oklch(0.64 0.14 165)` | `oklch(0.82 0.14 165)` | good only (shapes; for text use `oklch(0.5 0.12 165)` in light) |
| `--error` | `oklch(0.5 0.2 25)` | `oklch(0.62 0.2 22)` | bad only |

How to map with only these:
- Comparing two entities (old vs new, before vs after, A vs B): slate vs blue. A third entity uses a darker or lighter blue step, never a new hue.
- UI chrome (progress bar, links, focus rings) uses `--fg` or neutrals when blue is taken by an entity, so blue keeps one meaning.
- When blue is taken by an entity, the moving thing to follow (traffic, a request, data) is drawn in `--fg` with a soft glow: bright ink on dark, dark ink on light. Never the same hue as an entity.
- In-progress or pending (probing, starting, terminating): neutral with a dashed border, a spinner or a pulse. Not yellow, not green.
- Good and bad: only `--ok` and `--error`, always paired with an icon or label (check, cross, "Ready", "Failed").
- No yellow or orange warning unless the user's palette has one. Something risky but not failed is described in words or shown in red at low emphasis.

**Using a user or brand palette.** The user's palette overrides every default rule in this section, including tinted neutrals and "green and red only mean good and bad". Map it like this:
- If the brand defines semantics (success, error, warning), use them exactly, even when they differ from convention (for example success = teal, error = purple). Then the brand's red is not "bad".
- The brand primary becomes `--accent`: the one thing to follow, the key moment, primary actions.
- Colors the brand reserves for gradients or highlights are not used for state; use them sparingly or not at all.
- Brand greys are the neutrals as given, even if pure grey.
- Brand black and white are marks, not surfaces. For UI use the grey steps: in dark, surfaces from the darkest greys (900, 800) and text from the light greys (50 to 300); in light, white or grey 50 surfaces and grey 900 text. Large pure-black areas with pure-white text and white glows look blown out.
- Check text contrast. If the accent is under 4.5:1 as text, add `--accent-ink` (a darker step of the same hue in light, a lighter step in dark) for text, links and code keys, and keep `--accent` for shapes.

**b. Build a color map before any token.** List every entity and state the viewer must tell apart, and give each a role:
- **Entities** that are compared (old vs new version, source vs target, team A vs team B) get distinct hues. Never the same accent at different opacities.
- **States** (ready, pending, failing, terminating) use the semantic set: ok and error for good and bad, neutral (dashed, pulsing) for in-progress. Same meaning everywhere on the page.
- **Accent** is only for the one thing the eye must follow (the traffic, the active path, the insight). If everything is accent, nothing is.
Write the map as a comment above the tokens, for example `/* v1 = neutral slate, v2 = accent blue, traffic = fg glow, ready = ok + check icon, probing = neutral dashed, failed = error + cross */`.

**c. Then the tokens:** `--bg --fg --card --muted --muted-fg --border --neutral --accent --accent-soft --ok --error`, plus entity tokens only when the user's palette provides them. Rules:
- Directions change neutrals' warmth, type, layout and motion. Without a user palette, hues stay inside the default palette.
- Neutrals carry a slight hue. No pure mid-grey.
- Color is never the only signal: entities and states also differ by label, icon or border style.
- Contrast 4.5:1 or better for text, 3:1 for meaningful shapes. Check entity hues against each other too, not only against the background.
- Themes: light and dark by default. A direction that is dark by nature (Cinematic, Night ops) may ship dark only; then put the dark values on bare `:root` with `color-scheme: dark` and drop the light and `data-theme` blocks.

## 4. Typography

Write the scale as classes before markup and use only these roles: seven roles, and `h3` may share the `body` size. Code blocks use the `small` size. Prefix the classes (`t-display`, `t-body`, `t-small` ...) so they never clash with component names.

| Role | Size / line-height | Weight | Use |
|---|---|---|---|
| `display` | 40 to 48px / 1.05, tracking -0.02em | 600 | the thesis headline, once |
| `stat` | 64 to 88px / 1, tabular | 600 | the one hero number, once |
| `h2` | 22px / 1.25 | 600 | section headings |
| `h3` | 16px / 1.35 | 600 | card titles |
| `body` | 15 to 16px / 1.6 | 400 | paragraphs, max 65ch |
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
