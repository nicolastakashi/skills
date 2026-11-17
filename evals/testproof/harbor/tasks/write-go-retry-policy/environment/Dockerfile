FROM golang:1.24-bookworm
WORKDIR /app
COPY app/ /app/
RUN if [ -f /app/.setup.sh ]; then sh /app/.setup.sh && rm /app/.setup.sh; fi
