# Rust Template

A general-purpose Rust project template.

## Structure

```
rust/
├── src/
│   └── lib.rs      # Library entry point
├── tests/          # Integration tests
├── docs/
│   ├── actions/       # Action logs
│   ├── adr/           # Architecture Decision Records
│   ├── architecture/  # Architecture documentation
│   ├── guides/        # Development guides
│   └── poc/           # Proof of Concept documents
└── Cargo.toml
```

## Setup

```bash
cargo build
```

## Commands

- Build: `cargo build`
- Test: `cargo test`
- Format: `cargo fmt`
- Lint: `cargo clippy`

## Documentation

- [Architecture](./docs/architecture/README.md) - Project structure and design
- [Guides](./docs/guides/README.md) - Development guidelines and conventions
- [Actions](./docs/actions/README.md) - Action logs and records
- [ADR](./docs/adr/README.md) - Architecture Decision Records
- [PoC](./docs/poc/README.md) - Proof of Concept documents

## Development

See [CLAUDE.md](./CLAUDE.md) for coding conventions and guidelines.
