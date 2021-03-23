# FROM node:14
FROM node:13.12.0-alpine
WORKDIR /srv/app
ENV PATH /node_modules/.bin:$PATH
COPY package*.json ./
RUN npm install
COPY . .
EXPOSE 3000
# CMD [ "npm" ,"run", "start" ]