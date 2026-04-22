# Explain / Review PR

> Analyze a GitHub Pull Request — explain it in plain English, do a critical code review, or both.

## Variables

| Variable         | Description                         | Example                        |
| ---------------- | ----------------------------------- | ------------------------------ |
| `{{MODE}}`       | What kind of analysis you want      | `explain`, `review`, or `both` |
| `{{PR_CONTENT}}` | The PR title, description, and diff | Paste from GitHub              |

## Modes

| Mode      | What it does                                                                  |
| --------- | ----------------------------------------------------------------------------- |
| `explain` | Plain English breakdown — problem, what was built, why it matters. No jargon. |
| `review`  | Critical technical review — bugs, security issues, missing tests, verdict.    |
| `both`    | Full explanation first, then full review.                                     |

## Prompt

```
You are analyzing a GitHub Pull Request. The requested mode is: {{MODE}}

Follow the instructions for the selected mode below. Only output the sections for the selected mode — skip the others entirely.

---

MODE: explain
If the mode is "explain" or "both", structure your response with these sections:

## The Problem
What gap or need existed before this PR? Why was it needed?

## What Was Built
What does the PR actually add or change? Explain it like you're describing it to a smart non-expert. No code blocks.

## Step by Step
Walk through the key changes in order, in plain English.

## Why It Matters
What does this unlock? What can you do now that you couldn't before?

---

MODE: review
If the mode is "review" or "both", structure your response with these sections:

## Summary
One paragraph: what this PR does and whether it looks correct at a high level.

## Issues Found
List any bugs, logic errors, edge cases, or missing error handling. Be specific.

## Security Concerns
Flag anything that could introduce a vulnerability.

## Code Quality
Note anything hard to follow, poorly named, or inconsistent with good practice.

## Missing Tests
What should be tested that isn't?

## Verdict
One of: Approve / Request Changes / Needs Discussion — and why in one sentence.

---

Here is the PR to analyze:

{{PR_CONTENT}}
```
