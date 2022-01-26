FROM golang:1.17.6-buster
RUN go env -w GOPROXY=direct
WORKDIR /mydiary
COPY go.mod go.mod
RUN go mod download
RUN go mod tidy
RUN go install golang.org/x/tools/gopls@latest