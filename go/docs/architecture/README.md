# Architecture Documentation

This directory maintains documents that reflect the structure and design of source code and project files.

## Purpose

- Document the current state of the codebase architecture
- Explain design decisions and patterns used in the project
- Provide reference for understanding code structure

## Documents

| Document | Description |
|:---------|:------------|
| [config.md](config.md) | Config package structure, priority order, env var mapping, validation |
| [cli.md](cli.md)       | CLI entry point structure, Cobra command pattern, config integration  |

## Guidelines

- **Sync with code**: Any change to code or files must be immediately reflected in the corresponding architecture document
- **Current state only**: Documents always represent the current state of the actual source code and files
- **PLANNED tag**: Use `PLANNED` tag to mark planned but not yet implemented content

## File Naming

- Use descriptive names: `module-name.md`, `component-name.md`
- Group related content in single files when appropriate
