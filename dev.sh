# development script, this is kinda redundant at the moment because our
# dev and prod environments are exactly the same. Soon we'll separate out our
# dev and prod into their own docker images, and then that will require a dev.sh
# echo "GO TO localhost:8081 🫡🫡🫡🫡🫡🫡🫡🫡🫡🫡🫡🫡🫡🫡🫡🫡🫡🫡🫡🫡"
# docker compose up --build

docker compose -f docker/docker-compose.yaml up webserver-dev tailwindcss --build
