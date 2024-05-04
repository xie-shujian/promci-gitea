package main

type Config struct {
	GitServer GitServer `yaml:"git_server"`
	PromciApp PromciApp `yaml:"promci_app"`
}
type Repository struct {
	Name      string `yaml:"name"`
	Directory string `yaml:"directory"`
	ReloadUrl string `yaml:"reload_url"`
}
type GitServer struct {
	GroupURL     string       `yaml:"group_url"`
	User         string       `yaml:"user"`
	AccessToken  string       `yaml:"access_token"`
	Repositories []Repository `yaml:"repositories"`
}

type PromciApp struct {
	AccessToken string `yaml:"access_token"`
	LogFile     string `yaml:"log_file"`
}

func (c *Config) FindAndSetRepo(event_repo_url string) {
	for _, repository := range c.GitServer.Repositories {
		conf_repo_url := c.GitServer.GroupURL + "/" + repository.Name
		if conf_repo_url == event_repo_url {
			Repo = &repository
			return
		}
	}
}

type GitEvent struct {
	Repository RepositoryEvent
}
type RepositoryEvent struct {
	HtmlUrl string `json:"html_url"`
}
