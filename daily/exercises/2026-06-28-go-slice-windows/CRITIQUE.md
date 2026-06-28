# Critique — Go slice windowing (Take / Drop / Chunk)

**Outcome:** failed (Chunk red — `Take` and `Drop` green). Verdict: **3/5** for windowing.

## The win — `Take` and `Drop` are solid
Both were red in the previous exercise's territory; now both pass cleanly. Your bound-clamping
reads well:

```go
func Take(s []int, n int) []int {
	if len(s) <= n { return s }   // n past the end → all
	if n <= 0 { return []int{} }  // nothing
	return s[0:n]
}
```

That's the right shape: handle the two out-of-range ends first, then the normal slice. `Drop` is
the mirror image and you nailed it. The single-bound arithmetic that tripped you up before is now
in hand — that's the whole point of the drill, and it landed.

## The miss — `Chunk`, two separate bugs
Your attempt:

```go
result := [][]int{}
idx := 0
for idx+size >= len(s) {            // bug #1: condition inverted
	result = append(result, s[idx:idx+size])
	idx += size
}
result = append(result, s[idx:])    // bug #2: unconditional tail
```

**Bug #1 — the loop condition points the wrong way.** Ask: *"when should I take another full
window?"* Answer: while a full window still fits — i.e. while `idx+size <= len(s)`. You wrote `>=`,
so for `[1,2,3,4], size 2` the test is `2 >= 4` → false → the loop never runs, and for
`[1,2], size 5` it's `5 >= 2` → true → `s[0:5]` panics (out of range). The `>=`/`<=` flip *is* the
bug both times. The loop continues while there's **more** to do, not while you're **past** the end.

**Bug #2 — the unconditional tail append.** Even with the condition fixed, appending `s[idx:]` after
the loop adds a spurious empty group on exact multiples: `[1,2,3,4], size 2` would loop twice
(idx 0, 2), exit at idx 4, then append `s[4:]` = `[]` → `[[1 2] [3 4] []]`. The trailing partial
chunk isn't a special case to bolt on; it's just the last iteration with a shorter window.

## The cleaner shape — clamp inside one loop
Don't split "full windows" from "the tail." Walk the input in steps of `size`, and clamp each
window's end so the last one is naturally short:

```go
for i := 0; i < len(s); i += size {
	end := i + size
	if end > len(s) {
		end = len(s)
	}
	result = append(result, s[i:end])
}
```

One loop, one rule, no trailing special-case. This is the same `Take`-style "clamp the bound"
move you already got right — just applied once per step. Notice it reuses exactly the instinct
from `Take`: compute the end, clamp it, slice.

## Carry-forward
The residual gap is **loop-termination reasoning** (`<` vs `<=`, "continue while work remains"),
not slice syntax. That's the same muscle the next track exercises immediately — the two-pointer
converge loop (`for l < r { ... }`) is the very next step — so this reinforces directly rather than
needing a third Chunk re-drill. Keep the habit: before writing a loop, say out loud the condition
under which it should *keep going*.
</content>
