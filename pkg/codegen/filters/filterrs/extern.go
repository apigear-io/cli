package filterrs

import (
	"github.com/apigear-io/cli/pkg/objmodel"
)

type RsExtern struct {
	Name    string
	Crate   string
	Version string
}

func rsExtern(xe *objmodel.Extern) RsExtern {
	name := xe.Meta.GetString("rs.type")
	crate := xe.Meta.GetString("rs.crate")
	version := xe.Meta.GetString("rs.version")
	if name == "" {
		name = xe.Name
	}
	return RsExtern{
		Name:    name,
		Crate:   crate,
		Version: version,
	}
}
