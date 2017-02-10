# cog-go-rundeck

[Cog](https://github.com/operable/cog) bundle for Rundeck.
Created with [go-rundeck](https://github.com/lusis/go-rundeck).

## Notes

All commands output JSON where appropriate, but most have templates
to display human-readable output. If you want to get values from the
JSON (e.g., to pipe somewhere else), pipe to `operable:raw`:

```
rundeck:list | operable:raw
```

## Docker container

You can also run this as a Docker container:

```
docker run -e RUNDECK_URL=http://rundeck:4440 \
  -e RUNDECK_TOKEN=somelongtokenhere \
  steeef/cog-go-rundeck \
  /bundle/cog-go-rundeck list
```