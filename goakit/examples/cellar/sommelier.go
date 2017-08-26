package cellar

import (
	"context"

	"github.com/go-kit/kit/log"
	"goa.design/plugins/goakit/examples/cellar/gen/sommelier"
)

// sommelier service example implementation.
// The example methods log the requests and return zero values.
type sommeliersvc struct {
	logger log.Logger
}

// NewSommelier returns the sommelier service implementation.
func NewSommelier(logger log.Logger) sommelier.Service {
	return &sommeliersvc{logger}
}

// Pick implements pick.
func (s *sommeliersvc) Pick(ctx context.Context, p *sommelier.Criteria) (sommelier.StoredBottleCollection, error) {
	var res sommelier.StoredBottleCollection
	s.logger.Log("msg", "sommelier.pick")
	return res, nil
}
