# Narration

A voice-over that drives the page's animation: the viewer watches and listens, and each line moves the stage to its step.

## Offer it, never add it on your own

Offer narration in one line at delivery only when both are true:

- The page has an animated stage with a storyboard (MOTION 6 or higher).
- Other people will watch it on their own time: a proposal, a review, a walkthrough for a team.

Never offer it for status trackers, reference pages, comparisons, or unattended sessions. Narration adds 0.5 MB of audio per minute, build time, and one more thing to update on every edit, and most viewers will never press play. The page must work fully without sound.

Example offer: "Want a narrated walkthrough of the stage? About 2 minutes, I'll show you the script first."

## Script

Write it from the storyboard, one line per step, before any audio exists. Show it to the user and get a yes: a line of text is cheap to change, regenerated audio and timing are not.

- One or two lines per step, one sentence each. Every step label (`s0`, `s1` ...) gets at least one line.
- The last line belongs to the last step, the pull-back to the whole system.
- Say what the viewer sees move, in the order it moves. Do not read the on-page text aloud.
- Spell out what a voice would misread: "P99" as "p ninety-nine", "k8s" as "Kubernetes", ids and versions as words.
- About 150 words per minute. A walkthrough over 3 minutes loses people; cut steps, not words per line.

Save it as `lines.json`: `[{"step": "s0", "text": "A request arrives at the load balancer."}, ...]`.

## Audio

```bash
uv run <skill-dir>/scripts/narrate.py lines.json narration/
```

- Kokoro (offline, Apache 2.0) with voice `am_michael` by default. The first run downloads about 350 MB of model files to `~/.cache/visual-explainer/kokoro`.
- Other Kokoro voices: the first letter is the language (`a` American English, `b` British English, `p` Brazilian Portuguese, `e` Spanish, `f` French). Match the page's language: for Portuguese use `pf_dora`, `pm_alex` or `pm_santa`.
- When Kokoro cannot run, it falls back to macOS `say` (`--engine say`, `--say-voice Samantha`).
- Output: `narration/narration.m4a` and `narration/cues.json` (`start`, `step`, `text` per line). Every line is loudness-matched and followed by a 0.5 s gap.

## Page

- Embed the audio as a `data:audio/mp4;base64,...` URI, because an Artifact cannot load files from other hosts. Insert it and the cues with a script after the page is written, never by hand.
- Wire it with the `narration` skeleton in `motion-choreography/references/skeletons.md`. The audio is the clock: each step plays at its own speed from its first line, then holds until the next step's line. Seeking, the step buttons, the progress bar and the transcript all move the audio.
- The current line shows as a subtitle under the stage, in an `aria-live` region. A folded transcript lists every line; a click jumps there.
- A speed button cycles 1×, 1.25×, 1.5× and 2×, and the page remembers the choice. Voice and animation speed up together.
- No autoplay: browsers block sound without a click.

## Check

You cannot hear the audio. Check what you can, then ask the user to listen once and report misread words.

- Every cue's `step` exists as a timeline label, and every label has a cue.
- Seek the audio to the middle of three lines: the stage shows that step and the subtitle shows that line.
- The last frame after the audio ends is the whole system.
- On an edit that changes a step, regenerate only if the script changed; the cues must match the timeline labels again.
