# Contributing

## Repo layout

```
1panel-sdk/
├── client/                # low-level HTTP client (cookie, CSRF, /api/v2)
│   ├── client.go
│   ├── auth_helpers.go
│   ├── types.go
│   └── client_test.go
├── scripts/
│   ├── gen-from-swagger.py   # swagger.json -> services_swagger.go
│   ├── split-services2.py    # one-time services2.go split
│   └── split-services.py     # one-time services.go split
├── *.go                   # service files (see one of them for the pattern)
├── services_swagger.go    # GENERATED — do not edit by hand
├── example/               # runnable demo
├── example_test.go        # godoc Example* functions
├── sdk_test.go            # smoke tests (httptest)
├── go.mod, go.sum
├── Makefile
├── .golangci.yml
├── README.md, CHANGELOG.md, CONTRIBUTING.md, LICENSE
```

## Code style

- MixedCaps only, no underscores (per Go convention).
- Errors are lowercase, no trailing punctuation, prefixed with the package
  name (e.g. `"client: marshal body: %w"`).
- Wrap with `%w` to preserve `errors.Is`/`errors.As` chains.
- Don't panic for expected conditions.
- Receiver names are 1-2 letter abbreviations of the type.
- Don't introduce a third-party dependency without a strong reason — the SDK
  is intentionally stdlib-only.

## How to add or change an endpoint

1. **If 1Panel upstream added it**:
   ```bash
   curl -sL https://raw.githubusercontent.com/1Panel-dev/1Panel/master/core/cmd/server/docs/swagger.json \
       > /tmp/swagger.json
   # Edit scripts/gen-from-swagger.py to point at /tmp/swagger.json
   python3 scripts/gen-from-swagger.py
   go build ./... && make test
   ```
   This regenerates `services_swagger.go` (the only file that needs to
   change). Existing typed methods are preserved.

2. **If you want a typed wrapper** (frequent use, hand-curated signature):
   - Add the method to the matching service file (e.g. `container.go` for a
     new container endpoint).
   - Use the existing `s.Get`/`s.Post` helpers; do not bypass them.
   - If the endpoint's request body has rich fields, add a typed struct to
     `types.go` (e.g. `ContainerCreate`).
   - Add a godoc comment that starts with the method name, e.g.
     `// Restart restarts a container by id.`
   - Run `make lint` to keep the suite at 0 issues.

3. **Naming**:
   - Don't repeat the service name in method names
     (`sdk.Container.List`, not `sdk.Container.ListContainers`).
   - Booleans get `is`/`has`/`can` prefix on **fields** and methods
     (`IsEnabled()`, not `Enabled()`).
   - Avoid stutter at the call site: `onepanel.New` is the only `New` in
     the package.

## Tests

- Use `httptest` for any test that needs to hit the network.
- Keep `services_swagger.go` excluded from lint via
  `issues.exclude-files` in `.golangci.yml`; it intentionally has unused
  parameters and gocritic nits.
- `make all` runs build + test + lint and is the single command to run
  before sending a PR.

## Releases

- Bump the version in `CHANGELOG.md` under a new heading.
- Tag with `git tag vX.Y.Z && git push --tags`.
- The Go module path is `github.com/zy84338719/1panel-sdk`; users can
  `go get github.com/zy84338719/1panel-sdk@vX.Y.Z`.
