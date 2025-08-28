# gosparse
Cross-platform sparse files for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/CAFxX/gosparse.svg)](https://pkg.go.dev/github.com/CAFxX/gosparse)

---

Quick sanity check:

```bash
while IFS=/ read -r os arch; do
    echo $os/$arch
    GOOS=$os GOARCH=$arch go build -trimpath -o obj
    GOOS=$os GOARCH=$arch go tool objdump -s '[Pp]unchHole' obj > $os-$arch.asm
    rm obj
done < <(go tool dist list)
```
