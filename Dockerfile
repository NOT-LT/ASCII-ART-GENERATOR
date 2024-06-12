# Base Image
FROM golang:latest

LABEL maintainer="Ameena, Zahraa, Khadija"
LABEL description="Ascii Art Generator"

WORKDIR /Ascii-Web
COPY . .

#Creates a binary file named main in the new environment
RUN go build -o main .

EXPOSE 8080

#Run the binary file called main
CMD ["./main"]