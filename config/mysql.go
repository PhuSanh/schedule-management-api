package config

type Mysql struct {
	Username 		string
	Password 		string
	DatabaseName 	string 	`mapstructure:"database_name"`
	MaxIdleConn 	int		`mapstructure:"max_idle_conn"`
	MaxOpenConn 	int		`mapstructure:"max_open_conn"`
}
