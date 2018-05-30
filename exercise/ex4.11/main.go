// The environment variables GITHUB_USER and GITHUB_PASS are used for authentication.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func search(query []string) {
	result, err := SearchIssues(query)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range result.Items {
		format := "#%-5d %12.11s %.99s\n"
		fmt.Printf(format, item.Number, item.User.Login, item.Title)
	}
}

func read(owner, repo, number string) {
	issue, err := GetIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}

	body := issue.Body
	if body == "" {
		body = "<empty>\n"
	}

	fmt.Printf("repo: %s/%s\nnumber: %s\nuser: %s\ntitle: %s\n\n%s",
		owner, repo, number, issue.User.Login, issue.Title, body)
}

func edit(owner, repo, number string) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}

	editorPath, err := exec.LookPath(editor)
	if err != nil {
		log.Fatal(err)
	}

	tempfile, err := ioutil.TempFile("", "issue_crud")
	if err != nil {
		log.Fatal(err)
	}

	defer tempfile.Close()
	defer os.Remove(tempfile.Name())

	issue, err := GetIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}

	encoder := json.NewEncoder(tempfile)
	err = encoder.Encode(map[string]string{
		"title": issue.Title,
		"state": issue.State,
		"body":  issue.Body,
	})

	if err != nil {
		log.Fatal(err)
	}

	cmd := &exec.Cmd{
		Path:   editorPath,
		Args:   []string{editor, tempfile.Name()},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	tempfile.Seek(0, 0)
	fields := make(map[string]string)
	if err = json.NewDecoder(tempfile).Decode(&fields); err != nil {
		log.Fatal(err)
	}

	_, err = EditIssue(owner, repo, number, fields)
	if err != nil {
		log.Fatal(err)
	}
}

func close_(owner, repo, number string) {
	_, err := EditIssue(owner, repo, number, map[string]string{"state": "closed"})
	if err != nil {
		log.Fatal(err)
	}
}

func open_(owner, repo, number string) {
	_, err := EditIssue(owner, repo, number, map[string]string{"state": "open"})
	if err != nil {
		log.Fatal(err)
	}
}

var usage string = `usage:
search QUERY
[read|edit|close|open] OWNER REPO ISSUE_NUMBER
`

func usageDie() {
	fmt.Fprintln(os.Stderr, usage)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		usageDie()
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	if cmd == "search" {
		if len(args) < 1 {
			usageDie()
		}
		search(args)
		os.Exit(0)
	}

	if len(args) != 3 {
		usageDie()
	}

	owner, repo, number := args[0], args[1], args[2]
	switch cmd {
	case "read":
		read(owner, repo, number)
	case "edit":
		edit(owner, repo, number)
	case "close":
		close_(owner, repo, number)
	case "open":
		open_(owner, repo, number)
	}
}

//+!
/*
$ ./4.11 search read golang go 5901

#5901    aholbreich The newest 2.0.0-alpha.53 Quick start tutorials requires Python to be installed first
#5901      asmeurer solve error with expression with derivative
#5901      smandell Font Smoothing flickers when auto-savings on OS X and Safari
#5901         Zero3 Wrong amount of time remaining during navigation
#5901   vorpal-buil Exceeded slow_query limit (156.2 > 5) in mysql:
#5901       fzwaard OrientDB 2.1.14-SNAPSHOT on Ubuntu: Failed Server KeepAlive and  NoKeepAliveTests
#5901   dwang201510 spacemacs/report-issue fails if I set any layer variables with periods in them
#5901        kimchy Improved bloom filter hashing
#5901       vkuznet Fix bug in AJAX done function
#5901    hyperlogic Threshold based walking/leaning while in HMD.
#5901      electrum Hide all non-annotation Jackson classes from plugins
#5901       dhensby Cleaning up GridFieldAddExistingAutocompleter
#5901      spadgett Use assets/config.local.js if present for development config
#5901    coderanger Remove the easy_install resource
#5901      aphearin Add note on pre-releasing in affiliated package release docs
#5901          sbko [Python] Increase range of valid status codes
#5901        angrox PUP-6164: Tech-debt: links.puppetlabs.com -> links.puppet.com
#5901       bryevdv pass HTTP request args more explicitly to session
#5901         pitaj Use Benchpress
#5901     HebaruSan Update folder path
#5901      lpinsivy Update lifecycle in documentation
#5901   zackJKnight Added a missing step to get started on Mac.
#5901      suez1224 [FLINK-9235][Security] Add integration tests for YARN kerberos integration.
#5901   nomanzafarm create process error 2
#5901   MarcinSzysz Query: Wrong data in included navigation when using Skip() method
#5901      Macludde USER ISSUE: Drag & Drop Item Stuck
#5901       hyarsan [Feature Request] Add option for opacity to the RetroArch window
#5901      joydance [As3]flash.utils.Object will be converted "*" type when compile to as3
#5901   koorellasur replace a closed issue
#5901         RoyGF CHT-Super Issue-Module Community

*/
//-!
