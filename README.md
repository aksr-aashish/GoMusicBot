# GoMusicBot

An example bot using [GoTGCalls](https://github.com/gotgcalls/gotgcalls).

---

## Configuring

Copy `example.env` to `.env` and add your credentials.

- `STRING_SESSION`: A GramJS/Telethon string session. You can generate one [here](https://rojserbest.github.io/bssg).
- `API_ID`: Telegram app ID.
- `API_HASH`: Telegram app hash.
- `BOT_TOKEN`: Telegram bot token.

## Running

1. [Install](https://github.com/gotgcalls/gotgcalls#installation) GoTGCalls.

2. Start:

```bash
go run .
```

## Commands

| Command | Description                   |
| ------- | ----------------------------- |
| stream  | Stream the replied audio file |
| pause   | Pause streaming               |
| resume  | Resume streaming              |
| skip    | Skip the current playback     |
