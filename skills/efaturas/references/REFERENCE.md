# eFaturas Skill - Examples and Reference

This file contains examples and reference materials for the eFaturas skill.

## Sample Configuration

Use `assets/config.example.json` as a template, then save your runtime config to `config.json` in the skill root.

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
    "cae_hints": {
      "47111": "Despesas gerais familiares",
      "47112": "Despesas gerais familiares",
      "4771": "Despesas gerais familiares",
      "4773": "Despesas gerais familiares",
      "4751": "Saúde",
      "4752": "Saúde",
      "4763": "Educação",
      "4520": "Reparação automóveis",
      "7500": "Despesas Veterinárias",
      "9311": "Ginásios"
    }
  }
}
```

### Configuration Options Explained

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `url` | string | required | Portal das Finanças URL |
| `mode` | string | "dry-run" | "dry-run" (report only) or "live-run" (save) |
| `confidence_thresholds.auto_classify` | number | 80 | Minimum confidence to auto-apply |
| `confidence_thresholds.flag_review` | number | 60 | Threshold to flag for review |
| `reports.path` | string | "reports/" | Output directory for reports |
| `reports.include_low_confidence` | boolean | true | Include low-confidence rows |
| `learning.min_samples_for_confidence` | number | 3 | Samples required before scoring |
| `learning.max_history_per_merchant` | number | 10 | Max history entries per merchant |
| `agent_browser.headed` | boolean | false | Show browser window for debugging |
| `agent_browser.debug` | boolean | false | Enable verbose debug output |
| `agent_browser.trace_on_failure` | boolean | true | Capture trace file on errors |
| `agent_browser.profile_path` | string | null | Path to persistent browser profile |
| `cae_hints` | object | {} | Map CAE codes to expense categories |

## Merchant Profiles (Authoritative)

If a merchant NIF exists in `merchant_profiles.json`, that category is the source of truth. Do not override it with CAE hints or manual guesses.

Use `assets/merchant_profiles.sample.json` as a template.

## Sample JSON Report

Example output after running the skill:

```json
{
  "run_id": "2026-02-09T14:32:15Z",
  "mode": "dry-run",
  "period": "2026-01-01 to 2026-01-31",
  "summary": {
    "total_invoices": 12,
    "auto_classified": 9,
    "flagged_for_review": 2,
    "batch_applied": 1
  },
  "invoices": [
    {
      "merchant_nif": "514038942",
      "merchant_name": "irmãdona supermercados, unipessoal, lda.",
      "cae_code": null,
      "date": "2026-01-28",
      "amount": 24.99,
      "category": "Despesas gerais familiares",
      "confidence": 95,
      "classification_source": "history",
      "needs_review": false,
      "batch_applied": true
    },
    {
      "merchant_nif": "507513851",
      "merchant_name": "healthinvest braga exploracao de health clubs s.a.",
      "cae_code": "9311",
      "date": "2026-01-15",
      "amount": 35.00,
      "category": "Ginásios",
      "confidence": 95,
      "classification_source": "history",
      "needs_review": false,
      "batch_applied": false
    },
    {
      "merchant_nif": "504671529",
      "merchant_name": "ornimundo",
      "cae_code": null,
      "date": "2026-01-10",
      "amount": 15.50,
      "category": "Outro",
      "confidence": 0,
      "classification_source": "fallback",
      "needs_review": true,
      "batch_applied": false
    }
  ],
  "unknown_caes": ["85320", "62010"],
  "timestamp": "2026-02-09T14:35:22Z"
}
```

## Debug Artifact Checklist

When a failure occurs, collect these artifacts:

### 1. Snapshot Info
- **URL at failure**: `https://faturas.portaldasfinancas.gov.pt/...`
- **Step name**: e.g., "login", "click_submit", "classify_invoice"
- **Selector/ref that failed**: e.g., "@e5", "#btn-continuar"

