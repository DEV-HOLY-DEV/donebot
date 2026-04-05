# donebot

`donebot` is a simple Go CLI utility that wraps the execution of any terminal command and can notify a Discord channel via webhook when the command finishes. It is highly useful for monitoring long-running tasks.

## Usage

You can prefix any terminal command with `donebot` to execute it:

```bash
donebot -dis npm run build
donebot -dis docker-compose up -d
```