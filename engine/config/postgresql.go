package config

type PostgresqlConfig struct {
	Host     string `mapstructure:"host"`
	DbName   string `mapstructure:"db_name"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Ssl      string `mapstructure:"ssl"`
}

type ConnectionSpec struct {
	DriverName     string
	DataSourceName string
}

// GetConnectionSpec ...
func (ths PostgresqlConfig) GetConnectionSpec() ConnectionSpec {
	return ConnectionSpec{
		DriverName:     "postgres",
		DataSourceName: "postgres://" + ths.Username + ":" + ths.Password + "@" + ths.Host + "/" + ths.DbName + "?sslmode=" + ths.Ssl}
}
