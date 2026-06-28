# Go slice windowing — recovery drill

**Goal:** rebuild the one skill that tripped up `Chunk` yesterday — **slice index arithmetic**:
the `s[i:j]` slice expression, clamping a bound so it never runs past the end, and stepping an
index by more than one.

**Today's slice (~15 min):** implement the three stubs in `windows.go`, in order:

| Function | What it drills |
|----------|----------------|
| `Take(s, n)` | one upper bound: `s[:something]`, clamped |
| `Drop(s, n)` | one lower bound: `s[something:]`, clamped |
| `Chunk(s, size)` | put it together: step the index by `size`, clamp the tail |

Do them top to bottom — `Take` and `Drop` are the two halves of the bound-clamping you need,
and `Chunk` is the same idea applied in a loop.

**Contracts (the tests enforce these):**
- `Take`: `n <= 0` → empty; `n >= len(s)` → all of `s`.
- `Drop`: `n <= 0` → all of `s`; `n >= len(s)` → empty.
- `Chunk`: last group may be shorter; `size <= 0` → empty `[][]int{}`; empty input → empty `[][]int{}`.

**Acceptance criteria:** `./daily/run.sh` is green.

Reminders from the critique (no full solution here — that's yours to write): `size` is the
*elements per group*, the result's outer length is the *number of groups*, and don't pre-size
the result — start empty and `append` each window. Empty results must be non-nil (`[]int{}` /
`[][]int{}`), which the tests check via `reflect.DeepEqual`.

**Larger arc:** once windowing is solid, `/daily continue` returns to the deferred next rung —
error handling & multiple returns — then structs & methods, then goroutines/channels.
