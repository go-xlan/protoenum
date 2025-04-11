# protoenum

`protoenum` provides utilities for managing Protobuf enum metadata in Go. It wraps Protobuf enum values with custom descriptions and offers a enums-map for easy lookup by code, name, description.

## Installation

Install the package with:

```sh
go get github.com/go-xlan/protoenum
```

## Usage

### Single Enum

Create an enum descriptor with a custom description:

```go
import "github.com/go-xlan/protoenum"

status := protoenum.NewEnum(yourpackage.StatusEnum_SUCCESS, "Success")
println(status.Code()) // Outputs: enum numeric code
println(status.Name()) // Outputs: SUCCESS
println(status.Desc()) // Outputs: Success
```

### Enums

Manage multiple enums:

```go
enums := protoenum.NewEnums(
    protoenum.NewEnum(yourpackage.StatusEnum_SUCCESS, "Success"),
    protoenum.NewEnum(yourpackage.StatusEnum_FAILURE, "Failure"),
)

// Lookup examples
println(enums.GetByCode(1).Desc())  // Outputs: Success
println(enums.GetByName("FAILURE").Desc()) // Outputs: Failure
```

## Key Features

- **Enum**: Wraps a Protobuf enum with a description.
    - `Code()`: Gets the numeric code.
    - `Name()`: Gets the enum name.
    - `Desc()`: Gets the description.
- **Enums**: A enums-map for lookup by code, name, or description.

## License

MIT License. See [LICENSE](LICENSE) for details.

## Support

Welcome to contribute to this project by submitting pull requests or reporting issues.

If you find this package helpful, give it a star on GitHub!

**Thank you for your support!**

**Happy Coding with `protoenum`!** 🎉

Give me stars. Thank you!!!

---

## Starring

[![starring](https://starchart.cc/go-xlan/protoenum.svg?variant=adaptive)](https://starchart.cc/go-xlan/protoenum)
