go-windows-dbg
==============

This package provides the functions to output text to debuggers on Windows

( Renamed and simplified from [go-outputdebug](https://github.com/zetamatta/go-outputdebug) )

+ On Windows, the functions call the OutputDebugStringW API
+ With the compile option `-tags=ndebug`, debug strings are discarded
+ On non-Windows operating systems, the functions do nothing
+ You can check if debugging is enabled or disabled by referring to `dbg.Enabled`

[Example-1](./example.go)
--------------------------------

```example.go
package main

import (
    "github.com/nyaosorg/go-windows-dbg"
)

func main() {
    dbg.Print("output", 1, "text")
    dbg.Printf("output<%d>text", 1)
    dbg.Println("output", 1, "text")

    println("Enabled=", dbg.Enabled)
}
```

```
> start dbgview.exe
> go run example.go
```

Screenshot of [DebugView for Windows](https://technet.microsoft.com/ja-jp/sysinternals/debugview.aspx)

![screen shot](./screenshot.png)

[Example-2](./example2.go)
--------------------------

`dbg.X` is the small function similar with [`dbg!` on Rust](https://doc.rust-lang.org/std/macro.dbg.html).

```example2.go
package main

import (
    "github.com/nyaosorg/go-windows-dbg"
)

func main() {
    println(dbg.X(1 + 2))
}
```

```
> start dbgview.exe
> go run example2.go
3
```

![screen shot](./screenshot2.png)
