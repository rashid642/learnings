FROM ubuntu

RUN apt-get update && \
    apt-get install -y \
    curl \
    gnupg2 \
    build-essential

RUN curl -sL https://deb.nodesource.com/setup_14.x | bash 
RUN apt-get install -y nodejs

COPY package.json package.json
COPY package-lock.json package-lock.json
COPY main.js main.js 

RUN npm install

ENTRYPOINT ["node", "main.js"]