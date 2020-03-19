FROM jguyomard/hugo-builder:latest AS builder

COPY frontend app
COPY raw app/static/raw
COPY gen app/static/gen
WORKDIR app/
RUN hugo

FROM nginx:latest
COPY --from=builder /src/app/public/ /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
