# asciiart-middleware
Golang http middleware for ASCII-art in headers

This is a simple package with a middleware that adds a line-splitted headers to the response

### Installation
```bash
❯ go get github.com/anvlad11/asciiart-middleware
```

### Usage example
```golang
package main

import (
	"log"
	"net/http"
	"strings"

	asciiartmiddleware "github.com/anvlad11/asciiart-middleware"
)

func main() {
	prefix := "x-motd-"
	message := "       .------..\n     -          -\n   /              \\\n /                   \\\n/    .--._    .---.   |\n|  /      -__-     \\   |\n| |                 |  |\n ||     ._   _.      ||\n ||      o   o       ||\n ||      _  |_      ||\n C|     (o\\_/o)     |O     Uhhh, this computer\n  \\      _____      /       is like, busted or\n    \\ ( /#####\\ ) /       something. So go away.\n     \\  `====='  /\n      \\  -___-  /\n       |       |\n       /-_____-\\\n     /           \\\n   /               \\\n  /__|  AC / DC  |__\\\n  | ||           |\\ \\"

	middleware := asciiartmiddleware.NewAsciiArtMiddleware(prefix, strings.Split(message, "\n"))

	mux := http.NewServeMux()
	finalHandler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("OK"))
	})
	mux.Handle("/", middleware(finalHandler))

	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
```

```basg
❯ curl -sSL -D - http://localhost:3000
HTTP/1.1 200 OK
X-Motd-100: #        .------..
X-Motd-101: #      -          -
X-Motd-102: #    /              \
X-Motd-103: #  /                   \
X-Motd-104: # /    .--._    .---.   |
X-Motd-105: # |  /      -__-     \   |
X-Motd-106: # | |                 |  |
X-Motd-107: #  ||     ._   _.      ||
X-Motd-108: #  ||      o   o       ||
X-Motd-109: #  ||      _  |_      ||
X-Motd-110: #  C|     (o\_/o)     |O     Uhhh, this computer
X-Motd-111: #   \      _____      /       is like, busted or
X-Motd-112: #     \ ( /#####\ ) /       something. So go away.
X-Motd-113: #      \  `====='  /
X-Motd-114: #       \  -___-  /
X-Motd-115: #        |       |
X-Motd-116: #        /-_____-\
X-Motd-117: #      /           \
X-Motd-118: #    /               \
X-Motd-119: #   /__|  AC / DC  |__\
X-Motd-120: #   | ||           |\ \
Date: Tue, 20 Oct 2020 13:20:57 GMT
Content-Length: 2
Content-Type: text/plain; charset=utf-8

OK%
```

### Why
Why not?
Inspired by [this tweet](https://twitter.com/thingskatedid/status/1316079949413928961), made in half an hour
