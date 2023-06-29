install:
	go install github.com/goreleaser/goreleaser@latest
make:
# Do not forget to import GITHUB_TOKEN, GITLAB_TOKEN or GITEA_TOKEN env variable
# Then you should push a new tag to the repo, e.g. git tag -a v0.1.0 -m "First release" && git push origin v0.1.0
	goreleaser release --clean
runw:
	go build -o ./bin/goreleasertest.exe main.go && ./bin/goreleasertest.exe
runl:
	go build -o ./bin/goreleasertest main.go && ./bin/goreleasertest