---
name: efaturas
description: Auto-classify Portugal e-Fatura invoices using agent-browser with configurable debugging
compatibility: Requires internet access and agent-browser access to Portal das Finanças.
allowed-tools: agent-browser
metadata:
  author: nicolastakashi
  version: "1.0"
  last_updated: "2026-02-10"
---

# eFaturas Skill

## What I Do

I automate the classification of invoices in Portugal's e-Fatura portal (Portal das Finanças) to maximize tax deductions. I use agent-browser to navigate the portal, fetch pending invoices, classify them based on merchant history and optional CAE hints, and generate JSON reports.

## When to Use Me

Use this when you need to:
- Fetch and classify pending e-Fatura invoices
- Apply batch classifications to reduce manual clicks
- Generate JSON reports for review or processing
- Debug portal navigation or classification issues

## Inputs

You must provide at runtime:
- Login credentials (NIF and password) - interactively, no credential manager
- Date range for invoice classification (optional, defaults to current period)
- Complete 2FA during login when prompted

## Strictness Contract

The agent MUST follow this skill exactly. Do not infer or invent steps.

- Do not guess categories. If no exact match exists, classify as `Outro` and set `needs_review: true`.
- Do not change category labels or casing. They must match the canonical list exactly.
- Do not override `merchant_profiles.json` categories with CAE hints or UI heuristics.
- Do not proceed with `live-run` unless the user explicitly requests it in the same session.
- If any required field is missing (NIF, merchant name, date, amount), flag the invoice for review and continue.
- If confidence is below `efatura.confidence_thresholds.auto_classify`, do not apply any classification in the UI.

## Configuration

All settings must be under `efatura.*` in `config.json`. Use `assets/config.example.json` as a template.

```json
{
  "efatura": {
    "url": "https://faturas.portaldasfinancas.gov.pt/home.action",
    "mode": "dry-run",
    "confidence_thresholds": {
      "auto_classify": 80,
      "flag_review": 60
    },
    "reports": {
      "path": "reports/",
      "include_low_confidence": true
    },
    "learning": {
      "min_samples_for_confidence": 3,
      "max_history_per_merchant": 10
    },
    "agent_browser": {
      "headed": false,
      "debug": false,
      "trace_on_failure": true,
      "profile_path": null
    },
    "cae_hints": {}
  }
}
```

- `url`: Portal URL
- `mode`: `dry-run` (report only) or `live-run` (save classifications)
- `confidence_thresholds.auto_classify`: Minimum confidence to auto-apply classification
- `confidence_thresholds.flag_review`: Threshold below which items are flagged for review
- `reports.path`: Output directory for reports (relative to skill root)
- `reports.include_low_confidence`: Include low-confidence rows in report output
- `learning.min_samples_for_confidence`: Samples required before scoring confidence
- `learning.max_history_per_merchant`: Max history entries to retain per merchant
- `agent_browser.headed`: Show browser window for debugging
- `agent_browser.debug`: Enable debug output
- `agent_browser.trace_on_failure`: Capture trace on error
- `agent_browser.profile_path`: Optional persistent profile path
- `cae_hints`: Map of CAE codes to categories (grows over time)

## Deterministic Classification Algorithm

Use this exact order and stop at the first match:

1. If `merchant_nif` exists in `merchant_profiles.json`, use `most_likely_category`.
2. Else if `cae_code` exists and `efatura.cae_hints` contains it, use that category.
3. Else use `Outro` and set `needs_review: true`.

Confidence rules:

- If step 1 is used, set confidence to the stored merchant confidence or 95 if missing.
- If step 2 is used, set confidence to 80.
- If step 3 is used, set confidence to 0.

Auto-apply rule:

- Only apply a classification in the UI if `confidence >= efatura.confidence_thresholds.auto_classify`.
- If below the threshold, leave the invoice unclassified, set `needs_review: true`, and include the candidate category in the report only.

## Skill Files

- `config.json`: Runtime configuration
- `merchant_profiles.json`: Authoritative merchant history and category mapping
- `assets/config.example.json`: Safe template configuration
- `assets/merchant_profiles.sample.json`: Minimal merchant profile template
- `references/REFERENCE.md`: Examples, debug recipes, and reference material

## Category Labels (Canonical)

Use exact category labels consistently across `merchant_profiles.json`, `cae_hints`, and report output. Do not change casing or accents.

- `Outro`
- `Despesas gerais familiares`
- `Despesas Veterinárias`
- `Alojamento, restauração e similares`
- `Ginásios`
- `Saúde`
- `Educação`
- `Reparação automóveis`

If the portal displays a different label, update all sources to match that exact UI label.

## Safety Gates

- Always run in `dry-run` unless the user explicitly requests `live-run` in this session.
- If a CAPTCHA appears, stop and ask the user to take over.
- If login fails after one retry, stop and report.
- If the pending invoices page is not reachable, stop and report.

## Workflow

1. Open portal: `agent-browser open <efatura.url>`
   - If `efatura.agent_browser.headed` is true, add `--headed` (otherwise omit it)
   - If `efatura.agent_browser.debug` is true, add `--debug` (otherwise omit it)
2. Login:
   - Click "Adquirente"
   - Click "NIF" tab
   - Fill NIF and password (user-provided)
   - Click "Continuar Login"
   - Wait for user to complete 2FA manually
