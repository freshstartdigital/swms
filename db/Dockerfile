# Use the official PostgreSQL image as the base image
FROM postgres:latest

# Set the environment variables for the PostgreSQL container
ENV POSTGRES_USER your_user
ENV POSTGRES_PASSWORD your_password
ENV POSTGRES_DB your_db_name

# Copy the seed.sql file to the container
COPY seed.sql /docker-entrypoint-initdb.d/

# Enable remote connections
RUN echo "listen_addresses = '*'" >> /usr/share/postgresql/postgresql.conf.sample

# Allow all IP addresses to connect with a password
RUN echo "host all all 0.0.0.0/0 md5" >> /usr/share/postgresql/pg_hba.conf.sample

# Expose the default PostgreSQL port (not necessary if you're using -p with docker run)
# EXPOSE 5432
