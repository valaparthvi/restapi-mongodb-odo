# Install application and service with odo:
1. Install percona distribution mongodb operator, and service binding operator on cluster
2. odo create namespace mongodb-restapi
3. odo init --devfile=go --name=restapi
4. kubectl apply -f https://raw.githubusercontent.com/percona/percona-server-mongodb-operator/v1.12.0/deploy/cr-minimal.yaml
5. odo add binding --service minimal-cluster --name my-binding --bind-as-files=false

# REST API with Go and MongoDB
REST API with Golang and MongoDB.
* HTTP router
* CRUD operations
* MongoDB supported driver for Go 
* ENV configurations
* Formatted logs

## API Endpoints
- GET `/api/places` - List all places
- POST `/api/places` - Add a new place
- PUT `/api/places` - Update a place
- GET `/api/places/<id>` - Fetch place with id `<id>`
- DELETE `/api/places/<id>` - Delete place with id `<id>`


## Resources
mux: A powerful HTTP router and URL matcher for building Go web servers. https://github.com/gorilla/mux

mongo-driver: The MongoDB supported driver for Go. https://github.com/mongodb/mongo-go-driver

viper: Viper is a complete configuration solution for Go applications including 12-Factor apps. https://github.com/spf13/viper

CompileDaemon: Watches your .go files in a directory and invokes go build if a file changed. https://github.com/githubnemo/CompileDaemon
