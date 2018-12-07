<br><br>

<h1 align="center">Casec</h1>

<p align="center">
  <a href="/LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg"/></a>
  <a href="https://godoc.org/github.com/mingrammer/casec"><img src="https://godoc.org/github.com/mingrammer/casec?status.svg"/></a>
  <a href="https://goreportcard.com/report/github.com/mingrammer/casec"><img src="https://goreportcard.com/badge/github.com/mingrammer/casec"/></a>
  <a href="https://travis-ci.org/mingrammer/casec"><img src="https://travis-ci.org/mingrammer/casec.svg?branch=master"/></a>
</p>


<p align="center">
A text case converter
</p>
<br><br><br>

casec is a text case converter for programmers. casec now supports  `upper`, `lower`, `title`, `camel`, `pascal`, `snake`, `kebab` (or `lisp`) cases.

It also provides case conversion library not only command line tool.

## Installation

### Using go get

> Go version 1.10 or higher is required.

```
go get github.com/mingrammer/casec/...
```

If you want to only download the `casec` library

```
go get github.com/mingrammer/case
```

### Using [homebrew](https://brew.sh)

```
brew tap mingrammer/casec
brew install casec
```

### Using .tar.gz archive

Download gzip file from [Github Releases](https://github.com/mingrammer/casec/releases/latest) according to your OS. Then, copy the unzipped executable to under system path.

## Usage

### CLI

Convert all words to snake case.

```bash
$ casec -t snake main.py
```

Convert all snake case to camel case.

```bash
$ casec -f snake -t camel internal.go
```

Show diff between before and after conversion without actually applying. (dry-run)

```bash
$ casec -f snake -t kebab -d match.lisp
```

Convert all camel case to snake case except for words you don't want to convert. It is useful for preventing the keywords (reserved words) or conventions from converting.

```bash
$ casec -f snake -t pascal -i '^package|var|const|if|for|range|return|func|go$' redis.go
```

You can pass multiple ignore expressions.

```bash
$ casec -f snake -t pascal -i '^package|var|const|if|for|range|return|func|go$' -i '^github|com$' redis.go
```

### Library

> See details in [GoDoc](https://godoc.org/github.com/mingrammer/casec)

```go
package main

import (
    "fmt"

    "github.com/mingrammer/casec"
)

func main() {
    fmt.Println(casec.IsSnakeCase("this_is_snake"))
    // Output: true
    fmt.Println(casec.IsCamelCase("thisIsNot_camelCase"))
    // Output: false
    fmt.Println(casec.Invert("Invert Me"))
    // Output: iNVERT mE
    fmt.Println(casec.ToSnakeCase("IPAddress"))
    // Output: ip_address
}
```

## Known issues

casec separates the words with non-letters (except `-` and `_`) including `.` and `/` letters. So, the `ignore` option of casec can not recognize the dot-or-slash separated word (ex. `"github.com/mingrammer/cfmt"`) as a single chunk. So if you want to prevent the import path of Go source code, for example, `import "github.com/mingrammer/cfmt"` from converting,  you should pass the ignore expression as `-i "^github|com|mingrammer|cfmt$"`.

Here are the solutions that I'll consider making it an option for solving this issue.

1. Treat the string surrounded with quotes ("" or '') as a single word optionally.
2. Make an option for specifying the line number ranges for applying the conversion.

## License

MIT
