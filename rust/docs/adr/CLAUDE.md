# ADR Writing Guide (Claude Code)

This guide provides instructions for Claude Code when creating and modifying ADR documents.

## File Naming Convention

```
ADR-NNNN-kebab-title.md
```

- NNNN: 4-digit number (starting from 0001)
- kebab-title: decision title connected with hyphens (-)

Example: `ADR-0001-caching-strategy.md`

## Procedure for Creating a New ADR

1. Copy template: Copy [`ADR-0000-template.md`](./ADR-0000-template.md) (English) or [`ADR-0000-템플릿.md`](./ADR-0000-템플릿.md) (Korean).
2. Rename file: Rename to `ADR-NNNN-kebab-title.md` format.
3. Write content: Fill in content according to each section.
4. Set status: Set initial status to `📝 proposed`.
5. Update README: Add an entry to the ADR list table in `./README.md`.

## Document Structure

ADR documents consist of the following sections:

| Section | Required | Description |
|:---|:---:|:---|
| Info | Yes | Metadata: ID, status, proposer, decision makers, dates, milestone, etc. |
| Background | Yes | Situation, problem, requirements, constraints, decision motivation |
| Alternatives | Yes | Expected outcomes (positive/negative/risk) for each alternative considered |
| Tests | No | Verification items linked to POC |
| Decision | Yes | Selected option, expected outcomes, reason for selection, details |
| Implementation Status | No | Implementation progress linked to Actions |
| Outcomes and Impact | No | Actual results after completion compared to expectations |
| References | No | External reference documents and resources |
| Changelog | Yes | Document versions and change history |

## Status Values

### ADR Status

| Status | Emoji | Code Value |
|:---|:---:|:---|
| proposed | 📝 | `proposed` |
| under-review | 🔍 | `under-review` |
| approved | ✅ | `approved` |
| deprecated | ❌ | `deprecated` |
| superseded | 🔄 | `superseded` |
| to-be-retired | 🗑️ | `to-be-retired` |

### Implementation Status

| Status | Emoji | Code Value |
|:---|:---:|:---|
| pending | ⏳ | `pending` |
| in-progress | 🔄 | `in-progress` |
| done | ✅ | `done` |
| cancelled | ❌ | `cancelled` |
| blocked | 🚫 | `blocked` |

### Test Status

| Status | Emoji | Code Value |
|:---|:---:|:---|
| pending | ⏳ | `pending` |
| in-progress | 🔄 | `in-progress` |
| done | ✅ | `done` |
| cancelled | ❌ | `cancelled` |
| blocked | 🚫 | `blocked` |

## Writing Alternatives

Write the following three items for each alternative:

```markdown
### Alternative N: [Alternative Name]

#### Expected Positive Outcomes
- Expected benefits and advantages
- Performance, cost, maintainability improvements

#### Expected Negative Outcomes
- Expected drawbacks and trade-offs
- Additional costs, increased complexity, etc.

#### Expected Risk Factors
- Risks that may occur during implementation
- Consider mitigation measures
```

## Writing Tests

Write items requiring verification in table format in the Tests section:

```markdown
## Tests

| Test Item | Milestone | Status | Completed | POC |
|:---|:---|:---:|:---:|:---|
| Test 1 | Phase 1 | ✅ done | YYYY-MM-DD | [POC-0001-01](../poc/POC-0001-01-load-test.md) |
| Test 2 | Phase 2 | 🔄 in-progress | - | [POC-0001-02](../poc/POC-0001-02-benchmark-comparison.md) |

### Test Status Values

| Status | Emoji |
|:---|:---:|
| `pending` | ⏳ |
| `in-progress` | 🔄 |
| `done` | ✅ |
| `cancelled` | ❌ |
| `blocked` | 🚫 |
```

## Writing the Decision Section

### Selected Option

Write a description of the adopted solution along with the following items:

```markdown
### Selected Option
[Description of the adopted solution]

#### Expected Positive Outcomes
...

#### Expected Negative Outcomes
...

#### Expected Risk Factors
...
```

