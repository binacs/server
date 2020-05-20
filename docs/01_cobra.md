# spf13 - cobra



[GoDoc](https://godoc.org/github.com/spf13/cobra)



## 1. 概述

略



## 2. 关键概念

3 个基本概念：

- 命令（Command）：就是需要执行的操作； 

- 参数（Arg）：命令的参数，即要操作的对象； 

- 选项（Flag）：命令选项可以调整命令的行为。 



下面示例中， `start` 是一个（子）命令， `--configFile` 是选项：

```shell
server start --configFile=./config.toml
```



下面示例中， `clone` 是一个（子）命令， `URL` 是参数， `--bare` 是选项：

```shell
git clone URL --bare
```



## 3. 命令 Command

通过以下方式定义一个命令：

```go
	RootCmd = &cobra.Command{
		Use:   "root",
		Short: "Root Command",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
      // ...
			return nil
		},
	}
```

定义了一个 RootCmd 命令，并指明 Use (如何使用)、Short (简要描述)、PersistentPreRunE (功能函数) 等。

> 命令执行 `Execute` 时会去执行 Run 字段定义的回调函数，此外 Cobra 还提供了四个函数：PersistentPreRun、PreRun、PostRun、PersistentPostRun，可以在执行这个回调函数 Run 之前和之后执行。
>
> 它们的执行顺序依次是：PersistentPreRun、PreRun、Run、PostRun、PersistentPostRun。而且对于 PersistentPreRun 和 PersistentPostRun ，子命令是继承的，即子命令如果没有自定义自己的 PersistentPreRun 和 PersistentPostRun ，那它就会执行父命令的这两个函数。 
>
> 一般而言，PersistentPreRunE 先于 Run 执行，作为根命令，只完成 PersistentPreRunE 指定的检查、初始化日志系统并缓存配置的功能，和 Run 指定的版本打印、命令帮助功能。
>
> 更多参阅：https://godoc.org/github.com/spf13/cobra#Command 中关于结构体成员的注释。



通过  `AddCommand()` 方法为命令添加子命令，通过 `Execute()` 执行命令。



## 4. 参数 Args



1. 先定义再引用：

```go
var OrderingEndpoint string
var clientAuth bool

// 1.1
func AddOrdererFlags(cmd *cobra.Command) {
  	// get flags, then xxxVarP
    flags := cmd.PersistentFlags()
    flags.StringVarP(&OrderingEndpoint, "orderer", "o", "", "Ordering service endpoint")
    flags.BoolVarP(&clientAuth, "clientauth", "", false, "Use mutual TLS when communicating with the orderer endpoint")
}

// 1.2
func rootCmdFlags(cmd *cobra.Command) {
  	// cmd.xxxFlags.xxVar
		cmd.PersistentFlags().StringVar(&configFile, "configFile", "config.toml", "config file (default is ./config.toml)")
}

// Eg: 
// StringVarP 用来接收类型为字符串变量的标志。
// 相较StringVar， StringVarP 支持标志短写。
//	以 OrderingEndpoint 例：在指定标志时可以用 --orderer ，也可以使用短写 -o。
```



更多变量类型使用请参考[这里](https://www.godoc.org/github.com/spf13/pflag#FlagSet)。



## 5. 选项 Flag

选项分为两种： 全局参数(Persistent Flags)、局部参数(Local Flags)。

前者具有持久性，将被该被分配的命令及其所有子命令生效；后者只适用于被分配的命令本身。



在创建 `cobra.Command` 时，可以使用 Args 选项自定义参数验证器。形如：

```go
var cmd = &cobra.Command{
  Short: "hello",
  Args: func(cmd *cobra.Command, args []string) error {
    if len(args) < 1 {
      return errors.New("requires at least one arg")
    }
    return nil
  },
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Hello, World!")
  },
}
```

更多内置验证器请参阅 godoc 。