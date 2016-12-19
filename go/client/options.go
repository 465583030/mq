package client

type Options struct {
	// Number of retry attempts
	Retries int
	// Server list
	Servers []string
	// Selector
	Selector Selector
}

type Option func(o *Options)

// WithRetries sets the number of retry attempts
func WithRetries(i int) Option {
	return func(o *Options) {
		o.Retries = i
	}
}

// WithSelector sets the server selector used by the client
func WithSelector(s Selector) Option {
	return func(o *Options) {
		o.Selector = s
	}
}

// WithServers sets the servers used by the client
func WithServers(addrs ...string) Option {
	return func(o *Options) {
		o.Servers = addrs
	}
}
