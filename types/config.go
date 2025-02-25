package types

type Config struct {
	Database struct {
		Username string `yaml:"user", envconfig:"DB_USERNAME"`
		Password string `yaml:"password", envconfig:"DB_PASSWORD"`
		Name     string `yaml:"name", envconfig:"DB_NAME"`
		Host     string `yaml:"host", envconfig:"DB_HOST"`
		Port     string `yaml:"port", envconfig:"DB_PORT"`
	} `yaml:"database"`
	Indexer struct {
		Enabled                     bool `yaml:"enabled", envconfig:"INDEXER_ENABLED"`
		FullIndexOnStartup          bool `yaml:"fullIndexOnStartup", envconfig:"INDEXER_FULL_INDEX_ON_STARTUP"`
		IndexMissingEpochsOnStartup bool `yaml:"indexMissingEpochsOnStartup", envconfig:"INDEXER_MISSING_INDEX_ON_STARTUP"`
		CheckAllBlocksOnStartup     bool `yaml:"checkAllBlocksOnStartup", envconfig:"INDEXER_CHECK_ALL_BLOCKS_ON_STARTUP"`
		Node                        struct {
			Port string `yaml:"port", envconfig:"INDEXER_NODE_PORT"`
			Host string `yaml:"host", envconfig:"INDEXER_NODE_HOST"`
		} `yaml:"node"`
	} `yaml:"indexer"`
	Frontend struct {
		Enabled  bool   `yaml:"enabled", envconfig:"FRONTEND_ENABLED"`
		Imprint  string `yaml:"imprint", envconfig:"FRONTEND_IMPRINT"`
		SiteName string `yaml:"siteName", envconfig:"FRONTEND_SITE_NAME"`
		Server   struct {
			Port string `yaml:"port", envconfig:"FRONTEND_SERVER_PORT"`
			Host string `yaml:"host", envconfig:"FRONTEND_SERVER_HOST"`
		} `yaml:"server"`
	} `yaml:"frontend"`
}
