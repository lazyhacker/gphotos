package main

import (
	"flag"
	"fmt"
	"log"

	"golang.org/x/oauth2"
	"lazyhacker.dev/gclientauth"
	_ "lazyhacker.dev/googlephotos"
	googlephotos "lazyhacker.dev/googlephotos/v1"
)

var (
	client_secret = flag.String("secret", "", "Client secret JSON file from Google.")
	tokencache    = flag.String("tokencache", "accesstoken.json", "Local cached access token.")
	browser       = flag.Bool("browser", false, "Try to automatically open a browser.")
	port          = flag.String("port", "8081", "localhost port")
)

func main() {

	flag.Parse()
	scopes := []string{googlephotos.PhotoslibraryReadonlyScope}

	ctx := oauth2.NoContext
	token, config, err := gclientauth.GetGoogleOauth2Token(ctx, *client_secret, *tokencache, scopes, *browser, *port)
	if err != nil {
		log.Fatalf("Fetching oAuth token failed.\n%v", err)
	}
	cfg := config.Client(ctx, token)
	defer cfg.CloseIdleConnections()
	if err != nil {
		log.Fatal(err)
	}
	gp, err := googlephotos.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	res, err := gp.Albums.List().Do()
	for _, a := range res.Albums {
		fmt.Printf("%v\n", a.Title)
	}
}
