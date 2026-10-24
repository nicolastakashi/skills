#!/usr/bin/env python3
"""Prints one row per task for a Harbor job: mean of each metric, cost, and how
often the agent opened the reviewproof skill and /code-review. Usage: ./summary.py jobs/<job> ..."""
import json
import re
import sys
from collections import defaultdict
from pathlib import Path

SKILL_OPENED = re.compile(r'"name": ?"Skill".{0,200}"skill": ?"reviewproof"')
CODE_REVIEW = re.compile(r'"name": ?"Skill".{0,200}"skill": ?"code-review"')
COST = re.compile(r'"total_cost_usd":([0-9.]+)')
SHOWN = ["score", "evidence", "format"]


def trial_rows(job: Path):
    for trial in sorted(p for p in job.iterdir() if (p / "verifier").is_dir()):
        task = trial.name.split("__")[0]
        try:
            details = json.loads((trial / "verifier" / "details.json").read_text())
        except (OSError, ValueError):
            details = {"score": 0}
        log = "".join(p.read_text(errors="ignore") for p in (trial / "agent").rglob("*") if p.is_file())
        costs = COST.findall(log)
        yield task, details, float(costs[-1]) if costs else 0.0, bool(SKILL_OPENED.search(log)), bool(CODE_REVIEW.search(log))


def main() -> None:
    for job in map(Path, sys.argv[1:]):
        config = json.loads((job / "config.json").read_text())
        agents = config.get("agents") or [{}]
        print(f"\n{job}  model={agents[0].get('model_name')}  skills={agents[0].get('skills')}")
        by_task = defaultdict(list)
        for task, details, cost, opened, code_review in trial_rows(job):
            by_task[task].append((details, cost, opened, code_review))
        print(f"{'task':<24} {'n':>2} {'score':>6} {'evid':>6} {'format':>6} {'opened':>6} {'cr':>4} {'cost$':>7}  other")
        for task, rows in sorted(by_task.items()):
            mean = lambda key: sum(row[0].get(key, 0) for row in rows) / len(rows)
            other = sorted({k for row in rows for k, v in row[0].items() if isinstance(v, int)} - set(SHOWN) - {"posted"})
            print(f"{task:<24} {len(rows):>2} {mean('score'):>6.2f} {mean('evidence'):>6.2f} {mean('format'):>6.2f}"
                  f" {sum(row[2] for row in rows):>4}/{len(rows)} {sum(row[3] for row in rows):>4} {sum(row[1] for row in rows):>7.3f}  "
                  + " ".join(f"{k}={mean(k):.2f}" for k in other))


if __name__ == "__main__":
    main()
