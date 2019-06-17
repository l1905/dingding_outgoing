package conf

import (
	"flag"
	"fmt"
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
	// redis

	// seq conf
	Seq *Seq
}

type Seq struct {
	BusinessID int64
	Token      string
}

// Host host.
type Host struct {
	API    string
	Search string
}

// MySQL represent mysql conf
type MySQL struct {
	Host      string
}


func Init() (err error) {
	fmt.Println(ConfPath)
	_, err = toml.DecodeFile(ConfPath, &Conf)
	return
}


func init() {
	flag.StringVar(&ConfPath, "conf", "", "config path")
}
