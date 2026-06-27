---
description: Generate today's bite-sized TDD practice exercise (Go/Python/TS) and track progress
argument-hint: "[concept | continue | (empty to recommend)]"
allowed-tools: Bash(./daily/run.sh:*), Bash(date:*), Bash(find:*), Bash(git diff:*), Bash(git status:*), Read, Write, Edit
---

You are running the **daily practice ritual** for this repo. The user owns the learning —
your job is ONLY to manufacture one well-scoped, ≤20-minute TDD exercise and to calibrate its
difficulty from how past attempts went. Do not teach the concept and do not write the solution.

Argument: `$ARGUMENTS`

Today's date: run `date +%F` and use that value as `<date>` everywhere below.

## Step 1 — Close out the previous exercise (always do this first)

1. Find the most recent exercise: `find daily/exercises -mindepth 1 -maxdepth 1 -type d | sort | tail -1`.
2. If one exists and `PROGRESS.md` shows it as `in-progress`:
   - Run its tests: `./daily/run.sh <its-date-prefix>`.
   - Look at the user's attempt: `git diff -- daily/exercises/<that-folder>/` and read the changed files.
   - Decide the outcome: **passed** (tests green), **failed** (red but attempted), or **abandoned** (untouched stub).
   - **Write a critique** to `CRITIQUE.md` in that exercise's folder (skip only if the stub was untouched).
     Read the user's actual solution and assess it: what they got right / already idiomatic, where to
     sharpen (idiom, stdlib choices, performance, edge cases), and a 1–5 confidence verdict for the
     concept. Be specific and show the better version inline. Critique the code, not the person; these
     are sharpening notes, not bug reports unless tests actually fail.
   - Update `PROGRESS.md`:
     - Set that log entry's Outcome and add a one-line note on how it went and how the next task should adjust.
     - Update the concept's row in the Concepts table: bump `Times`, set `Last result`, adjust `Confidence` (1–5),
       set `Last practiced` = its date, and compute `Next review` using the spaced-repetition rule
       (pass → step the interval up 1→3→7→16→35→90 days; fail/abandoned → reset to +1 day).
3. If there's no prior in-progress exercise, skip straight to Step 2.

## Step 2 — Decide what to give them today

Parse `$ARGUMENTS`:
- **A concept** (e.g. `go channels`, `python iterators`, `ts mapped types`): use it.
  Infer the language from the concept. If the concept doesn't imply one of Go/Python/TypeScript,
  ask the user which of the three to use.
- **`continue`**: take the most recent active concept from `PROGRESS.md` and give the next
  bite-sized piece in that arc.
- **Empty**: **recommend** — from the Concepts table pick a concept whose `Next review` is on or
  before today (lowest `Confidence` first). If none are due, continue the active goal. If there's
  no history at all, ask the user what concept they want to start with.

Only Go, Python, and TypeScript are in scope.

## Step 3 — Calibrate difficulty

Use the concept's history in `PROGRESS.md`:
- Passed easily last time → push harder: a new facet, an edge case, or fewer scaffolds.
- Failed or struggled → shrink it: isolate the exact sticking point in a simpler task.
- New concept → start with the smallest meaningful slice.
Keep the whole task completable in **≤20 minutes**. If the concept is bigger, scope today to one
slice and note the larger arc in the spec so `continue` can pick up the rest.

## Step 4 — Generate today's exercise folder

Create `daily/exercises/<date>-<lang>-<slug>/` with:

- **`SPEC.md`** — a short problem statement: the goal, the specific ≤20-min slice for today,
  concrete acceptance criteria, and a line naming the larger arc this ladders into. No hints,
  no solution outline.
- **A stub file** with `// TODO` / `# TODO` markers and the smallest signatures needed.
- **Failing tests** that encode the acceptance criteria. Tests MUST fail against the stub.

Per-language conventions (match existing repo style):

- **Go** — add a `go.mod` (`module daily/<date>-<lang>-<slug>`, `go 1.25.6`), a `<slug>.go`
  stub (package matching the test), and `<slug>_test.go` using stdlib `testing` (table-driven).
- **Python** — a `<slug>.py` stub and a `test_<slug>.py` using stdlib `unittest` (no third-party deps).
- **TypeScript** — a `<slug>.ts` stub and a `<slug>.test.ts` using `bun test`'s
  `import { test, expect } from "bun:test"`.

## Step 5 — Record and hand off

1. Prepend a new entry to the `PROGRESS.md` **Log** (Outcome: `in-progress`) and add/update the
   concept's row in the Concepts table (Last practiced = today, Next review = today, set a starting Confidence).
2. Tell the user, briefly: the concept, the folder path, and to run `./daily/run.sh` to see RED.
   Do not reveal how to implement it.

## Hard rules

- Never write or sketch the solution. Folders contain only the spec + failing tests.
- Tests must fail initially (true RED).
- ≤20 minutes; split bigger concepts across days.
- Only Go, Python, TypeScript.
