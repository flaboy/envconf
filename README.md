# env-conf-reader

1. Read system-env first.
2. Parse config file second.
3. Support nested structure with environment key concatenation.

## Basic Usage

```go
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

## Nested Structure Support

```go
type StorageConfig struct {
	Type  string `cfg:"TYPE" default:"local"`
	Local LocalStorage  `cfg:"LOCAL"`
	S3    S3Storage    `cfg:"S3"`
}

type LocalStorage struct {
	BasePath string `cfg:"BASE_PATH" default:"storage"`
	BaseURL  string `cfg:"BASE_URL" default:"/storage"`
}

type S3Storage struct {
	AccessKey string `cfg:"ACCESS_KEY"`
	SecretKey string `cfg:"SECRET_KEY"`
	Bucket    string `cfg:"BUCKET"`
	Region    string `cfg:"REGION"`
	Endpoint  string `cfg:"ENDPOINT"`
	PublicURL string `cfg:"PUBLIC_URL"`
}

type Config struct {
	Storage StorageConfig `cfg:"STORAGE"`
}
```

Environment variables are concatenated with underscore:

```
STORAGE_TYPE=s3
STORAGE_S3_ACCESS_KEY=your-access-key
STORAGE_S3_SECRET_KEY=your-secret-key
STORAGE_S3_BUCKET=your-bucket
STORAGE_S3_REGION=your-region
STORAGE_S3_ENDPOINT=your-endpoint
STORAGE_S3_PUBLIC_URL=your-public-url
```

## Configuration Files

### env.conf
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

## Environment Variables Override

Environment variables take precedence over config file values:

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