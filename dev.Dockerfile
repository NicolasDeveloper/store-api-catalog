FROM golang:latest

RUN apt update && apt upgrade -y && \
    apt install -y git \
    make openssh-client

WORKDIR /app 


RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

RUN JQ_URL="https://circle-downloads.s3.amazonaws.com/circleci-images/cache/linux-amd64/jq-latest" \
  && curl --silent --show-error --location --fail --retry 3 --output /usr/bin/jq $JQ_URL \
  && chmod +x /usr/bin/jq \
  && jq --version

RUN download_url=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url') \
    && curl -o /usr/bin/swagger -L'#' "$download_url" \
    && chmod +x /usr/bin/swagger

WORKDIR / 

RUN mkdir swagger-ui && cd swagger-ui && \
    curl -L https://github.com/swagger-api/swagger-ui/archive/master.tar.gz | tar -xz --strip=2 swagger-ui-master/dist/ && \
    sed -i "s/https\:\/\/petstore.swagger.io\/v2/\./g" index.html

EXPOSE 8080

WORKDIR /app

CMD air