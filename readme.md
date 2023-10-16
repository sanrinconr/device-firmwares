# Device Firmware API

The purpose of this API is to expose services for interacting with the firmwares associated with a device. You can:

- Create, modify, or delete a device (which can have zero associated firmwares).
- Create, modify, or delete a firmware (always must be associated with one device).
- Obtain data about firmwares and devices.

Every request to the API must be authenticated with an "Authorization" header that has a base64-encoded value in the format of `Basic user:password`.

For example, the client must send `Authorization: Basic dGVzdFVzZXI6c2VjdXJlUGFzc3dvcmQ=` (decoded as: `testUser:securePassword`).

For a more detailed description of the services, you can use the Postman collection provided in the following link:

[dc-nearshore Postman Collection](docs/dcnearshore.postman_collection.json)

## How to Run

The best way to run the app is by using Docker. To run the app, execute the following steps:

#### Create .env file
You can copy and paste .env.template, the default value are enough, remember to rename to `.env`

#### Run container
Execute in your terminal
```bash
sh run.sh
```

(docker and docker-compose are required)

Wait 30s-50s until container are up.

This script will start the GoLang and Postgres containers, you can access the app at `localhost:8080``. Please refer to the Postman collection for interaction details: [postman collection](docs/dcnearshore.postman_collection.json)

## Architecture

The application is based on the Domain-Driven Design architecture (inspired by the proposal of Ben Johnson). There are the following created packages:

- dc-nearshore (alias domain): Defines the business logic and its behavior.
- auth: In-memory implementation of authentication (only used for "login" in the services). It uses bcrypt as the hashing algorithm.
- cmd
    - api: Initializes the web framework.
    - config: Reads configuration from the .env file.
    - controller: Defines the behavior of every exposed service.
    - dependencies: Manages dependency injection.
    - middlewares: Middleware applied to the web framework.
- database: Provides methods for saving, updating, getting, and deleting domain objects (abstraction of gorm).
- devicer: Implements the necessary behavior of a device.
- firmwarer: Implements the necessary behavior of a firmware.
- log: Provides tools for creating logs in the application.

### TODOs
- Create tests.
- Make available online.
- Create a swagger file.
- Add comments to the core parts.
- Comments to every package.



