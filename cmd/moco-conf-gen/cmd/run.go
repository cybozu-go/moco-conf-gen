package cmd

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

const (
	// PodNameEnvName is a name of the environment variable of a pod name.
	PodNameEnvName = "POD_NAME"

	// MySQLConfTemplatePath is
	MySQLConfTemplatePath = "/etc/mysql_template"

	// MySQLConfName is a filename for MySQL conf.
	MySQLConfName = "my.cnf"

	// MySQLConfPath is a path for MySQL conf dir.
	MySQLConfPath = "/etc/mysql"
)

// MyConfTemplateParameters define parameters for a MySQL configuration template
type MyConfTemplateParameters struct {
	// ServerID is the value for server_id of MySQL configuration
	ServerID uint32
	// AdminAddress is the value for admin_address of MySQL configuration
	AdminAddress string
}

func subMain() error {
	serverID, err := confServerID(os.Getenv(PodNameEnvName))
	if err != nil {
		return err
	}
	parameters := MyConfTemplateParameters{ServerID: serverID, AdminAddress: os.Getenv(PodNameEnvName)}

	tmpl, err := template.ParseFiles(filepath.Join(MySQLConfTemplatePath, MySQLConfName))
	if err != nil {
		return err
	}

	file, err := os.Create(filepath.Join(MySQLConfPath, MySQLConfName))
	if err != nil {
		return err
	}
	defer file.Close()

	err = tmpl.Execute(file, parameters)
	if err != nil {
		return err
	}

	return nil
}

func confServerID(podNameWithOrdinal string) (uint32, error) {
	s := strings.Split(podNameWithOrdinal, "-")
	if len(s) < 2 {
		return 0, errors.New("podName should contain an ordinal with dash, like 'podname-0', at the end: " + podNameWithOrdinal)
	}

	ordinal, err := strconv.ParseUint(s[len(s)-1], 10, 32)
	if err != nil {
		return 0, err
	}

	return uint32(ordinal) + serverIDBase, nil
}
