# POC Log Files

This directory stores log files generated during POC execution.

## File Naming Convention

```
POC-NNNN-XX.log
POC-NNNN-XX-{description}.log
```

- **NNNN**: Linked ADR number (4 digits)
- **XX**: POC sequence number (2 digits)
- **description**: Log type (optional)

## Examples

| Filename | Description |
|:---|:---|
| `POC-0001-01.log` | Default log for POC-0001-01 |
| `POC-0001-01-load-test.log` | Load test log |
| `POC-0001-01-error.log` | Error log |
| `POC-0001-01-debug.log` | Debug log |

## Log File Writing Guide

- **Timestamp**: Record accurate time for each log entry
- **Log levels**: `DEBUG`, `INFO`, `WARN`, `ERROR`
- **Output format**: Text log or structured JSON

## Related Documents
- [ADR Documents](../../adr/README.md) - Architecture decision records
- [POC Documents](../README.md) - Proof of concept documents
- [POC Datasets](../datasets/README.md) - Data storage

## Referencing in POC Documents

```markdown
### Raw Data
- [Log File](./logs/POC-0000-00.log)
```

## File Retention and Archiving

- **In progress**: Keep files while the POC is in progress
- **After completion**: Compress and archive if needed after POC completion
- **Deletion**: Files no longer needed can be deleted (archiving recommended if referenced in POC documents)
