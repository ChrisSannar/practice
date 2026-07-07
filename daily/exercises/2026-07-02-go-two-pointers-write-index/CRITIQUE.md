# Critique — write-index (2b)

All four green: `KeepPositives`, `RemoveElement`, `MoveZeroes`, `RemoveDuplicates`.

## What's solid

- **`KeepPositives`, `RemoveElement`, `MoveZeroes`** are the textbook idiom, done right: `write` means
  "next open slot," advances only on keep, `read` scans everything with `for read := range s`. Clean,
  no off-by-one, no wasted allocation. `MoveZeroes`'s two-phase pack-then-fill is exactly the intended shape.
- Good defensive `len(s) < 2` / `len(s) < 1` guards, though for these loops they're not strictly
  needed — `range` over an empty or single-element slice already does the right thing (0 or 1
  iterations, loop just doesn't fire). Harmless here, worth knowing you can drop them.

## One thing to notice — `RemoveDuplicates` uses a *different* pointer convention

Look closely: `write` here doesn't mean "next open slot" like in the other three. It means
**"index of the last value already kept."** That's why the comparison is `s[write] != s[read]`
(compare against the last kept value, not just test `s[read]` alone) and why you `write++` *before*
writing, and `return write + 1` (count = last index + 1) instead of `return write`.

Both conventions are legitimate and both show up in real code — but mixing them up mid-problem is
the classic write-index bug. The tell: if your keep-rule needs to ask "what did I keep last?", you're
in the `write`-as-last-kept-index convention; if it only asks "is this element good on its own?",
you're in the `write`-as-next-slot convention. You picked correctly for each function — that's the
sign this actually clicked, not just pattern-matched from the warm-up.

## Verdict

**Confidence: 4/5.** Ships correct, idiomatic code for both write-index conventions and articulates
(functionally, in the code) why `RemoveDuplicates` needs the other one. Two pointers (2a converge +
2b write-index) is solid — closing out the primitives here per your call to skip 2c and move on.
