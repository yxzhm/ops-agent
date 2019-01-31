package main

import (
	"github.com/jasonlvhit/gocron"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	Info.Printf("ops agent start")
	profilePath := os.Getenv("ProfilePath")
	if profilePath == "" {
		profilePath = "/var/opt/nuance/ncs/userdata/"
	}
	Info.Printf("Path is %s", profilePath)

	runTime := os.Getenv("RunTime")
	if runTime == "" {
		runTime = "16:00"
	}
	Info.Printf("RunTime is %s", runTime)

	delProfile(profilePath, true)
	gocron.Every(1).Day().At(runTime).Do(delProfile, profilePath, false)
	<-gocron.Start()
}

func delProfile(profilePath string, dryRun bool) {
	Info.Printf("Run Del Profile task, the path is %s, the DryRun is %t", profilePath, dryRun)
	err := delete(profilePath, dryRun)
	if err != nil {
		Error.Printf("Error happens on %s,", err)
	}
}

func delete(pathname string, dryRun bool) error {
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			delete(pathname+fi.Name()+"/", dryRun)
		} else {
			if isProfile(pathname + fi.Name()) {
				if dryRun {
					Info.Printf("DryRun Delete %s", pathname+fi.Name())
				} else {
					err = os.Remove(pathname + fi.Name())
					if err != nil {
						Error.Printf("Error happens when del %s", pathname+fi.Name())
					} else {
						Info.Printf("Delete %s", pathname+fi.Name())
					}
				}
			}
		}
	}
	return err
}

func isProfile(path string) bool {
	return strings.HasSuffix(path, "_profile.zip")

}
