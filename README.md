# Grace
[![Go Reference](https://pkg.go.dev/badge/github.com/vite-cloud/grace.svg)](https://pkg.go.dev/github.com/vite-cloud/go-grace)
[![codebeat badge](https://codebeat.co/badges/12534986-c871-4e80-8c21-364abae97ce2)](https://codebeat.co/projects/github-com-vite-cloud-grace-main)
[![Tests](https://github.com/vite-cloud/grace/actions/workflows/tests.yml/badge.svg)](https://github.com/vite-cloud/go-grace/actions/workflows/tests.yml)
[![codecov](https://codecov.io/gh/vite-cloud/grace/branch/main/graph/badge.svg?token=2EBL0P4UN6)](https://codecov.io/gh/vite-cloud/grace)

Graceful shutdown for HTTP(S) servers.

## Quick start

```go
// main.go
package main

import (
	"fmt"
	"github.com/vite-cloud/grace"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		fmt.Fprintln(w, "world")
	})

	server := &http.Server{Addr: ":8000"}

	manager := grace.New(
		grace.WithServer(server),
	)

	go func() {
		_ = server.ListenAndServe()
	}()

	manager.Wait()
}
```

Start the server:

```bash
go run .
```

Now, once the server is running, make a GET request:

```bash
curl localhost:8000/hello
```

This will print "world" after 5 seconds.

You may terminate the server with Ctrl+C, it will stop accepting new connections and finish existing ones. If the
request takes more than the defined timeout (10 seconds in recommended) the server will close the connection.