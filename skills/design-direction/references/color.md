# Color

Read this whole file before choosing any color. It is the source of truth for palettes in `design-direction`.

## Ask, then assign by meaning

Color is a UX decision, not decoration. Never pick it silently from the direction table.

**a. Find or ask for a palette.** Look for one first: the project's tokens or theme file, a brand guide, CLAUDE.md, earlier messages. If there is none and the user is present, ask once before building (a multiple-choice question when the tool exists):
- "Use my brand palette" (they paste hex values or a link),
- two or three proposed palettes, each named and with a one-line rationale tied to the subject (for example "Signal: neutral slate, violet for the new version, green and amber for health").
If the user has no palette, or the session is unattended, use the default palette below and say so in one line.

**Default palette (no user palette).** Blue and grey carry the page; green and red are reserved for good and bad; amber and plum are for a third and fourth entity. Values were checked for WCAG contrast and for separation under protanopia and deuteranopia (green is lighter and slightly teal, red darker, so the pair differs in lightness, not only hue).

| Role | Light | Dark | Use |
|---|---|---|---|
| `--bg` | `oklch(0.985 0.003 247)` cool off-white | `oklch(0.17 0.015 255)` | page ground |
| `--card` | `oklch(1 0 0)` | `oklch(0.21 0.017 255)` | surfaces |
| `--fg` | `oklch(0.22 0.02 255)` | `oklch(0.95 0.006 255)` | text, 16:1 |
| `--muted-fg` | `oklch(0.5 0.02 255)` | `oklch(0.7 0.018 255)` | secondary text, 5.7:1 / 7:1 |
| `--border` | `oklch(0.9 0.008 255)` | `oklch(0.31 0.02 255)` | lines |
| `--neutral` | `oklch(0.53 0.02 255)` slate | `oklch(0.62 0.02 255)` | old, inactive, baseline, "the other one" |
| `--accent` | `oklch(0.55 0.18 258)` blue | `oklch(0.72 0.15 252)` | new, active, the thing to follow |
| `--ok` | `oklch(0.64 0.12 170)` | `oklch(0.82 0.14 165)` | good only (shapes; for text use `oklch(0.54 0.1 170)` in light) |
| `--error` | `oklch(0.5 0.2 25)` | `oklch(0.62 0.2 22)` | bad only |
| `--cat-1` | `oklch(0.62 0.13 75)` amber | `oklch(0.78 0.16 75)` | third entity, never a state |
| `--cat-2` | `oklch(0.42 0.1 315)` plum | `oklch(0.6 0.19 315)` | fourth entity, never a state |

Every pair in this table that the viewer must tell apart (accent, neutral, ok, error, cat-1, cat-2) keeps an OKLab distance of 10 or more, in normal vision and simulated protanopia and deuteranopia, in both themes, at 3:1 or better against `--bg`.

How to map with only these:
- Comparing two entities (old vs new, before vs after, A vs B): slate vs blue. A third entity is `--cat-1`, a fourth `--cat-2`. More than four: group them, or tell them apart by label and icon on one neutral, never by more hues or by steps of blue.
- `--cat-1` and `--cat-2` mark entities only. Amber is not a warning here, so always pair it with the entity's label.
- No section is one hue. If every element in a section is blue, or every element is grey, the color carries no meaning: keep the accent for the one thing to follow, give compared entities their entity colors, and leave the rest neutral.
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

