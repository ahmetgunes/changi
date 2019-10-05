package configs

type Configuration struct {
	MemcachedConnectionString string `json:"memcached_connection_string"`
	TestDataPath              string `json:"test_data_path"`
	TestDataKey               string `json:"test_data_key"`
	LogFile                   string `json:"log_file"`
}