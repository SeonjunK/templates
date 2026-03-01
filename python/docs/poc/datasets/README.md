# POC Datasets

This directory stores raw datasets generated from POC execution results.

## File Naming Convention

```
POC-NNNN-XX-{datatype}.{extension}
```

- **NNNN**: Linked ADR number (4 digits)
- **XX**: POC sequence number (2 digits)
- **datatype**: Data type (metrics, results, raw, summary, etc.)

## Supported Data Types

| Data Type | Description |
|:---|:---|
| `metrics` | Performance measurement data |
| `results` | Experiment result data |
| `raw` | Raw data |
| `summary` | Summary data |

## Examples

| Filename | Description |
|:---|:---|
| `POC-0001-01-metrics.csv` | Performance measurement data (CSV) |
| `POC-0001-01-results.json` | Experiment results (JSON) |
| `POC-0001-01-summary.xlsx` | Summary table (Excel) |
| `POC-0001-01-raw.txt` | Raw output data |

## Data File Formats

| Format | Extension | Purpose |
|:---|:---|:---|
| CSV | `.csv` | Tabular data, measurement metrics |
| JSON | `.json` | Structured data, API responses |
| Excel | `.xlsx` | Analysis reports, summary tables |
| Text | `.txt`, `.log` | Unstructured output |

## Related Documents
- [ADR Documents](../../adr/README.md) - Architecture decision records
- [POC Documents](../README.md) - Proof of concept documents
- [POC Logs](../logs/README.md) - Execution log storage

## Referencing in POC Documents

```markdown
### Raw Data
- [Metrics CSV](./datasets/POC-0000-00-metrics.csv)
- [Results JSON](./datasets/POC-0000-00-results.json)
```

## File Retention and Archiving

- **In progress**: Keep files while the POC is in progress
- **After completion**: Compress and archive if needed after POC completion
- **Deletion**: Files no longer needed can be deleted (archiving recommended if referenced in POC documents)
