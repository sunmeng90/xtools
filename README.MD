## Example

```cmd
./x --config ../config/x.yaml stringer reverse abc
```

## Concept(Cobra)

**Commands** represent actions, **Args** are things and **Flags** are modifiers for those actions.  
The pattern to follow is `APPNAME VERB NOUN --ADJECTIVE`. or `APPNAME COMMAND ARG --FLAG`

### Commands

Command is the central point of the application. Each interaction that the application supports will be contained in a
Command. A command can have children commands and optionally run an action.

### Flags

A flag is a way to modify the behavior of a command. A Cobra command can define flags that persist through to children
commands and flags that are only available to that command.

#### Persistent Flags

A flag can be ‘persistent’, meaning that this flag will be available to the command it's assigned to as well as every
command under that command. For global flags, assign a flag as a persistent flag on the root.

#### Local Flags

A flag can also be assigned locally, which will only apply to that specific command.

### Args

Validation of positional arguments can be specified using the Args field of Command.

The following validators are built in:

* `NoArgs` - the command will report an error if there are any positional args.
* `ArbitraryArgs` - the command will accept any args.
* `OnlyValidArgs` - the command will report an error if there are any positional args that are not in the ValidArgs
  field of
  Command.
* `MinimumNArgs(int)` - the command will report an error if there are not at least N positional args.
* `MaximumNArgs(int)` - the command will report an error if there are more than N positional args.
* `ExactArgs(int)` - the command will report an error if there are not exactly N positional args.
* `ExactValidArgs(int)` - the command will report an error if there are not exactly N positional args OR if there are
  any
  positional args that are not in the ValidArgs field of Command
* `RangeArgs(min, max)` - the command will report an error if the number of args is not between the minimum and maximum
  number of expected args.

## Config(Viper)

Viper uses the following precedence order. Each item takes precedence over the item below it:

* explicit call to `Set`
* flag
* env
* config
* key/value store
* default

Config Overrides

```go
viper.Set("Verbose", true)
viper.Set("LogFile", LogFile)
```

Bind to flags(Pflags)
Viper has the ability to bind to flags. Specifically, Viper supports Pflags as used in the Cobra library.

```go 
serverCmd.Flags().Int("port", 1138, "Port to run Application server on")
viper.BindPFlag("port", serverCmd.Flags().Lookup("port"))
```

Get value from viper
In Viper, there are a few ways to get a value depending on the value’s type. The following functions and methods exist:

```go
Get(key string) : interface{}
GetBool(key string) : bool
GetFloat64(key string) : float64
GetInt(key string) : int
GetIntSlice(key string) : []int
GetString(key string) : string
GetStringMap(key string) : map[string]interface{}
GetStringMapString(key string) : map[string]string
GetStringSlice(key string) : []string
GetTime(key string) : time.Time
GetDuration(key string) : time.Duration
IsSet(key string) : bool
AllSettings() : map[string]interface{}
```

One important thing to recognize is that each Get function will return a zero value if it’s not found. To check if a
given key exists, the `IsSet()` method has been provided.

Viper can get subset of the configuration

```go 
cache1Config := viper.Sub("cache.cache1")
if cache1Config == nil { // Sub returns nil if the key cannot be found
	panic("cache configuration not found")
}

cache1 := NewCache(cache1Config)
```

## Build(Cross platform)

```cmd
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64

go build -o ./bin/x main.go
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