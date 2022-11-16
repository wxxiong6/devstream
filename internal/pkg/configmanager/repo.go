package configmanager

type (
	Repo struct {
		// all fields in RepoCommon will be flattened
		*RepoCommon `yaml:",inline" mapstructure:",squash"`
		ApiURL      string `yaml:"apiURL" mapstructure:"apiURL"`
	}

	RepoTemplate struct {
		// all fields in RepoCommon will be flattened
		*RepoCommon `yaml:",inline" mapstructure:",squash"`
	}

	RepoCommon struct {
		ScmType string `yaml:"scmType" mapstructure:"scmType"`
		Owner   string `yaml:"owner" mapstructure:"owner"`
		Org     string `yaml:"org" mapstructure:"org"`
		Name    string `yaml:"name" mapstructure:"name"`
		URL     string `yaml:"url" mapstructure:"url"`
	}
)

func (repo *RepoCommon) FillAndValidate() error {
	if err := repo.fill(); err != nil {
		return err
	}
	if err := repo.validate(); err != nil {
		return err
	}
	return nil
}

func (repo *RepoCommon) fill() error {
	scmType, owner, repoName, err := getRepoCommonFromUrl(repo.URL)
	if err != nil {
		return err
	}
	if repo.ScmType == "" {
		repo.ScmType = scmType
	}
	if repo.Owner == "" {
		repo.Owner = owner
	}
	if repo.Name == "" {
		repo.Name = repoName
	}
	return nil
}

func (repo *RepoCommon) GetOwner() string {
	if repo.Org != "" {
		return repo.Org
	}
	return repo.Owner
}

func getRepoCommonFromUrl(url string) (scmType, owner, repoName string, err error) {
	// TODO(aFlyBird0): complete this function
	return
}

func (repo *RepoCommon) validate() error {
	// TODO(aFlyBird0): complete this function
	return nil
}
