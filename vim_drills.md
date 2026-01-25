# Vim / Neovim — 2‑Week Daily Drill Plan

**Goal:** Reach comfortable, practical proficiency by building muscle memory.

**Rule:** No arrow keys. No mouse. Use Vim *only* for the drill target that day.

**For Proficiency:** Do the **In-File Exercises** in a regressive manner each day:
 - If you're on `Day 7` then you would practice `Day 7` exercises 4 times, `Day 6` exercises 3 times, `Day 5` 2 times, and `Day 4` one time.
 - This will help solidify the concepts in your mind
 - Since you'll be editing this file, consider copying it before beginning
 - Current Day: 6

---

## Day 0 - Exiting Vim
**Learn**
- Modes: Command
- `:`

**Drill**
- Enter/Exit Command mode
- Exit out of the editor

**In/Out-File Exercise**
1. Exit out of the editor using `:q` followed by `Enter`
2. Save and exit by using `:wq`

---

## Day 1 — Movement Basics
**Learn**
- `hjkl`
- `w b e`
- `0 ^ $`
- `{ }`
- `g`

**Drill**
- Navigate a file without editing
- Jump word-to-word and line ends only

**In-File Exercise**
1. Place cursor at top of this file using only `hjkl` (NO arrow keys)
2. Reach the word `Win` using only `w/b`
3. Go back to the word `Movement` using only `w/b`
4. Navigate to the word `Exercise` using only `e`
5. Jump to the beginning and end of each line in this list using `0`, `^`, and `$`
6. Jump to the top using using `{ }`
7. Jump to the end of this section using `{ }`
8. Return to the top with `gg`

**Win condition:** You stop thinking about movement keys

---

## Day 2 — Modes & Insert Control
**Learn**
- Modes: normal / insert
- `i a o O`
- `Esc` discipline

**Drill**
- Edit text using *only* `o/O/i/a`
- Exit insert mode immediately after edits

**In-File Exercise**
1. Add a new line *below* this sentence using `o`
2. Add a new line *above* using `O`
3. Insert a word in the middle of this sentence using `i`
4. Append punctuation to this sentence using `a`
5. Starting with the cursor at `5.`, add a message at the end of this line using `A`:
6. Starting at the end of this line, add a message before the `6.` using `I`

**Note:** Many developers prefer to keybind `jk` or something similar to exit modes quicker (it keeps your fingers on the home row).

---

## Day 3 — Delete, Change, Yank
**Learn**
- `d c y`
- `dd cc yy`
- `p P`

**Drill**
- Rewrite a paragraph using delete/change

**In-File Exercise**
1. Using `de`, cut the last word of this line.
2. Use `p` to put it back.
3. Using `db` cut the first word of this line then use `P` to put it back.
4. Yank this entire line using `yy`.
5. Paste it above this line using `p`.
6. Delete this entire line using `dd` then replace it with `p`.
7. Change a word in this sentence with `cw`.
8. Paste the word you changed using `p`: 

---

## Day 4 — Operators + Motions
**Learn**
- Operator + motion logic
- `dw`, `d$`, `c0`, `yG`
- `u`, `Ctrl-r`

**Drill**
- Never select text visually

**Practice Sentence**
I am a practice sentence.

**In-File Exercise**
1. Delete from "a" in the practice sentence to end of line using `d$`
2. Replace it with whatever you like.
3. From the middle of the practice sentence, change the first half using `c0`
4. Yank from the cursor to end of file using `yG`
5. Paste a duplicate of the file after `7.` using `p`
6. Undo the action with `u`
7. Redo the action with `Ctrl-r` (then undo it again)

---

## Day 5 — Text Objects (Huge Leap)
**Learn**
- `iw aw ip ap`
- `i" i' i(`

**Drill**
- Use `ciw`, `dap`, `yi"` repeatedly

**Practice Sentences**
I am a "testing" sentence.
I am a (particularly distinct) type of sentence.
I am the {final} sentence to [test] on.

**In-File Exercise**
1. Place cursor inside a word → run `ciw` and change it
2. Paste the word below using `p`
3. Place cursor inside "these quotation marks" → run `ci"`
4. Place the cut text above using `P`
5. Place cursor inside (these parentheses) → run `yi(`
6. Paste the "yanked" text under "**Practice Sentence**"
7. "Yank" the word {final} using `ya{` and paste it here: 
7. Place cursor inside this paragraph →  `dap`

**Bonus Exercise**: "Yank" 10 words, 5 lines, 3 quotes, and 2 paragraphs using `yw`, `yy`, `yi"`, and `yap`. Then paste them below

---

## Day 6 — Search
**Learn**
- `/ ? n N`
- `* #`

**Drill**
- Navigate files *only* via search

**In-File Exercise**
1. Search for the word `Exercise` using `/`, then press `Enter`
2. Jump forward and backwards with `n` / `N`
3. Escape your current search using `Esc`
4. Search again for `Exercise`, but use `?` this time (searches backwards)
5. Place cursor on a word and press `*`
6. Navigate word instances using `n` / `N`
7. Place cursor on a word and press `#` (searches backwards)

**Note:** Searching with `/` or `?` uses regex patterns while `*` and `#` only search for the selected word

---

## Day 7 — Replace & Repeat
**Learn**
- `:%s/foo/bar/gc`
- `.` repeat

**Drill**
- Make a change once, repeat everywhere

**In-File Exercise**
1. Change a word using `cw`
2. Move to another occurrence
3. Press `.` to repeat the change

---

## Day 8 — Visual Mode (Controlled)
**Learn**
- `v V Ctrl-v`
- `>` `<`

**Drill**
- Block edit columns
- Indent text blocks

**In-File Exercise**
1. Select multiple lines with `V`
2. Indent them using `>`
3. Enter block visual mode and edit a column

---

## Day 9 — Buffers & Files
**Learn**
- `:w :q :wq`
- `:ls :b :bd`

**Drill**
- Open multiple files, switch without tabs

**In-File Exercise**
1. Open another file with `:e`
2. List buffers with `:ls`
3. Switch buffers using `:b`

---

## Day 10 — Splits & Window Control
**Learn**
- `:sp :vsp`
- `Ctrl-w h j k l`

**Drill**
- Edit across 2–3 splits

**In-File Exercise**
1. Split this file horizontally
2. Move between splits using `Ctrl-w`
3. Close a split

---

## Day 11 — Registers
**Learn**
- `"0` yank register
- `"+` system clipboard

**Drill**
- Paste from specific registers intentionally

**In-File Exercise**
1. Yank a line
2. Paste from yank register using `"0p`
3. Copy text to system clipboard using `"+y`

---

## Day 12 — Macros
**Learn**
- `qa … q`
- `@a`, `@@`

**Drill**
- Automate repetitive edits

**In-File Exercise**
1. Record a macro that appends `;` to a line
2. Replay it on multiple lines

---

## Day 13 — Neovim Power
**Learn**
- LSP basics
- Telescope file search

**Drill**
- Navigate project without file tree

**In-File Exercise**
1. Use Telescope to open a file
2. Jump to a symbol using LSP

---

## Day 14 — Personalization
**Learn**
- Basic `init.lua`
- Keymaps

**Drill**
- Add 2 mappings you actually need

**In-File Exercise**
1. Open your `init.lua`
2. Add one normal-mode mapping
3. Reload config and test it

---

## After 2 Weeks
You should:
- Think in *motions*, not selections
- Rarely use visual mode
- Edit faster than typing

**Next upgrade paths:**
- Treesitter text objects
- Git plugins
- Custom motions

