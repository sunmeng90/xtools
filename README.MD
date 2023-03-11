
## Build(Cross platform)

```cmd
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64

go build -o /bin/x main.go
```

```bash
# print a list of the possible platforms that Go can build on.
go tool dist list

# compile for windows 64-bit
GOOS=windows GOARCH=amd64 go build -o bin/x main.go

# 64-bit
GOOS=linux GOARCH=amd64 go build -o bin/x main.go

# 32-bit
GOOS=linux GOARCH=386 go build -o bin/x32 main.go
```

set version at compile time
```bash
# build the xtools CLI in version 0.0.2
go build -o ./bin/x -ldflags="-X 'github.com/sunmeng90/go/xtools/cmd.version=0.0.2'" main.go
```

ref:
https://umarcor.github.io/cobra/