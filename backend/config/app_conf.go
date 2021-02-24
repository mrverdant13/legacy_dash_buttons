package config

type appConf struct {
	GraphQLServerConf GraphQLServerConf `mapstructure:"gql_server_conf"`
	DbConf            DbConf            `mapstructure:"db_conf"`
}