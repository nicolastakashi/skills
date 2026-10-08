---
name: motion-choreography
description: 'Use when an HTML page or explainer needs animation that explains something: camera moves over a diagram, traces travelling along connections, staged reveals, scrollytelling, or step-by-step timelines with player controls. Modern GSAP motion with tested skeletons.'
---

# Motion choreography

Motion here explains things. Every animation must answer "what does this show?" (hierarchy, cause and effect, sequence, state change). "It looks cool" does not count.

For GSAP API details, performance and framework integration, rely on the official GreenSock skills (`npx skills add https://github.com/greensock/gsap-skills`) when they are installed. This skill is the vocabulary and the taste on top.

## Feel (non-negotiable)

- Easing: `expo.out` or `power3.out` for entrances, `power2.inOut` for travel, `power3.inOut` for camera moves. Never `bounce`, `elastic` or `back`.
- Duration: 0.15 to 0.25s for micro-interactions, 0.4 to 0.8s for entrances, 0.9 to 1.4s for camera moves.
- Entrances: small offsets (8 to 24px), opacity, and blur (`filter: blur(6px)` to `0`). Stagger 0.04 to 0.08s.
- Animate transforms, opacity and filter only. Never width, height, top or left.
- One orchestrated moment per page, tied to the brief's "one thing". Everything else stays calm.

## Patterns

Pick the ones that serve the brief. Code for each is in `references/skeletons.md`. Read only the sections you use.

| Pattern | Use it to show | Skeleton |
|---|---|---|
| **Master timeline + player** | any multi-step story | `timeline`, `player` |
| **Camera** | where to look now: zoom and pan to the active part, then pull back to reveal the whole | `camera` |
| **Depth of field** | focus: blur and dim what is not active | `camera` |
| **Trace** | something travelling: a request, a message, data | `wires`, `trace` |
| **Linked visuals** | the same event in two views (diagram plus waterfall, chart plus table) | `trace` |
| **Headline reveal** | the page's thesis landing | `reveal` |
| **Flip** | layout change: tab switch, reorder, filter | GSAP Flip docs |
| **Scrollytelling** | a long story where a pinned visual evolves with scroll | ScrollTrigger with `pin: true`, `start: "top top"`, `scrub: 1` |

## Choosing patterns

Not every page gets a camera. Pick by what the content needs:

- **Camera**: the diagram is bigger than the viewport can show legibly, and attention moves between distant parts (a request crossing many services). Skip it when the whole diagram fits and stays readable at fit-all.
- **No camera, highlight in place**: compact diagrams, state machines, before/after. Use depth of field, rings and traces without moving the view. Give the diagram a second layout for narrow stages (under about 820px): rearrange the nodes into a vertical or tighter grid from the same data, instead of shrinking the desktop layout.
- **Scrollytelling**: long narratives where reading pace should be the user's.
- **Flip or reveal only**: comparisons, reference pages, dashboards.
- **Time axis**: processes over time (buffers, windows, retries, sampling). Rows per entity on a shared time axis, positions in percent of the axis, a playhead line moved by transform, and a simulated clock mapped onto the timeline (with holds at decisions). The `wires` and `camera` skeletons do not apply.

State the choice and the reason in the brief line.

## The camera, in short

The diagram lives in a fixed-size "world" inside an `overflow: hidden` viewport. The camera is a transform on the world (`x`, `y`, `scale`, origin `0 0`). For each step, compute the bounding box of the nodes involved and fit it in the viewport with padding, clamped between the fit-all scale and about 1.6 times it. Between steps the camera glides with `power3.inOut`. The last step pulls back to fit-all, so the viewer sees the whole system with everything they just watched.

On phones, framing two distant nodes makes both unreadable. When the viewport is narrow, frame only the node that changes in that step and let the caption name the other.

Pair it with depth of field: nodes outside the current focus get `blur(3px)` and lower opacity, and come back sharp on the pull-back. Put a soft vignette on the viewport edges.

## Rules for timelines

- One master `gsap.timeline({ paused: true })` with a label per step (`s0`, `s1` ...). All tweens go on it, so scrubbing and stepping always work, backwards too.
- Always give controls (see `player`): play and pause, previous and next step, a segmented progress bar (one segment per step, click or drag to seek), keyboard (space, arrows).
- Autoplay once, about one second after the entrance reveal.
- Rest state is readable: before the timeline runs, nodes, base connections and ghost bars are visible. Animation adds the accent layer on top.
- Positions come from layout, never hard-coded pixels. On resize, recompute and rebuild the timeline, keeping its progress. Rebuild only when the width changes, so a layout that sets the stage height does not loop.
- Text that changes over time (counters, status chips, step captions) is derived from the playhead in `sync()`, never set by `tl.call()`. Then scrubbing backwards and jumping between steps always show the right values.
- Nothing leaves residue. When an element leaves (flies away, fades out), its outline, placeholder or border leaves with it, or turns into an intentional end state (for example the window turning into a "kept" block). Check the last frame for orphan outlines.
- Marks sit exactly on their anchor: a mark at time t starts at t (left edge), not centered on it, so it never spills outside its window or row.
- Reset with `gsap.set(el, {clearProps:'...'})` only on elements whose position does not live in their inline style. `clearProps:'all'` wipes inline `left`/`width` and collapses anything positioned that way.
- `prefers-reduced-motion`: no autoplay, jump to `progress(1)`, no camera moves (fit-all), keep the controls.
