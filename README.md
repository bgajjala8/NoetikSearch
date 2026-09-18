# NoetikSearch

Prompts I reuse, kept as plain markdown so they work in whatever tool I'm using.

```
prompts/           the prompt itself
.github/prompts/   VS Code slash command version
AGENTS.md          notes for agents working in this repo
```

## The format

One markdown file per prompt:

```
# Prompt Name

> What it does, in one line.

## Variables
What you need to fill in.

## Modes (optional)
The modes it supports.

## Prompt
The text you actually send.
```

Variables look like `{{VARIABLE_NAME}}`, always caps. List every one you use under
`## Variables` so it's clear what has to be filled in before running it.

Modes are just a value passed in as `{{MODE}}`. The prompt body spells out what each
mode should do and the model handles the branching. No conditionals, no templating, no
tool-specific syntax.

## Using them

In VS Code, use the copy in `.github/prompts/` and it shows up as a slash command.
Anywhere else, open the file, copy the `## Prompt` section, fill in the variables and
paste it.

[`prompts/explain-pr.md`](prompts/explain-pr.md) is a working example.
