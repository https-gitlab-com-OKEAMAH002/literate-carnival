// Application exportschema exports the expected schema as a serialized schema.Description.
package main

import (
	"os"

	"go.skia.org/infra/go/sklog"
	"go.skia.org/infra/go/sql/schema/exportschema"
	"go.skia.org/infra/machine/go/machine/store/cdb"
)

func main() {
	err := exportschema.Main(os.Args, cdb.Tables{}, cdb.Schema)
	if err != nil {
		sklog.Fatal(err)
	}
}
