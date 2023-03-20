# confmap
Configuration mapping on the Go struct based on [viper](https://github.com/spf13/viper)

# Example usage

Configuration struct:
```go
type HttpClientConfig struct {
    Url     string `mapstructure:"url"`
    Timeout string `mapstructure:"timeout"`
}
```

### Map parameter from environment variables 
Environment variables:

```
export PREFIX_URL=example.com
export PREFIX_TIMEOUT=200ms
```

Init code:
```go
var config HttpClientConfig

if err := confmap.MapEnvs("PREFIX", &config); err != nil {
    panic(err)
}

client := NewHttpClient(config)
```

With Generics:
```go
config, err := confmap.MapEnvsTyped[HttpClientConfig]("PREFIX")
if err != nil {
    panic(err)
}

client := NewHttpClient(config)
```
### Default parameters
```go
type HttpClientConfig struct {
    Url     string `mapstructure:"url"`
    Timeout string `mapstructure:"timeout"`
}

func (c Http) DefaultEnvs() map[string]any {
    return map[string]any{
        "timeout": "1s",
    }
}
```