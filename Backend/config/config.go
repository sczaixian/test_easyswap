package config

import (
	"strings"

	"github.com/ProjectsTask/Base/evm/erc"
	logging "github.com/ProjectsTask/Base/logger"
	"github.com/ProjectsTask/Base/stores/gdb"

	"github.com/spf13/viper"
)

/*
mapstructure 用于配置解析：常用于 Viper 等配置库，将配置文件（如 YAML、JSON）中的数据映射到结构体
如果使用 Viper 读取配置，需要 mapstructure 标签，因为 Viper 内部使用 mapstructure 进行解码
ProjectCfg 可能需要从环境变量、命令行参数等其他来源读取，这些来源通常使用 mapstructure 标签
使用 Viper 读取配置文件
自定义 map 到结构体的转换
需要复杂解码钩子或元数据
*/
type Config struct {
	Api            `toml:"api" json:"api"`
	ProjectCfg     *ProjectCfg       `toml:"project_cfg" mapstructure:"project_cfg" json:"project_cfg"`
	Log            logging.LogConf   `toml:"log" json:"log"`
	DB             gdb.Config        `toml:"db" json:"db"`
	Kv             *KvConf           `toml:"kv" json:"kv"`
	Evm            *erc.NftErc       `toml:"evm" json:"evm"`
	MetadataParse  *MetadataParse    `toml:"metadata_parse" mapstructure:"metadata_parse" json:"metadata_parse"`
	ChainSupported []*ChainSupported `toml:"chain_supported" mapstructure:"chain_supported" json:"chain_supported"`
}

type Api struct {
	Port   string `toml:"port" json:"port"`
	MaxNum int64  `toml:"max_num" json:"max_num"`
}

type ProjectCfg struct {
	Name string `toml:"name" json:"name"`
}

//type KvConf struct {
//	Port   string `toml:"port" json:"port"`
//	MaxNum int64  `toml:"max_num" json:"max_num"`
//}

type KvConf struct {
	Redis []*Redis `toml:"redis" mapstructure:"redis" json:"redis"`
}

type Redis struct {
	MasterName string `toml:"master_name" mapstructure:"master_name" json:"master_name"`
	Host       string `toml:"host" json:"host"`
	Type       string `toml:"type" json:"type"`
	Pass       string `toml:"pass" json:"pass"`
}

type MetadataParse struct {
	NameTags       []string `toml:"name_tags" mapstructure:"name_tags" json:"name_tags"`
	ImageTags      []string `toml:"image_tags" mapstructure:"image_tags" json:"image_tags"`
	AttributesTags []string `toml:"attributes_tags" mapstructure:"attributes_tags" json:"attributes_tags"`
	TraitNameTags  []string `toml:"trait_name_tags" json:"trait_name_tags"`
	TraitValueTags []string `toml:"trait_value_tags" json:"trait_value_tags"`
}

type ChainSupported struct {
	Name     string `toml:"name" mapstructure:"name" json:"name"`
	ChainID  int    `toml:"chain_id" mapstructure:"chain_id" json:"chain_id"`
	Endpoint string `toml:"endpoint" mapstructure:"endpoint" json:"endpoint"`
}

func UnmarshalConfig(configFilePath string) (*Config, error) {
	viper.SetConfigFile(configFilePath)
	viper.SetConfigType("toml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("CNFT")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	config, err := DefaultConfig()
	if err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}
	return config, nil
}

func DefaultConfig() (*Config, error) {
	return &Config{}, nil
}
