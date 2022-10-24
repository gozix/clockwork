// Copyright 2018 Sergey Novichkov. All rights reserved.
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package clockwork

import (
	"github.com/gozix/di"
	"github.com/gozix/glue/v3"

	"github.com/jonboulle/clockwork"
)

// Bundle implements glue.Bundle interface.
type Bundle struct{}

// BundleName is bundle name.
const BundleName = "clockwork"

// Bundle implements glue.Bundle interface.
var _ glue.Bundle = (*Bundle)(nil)

// NewBundle create bundle instance.
func NewBundle() *Bundle {
	return new(Bundle)
}

func (b *Bundle) Name() string {
	return BundleName
}

func (b *Bundle) Build(builder di.Builder) error {
	return builder.Provide(clockwork.NewRealClock)
}
