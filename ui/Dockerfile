FROM node:9.4.0-alpine

WORKDIR /ui

COPY package*.json ./
RUN npm install -qy

COPY . ./
RUN npm run build

EXPOSE 3000

CMD ["npm", "start"]