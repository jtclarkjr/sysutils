
# Sysinfo

## Example

```go
package main

import (
    "fmt"

    "github.com/yourusername/sysutils/sysinfo"
)

func main() {
    fmt.Println("Host Information:")
    fmt.Println(sysinfo.GetHostInfo())

    fmt.Println("CPU Information:")
    fmt.Println(sysinfo.GetCPUInfo())

    fmt.Println("Memory Information:")
    fmt.Println(sysinfo.GetMemoryInfo())
}
```
