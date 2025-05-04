package cli

type CliCommand interface {
	Name() string
	Description() string
	Execute() error
}
