# Critique — 2026-06-25 Go basics

**Result:** all six tests green, including the Unicode `Reverse` and the empty-map
`WordCount` edge cases. Correct and readable across the board. Notes below are about
*idiom and sharpening*, not bugs.

## What was already idiomatic
- **`Max`** — textbook early-return; nothing to add.
- **`Sum`** — clean `range` accumulation with `_` for the unused index. 
- **`FizzBuzz`** — the `div3 := i%3 == 0` booleans plus `continue` read really well. Good structure.
- **`Reverse`** — correctly reached for `[]rune` instead of indexing bytes, which is exactly why `héllo` → `olléh` passed. Right instinct.

## Where to sharpen

1. **`WordCount` — the map zero-value trick (most valuable takeaway).**
   Reading a missing map key returns the value type's zero, so the existence check is
   unnecessary. This whole block:
   ```go
   _, mapVal := mappers[val]
   if mapVal {
       mappers[val]++
   } else {
       mappers[val] = 1
   }
   ```
   collapses to one line:
   ```go
   counts[val]++   // missing key reads as 0, then becomes 1
   ```
   Also: drop the leftover `// TODO`, the double parens in `strings.Fields((s))`, and prefer
   `counts := make(map[string]int)` over `var mappers map[string]int = make(...)`.

2. **`FizzBuzz` — `strconv` over `fmt`.**
   `fmt.Sprintf("%d", i)` works, but `strconv.Itoa(i)` is the direct tool for int→string —
   faster and lets you drop the `fmt` import entirely.

3. **`Reverse` — string building.**
   `result += string(runes[i])` re-allocates the whole string every iteration (O(n²)).
   Fine at this size, but the idiomatic pattern is to reverse in place:
   ```go
   for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
       runes[i], runes[j] = runes[j], runes[i]
   }
   return string(runes)
   ```
   Worth internalizing the two-index swap — it shows up everywhere.

4. **`CountVowels` — minor.**
   `strings.Contains(vowels, string(val))` allocates a string per rune. `strings.ContainsRune(vowels, val)`
   avoids that. A `switch val { case 'a','e',...: }` is also common.

## Verdict
Confidence: **4/5** for Go basics. You're clearly back in the flow — the gaps are idiom
(map zero-values, picking the right stdlib helper), not fundamentals. Tomorrow's `continue`
will go one rung up into slices/maps with a bit less scaffolding.
