# env-conf-reader

1. Read system-env first.
1. Parse config file second.

```
import "github.com/flaboy/envconf"

type config struct {
	STR_CFG   string
	INT_CFG   int
	FLOAT_CFG float64

	BOOL_CFG1 bool
	BOOL_CFG2 bool
	BOOL_CFG3 bool
	BOOL_CFG4 bool
	BOOL_CFG5 bool

	CustomVar string            `cfg:"CUSTOM_CFG"`
	JsonVar   map[string]string `cfg:"EXAMPLE_JSON_CFG"`
	JsonVar2  []string          `cfg:"EXAMPLE_JSON_CFG2"`
}


var Config *config

func main() {
	Config = &config{}
	err := envconf.Load("env.conf", Config)
}

```

env.conf
```
# comment
STR_CFG =   str-value
INT_CFG =   123
FLOAT_CFG= 12.2 #inline-comment

BOOL_CFG1=yes
BOOL_CFG2=1
BOOL_CFG3=true
BOOL_CFG4=on
BOOL_CFG5=blabla

CUSTOM_CFG = foo-value

# if not string/int/float/bool then parse as JSON
EXAMPLE_JSON_CFG={"a":"b","c":"d","e":"f"}
EXAMPLE_JSON_CFG2=["aaa", "bbb", "ccc"]
```


```
STR_CFG=3333 ./example

{
    "STR_CFG": "3333",
    "INT_CFG": 123,
    "FLOAT_CFG": 12.2,
    "CustomVar": "foo-value",
    "JsonVar": {
        "a": "b",
        "c": "d",
        "e": "f"
    },
    "JsonVar2": [
        "aaa",
        "bbb",
        "ccc"
    ]
}
```