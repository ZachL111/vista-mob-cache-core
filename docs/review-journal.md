# Review Journal

The cases below are the review handles I would use before changing the implementation.

The local checks classify each case as `ship`, `watch`, or `hold`. That gives the project a small review vocabulary that matches its mobile workflows focus without claiming live deployment or external usage.

## Cases

- `baseline`: `form pressure`, score 141, lane `ship`
- `stress`: `sync drift`, score 144, lane `ship`
- `edge`: `local state`, score 182, lane `ship`
- `recovery`: `conflict cost`, score 215, lane `ship`
- `stale`: `form pressure`, score 188, lane `ship`

## Note

A future change should add new cases before it changes the scoring rule.
