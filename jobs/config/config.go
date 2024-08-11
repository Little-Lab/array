package config

type Mysql struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"db_name"`
}

type AppConfig struct {
	Name  string   `json:"name"`
	Host  string   `json:"host"`
	Port  int      `json:"port"`
	Tags  []string `json:"tags"`
	Mysql `json:"mysql"`
}

type NacosConfig struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Dataid    string `json:"dataid"`
	Group     string `json:"group"`
	Namespace string `json:"namespace"`
}
