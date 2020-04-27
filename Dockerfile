# Start from the latest golang base image
FROM golang:alpine

#RUN useradd -ms /bin/bash admin
#USER admin

RUN mkdir /app
ADD . /app
WORKDIR /app

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .


# Build the application
RUN go build -o main .

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["/app/main"]