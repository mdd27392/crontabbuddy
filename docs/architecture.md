# Architecture

CrontabBuddy is organized around a small set of focused modules. The goal is to keep the core small and let plugins, hooks, and configuration do the heavy lifting.

## Layout

- `src/crontabbuddy/` - primary package source
- `tests/` - unit and integration tests
- `examples/` - runnable examples
- `docs/` - this documentation

## Design notes

- Configuration is read once at startup and kept immutable during a run.
- Errors are surfaced as typed exceptions so callers can react precisely.
- Everything I/O related lives behind small adapters to make testing simple.

## Extension points

The project exposes a hook system. New behavior can be registered without changing core code, which keeps the surface area stable as features are added.