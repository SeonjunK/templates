# ACT-0000-00. Action Title

## Info
<!--
| Field | Required | Description |
|:---|:---:|:---|
| ID | Yes | Unique ID (ACT-NNNN-XX format) |
| ADR | Yes | Linked ADR |
| Status | Yes | Action status |
| Type | Yes | Action type |
| Priority | Yes | Priority |
| Assignee | Yes | Person responsible for the task |
| Proposer | Yes | Person who proposed the Action |
| Created | Yes | Action creation date |
| Updated | Yes | Last modified date |
| Started | No | Task start date |
| Target Date | No | Target completion date |
| Completed | No | Actual completion date |
| Milestone | No | Target milestone for implementation (date/phrase/version, free text) |
| Tags | No | Tags for search/filtering |

### Status Values

| Status | Emoji |
|:---|:---:|
| `pending` | ⏳ |
| `in-progress` | 🔄 |
| `done` | ✅ |
| `cancelled` | ❌ |
| `blocked` | 🚫 |

### Action Types

| Type | Emoji | Description |
|:---|:---:|:---|
| `feature` | ✨ | New feature implementation |
| `bugfix` | 🐛 | Bug fix |
| `refactor` | 🔧 | Refactoring |
| `docs` | 📚 | Documentation writing/modification |
| `test` | 🧪 | Test code writing |
| `chore` | 🧹 | Miscellaneous tasks |

### Priority

| Priority | Color | Description |
|:---|:---:|:---|
| `critical` | 🔴 | Requires immediate attention |
| `high` | 🟠 | High priority |
| `medium` | 🟡 | Normal priority |
| `low` | 🟢 | Low priority |
-->

| Field | Value |
|:---|:---|
| ID | ACT-0000-00 |
| ADR | [ADR-0000](../adr/ADR-0000-title.md) |
| Status | ⏳ pending |
| Type | ✨ feature |
| Priority | 🟠 high |
| Assignee | Assignee Name |
| Proposer | Proposer Name |
| Created | YYYY-MM-DD |
| Updated | YYYY-MM-DD |
| Started | YYYY-MM-DD (optional) |
| Target Date | YYYY-MM-DD (optional) |
| Completed | YYYY-MM-DD (optional) |
| Milestone | 2025-Q1, v1.2.0, Phase 1 (optional) |
| Tags | tag1, tag2, tag3 |

## Overview

### Purpose
<!--
Describe the purpose of this Action and what it aims to achieve.
-->

## Details

### Description
<!--
Describe the details of the Action.
- Features to implement
- Bugs to fix
- Refactoring scope, etc.
-->

### Acceptance Criteria
<!--
Clearly define the completion criteria.
- [ ] Criterion 1
- [ ] Criterion 2
- [ ] Criterion 3
-->

### Implementation Scope
<!--
Describe the scope of work included in this Action.
-->

### Out of Scope
<!--
Clearly specify what is excluded from this Action.
-->

## External Links

### GitHub Repository

| Repository | Description | Link |
|:---|:---|:---|
| org/repo-1 | Main repository | [github.com/org/repo-1](https://github.com/org/repo-1) |

### Jira Ticket

| Ticket ID | Title | Status | Link |
|:---|:---|:---:|:---|
| TICKET-001 | Ticket Title | In Progress | [jira.example.com/browse/TICKET-001](https://jira.example.com/browse/TICKET-001) |

### Jira Status Values

| Status | Description |
|:---|:---|
| `Open` | Open |
| `Reopen` | Reopened |
| `Holding` | On hold |
| `In Progress` | In progress |
| `Resolved` | Resolved |
| `Closed` | Closed |

### GitHub Pull Request

| PR # | Title | Status | Link |
|:---:|:---|:---:|:---|
| 123 | Implementation PR | Open | [github.com/org/repo/pull/123](https://github.com/org/repo/pull/123) |

### GitHub PR Status Values

| Status | Description |
|:---|:---|
| `Open` | Open, awaiting review |
| `Draft` | Draft, preparing for review |
| `In Review` | Under review |
| `Approved` | Approved |
| `Merged` | Merged |
| `Closed` | Closed (not merged) |

## References

> External resources referenced when writing this Action.

### Implementation Related
- [Test Code](https://github.com/org/repo/tree/main/tests)
- [API Documentation](https://docs.example.com)

### Technical Resources
- [Tech Blog](https://example.com/blog)
- [Related Spec](https://example.com/spec)

## Changelog

| Version | Date | Changes | Author |
|:---:|:---:|:---|:---|
| 1.0 | YYYY-MM-DD | Initial draft | John Doe |
| 1.1 | YYYY-MM-DD | Acceptance criteria added | Jane Smith |
| 1.2 | YYYY-MM-DD | PR link added | Bob Johnson |
