# Use the official PostgreSQL image as the base image
FROM postgres:latest

# Set the environment variables for the PostgreSQL container
ENV POSTGRES_USER your_user
ENV POSTGRES_PASSWORD your_password
ENV POSTGRES_DB your_db_name

# Copy the seed.sql file to the container
COPY seed.sql /docker-entrypoint-initdb.d/

# RUN apt-get update && apt-get install -y <package-name>

# Expose the default PostgreSQL port
EXPOSE 5432
