install:
	go install github.com/goreleaser/goreleaser@latest
make:
# Do not forget to import GITHUB_TOKEN, GITLAB_TOKEN or GITEA_TOKEN env variable
	goreleaser release
runw:
	go build -o ./bin/goreleasertest.exe main.go && ./bin/goreleasertest.exe
runl:
	go build -o ./bin/goreleasertest main.go && ./bin/goreleasertest