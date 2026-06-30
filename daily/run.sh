#!/usr/bin/env bash
#
# Universal "run today's exercise" runner for the daily TDD practice system.
#
# Usage:
#   ./daily/run.sh            # run today's exercise (by date, e.g. 2026-06-25)
#   ./daily/run.sh 2026-06-24 # run the exercise for a specific date
#   ./daily/run.sh <folder>   # run a specific exercise folder name
#
# It finds the matching folder(s) under daily/exercises/, detects the language
# from the files present, and dispatches to the right test runner.
#
# Exit codes:
#   0  GREEN — all tests passed
#   1  RED   — tests ran but failed
#   2  ERROR — could not run (missing tool, no exercise, or unknown test setup)

set -uo pipefail

# Resolve the directory this script lives in, so it works from anywhere.
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
EXERCISES_DIR="$SCRIPT_DIR/exercises"

selector="${1:-$(date +%F)}"

if [[ ! -d "$EXERCISES_DIR" ]]; then
  echo "No exercises directory yet ($EXERCISES_DIR)."
  echo "Run /daily <concept> to generate your first exercise."
  exit 2
fi

# Collect matching folders: exact folder name, or any folder whose name starts
# with the selector (so a bare date matches "<date>-<lang>-<slug>").
matches=()
if [[ -d "$EXERCISES_DIR/$selector" ]]; then
  matches+=("$EXERCISES_DIR/$selector")
else
  while IFS= read -r dir; do
    matches+=("$dir")
  done < <(find "$EXERCISES_DIR" -mindepth 1 -maxdepth 1 -type d -name "${selector}*" | sort)
fi

if [[ ${#matches[@]} -eq 0 ]]; then
  echo "No exercise found for '$selector'."
  echo "Run /daily <concept> to generate one."
  exit 2
fi

# Fail early with a clear message if the toolchain a folder needs isn't installed.
# Go has no try/catch — we error-handle by checking exit codes, so do the same here.
require_tool() {
  if ! command -v "$1" > /dev/null 2>&1; then
    echo "ERROR: '$1' is not installed or not on PATH — cannot run $2."
    return 1
  fi
}

# Run one exercise folder.
# Exit codes: 0 = tests passed, 1 = tests ran but failed, 2 = could not run (tooling/setup).
run_one() {
  local folder="$1"
  local name status
  name="$(basename "$folder")"
  echo "=============================================="
  echo "Running: $name"
  echo "=============================================="

  if compgen -G "$folder/go.mod" > /dev/null; then
    require_tool go "$name" || return 2
    # -v lists each test (PASS/FAIL) by name, not just an overall ok/FAIL.
    ( cd "$folder" && go test -v ./... )
    status=$?
  elif compgen -G "$folder"/*_test.py > /dev/null || compgen -G "$folder"/test_*.py > /dev/null; then
    require_tool python "$name" || return 2
    ( cd "$folder" && python -m unittest -v )
    status=$?
  elif compgen -G "$folder"/*.test.ts > /dev/null || compgen -G "$folder"/*.spec.ts > /dev/null; then
    require_tool bun "$name" || return 2
    bun test "$folder"
    status=$?
  else
    echo "Could not detect a known test setup (go.mod / *_test.py / *.test.ts) in $name."
    return 2
  fi

  # A non-zero test command can mean a real failure OR a build/crash. go test, unittest
  # and bun test all surface that in the output above; we just classify the exit here.
  if [[ $status -ne 0 ]]; then
    echo "(exercise '$name' did not pass — exit $status)"
  fi
  return $status
}

failed=0   # at least one exercise ran but failed
errored=0  # at least one exercise could not be run at all
for folder in "${matches[@]}"; do
  run_one "$folder"
  case $? in
    0) ;;
    2) errored=1 ;;
    *) failed=1 ;;
  esac
  echo
done

if [[ $errored -eq 1 ]]; then
  echo "ERROR — one or more exercises could not be run (missing tool or no test setup). See above."
  exit 2
elif [[ $failed -eq 1 ]]; then
  echo "RED — tests failing. Keep going."
  exit 1
else
  echo "GREEN — all tests passed."
  exit 0
fi