### 2. Error Details
- **Error message**: Full text of the error
- **Last successful action**: What worked right before the failure
- **Timestamp**: ISO timestamp in UTC

### 3. Screenshots & Traces
- **Screenshot file**: `screenshots/efaturas_failure_2026-02-09T14:32:15Z.png`
- **Trace file**: `traces/efaturas_failure_2026-02-09T14:32:15Z.zip`

### 4. Console & Errors
- **Console output**: Last 50 lines from `agent-browser console`
- **Page errors**: Output from `agent-browser errors`
- **Network errors**: Any failed network requests

### 5. State
- **Current page title**: From `agent-browser get title`
- **Current URL**: From `agent-browser get url`
- **Last snapshot**: Output from `agent-browser snapshot -i`

## Debug Commands Reference

```bash
# Open in headed mode for debugging
agent-browser open https://faturas.portaldasfinancas.gov.pt/home.action --headed --debug

# Start trace recording
agent-browser trace start /tmp/efaturas_trace.zip

# Take screenshot
agent-browser screenshot /tmp/efaturas_state.png

# Get console logs
agent-browser console

# Get page errors
agent-browser errors

# Get current state
agent-browser snapshot -i

# Stop trace
agent-browser trace stop /tmp/efaturas_trace.zip

# Close browser
agent-browser close
```

## Typical Debug Session

```bash
# 1. Enable debug in config
# Set "efatura.agent_browser.headed": true
# Set "efatura.agent_browser.debug": true

# 2. Run skill (you'll see browser window)
# 3. When failure occurs:
agent-browser trace start /tmp/failure.zip
agent-browser screenshot /tmp/failure.png
agent-browser console > /tmp/console.log
agent-browser errors > /tmp/errors.log
agent-browser snapshot -i > /tmp/snapshot.txt
agent-browser trace stop /tmp/failure.zip

# 4. Review artifacts and update config if needed
# 5. Retry with updated configuration
```

## Common Debug Scenarios

### Login Fails
- Check: NIF and password are correct
- Check: 2FA completed successfully
- Check: No CAPTCHA blocking access
- Action: Retry login or ask user to complete manually

### Selector Not Found
- Check: Portal UI changed (common)
- Check: Use `agent-browser snapshot -i` to find new refs
- Action: Update selectors in workflow

### Session Timeout
- Check: Long delay between actions
- Check: Portal logs out inactivity
- Action: Reduce delays or use persistent profile

### Batch Dialog Not Detected
- Check: "Todas" button text or selector
- Check: Dialog text pattern
- Action: Update detection logic in skill

## Operational Notes

- The portal may require an active agent-browser session; use `agent-browser connect 9222` before `open` if you see "Browser not launched".
- After clicking "Guardar", the list can refresh/reorder; always re-snapshot and re-extract rows before continuing.
- The "Aviso" batch dialog appears for same-merchant NIFs; always click "Todas".
- Selected category state is not reliably exposed in the DOM (no consistent `aria-pressed`/`.selected`), so rely on intended selection plus save.
- "Guardar" (or form submit) persists selections for the current page.

## File Locations

- **Config**: `config.json`
- **Merchant Profiles**: `merchant_profiles.json`
- **Config Template**: `assets/config.example.json`
- **Merchant Profiles Template**: `assets/merchant_profiles.sample.json`
- **Reports**: `reports/`

## Example Prompts

### Basic Usage
```
Classify my pending e-Fatura invoices for January 2026
```

```
Run eFaturas in dry-run mode and generate a report
```

### Debug Mode
```
Enable headed mode and debug e-Fatura classification, then generate a report
```

```
Investigate why e-Fatura login is failing and collect debug artifacts
```

### Live Run
```
Classify and save all pending e-Fatura invoices for this month
```

```
Apply batch classifications to all invoices from the same merchant
```

### Configuration
```
Update my eFaturas config to add CAE hints for category "Ginásios"
```
