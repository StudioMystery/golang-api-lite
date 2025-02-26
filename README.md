# Golang API Lite

This project is a template to get you started with a golang + htmx + sqlite application.

The template follows my personal preferences when it comes to swagger generation, package organization, environment variables, and overall structure to make it easier to start new full stack project quickly.

## Get Started:

#### To set it up locally:

1. Install Golang
2. [OPTIONAL] `go get -u ./...` to update all packages at once
3. [OPTIONAL] `go mod tidy` to fix package organization
4. Create a sqlite database somewhere.
5. use "./SQL/foo_bars.sql" to add the schema to that db.
6. Create "config.json" files in _golang_api_lite_ and _tests_ folders.
7. [OPTIONAL] add more foo_bars records with tests.

#### To run the server:

- `go main.go`

## Customizations:

1. More environments can be defined in Config.go if required (prod, test, dev, etc.)
2. A config.json file should be created from example_config.json and added to the root project folder AND tests folder so tests run.
3. Dockerfile is just a sample for a Golang app.
4. Certs folder is where you store your .PEM files for encryption / authentication.
5. You can swap out the sqlite connection for different db if you want.

## Limitations:

1. Architectures is based on using an embedded sqlite db, so keep that in mind, as it may not be the best use-case for what you want.
2. I only listed a BasicAuth method because I wanted to keep this simple. Maybe in the future I'll add other auth schemes.
3. PageHandlers probably isn't organized the best. Perhaps I could combine the route methods and the html file / templating code to create modular components. For now, the limited functionality is enough.
4. I hate writing the swagger definitions by hand in JSON, but I also hate inline swagger comments even more. I also hate modifying existing, working code with swaggerest just to have explicit code-side swagger definitions. So for Golang, I haven't found an optimal solution. In light of all these annoying options, I decided to just use the raw JSON.
5. I much prefer the styling of jsend wrapping every API request to randomness.
6. I'm considering ripping out the gorm stuff and just hitting sqlite raw.

## Credits:

[Jsend by Clever Go](https://github.com/clevergo/jsend)
[Sqlite for Prod](https://www.youtube.com/watch?v=XcAYkriuQ1o)
[HTMX + Golang](https://www.youtube.com/watch?v=r-GSGH2RxJs)
[Swagger Editor](https://editor.swagger.io/)
[Postman to Swagger](https://metamug.com/util/postman-to-swagger/)
