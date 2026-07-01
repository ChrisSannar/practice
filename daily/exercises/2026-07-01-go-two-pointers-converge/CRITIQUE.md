# Critique — Two pointers, the converge (PairWithTarget)

**Outcome:** passed (green). Verdict: **4/5** for the converge loop.

## What's right
The core idea is exactly on the mark:

```go
for left, right := 0, len(s)-1; left < right; {
	sum := s[left] + s[right]
	if sum == target {
		return []int{left, right}
	}
	if sum < target {
		left++
	} else {
		right--
	}
}
return []int{}
```

- **The three-way decision is correct and minimal** — `==` returns, `<` raises the low end, else
  lowers the high end. That's the whole pattern, and you moved *only one* pointer per step, which
  was the new idea for this rung.
- **`for left < right`** is the right boundary — equal indices would double-count one element, and
  crossing means no pair exists.
- **`return []int{}`** (not `nil`) after the loop — you internalized the non-nil empty contract.
- You also **caught a bug in my test** (the negatives case had two valid pairs; a converge returns
  the *outer* one it reaches first). Reading the algorithm against the data closely enough to spot
  that is exactly the interview skill this track is for.

## Sharpen
- **Drop the debug line.** `fmt.Println("PairWithTarget test")` and its `import "fmt"` are still in
  — that's what printed three times in the test run. Before a step is "done," the file should be
  just the solution.
- **Style nit (optional):** the two `if`s can read as a single `switch` if you like it, which some
  find clearer for a three-way branch:
  ```go
  switch {
  case sum == target:
      return []int{left, right}
  case sum < target:
      left++
  default:
      right--
  }
  ```
  Your `if`/`if`/`else` is equally fine — this is taste, not correctness.

## The pattern, banked
You now have the full two-pointer *converge* family: both pointers moving (palindrome/reverse) and
one pointer moving on a comparison (this). Next is a different two-pointer shape — **same-direction**
(both start at the left, a slow "write" pointer trailing a fast "read" pointer), which is how in-place
array edits like remove-duplicates and move-zeroes work. Same family, new motion.
</content>
