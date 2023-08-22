package conf

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var Config *Conf

type Conf struct {
	System        System        `yaml:"system"`
	Database      Database      `yaml:"database"`
	Redis         Redis         `yaml:"redis"`
	Cache         Cache         `yaml:"cache"`
	Email         Email         `yaml:"email"`
	EncryptSecret EncryptSecret `yaml:"encryptSecret"`
	PhotoPath     PhotoPath     `yaml:"photoPath"`
}

type System struct {
	HttpPort    string  `yaml:"HttpPort"`
	Host        string  `yaml:"Host"`
	UploadModel string  `yaml:"UploadModel"`
	Domain      string  `yaml:"domain"`
	Version     float64 `yaml:"version"`
	Env         string  `yaml:"env"`
}

type Database struct {
	Dialect  string `yaml:"dialect"`
	DbHost   string `yaml:"dbHost"`
	DbPort   string `yaml:"dbPort"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbName"`
}

type Redis struct {
	RedisDbName   int    `yaml:"redisDbName"`
	RedisHost     string `yaml:"redisHost"`
	RedisPort     int    `yaml:"redisPort"`
	RedisPassword int    `yaml:"redisPassword"`
	RedisNetwork  string `yaml:"redisNetwork"`
}

type Cache struct {
	CacheEmpires int         `yaml:"cacheEmpires"`
	CacheWarmUp  interface{} `yaml:"cacheWarmUp"`
	CacheServer  interface{} `yaml:"cacheServer"`
	CacheType    string      `yaml:"cacheType"`
}

type Email struct {
	SmtpHost  interface{} `yaml:"smtpHost"`
	SmtpEmail interface{} `yaml:"smtpEmail"`
	SmtpPass  interface{} `yaml:"smtpPass"`
	Address   string      `yaml:"address"`
}

type EncryptSecret struct {
	JwtSecret   string `yaml:"jwtSecret"`
	EmailSecret string `yaml:"emailSecret"`
	PhoneSecret string `yaml:"phoneSecret"`
}

type PhotoPath struct {
	PhotoHost   string `yaml:"photoHost"`
	ProductPath string `yaml:"ProductPath"`
	AvatarPath  string `yaml:"AvatarPath"`
}

func loadYamlConfig(confPath string) {
	conf_to_load := Conf{}
	absPath, err := filepath.Abs(confPath)
	dataBytes, err := os.ReadFile(absPath)
	if err != nil {
		panic("读取配置文件失败" + err.Error())
	}

	err = yaml.Unmarshal(dataBytes, &conf_to_load)
	if err != nil {
		panic("解析配置文件失败" + err.Error())
	}

	Config = &conf_to_load
}

func init() {
	loadYamlConfig("conf/config.yaml")
}
