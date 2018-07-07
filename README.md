# gock

## Install

```
go get github.com/mitubaEX/gocrk
```

## Required Slack permission

- channels:read
- chat:write:user
- incoming-webhook

## Usage

Please export your slack token as env variable.

```
export SLACK_TOKEN=YOUR_SLACK_TOKEN
```

```
gocrk <command>
```

## Example

```
$ ./gocrk ls -la
total 13392
drwxr-xr-x  12 mituba  staff      384  7  8 01:45 .
drwxr-xr-x   8 mituba  staff      256  7  1 20:30 ..
drwxr-xr-x  15 mituba  staff      480  7  8 01:45 .git
-rw-r--r--   1 mituba  staff        8  6 30 02:17 .gitignore
drwxr-xr-x   8 mituba  staff      256  7  8 01:42 .idea
-rw-r--r--   1 mituba  staff      581  7  8 00:26 Gopkg.lock
-rw-r--r--   1 mituba  staff      666  7  8 00:21 Gopkg.toml
-rw-r--r--   1 mituba  staff      190  7  8 01:37 README.md
-rwxr-xr-x   1 mituba  staff  6833836  7  8 01:45 gocrk
-rw-r--r--   1 mituba  staff     1103  7  8 01:43 main.go
drwxr-xr-x   3 mituba  staff       96  7  8 00:26 vendor
```
