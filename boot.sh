docker build -t metric-broadcaster -f ./metric-broadcaster/Dockerfile ./metric-broadcaster && \
docker build -t metric-generator -f ./metric-generator/Dockerfile ./metric-generator && \
docker build -t metric-frontend -f ./metric-frontend/Dockerfile ./metric-frontend && \
docker build -t metric-proxy -f ./metric-proxy/Dockerfile ./metric-proxy && \
docker-compose up
