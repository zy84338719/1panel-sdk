# example/verify — live integration test against a real 1Panel server

Build and run a small CLI that exercises the SDK against an actual
1Panel v2 server using the API key (header-based) auth documented at
https://1panel.cn/docs/v2/dev_manual/api_manual/#22-token.

## Usage

```bash
go build -o verify .
./verify \
  -url http://192.168.108.235:55555 \
  -key 0HeE0VPfS2TY5N1ILGgE322McU7hRdmz
```

Flags:

- `-url`     1Panel base URL (default `http://192.168.108.235:55555`)
- `-key`     API key from the panel (required)
- `-sign`    `hmac-sha256` (default, recommended) or `md5` (legacy)
- `-node`    target node id (default: local)
- `-skip-noauth` skip the unauthenticated baseline probe

## What it probes

1. `POST /toolbox/device/base` — device info (object)
2. `GET  /dashboard/base/os`   — OS info (object)
3. `GET  /hosts/search`        — host list (object)
4. `OnNode("0")` override      — confirms the `CurrentNode` header is
                                 honored (server returns 500 for unknown
                                 node id).
5. `GET  /core/auth/current`   — expected 401 (proves `/core/*` requires
                                 session cookies, not the API key).
6. `GET  /dashboard/app/launcher` — array response, exercises GetList.
7. `POST /groups/search`       — array response, exercises PostList.

A successful run looks like:

```
[device/base] OK
[dashboard/base/os] OK
[hosts/search] OK
[hosts/search node=0] ERR  1panel POST .../hosts/search: code=500 message="未能找到该节点信息..."
[auth/current (expect 401)] ERR  1panel GET .../core/auth/current: code=401 ...
[dashboard/app/launcher] OK (len=6)
[groups/search] OK (len=2)
```

`debug.go` is a quick raw-body probe for any endpoint — build it with
`go run debug.go` (it has a `//go:build ignore` tag so it doesn't
interfere with `go build .`).
