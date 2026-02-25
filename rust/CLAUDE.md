# Rust Project Template

## Project Structure
```
rust/
├── src/
│   └── lib.rs          # Library entry point
├── tests/              # Integration tests
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
