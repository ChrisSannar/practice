# Critique — Sliding Window: the seen-set window (4b)

**Verdict: tests green, but there's a real correctness bug in tasks 2 and 3 that the test suite
didn't catch. Confidence: 2/5** — task 1 (the pure membership primitive) is solid; the window-driving
mechanic in tasks 2/3 isn't the actual shrink loop yet, it's a lookalike that happens to pass the
given cases.

## What's right

`FirstRepeatedIndex` is exactly the primitive it asks for: one map, one pass, early return on
membership. Clean.

## The bug

In `LongestUniqueLen` (and the copy-pasted `LengthOfLongestSubstring`), on hitting a duplicate you do:

```go
delete(uniqueSet, a[idx])   // remove the *incoming* value
leftIdx = idx               // jump left straight to here
```

But the spec's mechanic is: remove `a[left]` and advance `left`, **one element at a time, repeating
until the duplicate is cleared** — because the duplicate might not be the element currently at
`left`. Your version deletes the wrong element (the one that just came in, which you're about to
re-add anyway) and teleports `left` forward, silently dropping every element between the old `left`
and `idx` from the set without ever checking whether *they* were still legitimately in the window.

That's not a style nitpick — it produces wrong answers. Concretely:

```go
LongestUniqueLen([]int{0, 1, 2, 0, 3})
// window walk: 0,1,2 (set={0,1,2}), hit second 0 → your code jumps left to idx 3,
// discarding 1 and 2 from the set even though they're still inside [1,3]
// → returns 3
// correct answer is 4: "1,2,0,3" (b,c,a,d) is a valid 4-long unique run
```

Same bug, string form — this is the textbook LC3 case that exposes exactly this shortcut:

```go
LengthOfLongestSubstring("dvdf")
// your code returns 2
// correct answer is 3 ("vdf")
```

None of the provided test cases happened to have a *non-adjacent* repeat that needed elements
between the old and new `left` to survive — that's a gap in the tests I wrote for this day, not
just your bug. Worth flagging so future test sets for this pattern include a `"dvdf"`-shaped case.

## The fix

The actual shrink loop only ever touches `left`:

```go
for uniqueSet[a[idx]] {
    delete(uniqueSet, a[left])
    left++
}
uniqueSet[a[idx]] = true
if idx-left+1 > maxLen {
    maxLen = idx - left + 1
}
```

Notice this also kills the awkward post-loop special case
(`if rightIdx-leftIdx == len(uniqueSet) { return len(uniqueSet) }`) you needed to bolt on —
that hack exists because `maxLen` was only ever updated inside the "found a duplicate" branch,
so the trailing (often-largest) window never got recorded on its own. Updating `maxLen`
unconditionally every iteration removes the need for it entirely, and removes `rightIdx` as a
separate variable — `idx` already is the right edge.

## Smaller notes

- Leftover `fmt.Println("FirstRepeatedIndex")` (and the unused `"fmt"` import it requires) —
  debug print left in the committed function.
- `if alreadyAppeared { return idx } else { set[val] = true }` — drop the `else`; after a `return`
  it's dead weight in Go style.

## Next drill

Today's exercise re-isolates the shrink primitive directly (given a window and a right-edge value,
shrink `left` step-by-step until the duplicate clears) before re-attempting the two driving
functions with test cases shaped like `dvdf` that specifically break the jump-shortcut. 4c
(minimum window substring) waits until this one is airtight — it escalates the same shrink loop to
need-counts vs window-counts, so an unsound shrink here would just compound there.
