FROM golang:1.16.5-buster

#Selects everything in current directory and copies 
ADD ./ /asciicontainer

# Move to working directory /asciicontainer
WORKDIR /asciicontainer

# Build the application
RUN go build -o containerapp

# Metadata
LABEL meta-data.app-name="ascii-art-web"
LABEL meta-data.author.autor1="Robert (nimi25820)"

# Export necessary port
EXPOSE 8080

# Command to run when starting the container
CMD ["./containerapp"]
