## disgohook

![GitHub Tag](https://img.shields.io/github/v/tag/orewaee/disgohook?style=flat&color=5865F2) ![GitHub License](https://img.shields.io/github/license/orewaee/disgohook?style=flat&color=5865F2) ![Discord](https://img.shields.io/discord/1172841532827635742?style=flat&label=Singularity%20R%26D%20%F0%9F%91%BD&color=5865F2)

The disgohook package allows you to send [Discord webhooks](https://discord.com/developers/docs/resources/webhook) quickly and easily. It does not use any external libs and is implemented in pure [Go](https://go.dev/).


## 📦 Installation

```bash
go get github.com/orewaee/disgohook
```


## 🤖 Simple example

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
