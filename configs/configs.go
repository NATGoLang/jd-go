package configs

import (
	"bytes"
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		Credentials struct {
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		} `yaml:"credentials"`
		Net       string `yaml:"net"`
		Addr      string `yaml:"addr"`
		DBName    string `yaml:"db_name"`
		Charset   string `yaml:"charset"`
		Loc       string `yaml:"loc"`
		ParseTime bool   `yaml:"parse_time"`
	} `yaml:"database"`
}

func InitConfig() (Config, error) {
	f, err := os.Open("configs/private_config.yaml")
	if err != nil {
		// return empty Config here is weird
		return Config{}, err
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

// referencing the Config.FormatDSN() method in [github.com/go-sql-driver/mysql]
// just added the params we are using, may need to add some more later
// more details: [https://github.com/go-sql-driver/mysql#dsn-data-source-name]
// "root:26711153@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
func (cfg *Config) GetDSNString() (string, error) {
	dbCfg := cfg.Database

	var buf bytes.Buffer
	
	// [username[:password]@]
	if dbCfg.Credentials.Username == "" {
		return "", errors.New("username not found in config file")
	}
	buf.WriteString(dbCfg.Credentials.Username)

	if dbCfg.Credentials.Password == "" {
		return "", errors.New("password not found in config file")
	}
	buf.WriteByte(':')
	buf.WriteString(dbCfg.Credentials.Password)
	buf.WriteByte('@')
	
	// [protocol[(address)]]
	if dbCfg.Net == "" {
		return "", errors.New("net settings not found in config file")
	}
	buf.WriteString(dbCfg.Net)
	if dbCfg.Addr != "" {
		buf.WriteByte('(')
		buf.WriteString(dbCfg.Addr)
		buf.WriteByte(')')
	}

	if dbCfg.DBName == "" {
		return "", errors.New("database name not found in config file")
	}

	// db name
	buf.WriteByte('/')
	buf.WriteString(dbCfg.DBName)

	hasParam := false
	if dbCfg.Charset != "" {
		writeDSNParam(&buf, &hasParam, "charset", dbCfg.Charset)
	}

	if dbCfg.Loc != "" {
		writeDSNParam(&buf, &hasParam, "loc", dbCfg.Loc)
	}

	if dbCfg.ParseTime == true {
		writeDSNParam(&buf, &hasParam, "parseTime", "true")
	}

	return buf.String(), nil
}

func writeDSNParam(buf *bytes.Buffer, hasParam *bool, key string, value string) {
	if !*hasParam {
		*hasParam = true
		buf.WriteByte('?')
	} else {
		buf.WriteByte('&')
	}
	buf.WriteString(key)
	buf.WriteByte('=')
	buf.WriteString(value)
}
