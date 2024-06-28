FROM golang:1.22 as restaurant-base

LABEL stage="restaurant base"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./cmd/api/kitchen/main ./cmd/api/kitchen/main.go
RUN go build -o ./cmd/api/orders/main ./cmd/api/orders/main.go
RUN go build -o ./cmd/api/payment/main ./cmd/api/payment/main.go


FROM golang:1.22 as orders-app

LABEL stage="orders app"

WORKDIR /app

COPY --from=restaurant-base /app/cmd/api/orders/main /app/

EXPOSE 8001

CMD [ "./main" ]

FROM golang:1.22 as payment-app

LABEL stage="payment app"

WORKDIR /app

COPY --from=restaurant-base /app/cmd/api/payment/main /app/

EXPOSE 8002

CMD [ "./main" ]

FROM golang:1.22 as kitchen-app

LABEL stage="kitchen app"

WORKDIR /app

COPY --from=restaurant-base /app/cmd/api/kitchen/main /app/

EXPOSE 8003

CMD [ "./main" ]
