# syntax=docker/dockerfile:

FROM nginx:1.23-alpine

RUN rm /etc/nginx/conf.d/default.conf
RUN rm /etc/nginx/nginx.conf
COPY . /etc/nginx

VOLUME /var/log/nginx/user_service
VOLUME /var/log/nginx/drink_service
VOLUME /var/log/nginx/comment_service
VOLUME /var/log/nginx/purchase_service

VOLUME /var/log/nginx/user_service/access
VOLUME /var/log/nginx/user_service/error

VOLUME /var/log/nginx/drink_service/access
VOLUME /var/log/nginx/drink_service/error

VOLUME /var/log/nginx/comment_service/access
VOLUME /var/log/nginx/comment_service/error

VOLUME /var/log/nginx/purchase_service/access
VOLUME /var/log/nginx/purchase_service/error


EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
