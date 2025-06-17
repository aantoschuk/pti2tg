FROM golang:1.24.4-bullseye

RUN apt-get update && \
    apt-get install -y tesseract-ocr \
                       libleptonica-dev \
                       libtesseract-dev && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Set up the Go app
WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o pti2tg .

ENTRYPOINT ["./pti2tg"]
