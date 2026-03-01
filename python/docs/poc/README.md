# POC Documents

## Overview

Proof of concept documents for verifying ADR (Architecture Decision Records).
Records verification results such as performance measurements, load testing, benchmarks, etc.

> Note: Tests for Action implementations (unit tests, integration tests, etc.) are managed as code and are not separately documented.
> POCs are directly linked only to ADRs, not to Actions.

## POC List

| ADR | POC | Title | Tags | Status | Completed |
|:---:|:---:|:---|:---|:---:|:---:|
| - | - | - | - | - | - |

## Quick Reference

### Status Values
- ⏳ pending | 🔄 in-progress | ✅ done | ❌ cancelled | 🚫 blocked

### Tag Examples
- Verification purpose: `performance`, `security`, `feature`, `compatibility`, `scalability`
- Test type: `load`, `benchmark`, `stress`, `stability`
- Tech stack: `redis`, `postgresql`, `kubernetes`

### Related Documents
- [CLAUDE.md](./CLAUDE.md) - Writing guide
- [POC-0000-00-template.md](./POC-0000-00-template.md) - Document template (English)
- [POC-0000-00-템플릿.md](./POC-0000-00-템플릿.md) - Document template (Korean)
- [ADR Documents](../adr/README.md) - Architecture decision records
- [Action Documents](../actions/README.md) - Implementation item management
- [POC Logs](./logs/) - Execution log storage
- [POC Datasets](./datasets/) - Data storage

### Document Structure
POC documents follow the IMRaD (Introduction, Methods, Results, and Discussion) structure of academic papers:
- Introduction: Background, purpose, alternatives to verify, hypothesis
- Methods: Test environment, experimental design, procedure
- Results: Data, analysis
- Discussion: Interpretation, alternative comparison and selection
- Conclusion: Summary, recommendations (selected alternative and application scope)
