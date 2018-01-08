package main

import (
	"log"
	"net/http"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

// getClient uses a Context and Config to retrieve a Token
// then generate a Client. It returns the generated Client.

func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
	cacheFile, err := tokenCacheFile()
	if err != nil {
		log.Fatalln("Unable to get path to cached credential file. %v", err)
	}
	tok, err := tokenFromFile(cacheFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(cacheFile, tok)
	}
	return config.Client(ctx, tok)
}

// getTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token.
