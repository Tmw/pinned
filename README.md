# Pinned

In short; Pinned grabs pinned messages from slack and turns them
into a multiple choice quiz where you need to guess the original author.

It is still very much a work in progress :)

## Building

```bash
docker build . -t pinned
```

## Running

```bash
docker run --rm -d \
  -e "SLACK_TOKEN=XXX" \
  -e "SLACK_CHANNEL=general" \
  -p 4000:4000 \
  -t tiemenwaterreus/pinned:latest
```

## License

[MIT](./LICENSE)
