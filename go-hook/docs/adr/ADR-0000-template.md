# ADR-0000. Decision Title

## Info
<!--
| Field | Required | Description |
|:---|:---:|:---|
| ID | Yes | Unique ID (ADR-NNNN format) |
| Status | Yes | ADR status |
| Proposer | Yes | Person who initially proposed the ADR |
| Decision Makers | No | List of decision makers |
| Proposed Date | Yes | Date the ADR was proposed |
| Approved Date | No | Date of final approval |
| Milestone | No | Target milestone for ADR implementation (date/phrase/version, free text) |
| Superseded By | No | List of ADR IDs that supersede this ADR |
| Tags | No | Tags for search/filtering |

### Status Values

| Status | Emoji |
|:---|:---:|
| `proposed` | 📝 |
| `under-review` | 🔍 |
| `approved` | ✅ |
| `deprecated` | ❌ |
| `superseded` | 🔄 |
| `to-be-retired` | 🗑️ |
-->

| Field | Value |
|:---|:---|
| ID | ADR-0000 |
| Status | 📝 proposed |
| Proposer | Proposer Name |
| Decision Makers | Decision Maker 1, Decision Maker 2 |
| Proposed Date | YYYY-MM-DD |
| Approved Date | YYYY-MM-DD (optional) |
| Milestone | 2025-Q1, v1.2.0, Phase 1 (optional) |
| Superseded By | ADR-XXXX, ADR-XXXX, ADR-XXXX (optional) |
| Tags | tag1, tag2, tag3 |

## Background

### Situation
<!--
Describe the current state of the system/service.
-->
- System state: (description of the current system/service state)
- Operational state: (number of users, traffic, resource usage, etc.)

### Problem
<!--
Describe the problem being faced and its impact.
-->
- Issue: (specific problem and symptoms)
- Impact: (effects of the problem - user experience, cost, risk, etc.)

### Requirements

#### Technical Requirements
<!--
Describe requirements from a technical perspective.
-->
- Performance: (e.g., response time < 100ms, throughput > 1000 TPS)
- Scalability: (e.g., handle 3x traffic increase)
- Security: (e.g., data encryption, stronger authentication)
- Reliability: (e.g., 99.9% availability, recovery time)
- Compatibility: (e.g., compatibility with existing systems)

#### Non-Technical Requirements
<!--
Describe requirements from a business/operational perspective.
-->
- Cost: (e.g., reduce monthly operating cost by 20%)
- Time: (e.g., complete implementation within 2 weeks)
- Staffing: (e.g., reduce maintenance headcount by 1)
- Regulatory/Compliance: (e.g., comply with personal data protection laws)
- Organizational: (e.g., align with operations team's tech stack)

### Constraints
<!--
Describe constraints that affect the decision.
-->
- Technical constraints: (e.g., legacy system compatibility)
- Cost constraints: (e.g., no additional license costs)
- Time constraints: (e.g., complete by next sprint)
- Staffing constraints: (e.g., lack of specialists in a specific technology)

### Decision Motivation
<!--
Explain why this decision is needed now.
-->
- Business need: (e.g., performance improvement needed due to user growth)
- Technical debt: (e.g., reaching limits of legacy code)
- Market environment: (e.g., technical disadvantage compared to competitors)
- Risk management: (e.g., resolving security vulnerabilities)

## Alternatives

### Alternative 1: [Alternative Name]
<!--
Describe other options considered.
-->

#### Expected Positive Outcomes
...

#### Expected Negative Outcomes
...

#### Expected Risk Factors
...

### Alternative 2: [Alternative Name]

#### Expected Positive Outcomes
...

#### Expected Negative Outcomes
...

#### Expected Risk Factors
...

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

## Decision

### Selected Option
<!--
Describe the chosen solution.
-->

#### Expected Positive Outcomes
...

#### Expected Negative Outcomes
...

#### Expected Risk Factors
...

### Reason for Selection
<!--
Describe the specific reasons for choosing this option.
-->
- Technical reason: (e.g., technical advantages in performance, scalability, security)
- Operational reason: (e.g., operations team's tech stack, ease of maintenance)
- Cost reason: (e.g., license costs, infrastructure costs)
- Time reason: (e.g., faster implementation, market timing)
- Organizational reason: (e.g., team experience, training costs)
- Strategic reason: (e.g., long-term technical direction, standards compliance)

### Details
<!--
Describe the specifics of the decision:
- Architecture changes
- Technologies and tools used
- Implementation scope
-->

## Implementation Status

| Task | Milestone | Status | Completed | Action |
|:---|:---|:---:|:---:|:---|
| Task 1 | Phase 1 | ✅ done | YYYY-MM-DD | [ACT-0001-01](../actions/ACT-0001-01-implement-caching-layer.md) |
| Task 2 | Phase 2 | 🔄 in-progress | - | [ACT-0001-02](../actions/ACT-0001-02-documentation-cleanup.md) |

### Implementation Status Values

| Status | Emoji |
|:---|:---:|
| `pending` | ⏳ |
| `in-progress` | 🔄 |
| `done` | ✅ |
| `cancelled` | ❌ |
| `blocked` | 🚫 |

## Outcomes and Impact

> Write after task completion (leave empty at draft stage)
>
> - Decision section: describe expected benefits/drawbacks/risks
> - Outcomes and Impact section: record actual results after task completion

### Positive Outcomes
<!--
Describe the actual benefits gained after task completion.
- [Actual benefit 1]: description
- [Actual benefit 2]: description
-->

### Negative Outcomes
<!--
Describe the actual drawbacks or costs that occurred after task completion.
- [Actual drawback 1]: description (including mitigation measures)
- [Actual drawback 2]: description (including mitigation measures)
-->

### Risk Factor Occurrence
<!--
Describe whether risks occurred and how they were addressed after task completion.
- [Risk 1]: occurred/did not occur → response measures
- [Risk 2]: occurred/did not occur → response measures
-->

### Expected vs Actual

| Category | Expected (Decision Section) | Actual (Outcomes Section) | Notes |
|:---|:---|:---|:---|
| Positive outcomes | - | - | |
| Negative outcomes | - | - | |
| Risk occurrence | - | - | |

## References

> External resources referenced when writing this ADR.

### Technical Documentation
- [Document Title](https://example.com)
- [Spec Document](https://example.com/spec)

### Related Resources
- [Blog Post](https://example.com/blog)
- [GitHub Repository](https://github.com/org/repo)

## Changelog

| Version | Date | Changes | Author |
|:---:|:---:|:---|:---|
| 1.0 | YYYY-MM-DD | Initial draft | John Doe |
| 1.1 | YYYY-MM-DD | Decision approved, implementation status added | Jane Smith |
| 1.2 | YYYY-MM-DD | Alternative 2 added, outcomes section revised | Bob Johnson |
