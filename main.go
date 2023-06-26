package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	selfupdate "github.com/creativeprojects/go-selfupdate"
)

const version = "0.1.0"
const repoName = "test/c2"
const delay = 15 * time.Second

func main() {
	for {
		err := update(version)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(version)
		time.Sleep(delay)
	}
}

func update(version string) error {
	source, err := selfupdate.NewGiteaSource(selfupdate.GiteaConfig{
		BaseURL: "http://localhost:3000",
		APIToken: "test", // manually created "Applications" token /user/settings/applications
	})
	if err != nil {
		return err
	}

	giteaUpdater, err := selfupdate.NewUpdater(selfupdate.Config{
		Source: source,
	})
	if err != nil {
		return err
	}

	latest, found, err := giteaUpdater.DetectLatest(context.Background(), selfupdate.ParseSlug(repoName))
	if err != nil {
		return fmt.Errorf("error occurred while detecting version: %w", err)
	}
	if !found {
		return fmt.Errorf("latest version for %s/%s could not be found from github repository", runtime.GOOS, runtime.GOARCH)
	}

	if latest.LessOrEqual(version) {
		log.Printf("Current version (%s) is the latest", version)
		return nil
	}

	exe, err := os.Executable()
	if err != nil {
		return errors.New("could not locate executable path")
	}
	if err := giteaUpdater.UpdateTo(context.Background(), latest, exe); err != nil {
		return fmt.Errorf("error occurred while updating binary: %w", err)
	}
	log.Printf("Successfully updated to version %s", latest.Version())
	return nil
}