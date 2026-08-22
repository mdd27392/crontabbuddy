# CrontabBuddy guide 0

Practical notes collected while working on CrontabBuddy.

## What changed in this round

Small improvements and fixes that keep the project moving. Nothing here changes the public contract; it is internal polish that makes the codebase easier to maintain.

## Things to know

- Run the test suite before pushing changes.
- Keep new code in its own module so the core stays small.
- Update the changelog when behavior changes.

## Next steps

The roadmap stays conservative: finish the current module, then look at reducing allocations in the hot path.