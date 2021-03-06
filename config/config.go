package config

func DefaultConfig() *Config {
	return &Config{
		Servers: []*Server{
			&Server{
				BindAddr:  "0.0.0.0:25",
				Hostname:  "mailhog.example",
				PolicySet: DefaultSMTPPolicySet(),
			},
			&Server{
				BindAddr:  "0.0.0.0:587",
				Hostname:  "mailhog.example",
				PolicySet: DefaultSubmissionPolicySet(),
			},
		},
	}
}

type Config struct {
	Servers []*Server
}

type Server struct {
	BindAddr  string
	Hostname  string
	PolicySet PolicySet
}

type PolicySet struct {
	RequireAuthentication bool
	RequireLocalDelivery  bool
	MaximumRecipients     int
	EnableTLS             bool
	RequireTLS            bool
	MaximumLineLength     int
	MaximumConnections    int
}

func DefaultSubmissionPolicySet() PolicySet {
	return PolicySet{
		RequireAuthentication: true,
		RequireLocalDelivery:  false,
		MaximumRecipients:     500,
		RequireTLS:            true,
		EnableTLS:             true,
		MaximumLineLength:     1024000,
		MaximumConnections:    1000,
	}
}

func DefaultSMTPPolicySet() PolicySet {
	return PolicySet{
		RequireAuthentication: false,
		RequireLocalDelivery:  true,
		MaximumRecipients:     500,
		RequireTLS:            false,
		EnableTLS:             true,
		MaximumLineLength:     1024000,
		MaximumConnections:    1000,
	}
}

var cfg = DefaultConfig()

func Configure() *Config {
	return cfg
}

func RegisterFlags() {
	// TODO
}
