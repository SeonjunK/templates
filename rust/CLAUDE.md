# Rust Project Template

## Project Structure
```
rust/
├── src/
│   └── lib.rs          # Library entry point
├── tests/              # Integration tests
├── docs/
│   ├── architecture/   # Architecture documentation
│   └── guides/         # Development guides
└── Cargo.toml          # Package manifest
```

## Commands

- Format: `cargo fmt`
- Lint: `cargo clippy`
- Test: `cargo test`
- Test (verbose): `cargo test -- --nocapture`
- Build: `cargo build`
- Build (release): `cargo build --release`
- Add dependency: `cargo add <package>`
- Add dev dependency: `cargo add --dev <package>`

## Code Style

- `cargo fmt` formatting required (applied automatically on save)
- Clippy warnings must be resolved before commit
- Public functions and types require doc comments (`///`)

## Required Tools

```bash
# rust-analyzer (for VS Code)
# Install via VS Code extension marketplace
```

VS Code extensions:
- `rust-lang.rust-analyzer`

## Documentation

### docs/architecture/
Maintains documents that reflect the structure and design of source code and project files.

- Any change to code or files must be immediately reflected in the corresponding architecture document.
- Documents always represent the current state of the actual source code and files.
- Planned but not yet implemented content is marked with the `PLANNED` tag.

### docs/guides/
Maintains documents containing rules and guidelines that must be followed during development.

- Includes coding conventions, collaboration rules, and workflow processes.
- Always read the relevant guides before starting work and follow them.
