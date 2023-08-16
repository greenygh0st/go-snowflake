# snowflake
[![Build](https://github.com/greenygh0st/go-snowflake/actions/workflows/go.yml/badge.svg)](https://github.com/greenygh0st/go-snowflake/actions/workflows/go.yml) [![CodeQL](https://github.com/greenygh0st/go-snowflake/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/greenygh0st/go-snowflake/actions/workflows/github-code-scanning/codeql)

A library for creating Snowflake identifiers in Go. Snowflakes are 64 bit identifiers that are guaranteed to be unique across all processes and machines for a given time interval. They are comprised of:
* 41 bits for time in milliseconds (gives us 69 years with a custom epoch)
* 10 bits for a machine id (gives us up to 1024 machines)
* 12 bits of sequence (in case more than one identifier is requested in the same millisecond)

## Usage
#### Importing
```go
package main

import (
    "fmt"
    "github.com/greenygh0st/snowflake"
)

```

#### Usage
```go
node := NewSnowflake(1, 1)
id := node.Generate()
```
