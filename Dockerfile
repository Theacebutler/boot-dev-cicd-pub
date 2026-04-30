FROM  debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates

ADD notely /usr/bin/notely
ENV PORT=3000
CMD ["notely"]
