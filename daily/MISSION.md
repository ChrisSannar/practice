# Mission: Coding Interview Patterns (Go)

## Why

To become fluent in the recurring algorithmic patterns behind most coding-interview problems, so an unseen
question decomposes into "this is pattern X" instead of a blank page. The immediate driver is the daily
TDD practice system in this folder (`INTERVIEW_PATTERNS.md`), which walks a 15-pattern syllabus in Go.

## Success looks like

- Encounter a subarray-sum, range-query, or "find a contiguous slice that..." problem and immediately reach
  for the prefix-sum + hash-map technique without hesitation.
- Implement `HasSubarraySum` and `CountSubarraysWithSum` (LeetCode 560) from memory in Go, correctly
  initializing the seen-set with `{0}` (or `{0: 1}` for counting).
- Explain *why* `running - target` in the seen-set means a valid subarray exists — in terms of the
  `pre[j+1] - pre[i]` identity already mastered.
- Eventually close out all 15 patterns in `INTERVIEW_PATTERNS.md` and tackle unseen LeetCode problems
  by pattern-matching them to the right technique.

## Constraints

- Primary language: Go (interview track). Python and TypeScript remain available for cross-language drills.
- ~15–20 minutes per exercise session; small and daily beats big and occasional.
- No solutions in exercise folders — only specs and failing tests. Learning happens by doing.
- Only Go, Python, and TypeScript are in scope for the daily practice system.

## Out of scope

- System design interview prep (different track).
- Competitive programming (Codeforces-style optimization) — the goal is interview fluency, not sport coding.
- Non-Go language deep dives unless a concept is better illustrated in another language.