package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jmhobbs/whats-checked-out/biblionix"
	"github.com/peterbourgon/ff/v3"
)

func main() {
	fs := flag.NewFlagSet("whats-checked-out", flag.ExitOnError)
	accountId := fs.String("account", "", "The account ID (required)")
	password := fs.String("password", "", "The account password (required)")
	subdomain := fs.String("subdomain", "", "The library subdomain, i.e. <subdomain>.biblionix.com (required)")
	linked := fs.Bool("linked", true, "Fetch linked accounts")
	_ = fs.String("config", "", "Path to config file")

	if err := ff.Parse(fs, os.Args[1:],
		ff.WithConfigFileFlag("config"),
		ff.WithConfigFileParser(ff.PlainParser),
	); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if accountId == nil || *accountId == "" {
		fmt.Fprintf(os.Stderr, "error: account is required\n")
		os.Exit(1)
	}

	if password == nil || *password == "" {
		fmt.Fprintf(os.Stderr, "error: password is required\n")
		os.Exit(1)
	}

	if subdomain == nil || *subdomain == "" {
		fmt.Fprintf(os.Stderr, "error: subdomain is required\n")
		os.Exit(1)
	}

	client := biblionix.New(*subdomain)
	session, err := client.Login(*accountId, *password)
	if err != nil {
		log.Fatalf("error logging in: %v", err)
		os.Exit(1)
	}

	account, err := client.Account(session)
	if err != nil {
		log.Fatalf("error getting account: %v", err)
		os.Exit(1)
	}

	accounts := make(map[string]*biblionix.AccountResponse)
	accounts[*accountId] = account

	accountsToCheck := []string{}
	if *linked {
		for _, linked := range account.LinkedCard {
			accountsToCheck = append(accountsToCheck, linked.CardNumber)
		}
	}

	for _, id := range accountsToCheck {
		linkedSession, err := client.Login(id, *password)
		if err != nil {
			log.Printf("error logging into linked account %q: %v", id, err)
			continue
		}
		linkedAccount, err := client.Account(linkedSession)
		if err != nil {
			log.Fatalf("error getting linked account %q: %v", id, err)
			continue
		}
		accounts[id] = linkedAccount
	}

	for id, account := range accounts {
		fmt.Printf("\n=== %s [%s] ==============\n\n", account.Name.Printable, id)

		if len(account.Item) == 0 {
			fmt.Println("Nothing Checked Out")
			continue
		}

		maxTitle := 0
		maxAuthor := 0
		for _, item := range account.Item {
			maxTitle = max(len(item.Title), maxTitle)
			maxAuthor = max(len(item.Author), maxAuthor)
		}

		fmtString := fmt.Sprintf("%%-%ds | %%-%ds | %%s | %%s\n", maxTitle, maxAuthor)

		for _, item := range account.Item {
			renewable := "❓"
			switch item.Renewable {
			case "1":
				renewable = "✅"
			case "0":
				renewable = "❌"
			}
			fmt.Printf(fmtString, item.Title, item.Author, item.Due, renewable)
		}
	}

}
