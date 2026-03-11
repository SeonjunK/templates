# Action Writing Guide (Claude Code)

This guide provides instructions for Claude Code when creating and modifying Action documents.

## File Naming Convention

```
ACT-NNNN-XX-kebab-title.md
```

- NNNN: Linked ADR number (4 digits)
- XX: Action sequence number (2 digits, starting from 01)
- kebab-title: decision title connected with hyphens (-)

Example: `ACT-0001-01-implement-caching-layer.md`

## Procedure for Creating a New Action

1. Copy template: Copy [`ACT-0000-00-template.md`](./ACT-0000-00-template.md) (English) or [`ACT-0000-00-템플릿.md`](./ACT-0000-00-템플릿.md) (Korean).
2. Rename file: Rename to `ACT-NNNN-XX-kebab-title.md` format.
3. Write content: Fill in content according to each section.
4. Set status: Set initial status to `⏳ pending`.
5. Update README: Add an entry to the Action list table in `./README.md`.

## Document Structure

Action documents consist of the following sections:

| Section | Required | Description |
|:---|:---:|:---|
| Info | Yes | Metadata: ID, ADR, status, type, priority, assignee, dates, milestone, etc. |
| Overview | Yes | Purpose |
| Details | Yes | Description, acceptance criteria, implementation scope, out of scope |
| External Links | No | GitHub repository, Jira tickets, PR |
| References | No | External reference documents and resources |
| Changelog | Yes | Document versions and change history |

## Status Values

### Action Status

| Status | Emoji | Code Value |
|:---|:---:|:---|
| pending | ⏳ | `pending` |
| in-progress | 🔄 | `in-progress` |
| done | ✅ | `done` |
| cancelled | ❌ | `cancelled` |
| blocked | 🚫 | `blocked` |

### External Link Status (Jira/PR)

External links use each system's own status values.

#### Jira Ticket Status

| Status | Description |
|:---|:---|
| `Open` | Open |
| `Reopen` | Reopened |
| `Holding` | On hold |
| `In Progress` | In progress |
| `Resolved` | Resolved |
| `Closed` | Closed |

* Status may vary depending on project settings.

#### GitHub Pull Request Status

| Status | Description |
|:---|:---|
| `Open` | Open, awaiting review |
| `Draft` | Draft, preparing for review |
| `In Review` | Under review |
| `Approved` | Approved |
| `Merged` | Merged |
| `Closed` | Closed (not merged) |

## Action Types

| Type | Emoji | Code Value | Description |
|:---|:---:|:---|:---|
| New feature implementation | ✨ | `feature` | New feature implementation |
| Bug fix | 🐛 | `bugfix` | Bug fix |
| Refactoring | 🔧 | `refactor` | Refactoring |
| Documentation writing/modification | 📚 | `docs` | Documentation writing/modification |
| Test code writing | 🧪 | `test` | Test code writing |
| Miscellaneous tasks | 🧹 | `chore` | Miscellaneous tasks |

## Priority

| Priority | Color | Code Value | Description |
|:---|:---:|:---|:---|
| Requires immediate attention | 🔴 | `critical` | Requires immediate attention |
| High priority | 🟠 | `high` | High priority |
| Normal priority | 🟡 | `medium` | Normal priority |
| Low priority | 🟢 | `low` | Low priority |

## Writing Acceptance Criteria

Clearly define the completion criteria:

```markdown
### Acceptance Criteria
- [ ] Criterion 1
- [ ] Criterion 2
- [ ] Criterion 3
```

## README Update

When adding a new Action, add an entry to the Action list table in `./README.md`:

```markdown
| [ADR-NNNN](../adr/ADR-NNNN-kebab-title.md) | [ACT-NNNN-XX](./ACT-NNNN-XX-kebab-title.md) | Title | feature | high | ⏳ pending | Assignee | YYYY-MM-DD |
```

## External Link Management

Each Action can be linked to the following external resources:

### GitHub Repository
- Main repository
- Sub-repository
- Related libraries

### Jira Ticket
- Main ticket
- Sub-tasks
- Related tickets

### GitHub Pull Request
- Implementation PR
- Fix PR
- Refactoring PR

## Changelog

Always record in the `Changelog` section when modifying an Action document:

```markdown
| Version | Date | Changes | Author |
|:---:|:---:|:---|:---|
| 1.0 | YYYY-MM-DD | Initial draft | John Doe |
| 1.1 | YYYY-MM-DD | Acceptance criteria added | Jane Smith |
| 1.2 | YYYY-MM-DD | Status changed: in-progress → done | Bob Johnson |
```

---

## Directory Structure

```
docs/actions/
├── README.md                # Action list
├── CLAUDE.md                # This file (writing guide)
├── ACT-0000-00-template.md  # Action template (English)
└── ACT-0000-00-템플릿.md     # Action template (Korean)

# All Actions are stored together in this directory
```

---

## Related Document Links

Actions are managed in connection with ADR documents:

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

### Action Relationship Diagram

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
│ • Performance bench   │       │ • Bugfix              │
│ • Technical verify    │       │ • Refactor            │
└───────────────────────┘       ├───────────────────────┤
                                │ • GitHub repo         │
                                │ • Jira ticket         │
                                │ • GitHub PR           │
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

Reference the linked ADR document in the Info section of the Action.
