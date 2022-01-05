
# Gorude

CLI app like burpsuite's intruder attack.

Gorude is not a plugin for burpsuite.



## Features

- Supports only sniper attack for now
- Works standalone. It doesn't need burpsuite

## Usage/Examples

How to create an intercept file:

![intercept.txt](https://raw.githubusercontent.com/canack/sobot/main/get-intercept.gif)

Also you can add multiple payload positions everywhere you want

---
---

```go
package main

import (
	gorude "github.com/canack/gorude/pkg"
)

func main() {

	newTarget := gorude.New()

	newTarget.Config.Threads = 10
	newTarget.Config.SleepInterval = 0.1

	payloads, _ := gorude.LoadList("payloads.txt")

	newTarget.ParseFile("intercept.txt").Sniper(payloads)

}
```



## Roadmap

- Burpsuite's all intruder functions will be add
## Authors

- [@canack](https://www.github.com/canack)

