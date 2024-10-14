package assist

import "context"

type Runner interface {
	Run(ctx context.Context) error
	Stop() error
}
