package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := os.Args[1:]

	filename := fmt.Sprintf("%s.md", args[0])

	cmd := exec.Command("hugo", "new", fmt.Sprintf("items/%s", filename))
	err := cmd.Run()
	if err != nil {
		fmt.Println("error ", err.Error())
	}

	var title string
	var itemURL string
	var sites string
	var tags []string

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Title: ")
		title, err = reader.ReadString('\n')
		if err != nil {
			log.Printf("cannot read title: error as %v\n", err)
			continue
		}

		fmt.Print("URL: ")
		itemURL, err = reader.ReadString('\n')
		if err != nil {
			log.Printf("cannot read itemURL: error as %v\n", err)
			continue
		}

		fmt.Print("Site: ")
		sites, err = reader.ReadString('\n')
		if err != nil {
			log.Printf("cannot read sites: error as %v\n", err)
			continue
		}

		fmt.Print("Tags: input empty to quit ")
		for {
			tag, err := reader.ReadString('\n')
			if err != nil {
				log.Printf("cannot read tag: error as %v\n", err)
				continue
			}
			t := strings.TrimSpace(tag)
			if t == "" {
				break
			}
			tags = append(tags, t)
			fmt.Print("New Tag ")
		}
		break
	}

	tagsStr := ""

	for _, tag := range tags {
		tagsStr += tag + ","
	}

	tagsStr = tagsStr[:len(tagsStr)-1]

	content := fmt.Sprintf(`
---
title: "%s"
date: 2021-04-05T15:55:03+03:00
itemurl: "%s"
sites: "%s"
tags: [%s] 
draft: false
---
`, strings.TrimSpace(title), strings.TrimSpace(itemURL), strings.TrimSpace(sites), tagsStr)

	f, err := os.Create("./content/items/" + filename)
	if err != nil {
		log.Fatal("cannot create a file: %v", err)
	}
	defer f.Close()

	f.WriteString(content)

	fmt.Println("Content is created!")

	fmt.Println("Run Hugo to generate new page (y/n) ")
	comp, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("cannot read input for hugo: %v", err)
	}

	c := strings.ToLower(strings.TrimSpace(comp))

	if c == "y" {
		cmd := exec.Command("hugo", "-t", "hugonews")
		err := cmd.Run()
		if err != nil {
			fmt.Println("cannot run 'hugo -t hugonews': %v", err.Error())
		}
	}

}
