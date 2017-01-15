FROM golang

# Copy the app source
RUN mkdir -p /go/src/dicebot
WORKDIR /go/src/dicebot
ADD . /go/src/dicebot

# Build & Install the application
RUN go get
RUN go install

# Expose the default gin port
EXPOSE 8080
CMD ["dicebot"]
