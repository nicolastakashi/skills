# Agent Skills

Skills for AI coding agents, in the [Agent Skills](https://agentskills.io/) format. A skill is a set of instructions and reference files that an agent loads when the task matches.

## Install

```sh
npx skills add nicolastakashi/skills
npx skills add greensock/gsap-skills   # used by motion-choreography
```

## Skills

### Standalone

| Skill | Use it when | Details |
|---|---|---|
| `testproof` | You add, change, review, or delete tests. Gate mode reviews new tests; audit mode blocks deletions that lose coverage. Any language. | [`evals/testproof/EVIDENCE.md`](evals/testproof/EVIDENCE.md) |

### Visual bundle

`visual-explainer` is the entry point. It reads the brief, sets three dials (variance, motion, density), and runs the other three skills:

```
visual-explainer ━▶ design-direction ━▶ motion-choreography ━▶ visual-critique
```

Stack: Tailwind v4, GSAP 3.15, Lucide, Geist.

Install the bundle (the explainer needs all four):

```sh
npx skills add nicolastakashi/skills -s visual-explainer design-direction motion-choreography visual-critique
```

| Skill | Use it when | Details |
|---|---|---|
| `visual-explainer` | You want a concept, flow, or plan shown as an animated HTML page. | [`references/`](skills/visual-explainer/references/): pinned stack per surface (Artifact or repo) |
| `design-direction` | A page looks generic or "AI-made". Picks one visual direction and one bold move, and bans common AI design tells. | [`references/`](skills/design-direction/references/) |
| `motion-choreography` | A page needs motion that explains: camera moves, traces along connections, staged reveals, a step player. | [`references/skeletons.md`](skills/motion-choreography/references/skeletons.md). Pairs with [gsap-skills](https://github.com/greensock/gsap-skills). |
| `visual-critique` | A page is built and needs one review pass (light, dark, phone), or a refinement: bolder, quieter, polish, animate, surprise me. | — |

## Layout

```
skills/<skill>/
├── SKILL.md       instructions for the agent
└── references/    supporting docs and examples (optional)

evals/<skill>/     evals, not installed with the skill
```

## Evals

Evals live in `evals/<skill>/`, outside the skill folders, so they are not installed with the skill. `testproof` has agent-agnostic evals in [`evals/testproof/harbor/`](evals/testproof/harbor/).

## License

MIT
