# Grace

## Quick start

```go
package main

import (
	"time"
	"github.com/vite-cloud/grace"
	"net/http"
	"fmt"
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

Now, once you have started the server, make a GET request:
```bash
curl localhost:8000/hello
```

This will print "world" to the console after 5 seconds.

If you terminate the server with Ctrl+C, the server will be terminated after the current request is finished.