# Stack: Artifact or single HTML file

Use for claude.ai Artifacts and any standalone HTML. Artifact CSP rules:

- Scripts only from `cdn.jsdelivr.net/npm/`, `cdnjs.cloudflare.com`, `unpkg.com`, `cdn.tailwindcss.com`.
- Stylesheets only from Google Fonts. So Basecoat, DaisyUI or any component CSS cannot be linked. Use the component layer below.
- When publishing as an Artifact, do not write `<!doctype>`, `<html>`, `<head>` or `<body>`. Start with `<title>`.

## Head (pinned, copy verbatim, drop the plugins the page does not use)

```html
<title>Two To Four Words</title>
<link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Geist:wght@400;500;600;700&family=Geist+Mono:wght@400;500&display=swap">
<script src="https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4.3.3/dist/index.global.js"></script>
<script src="https://cdn.jsdelivr.net/npm/gsap@3.15.0/dist/gsap.min.js"></script>
<script src="https://cdn.jsdelivr.net/npm/gsap@3.15.0/dist/DrawSVGPlugin.min.js"></script>
<script src="https://cdn.jsdelivr.net/npm/gsap@3.15.0/dist/MotionPathPlugin.min.js"></script>
<script src="https://cdn.jsdelivr.net/npm/gsap@3.15.0/dist/SplitText.min.js"></script>
<script src="https://cdn.jsdelivr.net/npm/gsap@3.15.0/dist/Flip.min.js"></script>
<script src="https://cdn.jsdelivr.net/npm/gsap@3.15.0/dist/ScrollTrigger.min.js"></script>
<script src="https://cdn.jsdelivr.net/npm/lucide@1.52.0/dist/umd/lucide.min.js"></script>
<!-- optional, UI state (tabs, toggles) without hand-written JS -->
<script defer src="https://cdn.jsdelivr.net/npm/alpinejs@3.17.4/dist/cdn.min.js"></script>
```

The font pair above is the default. `design-direction` may swap it for the chosen direction.

## Tokens (replace values per direction, keep the shape)

Every color is a token. Light on bare `:root`, dark redefined twice so both the OS setting and an explicit `data-theme` work.

```html
<style>
/* default palette, see design-direction. Replace with the user's palette when given. */
:root{--bg:oklch(0.985 0.003 247);--card:oklch(1 0 0);--fg:oklch(0.22 0.02 255);--muted:oklch(0.96 0.005 255);
--muted-fg:oklch(0.5 0.02 255);--border:oklch(0.9 0.008 255);--neutral:oklch(0.62 0.02 255);
--accent:oklch(0.55 0.18 258);--accent-soft:oklch(0.55 0.18 258/.14);--ok:oklch(0.64 0.14 165);--error:oklch(0.5 0.2 25);
--grid:oklch(0.22 0.02 255/.05)}
@media (prefers-color-scheme:dark){:root:not([data-theme="light"]){/* DARK */}}
:root[data-theme="dark"]{/* DARK, same values */}
/* DARK = --bg:oklch(0.17 0.015 255);--card:oklch(0.21 0.017 255);--fg:oklch(0.95 0.006 255);--muted:oklch(0.25 0.018 255);
   --muted-fg:oklch(0.7 0.018 255);--border:oklch(0.31 0.02 255);--accent:oklch(0.72 0.15 252);--accent-soft:oklch(0.72 0.15 252/.18);
   --ok:oklch(0.82 0.14 165);--error:oklch(0.62 0.2 22);--grid:oklch(1 0 0/.045);color-scheme:dark */
body{background:var(--bg);color:var(--fg);font-family:Geist,ui-sans-serif,system-ui,sans-serif}
</style>
```

A dark-only direction puts the dark values on bare `:root` with `color-scheme:dark` and drops the media and `data-theme` blocks.

## Component layer (copy verbatim, then use only classes)

