# Prefix Sum & Subarray Detection Resources

## Knowledge

- [Article: "Subarray with Given Sum — Handles Negative Numbers" — GeeksforGeeks](https://www.geeksforgeeks.org/dsa/find-subarray-with-given-sum-in-array-of-integers/)
  The canonical explanation of the prefix-sum + hash-map approach for subarray sum detection. Covers the
  naive O(n²) approach as motivation, then the O(n) approach with a clear dry-run. Use for: understanding
  why `curr_sum - target` in the hashmap means a valid subarray exists.

- [Article: "Subarray with Given Sum" — GeeksforGeeks](https://www.geeksforgeeks.org/find-subarray-with-given-sum/)
  Parent article covering the same problem for non-negative integers (sliding window approach), then
  linking to the hashing approach for the general case with negatives. Use for: comparing sliding window
  vs prefix-sum + hashing and knowing when each applies.

- [Problem: "Subarray Sum Equals K" (LeetCode 560)](https://leetcode.com/problems/subarray-sum-equals-k/)
  The definitive interview problem this pattern unlocks — counting subarrays that sum to `k`. The counting
  variant requires a frequency map instead of a set. Use for: practicing the full pattern on a real
  interview question.

## Wisdom (Communities)

- [r/leetcode](https://reddit.com/r/leetcode)
  Active subreddit for interview prep discussion, pattern breakdowns, and solution approaches. Use for:
  seeing how others frame pattern recognition, troubleshooting stuck approaches.

## Gaps

- No high-quality visual/animated step-through of the prefix-sum + hashing algorithm has been found yet.
  The lesson in this workspace (`lessons/0001-subarray-sum-detection.html`) includes an interactive
  visualizer built from scratch to fill this gap.