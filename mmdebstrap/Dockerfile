FROM debian:bookworm

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        mmdebstrap \
        ca-certificates \
        wget \
        gnupg && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /workspace

ENTRYPOINT ["mmdebstrap"]
