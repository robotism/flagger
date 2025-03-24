# Flagger

[English](README.md) | [中文](README.cn.md)

Flagger is a library based on [`Viper`](https://github.com/spf13/viper) and [`Pflag`](https://github.com/spf13/pflag) for structured parameter handling.

## Usage

```bash
go get github.com/robotism/flagger
```

## Example Usage [example/cmd/root.go](example/cmd/root.go)

- Parameter Struct

  ```go
  type AppConfig struct {
      Debug    bool   `mapstructure:"debug" short:"d" description:"debug mode" default:"false"`
      Timezone string `mapstructure:"timezone" description:"timezone" default:"UTC"`
      Server Server `mapstructure:"server" group:"server"`
      Database map[string]Database `mapstructure:"database" group:"database" mapkey:"<dbkey>"`
  }

  type Server struct {
      Port int `mapstructure:"port" description:"port" default:"8080"`
  }

  type Database struct {
      Host string `mapstructure:"host" description:"host" default:"localhost"`
      Port int    `mapstructure:"port" description:"port" default:"3306"`
      User string `mapstructure:"user" description:"user" default:"root"`
      Pass string `mapstructure:"pass" description:"pass"`
  }
  ```

- Command Initialization

  ```go
  package cmd

  import (
      "log"
      "os"

      "github.com/robotism/flagger"
      "github.com/robotism/flagger/example/config"
      "github.com/spf13/cobra"
  )

  var (
      f = flagger.New()
      c = &config.AppConfig{}
  )

  // rootCmd represents the base command when called without any subcommands
  var rootCmd = &cobra.Command{
      Use:   "example",
      Short: "a flagger example",
      Long:  `a flagger example`,
      Run: func(cmd *cobra.Command, args []string) {
          log.Printf("%+v\n", c)
      },
  }

  func Execute() {
      err := rootCmd.Execute()
      if err != nil {
          os.Exit(1)
      }
  }

  func init() {
      f.UseFlags(rootCmd.Flags())
      f.UseConfigFileArgDefault()
      f.Parse(c)
  }
  ```

- Run: Display Help

  ```bash
  > go run main.go -h

  a flagger example

  Usage:
  example [flags]

  Flags:
  -c, --config string                  config file path
      --database.<dbkey>.host string   host (default "localhost")
      --database.<dbkey>.pass string   pass
      --database.<dbkey>.port int      port (default 3306)
      --database.<dbkey>.user string   user (default "root")
  -d, --debug                          debug mode
  -h, --help                           help for example
      --server.port int                port (default 8080)
      --timezone string                timezone (default "UTC")

  ```

- Run: Use Environment Variables and Command Line Parameters

  ```bash
  > set SERVER_PORT=9999
  > go run main.go -d=true --timezone=Asia/Shanghai --database.default.host=127.0.0.1

  2025/02/20 14:07:31 &{Debug:true Timezone:Asia/Shanghai Server:{Port:9999} Database:map[default:{Host:127.0.0.1 Port:3306 User:root Pass:}]}
  ```

- Run: Use Configuration File

  ```bash
  go run main.go -d=true --timezone=Asia/Shanghai --database.default.host=127.0.0.1 -c config.yaml

  2025/02/20 14:07:59 &{Debug:true Timezone:Asia/Shanghai Server:{Port:9999} Database:map[default:{Host:127.0.0.1 Port:4000 User:root Pass:12345678}]}

  go run main.go -c config.yaml

  2025/02/20 14:08:14 &{Debug:true Timezone:UTC Server:{Port:9999} Database:map[default:{Host:xxx.xxx.xxx.xxx Port:4000 User:root Pass:12345678}]}
  ```

## Loading Priority:

- **Command Line Flags (Flags)**: If a configuration option is provided via the command line, the value from the command line will take precedence.
- **Environment Variables**: If the configuration option is not provided via the command line, it will attempt to retrieve the value from environment variables.
- **Configuration File**: If the configuration option is not found in environment variables, it will load from the configuration file.
- **Default Values**: If none of the above methods provide the value, the default value set in the code will be used.
