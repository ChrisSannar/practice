Proposed Days 13 & 14
Day 13 — Precision Navigation & Quick Edits
Focus: Commands that make you dramatically faster at jumping to exact spots and making small targeted changes (these are the ones people say "blew my mind" after basics).
Learn

f F t T ; ,     (find char forward/back, till char, repeat ; / reverse ,)
%               (jump to matching brace/paren/bracket — huge for code)
Ctrl-a / Ctrl-x (increment / decrement number under cursor)
m<letter> / '<letter> / ``<letter>`   (marks: set, jump line-wise, jump exact position)
zz / zt / zb    (center / top / bottom current line)

Drill

Navigate solely with f/t/% for 5–10 minutes
Find every { or ( in a code block using % and jump back/forth
Change numbers quickly (e.g., turn version 1.2.3 → 1.3.0 using increments)

In-File Exercises (add these below Day 12 or in a new "## Day 13" section)

Place cursor on a ( somewhere in this document → press % to jump to matching )
From any {final} or similar → % back and forth 5 times
Find the next Exercise using /Ex then tE (till capital E), then ; to repeat
Find every Day heading using fD then ; / ,
Put a mark ma at the top of this section, scroll away, then 'a and ``a` to return
Find a number (e.g., any count like 3w or 10G) → Ctrl-a it several times, then Ctrl-x back
Center 3 different lines using zz, then move them to top/bottom with zt zb

Win condition: You instinctively reach for f<char>, %, and Ctrl-a instead of / or visual mode for nearby targets.
Day 14 — Quick Refactors & Efficiency Polish
Focus: Combine everything into real "refactoring flows" + small conveniences that save dozens of keystrokes daily.
Learn

gU / gu / ~     (uppercase / lowercase / toggle case — motion or visual)
J               (join lines)
> > / < <       (indent/outdent current line — repeat with counts or .)
==              (auto-indent current line — if syntax is on)
gv              (re-select last visual selection)
Quick :g global command basics (e.g., :g/foo/d delete all lines matching foo)
Bonus intro: :sort (sort lines / visual block)
Quick reminder of :help navigation and K (man/help on word under cursor)

Drill

Reflow messy text/code using joins, case changes, indents
Use . obsessively after small changes

In-File Exercises

Select a messy paragraph → gUap to uppercase it, then guap back
Join 3–4 lines into one using J (then undo)
Indent this entire Day 14 block with >> repeated or visual >ap
Re-select a previous visual area with gv and change it
Use :g/Day/d (carefully!) to delete all day headings — then undo
Sort the list of exercises in Day 1 alphabetically (visual select lines → :sort)
Find any word under cursor → press K (if help is working) or * + quick edit
Mini-challenge: Change every "foo" to "FOO" in the practice sentences using a quick global or macro + .

Win condition: You feel comfortable doing small-to-medium refactors (rename case, re-indent block, join/split lines, quick deletes) in <10 keystrokes average.
Why these two days (and not others)?

These fill the most common "missing after basics" gaps mentioned in recent (2024–2026) discussions/cheatsheets: precision char-wise jumps (f/t), brace matching (%), number tweaking (Ctrl-a/x), case ops, and quick line/block manipulation.
They stay vanilla (no plugins) — perfect continuation of your current style.
They emphasize composition (operator + motion + count + repeat) rather than isolated commands.
They lead naturally into plugin-land later (e.g., vim-surround, commentary, leap/flash, treesitter textobjects) if the user wants more.

Optional small changes to earlier days (to prepare for 13–14)

Day 4 or 5: Sneak in f<char> / t<char> (very quick win)
Day 5: Add % right after text objects (pairs are related)
Day 7: Mention gU / gu briefly with . repeat examples
Day 8: Add gv as a tiny note ("re-select last visual with gv")

If you add these, Days 13–14 become pure practice + combination rather than brand-new concepts.
What do you think — does this feel like the right next step, or would you prefer to lean more into ex commands (:g, :v, :sort family), quickfix, or something else?
