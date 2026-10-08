# Agent Skills

A collection of skills for AI coding agents. Skills are packaged instructions and scripts that extend agent capabilities.

Skills follow the [Agent Skills](https://agentskills.io/) format.

## Available Skills

### efaturas

Auto-classify Portugal e-Fatura invoices using agent-browser with configurable debugging.

See `skills/efaturas/references/REFERENCE.md` for detailed usage, examples, and configuration.

### testproof

Checks that tests earn their place. Gate mode reviews new or changed tests; audit mode blocks test deletions that lose coverage. Any language.

Rules come from measured data. See `evals/testproof/EVIDENCE.md` for the data and `evals/testproof/harbor/` for agent-agnostic evals.

### visual-explainer

Turns explanations into beautiful, animated, token-efficient HTML pages. Entry point that reads the brief, sets three dials (variance, motion, density) and pulls in the three skills below. Stack: Tailwind v4, GSAP 3.15, Lucide, Geist. See `skills/visual-explainer/references/` for the pinned stack per surface (Artifact or repo).

### design-direction

Picks a distinctive visual direction (Instrument, Cinematic, Night ops, Editorial, Blueprint), one bold move, tokens first, and a banned list of AI design tells.

### motion-choreography

Purposeful GSAP motion: camera with depth of field, traces along connections, linked views, staged reveals and a player with segmented progress. Tested skeletons in `skills/motion-choreography/references/skeletons.md`. Pairs with the official [gsap-skills](https://github.com/greensock/gsap-skills).

### visual-critique

One screenshot review pass (light, dark, phone) against the design and motion rules, plus refinement commands: bolder, quieter, polish, animate, surprise me.

## Installation

npx skills add nicolastakashi/skills

## Skill Structure

Each skill contains:

- `SKILL.md` - Instructions for the agent
- `references/` - Supporting documentation and examples (optional)

Evals live outside the skill folders, in `evals/<skill>/`, so they are not installed with the skill.

## License

MIT
