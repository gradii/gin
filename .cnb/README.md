# CNB test modules

The CNB pipelines test this route driver together with the framework in an
explicit Go workspace, so a change here is checked against the framework's
`main` rather than against the released tag this module pins in `go.mod`.

The current repository always uses the checkout supplied by CNB. The framework
is cloned from `GRADII_FRAMEWORK_URL` at `GRADII_FRAMEWORK_REF`, both provided
by the shared refs file `https://cnb.cool/gradii/secret/-/blob/main/ci/framework-refs.yml`
that every test job imports. The URL defaults to `https://cnb.cool/gradii/framework`
and, with the ref unset, the checkout follows that repository's default branch.
The GitHub home, `https://github.com/gradii/framework`, is written by the mirror
pipelines described below and never read here. The checkout is validated against
its expected `github.com/gradii/framework` module name before testing.
Authentication uses the runner's Git configuration; URLs should not contain
credentials.

`bash .cnb/ci.sh unit` tests both root modules, this one and the framework's.
`bash .cnb/ci.sh workspace` builds the workspace and prints `GOWORK` for the
lint and coverage jobs to export. The `database`, `integration` and `specified`
modes belong to the database modules and fail fast here.

The helper creates an isolated temporary workspace and removes it on exit. It
does not depend on a developer's ignored `go.work`, rewrite module declarations,
or run `go mod tidy`. Every repository of the workspace carries the same
`ci.sh`, so its CNB checkout can prepare the workspace without fetching an
unpinned helper script.

## GitHub mirror

This repository mirrors itself to its GitHub home, `github.com/gradii/<module>`,
from the CNB pipeline. Only two things are published, and only from `main`: a
push to `main` publishes that branch together with every `v*` tag the CNB
repository holds, and a `v*` tag push publishes that tag. No other branch is
mirrored.

The destination is derived from the module path in `go.mod` and the credential
is `DEPLOY_GITHUB_GRADII_<MODULE>_TOKEN`, imported from
`https://cnb.cool/gradii/secret/-/blob/main/deploy-github/gradii.yml`. Without
the token the mirror job skips rather than fails. Both stages push with
`--force --atomic`, so GitHub is overwritten to match CNB and a failure leaves
no half-applied update. Tags below `v1.20.0` are upstream release history and
are never published.
