# golens 🌿 
## What's this ?
`golens` is a wrapper for Lens Protocol written in Go.
 This is an proof of concept implementation of a simple lens library for Go.
## POC Goals
- [x] Search for a lens profile by handle
- [x] Get profile follower revenue by profileId
- [x] Explore profiles based on criteria with cursor flag
- [ ] List all lens profiles
- [ ] List all lens followers for a given profile
- [ ] JSON export for results
## Project Status
This project is currently in a proof of concept state. It is not ready for production use.
## Sources
- Graphql examples https://github.com/lens-protocol/api-examples/tree/master/src/graphql
- Lens graphql https://api.lens.dev/ 
## Local Development
```bash
# Run tests
go test -cover ./...
# Run cli 
go run cmd/golens/main.go
```


# Lens Graphql Schema
## setting up dev environment
```bash
npm install -g apollo
npm install -g graphql
npx apollo client:download-schema --endpoint=https://api.lens.dev/ schema.graphql
```
## Generate go code from schema
```bash
cd genqlientFiles
# edit genqlient.graphql to add the schema
vim genqlient.graphql
# run genqlient to generate the go code
go run github.com/Khan/genqlient
# add missing bindings in genqlient.yaml if any
vim genqlient.yaml
```