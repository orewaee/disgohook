## disgohook

The disgohook package allows you to send [Discord webhooks](https://discord.com/developers/docs/resources/webhook) quickly and easily. It does not use any external libs and is implemented in pure [Go](https://go.dev/).


## Installation

```bash
go get -u github.com/orewaee/disgohook
```

## Simple example

The disgohook package uses builders to generate the correct webhook structure. For example, to send a message with an embed, you can do this:

```go
package main

import (
	"github.com/orewaee/disgohook"
	"log"
)

func main() {
    embed := disgohook.
        NewEmbedBuilder().
        SetColor(0x5865f2).
        SetTitle("something very important").
        SetDescription("disgohook is already here :alien:").
        Build()
    
    webhook := disgohook.
        FromIdAndToken("your webhook id", "your webhook token").
        SetEmbeds(embed).
        Build()
    
    if err := webhook.Send(); err != nil {
        log.Fatalln(err)
    }
}
```
