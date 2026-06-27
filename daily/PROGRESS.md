# Daily Practice — Progress

This file is the memory of the daily TDD practice system. `/daily` reads it to decide
what to give you next and to space out reviews; it appends to it after each session.
You can edit it by hand any time (bump a confidence score, retire a concept, add a note).

## Spaced-repetition rule

Each concept has a **Next review** date. When you pass an exercise, its interval grows;
when you fail, it resets so the concept comes back fast.

| Outcome | Next interval |
|---------|---------------|
| pass    | 1 → 3 → 7 → 16 → 35 → 90 days (step up one level each pass) |
| fail / abandoned | reset to 1 day |

`Confidence` is a 1–5 self/observed rating (1 = shaky, 5 = solid). In **recommend** mode,
`/daily` prefers concepts whose `Next review` is on or before today (lowest confidence first).

## Concepts

<!-- One row per concept. /daily maintains this table. -->

| Concept | Lang | Times | Last result | Confidence | Last practiced | Next review |
|---------|------|-------|-------------|------------|----------------|-------------|
| Go basics (vars/slices/strings/maps/loops) | Go | 1 | passed | 4 | 2026-06-25 | 2026-06-28 |

## Log

<!-- Reverse-chronological. Newest entry on top. /daily prepends one entry per exercise. -->

### 2026-06-25 — Go basics (Go)
- Folder: `exercises/2026-06-25-go-basics/`
- Outcome: passed (all six green, including Unicode `Reverse` and empty-map `WordCount`)
- Critique: `exercises/2026-06-25-go-basics/CRITIQUE.md`
- Notes: Back in the flow cleanly. Gaps are idiom, not fundamentals — map zero-value
  increment (skip the existence check), `strconv.Itoa` over `fmt.Sprintf`, in-place
  two-index reverse, `strings.ContainsRune`. Next task (`continue`): deeper slices/maps,
  less scaffolding.

<!--
### YYYY-MM-DD — <concept> (<lang>)
- Folder: `exercises/YYYY-MM-DD-<lang>-<slug>/`
- Outcome: passed | failed | abandoned | in-progress
- Notes: <what went well / where you struggled / how the next task should adjust>
-->
