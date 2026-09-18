# AGENTS.md

Instructions for AI coding agents working in this repo.

## What this repo is

A library of reusable prompts for AI coding tools. There is no application code, no
build, no test suite and no dependencies. Every change is markdown.

## Layout

```
prompts/                    portable version of each prompt
.github/prompts/            VS Code Copilot slash commands
README.md                   the format spec
```

Each prompt exists twice. `prompts/explain-pr.md` is the portable copy that works
anywhere by copy-paste. `.github/prompts/explain-pr.prompt.md` is the same prompt using
VS Code's native frontmatter and `${if input:...}` syntax. The base filename must match
across both.

## Adding or changing a prompt

Write the portable copy first, then port it to the VS Code version. If you change one,
change the other so they do not drift.

The portable copy follows the structure in README.md: an H1 name, a one-line blockquote
description, a `## Variables` section, an optional `## Modes` section, then `## Prompt`
with the prompt body in a fenced block.

Variables are `{{SCREAMING_SNAKE_CASE}}` and every one used in the body must be listed
in `## Variables`. Modes are passed as a `{{MODE}}` value and the prompt body explains
what each mode does. Do not use tool-specific syntax in the portable copy, that is what
the `.github/prompts/` version is for.

## Conventions

Keep prompts plain markdown. No templating engines, no build step, no scripts.

Prompt text should be direct and specific about the output it wants. If a prompt asks
for structured output, spell out the exact sections.

Recent commits use conventional commit prefixes, for example `feat: add prompt library`.

## Before you finish

There is nothing to run. Check that any prompt you touched still matches the format in
README.md, and that the portable and VS Code copies agree.
