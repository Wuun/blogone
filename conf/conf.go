package conf

import (
	"os"
	"time"
)

//Conf is config will be used.
var G_CONF *Conf

type Conf struct {
	MaxConnected     int
	ConnectTimeout   time.Duration
	MSHost           string
	MSPWD            string
	MSPort           string
	MSDBName         string
	MSUser           string
	MSListArtQuery   string
	MSGetArtConQuery string
}

func init() {
	G_CONF = newConf()
}

func newConf() *Conf {
	host := os.Getenv("MYSQL_HOST")
	pwd := os.Getenv("MYSQL_PWD")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DB_NAME")
	user := os.Getenv("MYSQL_USER")
	liArQuery := os.Getenv("LIST_ARTICLE_QUERY")
	getArConQuery := os.Getenv("GET_ART_CON_QUERY")
	if host == "" {
		host = "127.0.0.1"
	}
	if port == "" {
		port = "3306"
	}

	return &Conf{
		MSHost:           host,
		MSPWD:            pwd,
		MSPort:           port,
		MSDBName:         dbName,
		MSUser:           user,
		MSListArtQuery:   liArQuery,
		MSGetArtConQuery: getArConQuery,
	}
}
