package steps

import (
	"github.com/epik-protocol/epik-gateway-backend/query/linkedql"
)

func init() {
	linkedql.Register(&In{})
}

var _ linkedql.PathStep = (*In)(nil)

// In is an alias for ViewReverse.
type In struct {
	VisitReverse
}

// Description implements Step.
func (s *In) Description() string {
	return "aliases for ViewReverse"
}
