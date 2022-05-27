# Grace

## Quick start

```go
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

Start the server

```bash
go run .
```

Now, once the server is running, make a GET request:

```bash
curl localhost:8000/hello
```

This will print "world" to the console after 5 seconds.

You may terminate the server with Ctrl+C, it will stop accepting new connections and finish existing ones. If the
request takes more than the defined timeout (10 seconds in recommended) the server will close the connection.