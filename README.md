# Lorem

CLI to generate lorem ipsum text.

## Installation

```bash
go install github.com/mskelton/lorem@latest
```

## Usage

### Generate paragraph

To generate a paragraph of lorem text, simply call the `lorem` command with an
optional count. If a count is specified, that many paragraphs will be printed.

```bash
lorem [count]
```

### Generate sentence

To generate a sentence of lorem text, add the `-s` flag with an optional count.
If a count is specified, that many sentences will be printed.

```bash
lorem -s [count]
```

