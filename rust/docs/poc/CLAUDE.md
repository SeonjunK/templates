# POC Writing Guide (Claude Code)

This guide provides instructions for Claude Code when creating and modifying POC documents.

## File Naming Convention

```
POC-NNNN-XX-kebab-title.md
```

- NNNN: Linked ADR number (4 digits)
- XX: POC sequence number (2 digits, starting from 01)
- kebab-title: decision title connected with hyphens (-)

Example: `POC-0001-01-caching-performance-test.md`

## Procedure for Creating a New POC

1. Copy template: Copy [`POC-0000-00-template.md`](./POC-0000-00-template.md) (English) or [`POC-0000-00-템플릿.md`](./POC-0000-00-템플릿.md) (Korean).
2. Rename file: Rename to `POC-NNNN-XX-kebab-title.md` format.
3. Write content: Fill in content according to each section.
4. Set status: Set initial status to `⏳ pending`.
5. Update README: Add an entry to the POC list table in `./README.md`.

## Document Structure

POC documents follow the IMRaD (Introduction, Methods, Results, and Discussion) structure of academic papers.

| Section | Required | Description |
|:---|:---:|:---|
| Info | Yes | Metadata: ID, ADR, status, assignee, dates, milestone, tags, etc. |
| Introduction | Yes | Background, purpose, alternatives to verify, hypothesis |
| Methods | Yes | Test environment, experimental design, procedure |
| Results | No | Data, analysis |
| Discussion | No | Interpretation, alternative comparison and selection |
| Conclusion | Yes | Summary, recommendations (selected alternative and application scope) |
| References | No | External reference documents and resources |
| Changelog | Yes | Document versions and change history |

## Status Values

### POC Status

| Status | Emoji | Code Value |
|:---|:---:|:---|
| pending | ⏳ | `pending` |
| in-progress | 🔄 | `in-progress` |
| done | ✅ | `done` |
| cancelled | ❌ | `cancelled` |
| blocked | 🚫 | `blocked` |

## Tags

Enter tags freely to indicate the nature of the POC:
- Verification purpose: `performance`, `security`, `feature`, `compatibility`, `scalability`
- Test type: `load`, `benchmark`, `stress`, `stability`
- Tech stack: `redis`, `postgresql`, `kubernetes`

Example: `performance, load, caching`

## README Update

When adding a new POC, add an entry to the POC list table in `./README.md`:

```markdown
| [ADR-NNNN](../adr/ADR-NNNN-kebab-title.md) | [POC-NNNN-XX](./POC-NNNN-XX-kebab-title.md) | Title | load | 🔄 in-progress | YYYY-MM-DD |
```

## Changelog

Always record in the `Changelog` section when modifying a POC document:

```markdown
| Version | Date | Changes | Author |
|:---:|:---:|:---|:---|
| 1.0 | YYYY-MM-DD | Initial draft | John Doe |
| 1.1 | YYYY-MM-DD | Results added | Jane Smith |
```

---

## Directory Structure

```
docs/
├── adr/          # ADR documents
├── actions/      # Action documents
└── poc/          # POC documents
    ├── logs/     # Log files
    └── datasets/ # Raw data
```

---

## Related Document Links

POCs are managed in connection with ADR documents:

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
│   ├── CLAUDE.md
│   ├── POC-0000-00-template.md
│   ├── POC-0001-01-load-test.md
│   ├── logs/                    # POC logs
│   └── datasets/                # POC datasets
```

### POC Relationship Diagram

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

Reference the linked ADR in the Info section of the POC. Actions are not directly linked to POCs.

---

## Data and Log Management

### logs/ Directory

Stores log files generated during POC execution.

#### File Naming Convention

```
POC-NNNN-XX.log
POC-NNNN-XX-{description}.log
```

Examples:
- `POC-0001-01.log`
- `POC-0001-01-load-test.log`
- `POC-0001-01-error.log`

#### Log File Writing Guide

- Timestamp: Record accurate time for each log entry
- Log levels: `DEBUG`, `INFO`, `WARN`, `ERROR`
- Output format: Text log or structured JSON

#### Referencing in POC Documents

```markdown
### Raw Data
- [Log File](./logs/POC-0000-00.log)
```

### datasets/ Directory

Stores raw datasets generated from POC execution results.

#### File Naming Convention

```
POC-NNNN-XX-{datatype}.{extension}
```

Supported data types:
- `metrics` - Performance measurement data
- `results` - Experiment result data
- `raw` - Raw data
- `summary` - Summary data

Examples:
- `POC-0001-01-metrics.csv`
- `POC-0001-01-results.json`
- `POC-0001-01-summary.xlsx`

#### Data File Formats

| Format | Extension | Purpose |
|:---|:---|:---|
| CSV | `.csv` | Tabular data, measurement metrics |
| JSON | `.json` | Structured data, API responses |
| Excel | `.xlsx` | Analysis reports, summary tables |
| Text | `.txt`, `.log` | Unstructured output |

#### Referencing in POC Documents

```markdown
### Raw Data
- [Metrics CSV](./datasets/POC-0000-00-metrics.csv)
- [Results JSON](./datasets/POC-0000-00-results.json)
```

### File Retention and Archiving

- In progress: Keep files in the directory while the POC is in progress
- After completion: Compress and archive if needed after POC completion
- Deletion: Files no longer needed can be deleted (archiving recommended if referenced in POC documents)
