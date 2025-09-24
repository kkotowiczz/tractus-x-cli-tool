### Tractus-X CLI Tool 
This application is created for midterm event.
It consists three main commands:

- createAsset.go
- createContractDefinition.go
- createPolicy.go

#### Installation for development

1. First install cobra library:
`go get -u github.com/spf13/cobra@latest`

2. Install cobra-cli:
`go install github.com/spf13/cobra-cli@latest`

#### Running program:

1. without compilation:

In order to run particular command without compilation process type: `go run main.go $commandName --arg argValue`

2. with compilation:

- Run `go build -o bin/tractus-x-cli-tool main.go` for UNIX based systems or `go build -o bin/tractus-x-cli-tool.exe main.go` for Windows.
- Then to run the program from root directory type: `./bin/tractus-x-cli-tool $commandName --arg argValue`

To get all available commands run:

`./bin/tractus-x-cli-tool --help`

#### Usage:

1. createAsset command requires --assetId argument, exemplary usage `./bin/tractus-x-cli-tool createAsset --assetId 200`
2. createPolicy command requires --policyId argument, exemplary usage `./bin/tractus-x-cli-tool createPolicy --policyId 200`
