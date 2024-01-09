# Use the official Node.js 16 image as a parent image
FROM node:16-bullseye-slim


ARG AWS_ACCESS_KEY_ID
ARG AWS_SECRET_ACCESS_KEY


ENV AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID
ENV AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY

# Set the working directory
WORKDIR /app

# Install Puppeteer dependencies
RUN apt-get update \
    && apt-get install -y wget gnupg ca-certificates procps libxshmfence1 libatk-bridge2.0-0 libdrm2 libxkbcommon0 libxcomposite1 libxdamage1 libxrandr2 libcups2 libgtk-3-0 libasound2 chromium \
    && rm -rf /var/lib/apt/lists/*

# Copy package.json and package-lock.json (or yarn.lock)
COPY package*.json ./

# Install app dependencies, including Puppeteer
RUN npm install --only=production

# Copy the rest of the app source code
COPY . .

# Set permissions for the "node" user
RUN chown -R node:node /app

# Use non-root user
USER node

# Start the application
CMD ["node", "index.js"]
