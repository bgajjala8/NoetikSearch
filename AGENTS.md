# Prompt Library Format Spec

A simple, portable prompt format that works across VS Code, Cursor, Claude Code, and any other AI coding tool.

## Design Goals

- Plain markdown — no proprietary syntax
- Works by copy-paste, slash command, or file reference
- Variables use `{{VARIABLE_NAME}}` — universally understood by every LLM
- Modes are passed as a variable, not baked into conditional blocks
- No dependencies, no framework, no build step

## File Structure

Each prompt is a single `.md` file with three sections:

```
# Prompt Name

> One-line description of what this prompt does.

## Variables
List of variables the user must fill in before running.

## Modes (optional)
List of available modes and what each one does.

## Prompt
The actual prompt text sent to the model.
```

## Variables

Use `{{VARIABLE_NAME}}` anywhere in the prompt body. Variable names are SCREAMING_SNAKE_CASE.

Always document variables in the `## Variables` section so it's clear what needs to be filled in.

## Modes

Modes are just a value passed into `{{MODE}}`. The prompt body tells the model what to do for each mode. No conditionals, no tool-specific syntax — the LLM handles the branching.

## How to Use in Each Tool

| Tool                | How to use                                                                                |
| ------------------- | ----------------------------------------------------------------------------------------- |
| **VS Code Copilot** | Use the matching `.prompt.md` in `.github/prompts/` for a native slash command experience |
| **Cursor**          | Paste into Cursor chat with variables filled in, or add to `.cursorrules` as context      |
| **Claude Code**     | Reference the file directly or paste into the conversation                                |
| **Any tool**        | Copy the `## Prompt` section, fill in `{{VARIABLES}}`, paste and run                      |

## Example

See [`prompts/explain-pr.md`](prompts/explain-pr.md) for a real example.
