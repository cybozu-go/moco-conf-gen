package cmd

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/cybozu-go/moco"
)

func subMain() error {
	serverID, err := confServerID(os.Getenv(moco.PodNameEnvName))
	if err != nil {
		return err
	}
	parameters := moco.MyConfTemplateParameters{ServerID: serverID, AdminAddress: os.Getenv(moco.PodIPEnvName)}

	tmpl, err := template.ParseFiles(filepath.Join(moco.MySQLConfTemplatePath, moco.MySQLConfName))
	if err != nil {
		return err
	}

	file, err := os.Create(filepath.Join(moco.MySQLConfPath, moco.MySQLConfName))
	if err != nil {
		return nil
	}
	defer file.Close()

	err = tmpl.Execute(file, parameters)
	if err != nil {
		return err
	}

	return nil
}

func confServerID(podNameWithOrdinal string) (uint, error) {
	// ordinal should be increased by 1000 because the case server-id is 0 is not suitable for the replication purpose
	const ordinalOffset = 1000

	s := strings.Split(podNameWithOrdinal, "-")
	if len(s) < 2 {
		return 0, errors.New("podName should contain an ordinal with dash, like 'podname-0', at the end: " + podNameWithOrdinal)
	}

	ordinal, err := strconv.ParseUint(s[len(s)-1], 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(ordinal + ordinalOffset), nil
}
