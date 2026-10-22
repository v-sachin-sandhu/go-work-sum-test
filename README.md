# go-work-sum-test

Reproduces dependabot/dependabot-core#16518 (golang/go#63901, fixed in Go 1.28): after updating a Go workspace
on Go < 1.28, `go.work.sum` is missing checksums unless `go list -m all` is run last.

Workspace: `./api`, `./worker`, Go 1.27.0, `k8s.io/apimachinery` and `k8s.io/client-go` at v0.37.0.

Manual repro: `go get k8s.io/apimachinery@v0.37.1 k8s.io/client-go@v0.37.1` in each module, `go work sync`,
`go mod tidy` in each module (no `go.work.sum` change), then `go list -m all` adds
`k8s.io/streaming v0.37.1/go.mod ...` to `go.work.sum`.
