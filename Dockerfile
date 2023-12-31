FROM ubuntu:latest

# Essential for using tls
RUN apt-get update
RUN apt-get install ca-certificates -y
RUN update-ca-certificates

# web port
EXPOSE 8080

ADD build/strategiesGo /app/strategiesGo
ADD networks.json /app/networks.json
ADD strategies.json /app/strategies.json
ADD substrate-strategies.json /app/substrate-strategies.json
RUN chmod +x /app/strategiesGo

WORKDIR /app
CMD /app/strategiesGo
