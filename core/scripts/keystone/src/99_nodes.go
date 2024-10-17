package src

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
)

type node struct {
	url            *url.URL
	remoteURL      *url.URL
	serviceName    string
	deploymentName string
	login          string
	password       string
}

func (n node) IsTerminal() bool {
	return false
}

func (n node) PasswordPrompt(p string) string {
	return n.password
}

func (n node) Prompt(p string) string {
	return n.login
}

func writeNodesList(path string, nodes []*node) error {
	fmt.Println("Writing nodes list to", path)
	var lines []string
	for _, n := range nodes {
		line := fmt.Sprintf(
			"%s %s %s %s %s %s",
			n.url.String(),
			n.remoteURL.String(),
			n.serviceName,
			n.deploymentName,
			n.login,
			n.password,
		)
		lines = append(lines, line)
	}

	return writeLines(lines, path)
}

func mustReadNodesList(path string) []*node {
	fmt.Println("Reading nodes list from", path)
	nodesList, err := readLines(path)
	helpers.PanicErr(err)

	var nodes []*node
	var hasBoot bool
	for _, r := range nodesList {
		rr := strings.TrimSpace(r)
		if len(rr) == 0 {
			continue
		}
		s := strings.Split(rr, " ")
		if len(s) != 6 {
			helpers.PanicErr(errors.New("wrong nodes list format: expected 6 fields per line"))
		}
		if strings.Contains(s[0], "boot") && hasBoot {
			helpers.PanicErr(errors.New("the single boot node must come first"))
		}
		if strings.Contains(s[0], "boot") {
			hasBoot = true
		}
		parsedURL, err := url.Parse(s[0])
		helpers.PanicErr(err)
		parsedRemoteURL, err := url.Parse(s[1])
		helpers.PanicErr(err)
		nodes = append(nodes, &node{
			url:            parsedURL,
			remoteURL:      parsedRemoteURL,
			serviceName:    s[2],
			deploymentName: s[3],
			login:          s[4],
			password:       s[5],
		})
	}
	return nodes
}
