1. Build the Docker Image

The first time you set up the project, or whenever you make changes to the Dockerfile, you’ll need to build the Docker image.

```bash
docker-compose build
```

2. Start the Application

Once the image is built, you can start the application and the PostgreSQL container using:

```bash
docker-compose up
```

This command will start the containers in the foreground and display the logs in your terminal. You can also add -d to run the containers in detached mode (in the background):

```bash
docker-compose up -d
```

3. View Logs

To view logs from your containers, use the following command:

```bash
docker-compose logs -f
```

This will show real-time logs from both the Gin backend and PostgreSQL containers. You can also specify a specific service (e.g., app or db) to see logs for only that service:

```bash
docker-compose logs -f app  # Logs for the Gin app
docker-compose logs -f db   # Logs for PostgreSQL
```

4. Stop the Application

To stop the application, use the following command:

```bash
docker-compose down
```

This will stop and remove the containers but will keep the volumes (database data) and networks intact.
5. Rebuild the Application

If you make changes to your code that require the Docker image to be rebuilt (such as changes in main.go, go.mod, go.sum, or the Dockerfile), you can rebuild and restart the application by using:

```bash
docker-compose up --build
```

This command rebuilds the Docker image and restarts the containers.
6. Update Dependencies

If you update dependencies in your go.mod or go.sum file, you’ll want to rebuild the image to install the new dependencies:

```bash
docker-compose build
docker-compose up -d
```

7. Restart Specific Services

If you want to restart only one service (e.g., the Gin application without restarting PostgreSQL), you can use the up command with the service name:

```bash
docker-compose up -d --build app
```

This will rebuild and restart only the app service, which is useful when you make code changes without needing to touch the database.
8. Removing Containers, Networks, and Volumes

If you want to completely clean up and remove containers, networks, and persistent data volumes, you can use:

```bash
docker-compose down -v
```

This will stop and remove everything associated with the application, including the database data stored in volumes.
9. Check the Status of Services

To check the status of running services in your Docker Compose setup:

```bash
docker-compose ps
```

This command shows which services are running and their statuses.
Quick Summary of Commands

    Build the Docker image: docker-compose build
    Start the application: docker-compose up or docker-compose up -d
    View logs: docker-compose logs -f
    Stop the application: docker-compose down
    Rebuild and restart: docker-compose up --build
    Update dependencies: docker-compose build && docker-compose up -d
    Restart specific service: docker-compose up -d --build app
    Remove containers, networks, and volumes: docker-compose down -v
    Check status of services: docker-compose ps