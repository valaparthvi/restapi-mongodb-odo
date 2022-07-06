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
