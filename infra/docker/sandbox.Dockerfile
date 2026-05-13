FROM ubuntu:24.04

RUN apt-get update && apt-get install -y --no-install-recommends \
    golang-go \
    python3 \
    python3-pip \
    nodejs \
    && rm -rf /var/lib/apt/lists/*

RUN groupadd -g 1000 sandbox && \
    useradd -u 1000 -g sandbox -m -s /bin/bash sandbox

WORKDIR /workspace
RUN chown sandbox:sandbox /workspace

USER sandbox
