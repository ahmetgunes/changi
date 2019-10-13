package configs

type Configuration struct {
	Port                      string `json:"port"`
	LogFile                   string `json:"log_file"`
	MemcachedConnectionString string `json:"memcached_connection_string"`
	TestDataPath              string `json:"test_data_path"`
	TestDataKey               string `json:"test_data_key"`
}
