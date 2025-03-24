# Flagger

[English](README.md) | [中文](README.cn.md)

Flagger 是一个基于 [`Viper`](https://github.com/spf13/viper) 和 [`Pflag`](https://github.com/spf13/pflag) 的用于结构化参数的库。

## 使用

```bash
go get github.com/robotism/flagger
```

## 使用示例 [example/cmd/root.go](example/cmd/root.go)

- 参数结构体

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

- 命令初始化

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

- 运行: 打印帮助

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

- 运行: 使用环境变量和命令行参数

  ```bash
  > set SERVER_PORT=9999
  > go run main.go -d=true --timezone=Asia/Shanghai --database.default.host=127.0.0.1

  2025/02/20 14:07:31 &{Debug:true Timezone:Asia/Shanghai Server:{Port:9999} Database:map[default:{Host:127.0.0.1 Port:3306 User:root Pass:}]}
  ```

- 运行: 使用配置文件

  ```bash
  go run main.go -d=true --timezone=Asia/Shanghai --database.default.host=127.0.0.1 -c config.yaml

  2025/02/20 14:07:59 &{Debug:true Timezone:Asia/Shanghai Server:{Port:9999} Database:map[default:{Host:127.0.0.1 Port:4000 User:root Pass:12345678}]}

  go run main.go -c config.yaml

  2025/02/20 14:08:14 &{Debug:true Timezone:UTC Server:{Port:9999} Database:map[default:{Host:xxx.xxx.xxx.xxx Port:4000 User:root Pass:12345678}]}

  ```

## 加载优先级:

- 命令行标志（Flags）：如果在命令行中提供了某个配置项的值，会优先使用该值。
- 环境变量（Environment Variables）：如果命令行中未提供该配置项的值，会尝试从环境变量中获取。
- 配置文件（Config File）：如果环境变量中也未找到该配置项的值，会从配置文件中加载。
- 默认值（Default Values）：如果上述三者都未提供该配置项的值，会使用预先设置的默认值。
