package configuration

type ElasticsearchConfig struct {
	EsAddress string
	EsCert    string
	EsKey     string
	UseTLS    bool
	// Timeout   time.Duration // JZ Timeout here to config
}