```html
<style type="text/tailwindcss">
@theme inline{--color-bg:var(--bg);--color-fg:var(--fg);--color-card:var(--card);--color-muted:var(--muted);
--color-muted-fg:var(--muted-fg);--color-border:var(--border);--color-accent:var(--accent);
--color-neutral:var(--neutral);--color-ok:var(--ok);--color-error:var(--error);--font-mono:"Geist Mono",ui-monospace,monospace}
@layer components{
.card{@apply rounded-xl border border-border bg-card shadow-[0_1px_2px_rgb(0_0_0/.04)]}
.btn{@apply inline-flex items-center gap-1.5 h-8 px-3 rounded-md text-[13px] font-medium border border-border bg-card hover:bg-muted transition-colors focus-visible:outline-2 focus-visible:outline-accent}
.btn-primary{@apply bg-fg text-bg border-fg hover:opacity-90}
.badge{@apply inline-flex items-center gap-1 rounded-full px-2 py-0.5 text-[11px] font-mono border border-border text-muted-fg}
.eyebrow{@apply text-[11px] uppercase tracking-[.12em] text-muted-fg font-medium}
/* player controls, see motion-choreography */
.ib{@apply size-8 grid place-items-center rounded-full text-muted-fg hover:text-fg hover:bg-card transition-colors focus-visible:outline-2 focus-visible:outline-accent}
.ib svg{width:15px;height:15px}
.ib-main{@apply size-9 bg-fg text-bg hover:text-bg hover:bg-fg hover:scale-105 shadow-sm transition-transform}
.ib-main svg{fill:currentColor;stroke-width:0}
.seg{@apply relative flex-1 h-1 rounded-full bg-border overflow-hidden transition-[height] duration-200}
.seg i{position:absolute;inset:0;transform-origin:left;transform:scaleX(0);border-radius:inherit;background:var(--accent)}
}
</style>
```

## Code blocks (syntax highlight)

Every code, config or command block is highlighted. Never hand-color code with `<span>`s. Prism, loaded from cdnjs, with only the languages the page uses:

```html
<script src="https://cdnjs.cloudflare.com/ajax/libs/prism/1.30.0/components/prism-core.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/prism/1.30.0/components/prism-yaml.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/prism/1.30.0/components/prism-bash.min.js"></script>
<!-- others: json, go, python, typescript (needs javascript first), sql, promql, docker, diff -->
```

Markup: `<pre class="code"><code class="language-yaml">...</code></pre>`. Write the code as plain escaped text. After the markup is in the DOM, call `Prism.highlightAll()`. Commands go in their own block with `language-bash`, not mixed into the YAML.

Prism's theme stylesheets cannot be linked, and they would not match anyway. Use this syntax theme, built from the palette (blue and grey only, so green and red stay free for good and bad). Contrast checked at 4.5:1 or better on `--card`:

```css
:root{--code-key:var(--accent);--code-str:oklch(0.47 0.09 225);--code-num:oklch(0.45 0.15 280);--code-com:oklch(0.55 0.02 255);--code-punc:oklch(0.55 0.02 255)}
/* dark blocks: --code-str:oklch(0.84 0.07 220);--code-num:oklch(0.78 0.11 285);--code-com:oklch(0.64 0.02 255);--code-punc:oklch(0.62 0.02 255) */
.code{font-family:var(--font-mono,ui-monospace,monospace);font-size:12.5px;line-height:1.65;color:var(--fg);tab-size:2;overflow-x:auto}
.token.comment,.token.prolog{color:var(--code-com);font-style:italic}
.token.key,.token.atrule,.token.property,.token.keyword,.token.function,.token.tag{color:var(--code-key)}
.token.string,.token.char,.token.attr-value{color:var(--code-str)}
.token.number,.token.boolean,.token.constant,.token.important{color:var(--code-num)}
.token.punctuation,.token.operator{color:var(--code-punc)}
.token.variable,.token.parameter{color:var(--fg)}
.token.deleted{color:var(--error)} .token.inserted{color:var(--ok)}
```

To point at specific lines, add a short comment in the code (`# never fewer than 3 ready`) instead of custom highlighting markup. When the user's palette is used, derive `--code-*` from it the same way: keys from the accent, everything else in its neutrals.

## Tailwind v4 gotchas

- `@apply` only accepts Tailwind utilities. Applying one of your own component classes inside another (`.row{@apply small}`) throws and drops all Tailwind styles. Compose in markup instead (`class="row t-small"`).
- Prefix your own classes (`t-`, `c-`) so they do not collide with each other or with Tailwind names.

## Icons

`<i data-lucide="credit-card"></i>` then `lucide.createIcons()` after the markup is in the DOM. Call it again after replacing an icon.

## Layout guards

- Page body never scrolls sideways. Wide stages live in `overflow-hidden` (camera) or `overflow-x-auto` containers.
- Side gutter of at least 16px. Works at 400px width.
- `tabular-nums` on every changing number.
