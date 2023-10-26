# Terraform Google Provider Schema Repository

This repository contains the generated Go files for the ModTm provider schemas, which are based on the Terraform ModTm Provider. These schema files can be used as a reference when writing tools, such as TFLint plugins, that interact with the ModTm provider.

## Repository Structure

Each tag version of the Terraform Google Provider has a corresponding tag in this repository. You can find the schema files for each provider version under the respective tag.

e.g.: to use `modtm`'s `0.1.8` schema, you could:

```shell
$ go get github.com/lonegunmanb/terraform-modtm-schema/@v0.1.8
```

Then you can read schemas like this:

```go
import (
"testing"

"github.com/lonegunmanb/terraform-modtm-schema/generated"
"github.com/stretchr/testify/assert"
)

func TestResourceSchema(t *testing.T) {
assert.NotEmpty(t, generated.Resources)
assert.NotEmpty(t, generated.DataSources)
}
```

## Generating Schema Files

The schema files are generated using the terraform provider schema -json command. This command retrieves the schema information and converts it into JSON format. The JSON files are then converted into Go files, which can be found in this repository.

If you encounter any issues or would like to contribute to this repository, please submit an issue or a pull request on GitHub.

## License

[MIT](LICENSE)
