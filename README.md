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

## Installation

npx skills add nicolastakashi/skills

## Skill Structure

Each skill contains:

- `SKILL.md` - Instructions for the agent
- `references/` - Supporting documentation and examples (optional)

Evals live outside the skill folders, in `evals/<skill>/`, so they are not installed with the skill.

## License

MIT
