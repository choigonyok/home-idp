package cli

import (
	"github.com/spf13/pflag"
)

type Flag interface {
	GetFlag() *pflag.FlagSet
	CreateFlag()
}
