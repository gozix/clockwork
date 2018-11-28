package clockwork

import (
	"github.com/jonboulle/clockwork"
	"github.com/sarulabs/di"
)

type (
	// Bundle implements the glue.Bundle interface.
	Bundle struct{}

	// Clock is type alias of clockwork.Clock
	Clock = clockwork.Clock
)

// BundleName is default definition name.
const BundleName = "clockwork"

// NewBundle create bundle instance.
func NewBundle() *Bundle {
	return new(Bundle)
}

// Key implements the glue.Bundle interface.
func (b *Bundle) Name() string {
	return BundleName
}

// Build implements the glue.Bundle interface.
func (b *Bundle) Build(builder *di.Builder) error {
	return builder.Add(di.Def{
		Name: BundleName,
		Build: func(_ di.Container) (interface{}, error) {
			return clockwork.NewRealClock(), nil
		},
	})
}
