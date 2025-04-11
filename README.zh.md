# protoenum

`protoenum` 是一个 Go 语言包，提供管理 Protobuf 枚举元数据的工具。它将 Protobuf 枚举值与自定义描述包装在一起，并通过 `Enums` 结构体支持按代码、名称或描述进行快速查找。

## 安装

通过以下命令安装包：

```sh
go get github.com/go-xlan/protoenum
```

确保您的 Go 环境已正确配置。

## 使用方法

### 单个枚举描述

创建带有自定义描述的枚举描述符：

```go
import "github.com/go-xlan/protoenum"

status := protoenum.NewEnum(yourpackage.StatusEnum_SUCCESS, "成功")
println(status.Code()) // 输出: 枚举的数值代码
println(status.Name()) // 输出: SUCCESS
println(status.Desc()) // 输出: 成功
```

### 枚举集合

管理多个枚举：

```go
enums := protoenum.NewEnums(
    protoenum.NewEnum(yourpackage.StatusEnum_SUCCESS, "成功"),
    protoenum.NewEnum(yourpackage.StatusEnum_FAILURE, "失败"),
)

// 查找示例
println(enums.GetByCode(1).Desc())  // 输出: 成功
println(enums.GetByName("FAILURE").Desc()) // 输出: 失败
```

## 主要功能

- **Enum**：将 Protobuf 枚举值与自定义描述包装。
    - `Code()`：获取枚举的数值代码。
    - `Name()`：获取枚举的名称。
    - `Desc()`：获取自定义描述。
- **Enums**：支持按代码、名称或描述查找枚举的集合。

## 许可证

MIT 许可证。详情请见 [LICENSE](LICENSE) 文件。

## 贡献与支持

欢迎通过提交 pull request 或报告问题来贡献此项目。

如果你觉得这个包对你有帮助，请在 GitHub 上给个 ⭐，感谢支持！！！

**感谢你的支持！**

**祝编程愉快！** 🎉

Give me stars. Thank you!!!
