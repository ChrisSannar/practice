# Daily Practice

A daily ritual for retaining and learning concepts in **Go, Python, and TypeScript** through
small test-driven exercises. You name a concept; the system hands you one bite-sized failing
test suite to make pass. You learn by doing — it does not teach or hand over solutions.

## The ritual

1. **Get a task.** Run the `/daily` slash command in Claude Code:
   - `/daily go channels` — a new exercise on a concept you name.
   - `/daily continue` — the next bite-sized piece of your current concept.
   - `/daily` — let it recommend: it resurfaces a concept that's due for review, or
     continues your active goal.
2. **See RED.** A new folder appears under `exercises/<date>-<lang>-<slug>/` with a
   `SPEC.md` and failing tests. Run:
   ```bash
   ./daily/run.sh
   ```
3. **Make it GREEN.** Edit the stub until `./daily/run.sh` passes. Aim for ≤20 minutes —
   if a concept is bigger, it's split across days.
4. **Next day**, run `/daily` again. It first checks how yesterday went (runs the tests,
   looks at your solution), writes a `CRITIQUE.md` into that exercise's folder (what was
   idiomatic, where to sharpen, a confidence verdict), records the outcome in `PROGRESS.md`,
   adjusts difficulty, and sets you up with the next task.

## Running tests

```bash
./daily/run.sh            # today's exercise
./daily/run.sh 2026-06-24 # a specific date
./daily/run.sh <folder>   # a specific exercise folder name
```

The runner detects the language from the files in the folder and dispatches to
`go test`, `python -m unittest`, or `bun test`. It exits non-zero while any test fails.

## How it remembers — `PROGRESS.md`

`PROGRESS.md` tracks every concept, your results, a confidence score, and a **next-review**
date for spaced repetition (pass → longer interval, fail → comes back tomorrow). It's plain
markdown — read it to see where you stand, or edit it by hand to retire/adjust a concept.

## Philosophy

- **You own the learning.** The system's only job is to manufacture well-scoped tasks and
  calibrate their difficulty from how your attempts go.
- **No solutions, no hints in the folder** — just a spec and failing tests. If you're truly
  stuck, ask in chat.
- **Small and daily** beats big and occasional.

## Layout

```
daily/
  README.md      ← this file
  PROGRESS.md    ← tracked log + spaced-repetition table
  run.sh         ← universal test runner
  exercises/
    <date>-<lang>-<slug>/   ← one folder per day
```
