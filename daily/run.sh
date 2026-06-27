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
# from the files present, and dispatches to the right test runner. Exits non-zero
# if any test fails, so RED/GREEN is unambiguous.

set -uo pipefail

# Resolve the directory this script lives in, so it works from anywhere.
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
EXERCISES_DIR="$SCRIPT_DIR/exercises"

selector="${1:-$(date +%F)}"

if [[ ! -d "$EXERCISES_DIR" ]]; then
  echo "No exercises directory yet ($EXERCISES_DIR)."
  echo "Run /daily <concept> to generate your first exercise."
  exit 1
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
  exit 1
fi

# Run one exercise folder; returns the runner's exit code.
run_one() {
  local folder="$1"
  local name
  name="$(basename "$folder")"
  echo "=============================================="
  echo "Running: $name"
  echo "=============================================="

  if compgen -G "$folder/go.mod" > /dev/null; then
    # -v lists each test (PASS/FAIL) by name, not just an overall ok/FAIL.
    ( cd "$folder" && go test -v ./... )
    return $?
  fi

  if compgen -G "$folder"/*_test.py > /dev/null || compgen -G "$folder"/test_*.py > /dev/null; then
    ( cd "$folder" && python -m unittest -v )
    return $?
  fi

  if compgen -G "$folder"/*.test.ts > /dev/null || compgen -G "$folder"/*.spec.ts > /dev/null; then
    bun test "$folder"
    return $?
  fi

  echo "Could not detect a known test setup (go.mod / *_test.py / *.test.ts) in $name."
  return 2
}

overall=0
for folder in "${matches[@]}"; do
  run_one "$folder" || overall=1
  echo
done

if [[ $overall -eq 0 ]]; then
  echo "GREEN — all tests passed."
else
  echo "RED — tests failing. Keep going."
fi
exit $overall
