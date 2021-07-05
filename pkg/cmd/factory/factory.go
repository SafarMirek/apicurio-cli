package factory

import (
	"github.com/apicurio/apicurio-cli/internal/config"
	"github.com/apicurio/apicurio-cli/pkg/iostreams"
	"github.com/apicurio/apicurio-cli/pkg/localize"
	"github.com/apicurio/apicurio-cli/pkg/logging"
)

// Factory is an abstract type which provides access to
// the root configuration and connections for the CLI
type Factory struct {
	// Type which defines the streams for the CLI
	IOStreams *iostreams.IOStreams
	// Interface to read/write to the config
	Config config.IConfig

	// Returns a logger to create leveled logs in the application
	Logger func() (logging.Logger, error)
	// Localizer provides text to the commands
	Localizer localize.Localizer
}
