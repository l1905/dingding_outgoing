package conf

import (
	"flag"
	"github.com/BurntSushi/toml"
)

// global var
var (
	ConfPath string
	Conf     *Config
)

type Config struct {

	// mysql
	MySQL *MySQL

}

// MySQL represent mysql conf
type MySQL struct {
	Host      string
}


func Init() (err error) {
	_, err = toml.DecodeFile(ConfPath, &Conf)
	return err
}


func init() {
	flag.StringVar(&ConfPath, "conf", "", "config path")
}
