---
name: visual-explainer
description: 'Use when asked to visualize, explain visually, "show me", diagram, or make an HTML explainer or page for a concept, flow, architecture, trace, plan or comparison. Builds a beautiful, animated, token-efficient HTML page instead of prose and orchestrates design-direction, motion-choreography and visual-critique.'
---

# Visual explainer

**Loading these skills:** read every SKILL.md and reference file in full with the file-reading tool, never with `cat` or `head` (long tool output gets truncated in the middle and silently drops rules). Load references only when the step that needs them comes up.

Entry point for visual pages. It decides whether to go visual, reads the brief, sets the dials, picks the stack for the surface, and pulls in the other skills:

- `design-direction`: how it should look (taste, palette, type, banned tells).
- `motion-choreography`: how it should move (camera, trace, reveal, controls).
- `visual-critique`: one review pass before delivery.

The goal is pages that surprise and still cost few output tokens. Spend tokens on content and on motion that teaches. Never spend them reinventing CSS.

## 1. Go visual only when it pays

Go visual for flows, architectures, request paths, pipelines, state machines, timelines, side-by-side comparisons and structured plans. If a paragraph or a small table answers the question, answer in text and stop.

## 2. Read the brief (one line, before any code)

State it in the reply or as a comment at the top of the file:

> Reading this as: [what is being explained] for [who], the one thing they must get is [insight], direction [name from design-direction].

The "one thing" decides the single orchestrated moment of the page. Example: "the payment call is the bottleneck", so the camera lands there and holds.

## 3. Set the three dials (1 to 10)

| Dial | Low | High | Default for explainers |
|---|---|---|---|
| `VARIANCE` | symmetric grid | asymmetric, editorial | 5 |
| `MOTION` | static, hover only | choreographed camera and timeline | 7 |
| `DENSITY` | one idea per screen | data-dense, dashboard-like | 5 |

Adjust from the brief without asking. A trace or pipeline pushes MOTION up. A comparison or reference sheet pushes DENSITY up and MOTION down. Write the chosen values in the brief line.

## 4. Storyboard, when it pays

A storyboard is 4 to 8 lines, one per frame: what is on screen, what moves, what the viewer learns. Example: "spans leave three services and enter the Collector box / each trace gets a compartment that fills, with a 10 s ring / when the ring closes, a scanner sweeps the compartment / a kept trace slides whole to the exporter".

**Write one when** the animation carries the explanation: MOTION 6 or higher, a multi-step mechanism where things move between places, or a diagram type you have not built before. If the user is present and the page is new, show it and get a yes before building; it is cheap to change three lines and expensive to change a built page. If unattended, write it as a comment at the top of the file and build from it.

**Skip it when** the page is static or reference material, a comparison, motion is only reveals and hover (MOTION 4 or lower), the change is an edit to an existing page, or the user already described the frames.

A storyboard for a mechanism answers three questions, each in a frame:

- **Trigger:** what starts it (a schedule, an event, a request, a poll).
- **Order:** how items move through it (FIFO, priority, fan-out, one at a time or in parallel).
- **Pressure:** what happens when it is overloaded or a part fails (items wait, drop, retry, back off).

A storyboard without a pressure frame is not finished.

Test every storyboard with one question: do the verbs describe movement between places (enters, fills, sweeps, slides, leaves)? If most frames say "appears", "highlights" or "fades", the motion is decoration. Redesign the diagram until the core concept is a place or a thing that changes.

## 5. Pick the stack for the surface

- **claude.ai Artifact, or a single shareable HTML file**: read `references/stack-artifact.md`. CDN scripts only, no external stylesheets except Google Fonts.
- **Inside a repo (React, Next, Vite)**: read `references/stack-repo.md`. Real shadcn, Tailwind and GSAP via npm.

Never improvise library versions. Use the pinned ones in the reference.

## 6. Content architecture

Plan the page as a reading order before any styling. A beautiful animation inside a broken layout is a broken page.

1. **Thesis**: headline plus one sentence.
2. **Stage**: the animated diagram, full content width. Nothing sits beside it.
3. **Live companions**: anything that changes with the timeline (counters, status, the current step text) lives inside the stage as a HUD overlay in a corner or in the caption, or in one row directly under the stage, aligned to the stage edges.
4. **Supporting detail**: static material (config, YAML, code, notes) below, in a grid whose columns share top edges and heights. Code is never clipped: full width, or collapsible with `<details>` when long.

Rules:

- For every element ask "does it change with the timeline?" Yes: inside or directly under the stage, synced. No: below, as supporting detail.
- No narrow sidebar next to an animated stage. Side-by-side columns only for content of equal weight that aligns top and bottom.
- Every block shares the same left and right edges (one container, one grid). Inner content aligns to a consistent baseline.
- No ambiguous mini-charts. A chart needs labels or an obvious reading; otherwise drop it.
- Every section after the stage starts with an `h2` heading. A section without a heading is just a gap.
- Stat groups: one hero number, then supporting stats as rows of `label` (sans, small, muted) on the left and value (tabular, same size for every row) on the right, separated by hairlines. Or a strip of equal tiles. Never a cramped column of mixed sizes. A card holding stats is at least 320px wide; if the row cannot give it that, stack the row.

**Spacing rhythm.** Distance encodes relationship. Use one scale and keep to it:

| Between | Gap |
|---|---|
| lines inside a block (label to value, title to text) | 4 to 8px |
| blocks inside a card | 12 to 16px |
| sibling cards | 16px |
| groups inside a section (stage to player to companion row) | 16 to 24px |
| sections (always with an `h2`) | 56 to 64px |

