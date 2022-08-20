# syntax=docker/dockerfile:1

FROM rust:1.62

COPY ./src ./src/
COPY ./*.lock ./
COPY ./*.toml ./

ENV POSTGRES_DB=ntp-purchase-service
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=root
ENV POSTGRES_HOST=purchase-service-database
ENV ROCKET_ADDRESS=0.0.0.0
ENV ROCKET_PORT=8084

RUN cargo build --release

EXPOSE 8084

CMD ["./target/release/purchase-service"]