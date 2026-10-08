---
name: visual-critique
description: 'Use after building a visual HTML page, or when asked to critique, polish, refine, make it bolder, quieter or "surprise me". One screenshot review pass in light, dark and phone width against the design and motion rules, then one fix pass.'
---

# Visual critique

One review pass, then deliver. This is not a loop.

## 1. Look at it

Use whatever this session offers, once:

- Artifact preview tool if available.
- Otherwise Playwright: screenshots at 1280px, 768px and 400px wide, light and dark (`colorScheme`), plus one frame mid-animation and one at the end (`tl.pause(); tl.progress(x)`). Collect console errors and check `document.documentElement.scrollWidth` equals the viewport width.
- If CDNs are blocked where you run, point the script tags at local copies from `npm pack` for the check only.

## 2. Check

**Breaks (fix always)**
- Console errors, a library that did not load, icons that did not render.
- Horizontal page scroll, clipped or overlapping text, a caption covering the focus.
- Unreadable text in either theme, a color that ignores the theme.
- Rest state is empty or invisible without the animation.

**Structure (fix always)**
- Reading order is thesis, stage, live companions, supporting detail. Nothing sits beside the animated stage.
- All blocks share the same left and right edges. Columns side by side align top and bottom. No clipped code.
- At 768px nothing is squeezed side by side.
- Every section after the stage has an `h2`. Gaps follow the spacing scale: nothing related sits far apart, nothing unrelated sits close.
- Stat cards: labels and values in consistent rows, same sizes, nothing cramped.
- Sections are separated by hairline dividers with an `h2`. Cards only wrap objects (stage, code, stats, interactive). Prose in a card is a bug.

**Type and icons (fix always)**
- Only the roles from the type scale are used, six sizes at most. Mono only for code, ids and numbers; units and words in sans.
- Sibling headings either all have icons (same size, neutral color) or none do. No emoji.
- Every code block is syntax highlighted with the palette theme, not hand-colored spans, and stays readable in both themes.

**Color (fix always)**
- There is a color map comment and the page follows it. Entities the viewer compares have distinct hues, states use the semantic set, the accent marks only the one thing to follow.
- In a grayscale screenshot, categories are still distinguishable by label, icon or border.

**Taste (fix when it fails)**
- Is there exactly one bold move, and does it serve the brief's "one thing"?
- Would this page look the same for any other subject? If so, the direction is not specific enough.
- Anything from the banned list in `design-direction`.
- Every animation answers "what does this show?" No bounce or elastic. Durations in range.
- Controls look designed (player pill, segmented progress), not native widgets.

**Mechanism (fix always)**
- Describe the animation in one sentence per step. If most verbs are "appears", "highlights" or "fades" instead of movement between places, the motion is decoration: redesign the diagram, do not polish it.
- The page's core concept is visible as a place or thing on the stage.

**Animation integrity (fix always)**
- Screenshot at least three frames: start, the key step, the last frame. In each, every element sits on its anchor (time, node, row) and nothing is left over from an element that already left (orphan outlines, placeholders, half-visible marks).
- Scrub backwards once and re-check the key step: the page must look the same as when played forward.

**Phone**
- The focus area is readable at 400px. Zoom caps may need to be higher on small screens.
- Controls wrap cleanly, progress bar gets its own row.

## 3. Fix once, then deliver

Make one pass of edits for what failed, in place. Do not take a second round of screenshots unless something was visibly broken.

## Refinement commands

When the user asks for one of these on an existing page, change only that dimension, in place:

| Ask | Do |
|---|---|
| **bolder** | Push the bold move further: bigger type contrast, stronger camera move, more saturated accent on the focus only. Do not add new effects elsewhere. |
| **quieter** | Remove a layer: fewer glows, shorter moves, lower contrast on secondary elements. Keep the one moment. |
| **polish** | Spacing rhythm, alignment of repeated elements, type scale consistency, focus states, reduced motion. |
| **animate** | Add motion only where it explains: a trace, a reveal of the thesis, a state change. Use `motion-choreography`. |
| **surprise me** | Switch to a different direction from `design-direction` and rebuild the look, keeping content and data. |
