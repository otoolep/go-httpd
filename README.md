go-httpd [![Circle CI](https://circleci.com/gh/otoolep/go-httpd/tree/master.svg?style=svg)](https://circleci.com/gh/otoolep/go-httpd/tree/master) [![GoDoc](https://godoc.org/github.com/otoolep/go-httpd?status.png)](https://godoc.org/github.com/otoolep/go-httpd) [![Go Report Card](https://goreportcard.com/badge/github.com/otoolep/go-httpd)](https://goreportcard.com/report/github.com/otoolep/go-httpd)
======

go-httpd is an example [Go](http://golang.org) project, showing you how to organise that most basic of systems -- a HTTP server that allows you to set and get values in a key-value store. Trivial _Hello, World_ programs don't demonstrate how to organise a real program, but go-httpd does. It also exhibits many other important principles of a good Go project.

## Project layout
`main.go`, within the _main_ package, is the entry point for the program, and this source file is at the top level of the tree.

There are other possible paths for this main package, depending on your requirements. One other suggested place is in `cmd/go-httpd/`. This can be convenient if your project builds more than one binary -- a CLI for example. Then the CLI can be located in `cmd/go-http-cli`. Two examples of projects using this structure are [rqlite](http://github.com/rqlite/rqlite) and [InfluxDB](https://github.com/influxdata/influxdb).

Otherwise the remainder of the project is contained in two sub-packages -- _http_ and _store_.

### Project name
It's best not to call your project `go-<something>`. Just call it `<something>`. I called this project `go-httpd` to make the connection with Go clear.

## Running go-httpd
*Building go-httpd requires Go 1.5 or later. [gvm](https://github.com/moovweb/gvm) is a great tool for installing and managing your versions of Go.*

Starting and running go-httpd is easy. Download and build it like so:
```
mkdir go-httpd # Or any directory of your choice
cd go-httpd/
export GOPATH=$PWD
go get github.com/otoolep/go-httpd
```
Some people consider using a distinct `GOPATH` environment variable for each project _doing it wrong_. In practise I, and many other Go programmers, find this actually most convenient.

Run it like so:
```
$GOPATH/bin/go-httpd
```

You can now set a key and read its value back as follows:
```
curl -XPOST localhost:8080/key -d '{"user1": "batman"}'
curl -XGET localhost:8080/key/user1
```

### Building and rebuilding
Once you've downloaded the code from GitHub, you'll probably want to change the code and rebuild it. The easiest way to do this is to execute the following commands:
```
cd $GOPATH/src/github.com/otoolep/go-httpd
go build
./go-httpd
```
To build __and__ install the binary in `$GOPATH/bin`, instead execute `go install`.

### Development environments
vim has good support for Go, via a [plugin](https://github.com/fatih/vim-go). Go support for Sublime 3 is also available via [GoSublime](https://github.com/DisposaBoy/GoSublime).

## Use the standard library
Go comes with a high-quality standard library, with support for much important functionality such as networking, compression, JSON serialization, encryption, IO, concurrency, and synchronization. It is better to use the standard library even if it means having to write a little bit of extra code, than to import a third-party library. You can learn more about this philosophy [here](https://blog.gopheracademy.com/advent-2014/case-against-3pl/).

You can see this principle at work in go-httpd. It uses the standard _logging_ package, _testing_ package, and does its own HTTP routing, even though there are literally hundreds of non-standard packages that claim to do a better job in each area. In my experience just stick with the standard library if at all possible.

## Logging
Each package prefixes its log output with an identifiable string, and sends the output to `stderr`. Many operators find this most convenient, as **they** can then decide where to send the log output of your program.

## Testing
Each package is tested via the framework that comes with the standard library, using simple test cases. A full test suite would involve many more tests.

To run the test suite execute the following command:
```
cd $GOPATH/src/github.com/otoolep/go-httpd
go test -v ./...
```

### Interfaces
Interfaces are a key concept within Go, and allow components to interact in a very natural way. The [_io.Reader_](https://golang.org/pkg/io/#Reader) and [_io.Writer_](https://golang.org/pkg/io/#Writer) interfaces are the most important examples in the standard library.

Interfaces are also very useful for testing. The HTTP service within go-httpd does not import the Store directly, but instead specifies the interface any Store must support. This makes it very easy to pass a mock store to the HTTP service, as part of this testing.

## Documentation
The _GoDoc_ standard is very easy to follow and results in nice, easy to read, documentation.

go-httpd is documented in the GoDoc style. The GoDoc link at the top of this file is automatically generated from the comments in the source. Check the source of the README file itself to see how to add these badges to your own projects.

## Pre-commit hooks
Within the `hooks` directory is a git _pre-commit hook_. Installing this hook means that your code is checked for [_go fmt_](https://golang.org/cmd/gofmt/) and _go vet_ errors before it can even be committed. While these checks can be run manually at any time, installing them as hooks ensures you don't forget.

To install the hook execute `make install` from within the `hooks directory`.

## CircleCI integration
[CircleCI](http://www.circleci.com) supports basic Go testing without any extra work on your part, but that testing can be more sophisticated. The example `yml` file included in this repository shows how to instruct CircleCI (version 1.0) to perform extra testing such as checking for formatting and linting issues. It also instructs CircleCI to run your code through Go's [race detection](https://blog.golang.org/race-detector) system.

## Go Programming References
Be sure to check out the following references:
* The [standard docs](https://golang.org/pkg/). You really don't need much else.
* [Effective Go](https://golang.org/doc/effective_go.html). You should read this when you first start programming in Go, and then read it again 3 months later. And then read it again 6 months after that.
* [How to Write Go](https://golang.org/doc/code.html).
* [The Go Programming Language](http://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing/dp/0134190440).
* The [Go Playground](https://play.golang.org/) -- a useful place to share snippets of Go code.
* [rqlite](http://github.com/rqlite/rqlite) -- a distributed database, with SQLite as its storage engine. rqlite is a much more sophisticated example of production-quality Go code, which strives to follow all the principles outlined above.
* [400 Days of Go](http://www.philipotoole.com/400-days-of-go/) -- a blog post I wrote on why I like Go.
