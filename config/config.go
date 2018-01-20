package config

type Config struct {
	GitHubTokens []string `mapstructure:"github_tokens"`
}
