
Here's an improved version of your README.md with better structure, formatting, and clarity:

Commands to Run This Project
1. Build and Start the Project
bash
Sao chép mã
docker compose up -d --build
Builds the project and starts the containers in detached mode.
Use this command when you make changes to the Dockerfile or docker-compose.yml to ensure everything is rebuilt from scratch.
2. Restart the Containers
bash
Sao chép mã
docker compose restart
Restarts all running containers without rebuilding the images.
3. Stop and Remove Containers
bash
Sao chép mã
docker compose down
Stops the running containers and removes them, but keeps the volumes intact to preserve data.
4. Stop, Remove Containers, and Clear Volumes
bash
Sao chép mã
docker compose down -v
Stops the containers, removes them, and deletes all associated volumes.
Use this if you want a clean state for your containers and volumes.