### Reason for Selection

Describe the specific reasons for choosing this option:

- Technical reason: technical advantages in performance, scalability, security
- Operational reason: operations team's tech stack, ease of maintenance
- Cost reason: license costs, infrastructure costs
- Time reason: faster implementation, market timing
- Organizational reason: team experience, training costs
- Strategic reason: long-term technical direction, standards compliance

## Implementation Status Table

```markdown
| Task | Milestone | Status | Completed | Action |
|:---|:---|:---:|:---:|:---|
| Task 1 | 2025-Q1 | ✅ done | YYYY-MM-DD | [ACT-0001-01](../actions/ACT-0001-01-implement-caching-layer.md) |
| Task 2 | v1.2.0 | 🔄 in-progress | - | [ACT-0001-02](../actions/ACT-0001-02-documentation-cleanup.md) |
```

## README Update

When adding a new ADR, add an entry to the ADR list table in `./README.md`:

```markdown
| [ADR-NNNN](./ADR-NNNN-kebab-title.md) | Title | 📝 proposed | YYYY-MM-DD | - |
```

## Changelog

Always record in the `Changelog` section when modifying an ADR document:

```markdown
| Version | Date | Changes | Author |
|:---:|:---:|:---|:---|
| 1.0 | YYYY-MM-DD | Initial draft | John Doe |
| 1.1 | YYYY-MM-DD | Status changed: proposed → approved | Jane Smith |
```

---

## Directory Structure

```
docs/adr/
├── README.md                # ADR list
├── CLAUDE.md                # This file (writing guide)
├── ADR-0000-template.md     # ADR template (English)
└── ADR-0000-템플릿.md        # ADR template (Korean)

# All ADRs are stored together in this directory
```

---

## Related Document Links

ADRs are managed in connection with Action and POC documents:

```
docs/
├── adr/                         # ADR documents
│   ├── README.md
│   ├── CLAUDE.md
│   ├── ADR-0000-template.md
│   └── ADR-0001-caching-strategy.md
│
├── actions/                     # Action documents (ADR implementation items)
│   ├── README.md
│   ├── CLAUDE.md
│   ├── ACT-0000-00-template.md
│   ├── ACT-0001-01-implement-caching-layer.md
│   └── ACT-0001-02-documentation-cleanup.md
│
├── poc/                         # POC (ADR verification and results)
│   ├── README.md
│   ├── POC-0000-00-template.md
│   ├── POC-0001-01-load-test.md
│   ├── logs/                    # POC logs
│   └── datasets/                # POC datasets
```

### ADR Relationship Diagram

```
┌─────────────────────────────────────────────────────────────────┐
│                        ADR                                      │
│              (Architecture decision record, approved)           │
└───────────────────────────┬─────────────────────────────────────┘
                            │
            ┌───────────────┴───────────────┐
            │                               │
            │ 1:N                           │ 1:N
            ▼                               ▼
┌───────────────────────┐       ┌───────────────────────┐
│      POC              │       │      Actions          │
│   (verification)      │       │   (implementation)    │
├───────────────────────┤       ├───────────────────────┤
│ • Load testing        │       │ • Feature impl        │
│ • Performance bench   │       │ • Bugfix             │
│ • Technical verify    │       │ • Refactor           │
└───────────────────────┘       ├───────────────────────┤
                                │ • GitHub repo         │
                                │ • Jira ticket        │
                                │ • GitHub PR          │
                                │                       │
                                │ (tests managed        │
                                │  in code)             │
                                │ • Unit tests          │
                                │ • Integration tests   │
                                │ • E2E tests           │
                                └───────────────────────┘
```

### Entity Relationships

| Relationship | Cardinality | Description |
|:---|:---:|:---|
| ADR → Actions | 1:N | One ADR can have multiple Actions |
| ADR → POC | 1:N | One ADR can have multiple POCs |
| Actions → POC | - | Actions and POCs are not directly linked |

Reference POCs in the Tests section and Actions in the Implementation Status section of the ADR.