Card inner padding is the same on every card (20 or 24px). Never use a bigger gap to "separate" things that have no heading.

**Dividers, not boxes.** Cards are for objects: the stage, a code block, a stat group, something interactive. Sections are separated by a full-width hairline (`border-top: 1px solid var(--border)`) with the `h2` 24 to 32px below it. Prose inside a section (notes, explanations, "what if" cases) sits directly on the page, in columns when there are siblings, with no card around each paragraph. Between sibling columns use a vertical hairline or just the gap, never a box per column. A page where every block is a card has no hierarchy.

Check 400px, 768px and 1280px. At tablet width nothing should be squeezed side by side.

### Docs with several parts: tabs

Use tabs when the page has three or more parts (how it works, what we decided, what is open, what ships when) and would need long scrolling, or when it will be updated over many sessions. With fewer parts, keep one page with `h2` sections.

- **One tab per part of the work, never per item.** `Open questions` is one tab that holds every question (Q1 to Q20 are rows in it, not tabs). Name tabs with the reader's word: `Architecture`, `Requirements`, `Open questions`, `Delivery`. Use the user's names when they gave them.
- **Five tabs at most.** A sixth means two tabs overlap: merge them, or move content out to the tracker.
- **The first tab is the overview:** thesis, the main animation, and one line per other tab that says what the reader finds there.
- **Each tab follows the reading order above** (thesis, stage, companions, detail), with at most one animated stage.
- **Every section opens with its answer** in one sentence: the proposal, the decision, or the status. Detail comes after.
- **Each fact lives in one place.** If two tabs need it, one owns it and the other links to it. Before you add a section, check that no other tab already says it.
- **Only what the reader acts on:** decisions, open questions, designs, and links. History, logs and long notes go to the tracker or notes, with a link.
- **Done items leave the main flow.** Closed questions, shipped tickets and old decisions go into one collapsed `<details>` at the end of their tab, one line each.
- **Status shows what blocks it.** Each tracked item has a state chip, and a blocked item names the blocker ("blocked by API delete"), not only a red mark.
- **Edits start with placement.** Decide which tab owns the new content before you write it. If no tab fits, propose a new tab. Never append to the end of the page.

Tab markup is in the stack reference.

## 7. Token budget rules

- Copy the base block from the stack reference verbatim, then use only classes. Do not write custom CSS that Tailwind or the component layer already covers.
- Generate repeated elements (nodes, rows, steps, cards) from a small JS data array plus a template string. Never copy-paste markup.
- Icons come from the icon library. Never hand-draw SVG icons or decorative SVG paths.
- One master GSAP timeline per page. No scattered one-off tweens.
- For a tweak, edit the existing file in place. Never regenerate the page.
- Target size: a rich page should fit in 12 to 20 KB. If it grows past 25 KB, something is being reinvented.

## 8. Content rules

- Real content, real units, real names from the subject (span names, ms, ids, versions). No lorem ipsum, no "Acme", no fake-precise numbers without a source.
- Component names come from the code or the system. When it has none, use the industry term (queue, worker, scheduler, load balancer). Never invent a name like "dispatch service".
- Every mark (dot, square, bar, chip) has a label or a legend entry. A mark with neither is removed. A section that does not serve the brief's "one thing" is removed too.
- Short headline (8 words or fewer is ideal), short paragraphs. The page explains, it does not market.
- **Text budget.** The animation carries the explanation; text only names what the viewer sees.
  - Step caption: one sentence.
  - Diagram nodes: a name and at most one short line (a value, a state). No sentences inside boxes.
  - Section prose: about 80 words. More than that becomes a table or a list, goes into a collapsed `<details>`, or moves out behind a link.
  - When the user asks for less text, cut words. Do not move the same words into more boxes.
- **No AI-sounding copy.** Say the thing directly. Banned:
  - Contrast framing: "Two new things, not a new system", "It's not X, it's Y".
  - Staged openers: "First, answer one question", "Here's the thing", "The key insight:".
  - One-line punchline closers after a paragraph that already made the point.
  - If a prose-cleaning skill is installed (for example `unslop` or `humanizer`), run the visible copy through it.
- No em dashes in visible copy. Use commas, periods or line breaks.
- Write in the user's language.
- Every code, config or command block gets syntax highlighting (see the stack reference). Never hand-color code.

## 9. Build order

1. Brief line and dials.
2. Sketch the content architecture (section 6) as a list of rows, and write the storyboard if section 4 asks for one.
3. List every arrow in the diagram as `from → to: what flows`, each with evidence: `file:line`, a doc link, or the user's words. An arrow without evidence is drawn dashed with an "unverified" label, or not drawn. Write the list as a comment at the top of the file.
4. Load `design-direction` and its `references/color.md`: choose a direction, take the palette from the user, the repo's design system, a question, or the default (in that order), then write the color map and tokens.
5. Load `motion-choreography` and pick the patterns that serve the "one thing".
6. Build from the stack reference.
7. Load `visual-critique` and do one review pass.
8. Deliver: publish as an Artifact when available (two to four word title, one sentence description), otherwise write the file next to the work. Reply with one or two lines on what the page shows.

**Edits to an existing page** skip steps 1, 2, 4 and 5. If the edit adds or changes an arrow, do step 3 for that arrow. Edit in place, then run the edit pass from `visual-critique` before you republish. Every edit, however small: most layout bugs come from edits, not first builds.
