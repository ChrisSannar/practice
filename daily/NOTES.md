# Notes

## Focus

- Work on the skills needed to understand and solve most coding chalenges: What algorithm or data structure is needed. The "flowchart" of what to use and how to use it

## Things to Change

- Each daily should have at least 3 exercises to practice on

## Things to Study

- Learn how golang packages work

## Done

- **Error-handle test runs (2026-06-27)** — `run.sh` now distinguishes a *tooling/setup error*
  (missing `go`/`python`/`bun`, or an unknown test layout) from a genuine *test failure*.
  Exit codes: `0` GREEN (passed), `1` RED (ran but failed), `2` ERROR (couldn't run). Missing
  tools are reported with a clear message instead of a raw "command not found".
- **Run a date other than today (2026-06-27)** — `run.sh` already accepts a date/folder argument
  (`./daily/run.sh 2026-06-24`). The `/daily` skill now also honors an explicit `YYYY-MM-DD`
  token in its argument, so you can generate or grade an exercise for a non-today date.
