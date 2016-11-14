## Developing

**Requirements:**

* Golang 1.7.3
* Dog (https://github.com/dogtools/dog)

Once you have the requirements, you can simply do:

```
dog deps
dog build
```

And now you should have a bunch of binaries in the `dist/` folder.

## build

```
go-bindata -o assets/assets.go -pkg assets redirect-to-home
go build -o quickwiki main.go
```
