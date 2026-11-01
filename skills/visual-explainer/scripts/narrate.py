# /// script
# requires-python = ">=3.10"
# dependencies = ["kokoro-onnx>=0.4", "soundfile"]
# ///
"""Turn narration lines into one audio file plus a cue list, offline.

    uv run narrate.py <lines.json> <out-dir> [--voice am_michael] [--say-voice Samantha] [--speed 1.0] [--gap 0.5] [--engine auto|kokoro|say]

lines.json is a list of {"step": "<timeline label>", "text": "<one sentence>"}.
Writes <out-dir>/narration.m4a and <out-dir>/cues.json ({"start", "step", "text"} per line,
start = seconds from the beginning). Prints a JSON summary. Needs ffmpeg and ffprobe.

Kokoro model files (about 350 MB) download once to ~/.cache/visual-explainer/kokoro.
--engine auto falls back to macOS `say` when Kokoro cannot run.
"""

import argparse
import json
import os
import shutil
import subprocess
import sys
import urllib.request
from pathlib import Path

MODELS = "https://github.com/thewh1teagle/kokoro-onnx/releases/download/model-files-v1.0"
CACHE = Path.home() / ".cache" / "visual-explainer" / "kokoro"
# Kokoro voice ids start with a language letter: am_michael is American English.
LANGS = {"a": "en-us", "b": "en-gb", "e": "es", "f": "fr-fr", "h": "hi", "i": "it", "j": "ja", "p": "pt-br", "z": "cmn"}


def run(*cmd):
    subprocess.run(cmd, check=True, capture_output=True)


def seconds(path):
    out = subprocess.check_output(["ffprobe", "-v", "error", "-show_entries", "format=duration", "-of", "csv=p=0", str(path)])
    return float(out)


def fetch(name):
    path = CACHE / name
    if not path.exists():
        CACHE.mkdir(parents=True, exist_ok=True)
        print(f"downloading {name} (once)", file=sys.stderr)
        part = path.with_suffix(".part")
        urllib.request.urlretrieve(f"{MODELS}/{name}", part)
        part.rename(path)
    return str(path)


def speak_kokoro(lines, out, voice, speed):
    import soundfile
    from kokoro_onnx import Kokoro

    kokoro = Kokoro(fetch("kokoro-v1.0.onnx"), fetch("voices-v1.0.bin"))
    lang = LANGS.get(voice[0], "en-us")
    files = []
    for index, line in enumerate(lines):
        samples, rate = kokoro.create(line["text"], voice=voice, speed=speed, lang=lang)
        files.append(out / f"raw-{index}.wav")
        soundfile.write(files[-1], samples, rate)
    return files


def speak_say(lines, out, voice):
    files = []
    for index, line in enumerate(lines):
        files.append(out / f"raw-{index}.aiff")
        run("say", *(["-v", voice] if voice else []), "-o", str(files[-1]), line["text"])
    return files


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("lines")
    parser.add_argument("out")
    parser.add_argument("--voice", default="am_michael")
    parser.add_argument("--say-voice", help="macOS voice for the say engine; default is the system voice")
    parser.add_argument("--speed", type=float, default=1.0)
    parser.add_argument("--gap", type=float, default=0.5, help="silence after each line, seconds")
    parser.add_argument("--engine", choices=["auto", "kokoro", "say"], default="auto")
    args = parser.parse_args()

    for tool in ("ffmpeg", "ffprobe"):
        if not shutil.which(tool):
            sys.exit(f"{tool} is required")
    lines = [line for line in json.load(open(args.lines)) if line.get("text", "").strip()]
    if not lines or not all(isinstance(line.get("step"), str) for line in lines):
        sys.exit('lines.json must be a non-empty list of {"step": str, "text": str}')
    out = Path(args.out)
    out.mkdir(parents=True, exist_ok=True)

    engine = args.engine
    if engine in ("auto", "kokoro"):
        try:
            raw = speak_kokoro(lines, out, args.voice, args.speed)
            engine = "kokoro"
        except Exception as error:
            if engine == "kokoro" or not shutil.which("say"):
                raise
            print(f"Kokoro failed ({error}); falling back to say", file=sys.stderr)
            engine = "say"
    if engine == "say":
        raw = speak_say(lines, out, args.say_voice)

    # Same loudness for every line, then a fixed gap, so cue times are exact sums.
    cues, elapsed, listing = [], 0.0, out / "list.txt"
    with open(listing, "w") as listing_file:
        for index, (line, source) in enumerate(zip(lines, raw)):
            norm = out / f"line-{index}.wav"
            run("ffmpeg", "-y", "-i", str(source), "-af", f"loudnorm=I=-20:TP=-2:LRA=7,afade=t=in:d=0.02,apad=pad_dur={args.gap}",
                "-ar", "24000", "-ac", "1", str(norm))
            cues.append({"start": round(elapsed, 2), "step": line["step"], "text": line["text"]})
            elapsed += seconds(norm)
            listing_file.write(f"file '{norm.resolve()}'\n")
    audio = out / "narration.m4a"
    run("ffmpeg", "-y", "-f", "concat", "-safe", "0", "-i", str(listing), "-c:a", "aac", "-b:a", "64k", "-ac", "1", str(audio))
    for path in [*raw, listing, *out.glob("line-*.wav")]:
        os.remove(path)
    json.dump(cues, open(out / "cues.json", "w"), indent=1)
    print(json.dumps({"engine": engine, "voice": args.voice if engine == "kokoro" else (args.say_voice or "system"), "seconds": round(elapsed, 1),
                      "bytes": audio.stat().st_size, "lines": len(cues), "audio": str(audio), "cues": str(out / "cues.json")}))


if __name__ == "__main__":
    main()
