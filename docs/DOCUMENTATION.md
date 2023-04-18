
## math/rand vs crypto/rand

This code generates a password with a length of 16 characters by creating a
byte slice of length passwordLength and calling rand.Read() to fill it with
random bytes.

```golang

package main

import (
    "crypto/rand"
    "fmt"
)

func main() {
    const passwordLength = 16

    password := make([]byte, passwordLength)
    _, err := rand.Read(password)
    if err != nil {
        panic(err)
    }

    fmt.Println(string(password))
}

```

The generated password is printed to the console using fmt.Println(). Note that
rand.Read() returns the number of bytes read and an error, which we ignore by
using the blank identifier \(_\). Also, rand.Read() returns a cryptographically
secure random byte slice, which is more suitable for generating passwords than
using the 'math/rand' package.

