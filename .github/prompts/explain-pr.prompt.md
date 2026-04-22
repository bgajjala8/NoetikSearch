---
agent: ask
description: Explain or review a PR — pick a mode when prompted
inputs:
  - id: mode
    description: "How do you want to analyze this PR?"
    type: pickString
    options:
      - explain     # plain English summary, no jargon
      - review      # critical technical review, find issues
      - both        # explain first, then review
    default: explain
---

${if input:mode == "explain" || input:mode == "both"}
You are explaining a GitHub Pull Request to someone who may not be deeply familiar with the codebase. Use plain, clear language — no unnecessary jargon.

Structure your response like this:

## The Problem
What gap or need existed before this PR? Why was it needed?

## What Was Built
What does the PR actually add or change? Explain it like you're describing it to a smart non-expert.

## Step by Step
Walk through the key changes in order, in plain English (no code blocks).

## Why It Matters
What does this unlock? What can you do now that you couldn't before?
${end}

${if input:mode == "review" || input:mode == "both"}
You are doing a critical technical review of a GitHub Pull Request. Be thorough and direct.

Structure your response like this:

## Summary
One paragraph: what this PR does and whether it looks correct at a high level.

## Issues Found
List any bugs, logic errors, edge cases, or missing error handling. Be specific — call out file/function names where relevant.

## Security Concerns
Flag anything that could introduce a vulnerability (input validation, auth, data exposure, etc.).

## Code Quality
Note anything that's hard to follow, poorly named, over-engineered, or inconsistent with good practice.

## Missing Tests
What should be tested that isn't?

## Verdict
Approve / Request Changes / Needs Discussion — and why in one sentence.
${end}

---

Here is the PR to analyze:

$PR_DESCRIPTION
