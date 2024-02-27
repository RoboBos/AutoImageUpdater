package main

import (
    "fmt"
    log "github.com/sirupsen/logrus"
	"github.com/Masterminds/semver/v3"
)


// "flag"
// "fmt"
// "html"
// "log"
// "net/http"
// "os"
// "runtime/debug"
// "strings"

// func usage() {
// 	fmt.Fprintf(os.Stderr, "usage: helloserver [options]\n")
// 	flag.PrintDefaults()
// 	os.Exit(2)
// }

// var (
// 	greeting = flag.String("g", "Hello", "Greet with `greeting`")
// 	addr     = flag.String("addr", "localhost:8080", "address to serve")
// )


type VersioningSchemesValidtor func(string) bool

func SemVersionVersioningSchemesValidtor(version string) bool {
	var isSemVersionVersioningSchemesValidtor = true
	_, err := semver.NewVersion(version)
	if err != nil {
		log.Debugf("could not parse input tag %s as semver: %v", version, err)
        isSemVersionVersioningSchemesValidtor = false
    }
	return isSemVersionVersioningSchemesValidtor
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)

	results, err := GetImageVersions("library", "zookeeper", "100", SemVersionVersioningSchemesValidtor)

	if err != nil {
		panic(err)
	}

	fmt.Println(results)
	// // Parse flags.
	// flag.Usage = usage
	// flag.Parse()

	// // Parse and validate arguments (none).
	// args := flag.Args()
	// if len(args) != 0 {
	// 	usage()
	// }

	// // Register handlers.
	// // All requests not otherwise mapped with go to greet.
	// // /version is mapped specifically to version.
	// http.HandleFunc("/", greet)
	// http.HandleFunc("/version", version)

	// log.Printf("serving http://%s\n", *addr)
	// log.Fatal(http.ListenAndServe(*addr, nil))
}

// func version(w http.ResponseWriter, r *http.Request) {
// 	info, ok := debug.ReadBuildInfo()
// 	if !ok {
// 		http.Error(w, "no build information available", 500)
// 		return
// 	}

// 	fmt.Fprintf(w, "<!DOCTYPE html>\n<pre>\n")
// 	fmt.Fprintf(w, "%s\n", html.EscapeString(info.String()))
// }

// func greet(w http.ResponseWriter, r *http.Request) {
// 	name := strings.Trim(r.URL.Path, "/")
// 	if name == "" {
// 		name = "Gopher"
// 	}

// 	fmt.Fprintf(w, "<!DOCTYPE html>\n")
// 	fmt.Fprintf(w, "%s, %s!\n", *greeting, html.EscapeString(name))
// }
