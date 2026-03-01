---
name: interview
description: Conduct structured interviews to gather requirements and create ADR documents. Use when user wants to build a new feature, implement something, or requests "interview".
user-invokable: true
---

# Interview Workflow

When a user wants to build [something], conduct a structured interview.
Ask about technical implementation, UI/UX, edge cases, concerns, and tradeoffs.
Do not ask obvious questions. Dig into difficult areas the user may not have considered.

## Interview Stages

### Stage 1: Understand the Problem

Use AskUserQuestion to identify:
- What is the current situation?
- What problem are you trying to solve?
- How often does this problem occur and what is its impact?
- Why does this problem need to be solved now?

### Stage 2: Technical Requirements

- Are there performance requirements? (response time, throughput)
- Are there scalability considerations?
- Are there security requirements?
- What about compatibility with existing systems?
- Are there technical constraints?

### Stage 3: Business Requirements

- Who are the users?
- What are the success criteria?
- Is there a deadline?
- Are there cost constraints?
- Are there regulatory/compliance requirements?

### Stage 4: Edge Cases

- What about unexpected situations?
- How should errors be handled?
- Are there concurrency issues?
- How is data consistency ensured?
- What are the rollback scenarios?

### Stage 5: Tradeoffs

- Simplicity vs Flexibility
- Performance vs Cost
- Fast release vs Perfect implementation
- Consistency vs Availability

## Interview Tips

**Don't:**
- Ask obvious questions ("Do we need to store data?")
- Repeat questions the user already answered
- Ask too many questions at once

**Do:**
- Ask deep questions ("Why Redis instead of Memcached?")
- Request opinions on tradeoffs
- Ask specifically about edge cases
- Provide context for each question

## Deliverables

After completing the interview, perform the following:

### 1. Create ADR Draft

- **Template**: Reference `docs/adr/ADR-0000-template.md`
- **Guide**: Follow writing rules in `docs/adr/CLAUDE.md`
- **Filename**: `docs/adr/ADR-NNNN-title.md` format

### 2. Update docs/adr/README.md

Reference the "README Update" section in `docs/adr/CLAUDE.md` to add a new entry to the ADR list table.

### 3. Add PLANNED Markers to Architecture Documents

Add PLANNED blocks to relevant architecture documents (`docs/architecture/**/*.md`):

```markdown
> **PLANNED**: [Summary of changes] (See [ADR-NNNN](../adr/ADR-NNNN-title.md))
> - Detail 1
> - Detail 2
```

### 4. Create Implementation Plan

Call EnterPlanMode to create an implementation plan.

## Reference Files

| File | Purpose |
|:---|:---|
| `docs/adr/ADR-0000-template.md` | ADR template |
| `docs/adr/CLAUDE.md` | ADR writing guide |
| `docs/adr/README.md` | ADR list |
| `docs/architecture/README.md` | Architecture structure |
