FROM golang:alpine

WORKDIR /

COPY . .

EXPOSE 11500

ENTRYPOINT ["go", "run", "main.go"]

# To run this dockerfile
# $ docker run -p 11500:11500 simple
# In a new tab $cd client; go run main.go