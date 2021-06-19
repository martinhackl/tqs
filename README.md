# TQS (Tmux Quick Sessions)

Quickly create predefined Tmux sessions with a config file.

## Usage

```
$ tqs -s user=myuser -s foo=bar ./examples/workspace.tqs.json
```

`-s` (optional) defines variables which will be replaced within window paths.
Within the config file each variable must be surrounded with `{}`.

## Config file

See `examples/workspace.tgs.json` for details.