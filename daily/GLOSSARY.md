# Prefix Sum Glossary

The vocabulary of the prefix sum pattern, as used across lessons and reference documents in this workspace.

## Terms

**Prefix sum array**:
An array `pre` where `pre[0] = 0` and `pre[i] = pre[i-1] + a[i-1]` for `i = 1..len(a)`. `pre[i]` is the sum of the first `i` elements of `a`. Length is `len(a) + 1`.
_Avoid_: cumulative array, running total array

**Running sum**:
The prefix sum at the current position during a single left-to-right pass — `running += a[i]` at each step. Conceptually the same as a prefix sum, but maintained as a scalar instead of stored in an array.
_Avoid_: current sum, cumulative sum

**Range sum**:
The sum of a contiguous subarray `a[i..j]`, computed in O(1) as `pre[j+1] - pre[i]`.
_Avoid_: subarray total, slice sum

**Seen-set**:
A hash set of prefix-sum (running-sum) values encountered so far during a single pass. Used to check whether `running - target` has occurred before. Initialized with `{0}` to account for the empty prefix.
_Avoid_: visited set, prefix cache

**Seen-map**:
A hash map from prefix-sum values to how many times each has occurred. The counting variant of a seen-set: instead of a boolean check, `counts[running - target]` is added to the result. Initialized with `{0: 1}`.
_Avoid_: frequency map, count map

**Subarray sum detection**:
Determining whether any contiguous subarray of `a` sums to `target` in O(n) time, using a running sum and a seen-set. The core insight: `pre[j+1] - pre[i] = target` means `pre[i] = pre[j+1] - target`, so at each step you ask "have I seen `running - target` before?"
_Avoid_: subset sum, subarray match

**Target**:
The integer sum that we are searching for among all contiguous subarrays of `a`.
_Avoid_: k, goal sum