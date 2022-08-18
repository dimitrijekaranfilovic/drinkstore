# syntax=docker/dockerfile:1

FROM node:lts-alpine as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build-proxy

# production stage
FROM nginx:stable-alpine as production-stage
RUN rm /etc/nginx/nginx.conf
COPY nginx.conf /etc/nginx/

COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]