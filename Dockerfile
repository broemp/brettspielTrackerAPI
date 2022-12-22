FROM golang
WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /brettspielTrackerAPI
CMD [ "/brettspielTrackerAPI"]
EXPOSE 8080
