# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this repo is

A personal polyglot practice monorepo. Each subdirectory is an isolated learning area for a different language or technology. There is no shared build system across the whole repo.

## Languages and how to run them

### Go (`go/`)

The root `go/` module (`module practice`) has a Google Gemini AI dependency (`google.golang.org/genai`). Sub-packages (`hello/`, `problems/`, `leetcode/`, `cookbook/`, `tour/`, `agent/`) each have their own `go.mod` and are independent modules.

```bash
# Run a package
go run go/problems/main.go

# Run tests (from within the module directory)
cd go/problems && go test ./...

# Run a single test
cd go/problems && go test -run TestEnglishWordToInteger
```

### .NET (`dotnet/MyApp/`)

Targets .NET 10.0 with nullable and implicit usings enabled.

```bash
cd dotnet/MyApp && dotnet run
dotnet build dotnet/MyApp
```

### TypeScript (`typescript/`)

Uses Bun (there is a `bun.lock` at the repo root).

```bash
bun run typescript/index.ts
```

### Python (`python/`)

Plain Python, no virtualenv needed.

```bash
python python/main.py
```

### Assembly (`asm/`)

x86 NASM-style. Build manually:

```bash
nasm -f elf64 asm/hello.asm -o asm/hello.o && ld asm/hello.o -o asm/hello
```

## Daily exercises (`daily/`)

A daily TDD practice system for Go, Python, and TypeScript. The user names a concept; the
`/daily` slash command generates one ≤20-min exercise folder under `daily/exercises/<date>-<lang>-<slug>/`
containing a `SPEC.md` and failing tests (no solution, no hints). The system tracks progress and
spaced repetition in `daily/PROGRESS.md` — on each `/daily` run it first grades the previous
exercise (runs its tests, inspects the diff) and calibrates the next task's difficulty.

```bash
# Generate today's exercise (in Claude Code)
/daily go channels      # a named concept
/daily continue         # next piece of the current concept
/daily                  # recommend / resurface a due concept

# Run today's exercise tests (auto-detects go test / unittest / bun test)
./daily/run.sh
./daily/run.sh 2026-06-24   # a specific date
```

See `daily/README.md` for the full ritual. Command logic lives in `.claude/commands/daily.md`.

## Architecture notes

- **Go modules are not nested** — `go/go.mod` does not govern `go/hello/`, `go/problems/`, etc. Each subdirectory with its own `go.mod` is a fully independent module. `go run ./...` from `go/` will not work across them.
- **`go/api.go`** is the entry point for Gemini AI experiments; the client code is commented out pending an API key.
- **`go/agent/agent.go`** is a placeholder for building a Go-based AI agent.
- **`course22/`** is a separate git repository (fastai course notebooks) — it has its own `.git`. Changes inside it are not tracked by the outer repo.
- **`android/`** contains two separate React Native projects: a plain RN app (`MyApp`) and an Expo app (`MyExpoApp`).
- **`claude/`** is reserved for Claude Code practice work.
