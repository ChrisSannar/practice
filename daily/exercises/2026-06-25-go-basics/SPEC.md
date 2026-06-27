# Go basics — warm-up

**Goal:** shake the rust off Go fundamentals by implementing six small functions until
the tests pass. No fancy patterns — just the everyday tools.

**Today's slice (~15 min):** implement the six stubs in `basics.go`:

| Function | Touches |
|----------|---------|
| `Max(a, b int) int` | conditionals |
| `Sum(nums []int) int` | slice iteration |
| `Reverse(s string) string` | strings & **runes** (Unicode-safe) |
| `CountVowels(s string) int` | loops, case handling |
| `WordCount(s string) map[string]int` | maps, splitting whitespace |
| `FizzBuzz(n int) []string` | loops, modulo, string conversion |

**Acceptance criteria:** `./daily/run.sh` is green — every case in `basics_test.go` passes.

Watch the edge cases the tests encode: empty inputs, an equal pair in `Max`, a multi-byte
rune in `Reverse` (`héllo` → `olléh`), blank-only input for `WordCount`, and `n == 0`
returning an empty (non-nil) slice.

**Larger arc:** Go fundamentals refresher. From here we can ladder into slices/maps in more
depth, then error handling, then structs & methods, then goroutines/channels. Run
`/daily continue` tomorrow to go one rung up.
