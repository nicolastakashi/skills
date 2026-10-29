FROM golang:1.24-bookworm
WORKDIR /app
COPY app/ /app/
# A real git history lets review tools such as /code-review diff the PR against main.
RUN git init -q -b main && git config user.email eval@example.com && git config user.name eval \
 && git apply -R PR.diff && git add -A && git commit -qm base \
 && git checkout -qb pr && git apply PR.diff && git add -A && git commit -qm pr
