FROM ghcr.io/celestiaorg/celestia-node:v0.8.2 AS celestia-node

FROM ghcr.io/celestiaorg/celestia-app:0.12.2

COPY --from=celestia-node /bin/celestia /

RUN apk update && apk --no-cache add curl jq libc6-compat

COPY celestia-devnet.entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh 

RUN cat /entrypoint.sh
RUN ls -la /entrypoint.sh

EXPOSE 26657 26659 9090

ENTRYPOINT ["/entrypoint.sh"]
