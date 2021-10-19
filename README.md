# go-ps
`go-bg` is a simple implementation of a banner-grabber in golang.

## Usage
To obtain the module, simply execute in shell:

```zsh
$ go get github.com/NovusEdge/go-bg
```

<br>

#### Example use-case:

* Importing the module:

```go
import gobg "github.com/NovusEdge/go-bg"
```

<br>
<br>


* Declaring an a variable of type `BannerGrabber`:
 
```go
// Decalring a BannerGrabber object.

/*
struct definition:
type BannerGrabber struct {
    URL string
}
*/

bg := gobg.BannerGrabber{
	URL: "http://scanme.nmap.org/",
}

```

<br>
<br>

* `Grab()` uses the [`curl`](https://curl.se/) command-line-tool and reports the banner thus obtained by _printing to stdout_.

Sample usage:
```
bg.Grab()
```


Output: 
```
[*] Grabbing banner for: scanme.nmap.org...
[*] Banner Grab Successful!
Banner:
HTTP/1.1 200 OK
Date: Wed, 20 Oct 2021 05:09:49 GMT
Server: Apache/2.4.7 (Ubuntu)
Accept-Ranges: bytes
Vary: Accept-Encoding
Content-Type: text/html
```

<br>
<br>

* `Grabs()` uses the [`curl`](https://curl.se/) command-line-tool and _returns_ the banner thus obtained.


```go
banner, _ := bg.Grabs()

fmt.Println(banner)
```

Output:
```
HTTP/1.1 200 OK
Date: Wed, 20 Oct 2021 05:09:49 GMT
Server: Apache/2.4.7 (Ubuntu)
Accept-Ranges: bytes
Vary: Accept-Encoding
Content-Type: text/html
```

***

### Sample program:
```go
package main

import (
    "fmt"
    "time"

    gobg "github.com/NovusEdge/go-bg"
)

func main() {
    bg := gobg.BannerGrabber{
	    URL: "http://scanme.nmap.org/",
    }

    // Scan and report banner in stdout:
    bg.Grab()

    // Scan and return the banner as a string:
    banner, _ := bg.Grabs()

    // Use in whatever way you like :)
    // For this sample, we'll just print it to stdout:

    fmt.Println(banner)
}
```


