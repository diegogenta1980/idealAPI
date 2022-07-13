FROM golang:latest
WORKDIR /app
EXPOSE 3000

# Download Go modules
# COPY go.mod .
# COPY go.sum .
# RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
# COPY *.go ./
# COPY *.json ./
# COPY asset/ ./

ADD . /app/

# Build
# RUN go build -o idealAPI
# CMD ["go", "build", "main.go"]

#CMD [ "/idealAPI" ]

CMD ["go", "run", "main.go"]
