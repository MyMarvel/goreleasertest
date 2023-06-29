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

const version = "0.1.6"
const repoName = "test/c2"
const delay = 60 * time.Second

func main() {
	for {
		fmt.Println("Before update")
		err := update(version)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(version)
		time.Sleep(delay)
		fmt.Println("After sleep")
	}
}

func update2(version string) error {
	selfupdate.SetLogger(log.New(os.Stdout, "", 0))

	source, err := selfupdate.NewGiteaSource(selfupdate.GiteaConfig{
		BaseURL: "http://host.docker.internal:3000",
		APIToken: "45ea01a6c677552ea94d557a35a2fd4afd32d218", // manually created "Applications" token /user/settings/applications
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

func update3(version string) error {
	latest, found, err := selfupdate.DetectLatest(context.Background(), selfupdate.ParseSlug("MyMarvel/goreleasertest"))
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
	if err := selfupdate.UpdateTo(context.Background(), latest.AssetURL, latest.AssetName, exe); err != nil {
		return fmt.Errorf("error occurred while updating binary: %w", err)
	}
	log.Printf("Successfully updated to version %s", latest.Version())
	return nil
}

func update(version string) error {
	//err := os.Remove(".goreleasertest.exe.old")
	//if err != nil {
//		return err
//	}

	selfupdate.SetLogger(log.New(os.Stdout, "", 0))
	latest, found, err := selfupdate.DetectLatest(context.Background(), selfupdate.ParseSlug("MyMarvel/goreleasertest"))
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

	re, err := selfupdate.UpdateSelf(context.Background(), version, selfupdate.ParseSlug("MyMarvel/goreleasertest"))
	if err != nil {
		return err
	}
	log.Printf("Successfully updated to version %s", re.Version())
	//os.Exit(0)

	return nil
}
