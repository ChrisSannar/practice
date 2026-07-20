# Prefix sum foundations established

The user has demonstrated working knowledge of building a prefix sum array (`BuildPrefixSum`) and using
it for O(1) range-sum queries (`RangeSum` via `pre[j+1] - pre[i]`). Both functions are green in
`exercises/2026-07-07-go-prefix-sum/`. The mechanical foundation — the index shift, the identity
`sum(a[i..j]) = pre[j+1] - pre[i]` — is solid.

What was **not** yet demonstrated: the running-sum + seen-set technique for subarray sum detection
(`HasSubarraySum`, `CountSubarraysWithSum`). Those stubs remain unimplemented. This is the zone of
proximal development — the user has the building blocks but needs the algorithmic insight that
`pre[i] = pre[j+1] - target` turns the O(n²) pair search into a single-pass O(n) check.

**Prior knowledge also confirmed** (from `PROGRESS.md`): two pointers (converge and write-index) are
solid (confidence 4/5). Slice windowing (Take/Drop) is solid; Chunk was a sticking point but resolved.
Map-as-counter idioms are comfortable from earlier Go basics work.