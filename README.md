# mercurylampe

A Telegram bot written in Go using [go-telegram/bot](https://github.com/go-telegram/bot).

## Commands

| Command | Behavior |
|---|---|
| `/cat` | Sends a random cat photo |
| `/gcat` | Sends a random cat gif |
| `/tcat` | Sends a cat photo captioned with the given text |

Any other message is echoed back by the default handler.

## Running

```sh
export TG_TOKEN=<telegram-bot-token>
go run .
```

With Nix:

```sh
nix develop   # go, gopls, nil in scope
go run .
```