3. Check for pending invoices banner
4. Click "Complementar Informação Faturas"
5. Navigate to pending invoices list
6. For each invoice:
    - Extract: merchant NIF, merchant name, CAE code (if available), date, amount
    - Resolve classification using: merchant_profiles.json (authoritative) > CAE hints > "Outro"
    - If merchant_profiles.json contains the merchant NIF, always use its most_likely_category and never override it with CAE hints or fallback
    - If confidence is below `efatura.confidence_thresholds.auto_classify`, do not apply any UI classification for this invoice
    - Apply batch "Todas" (All) when prompted for same NIF
    - Click "Guardar" (Save) button to persist categorization
    - Re-snapshot after save and verify the row still exists; if not, stop and report
7. Only click "Guardar" if at least one classification was applied on the current page
8. If zero classifications were applied on the current page, do not click "Guardar" and log the page as review-only
9. Generate JSON report to configured output path
10. Close browser

## Learnings

See `references/REFERENCE.md` for operational notes and portal behaviors.

## Refactor Opportunities

- Target rows by merchant NIF + invoice number instead of row index.
- Centralize classification resolution (history -> CAE hints -> fallback) to avoid hard-coded categories in actions.
- Add a pagination loop: process page -> save -> re-snapshot -> next page until total processed.
- Add selection verification fallback (e.g., compare row values before/after save) since DOM state is inconsistent.
- Wrap batch dialog handling in a helper that waits for "Aviso" then clicks "Todas".
- Capture structured debug artifacts on failure (URL, step, selector/ref, screenshot, trace).

## Batch Classification Rule

When the portal prompts: "Pretende atribuir o mesmo setor (X) às faturas da página com o NIF emitente Y?"

Always click "Todas" (All) button to apply classification to all invoices from the same merchant at once.

Detection: Look for dialog with heading "Aviso" containing text about "NIF emitente" and buttons "Todas" / "Apenas esta".

## Classification Order (Authoritative)

1. Merchant profiles are authoritative: if the merchant NIF exists in merchant_profiles.json, always use its most_likely_category
2. If no merchant profile exists, check efatura.cae_hints when CAE code is available
3. Fall back to "Outro" with review flag if no match

## Normalization Rules

- Normalize merchant NIF to digits-only (keep leading zeros)
- Normalize merchant name to lowercase and collapse whitespace before matching
- Preserve CAE codes as strings to avoid dropping leading zeros
- Treat category labels as case- and accent-sensitive exact matches
- Skip re-classification if the UI already shows the intended category

## Merchant Profiles Rules

- Treat merchant_profiles.json as the source of truth for any known merchant NIF
- Never override a merchant_profiles.json category with CAE hints or manual guessing
- If a known merchant has conflicting UI feedback, keep the merchant_profiles.json category and proceed

## Report Output

Generate JSON report with:

```json
{
  "run_id": "timestamp",
  "mode": "dry-run|live-run",
  "period": "date-range",
  "summary": {
    "total_invoices": 0,
    "auto_classified": 0,
    "flagged_for_review": 0,
    "batch_applied": 0
  },
  "invoices": [
    {
      "merchant_nif": "string",
      "merchant_name": "string",
       "cae_code": "string|null",
       "date": "ISO-date",
      "amount": 0.00,
      "category": "string",
      "confidence": 0,
      "classification_source": "history|cae_hints|fallback",
      "needs_review": false,
      "batch_applied": false
    }
  ],
  "unknown_caes": ["code1", "code2"],
  "timestamp": "ISO-timestamp"
}
```

## Debug Protocol

When `efatura.agent_browser.debug` is true or on failure:

1. Respect `efatura.agent_browser.headed` for whether the browser is visible
   - If you need to watch the UI to debug and `efatura.agent_browser.headed` is false, temporarily set it to true (or add `--headed` when opening)
2. Start trace: `agent-browser trace start <trace-path>`
3. Capture:
   - Current URL
   - Step name (e.g., "login", "navigate_to_invoices", "classify_invoice")
   - Selector or ref used (e.g., "@e2", "#submit-button")
   - Error text or message
   - Last successful action
   - Timestamp
   - Screenshot: `agent-browser screenshot <screenshot-path>`
   - Console logs: `agent-browser console`
   - Page errors: `agent-browser errors`
4. Stop trace: `agent-browser trace stop <trace-path>`
5. Include artifacts in report or separate debug log

Debug artifact checklist:
- URL at failure point
- Step name and description
- Selector/ref that failed
- Full error message
- Last successful action and state
- Screenshot file path
- Trace file path
- Console output (last 50 lines)
- Page errors (if any)
- Timestamp in UTC

## Common Failures

- 2FA requires manual intervention
- CAPTCHA blocks automation
- Session timeout
- Selector drift (portal changes)
- Network issues or portal downtime
- Login credential errors

On any failure:
- Log all debug artifacts
- Flag affected invoices for review
- Continue with next invoice if possible
- Stop and report if critical failure (login, navigation)

## Safe Stop Conditions

- User cancels
- Login fails after retry
- Portal is unavailable (500, timeout)
- Maximum retries exceeded
- Unknown critical error

## Examples

See `references/REFERENCE.md` for sample prompts, configuration examples, and report output.

## Common Edge Cases

See `references/REFERENCE.md` for common failures, selector drift handling, and debug workflows.
