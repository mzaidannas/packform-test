# Get Go image from DockerHub.
FROM golang:alpine AS api

# Set required build arguments for node
ARG NODE_ENV="production"

RUN apk update && apk add --update nodejs npm

# Set working directory.
WORKDIR /compiler

# Copy dependency locks so we can cache.
COPY go.mod go.sum ./

# Get all of our dependencies.
RUN go mod download

# Copt frontend dependency and lock files
COPY package.json package-lock.json ./

# Download frontend dependencies
RUN npm install

# Copy all of our remaining application.
COPY . .

# Build our application.
RUN CGO_ENABLED=0 GOOS=linux go build -o packform-test ./main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o import-data ./src/scripts/import_data.go

# Build frontend application
RUN npm run build-only

# Use 'scratch' image for super-mini build.
FROM scratch AS prod

# Set working directory for this stage.
WORKDIR /production

# Copy our compiled executable from the last stage.
COPY --from=api /compiler/packform-test /compiler/import-data /compiler/dist ./

# Run application and expose port 8080.
EXPOSE 3000

CMD ["./packform-test"]
