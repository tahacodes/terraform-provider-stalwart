#!/usr/bin/env bash
set -euo pipefail

IMAGE="${STALWART_IMAGE:-stalwartlabs/stalwart:v0.16.17}"
NAME="${STALWART_CONTAINER:-stalwart-acc}"
PORT="${STALWART_PORT:-18080}"
ADMIN_USER="admin"
ADMIN_PASSWORD="acctest123"
URL="http://127.0.0.1:${PORT}"

docker rm -f "${NAME}" >/dev/null 2>&1 || true
docker run -d --name "${NAME}" -p "${PORT}:8080" \
	-e "STALWART_RECOVERY_ADMIN=${ADMIN_USER}:${ADMIN_PASSWORD}" \
	"${IMAGE}" >/dev/null

wait_ready() {
	for _ in $(seq 1 60); do
		if curl -sf -o /dev/null --max-time 5 -u "${ADMIN_USER}:${ADMIN_PASSWORD}" "${URL}/jmap/session"; then
			return 0
		fi
		sleep 2
	done
	echo "stalwart did not become ready" >&2
	exit 1
}

wait_ready

curl -sf --max-time 60 -u "${ADMIN_USER}:${ADMIN_PASSWORD}" -X POST \
	-H "Content-Type: application/json" \
	-d '{"using":["urn:ietf:params:jmap:core"],"methodCalls":[["x:Bootstrap/set",{"accountId":"d333333","update":{"singleton":{"defaultDomain":"example.com","serverHostname":"mail.example.com","generateDkimKeys":false,"requestTlsCertificate":false,"blobStore":{"@type":"Default"},"dataStore":{"@type":"RocksDb","path":"/var/lib/stalwart/"},"directory":{"@type":"Internal"},"dnsServer":{"@type":"Manual"},"inMemoryStore":{"@type":"Default"},"searchStore":{"@type":"Default"},"tracer":{"@type":"Log","path":"/var/log/stalwart/"}}}},"c0"]]}' \
	"${URL}/jmap" >/dev/null

docker restart "${NAME}" >/dev/null
wait_ready

cat <<EOF
Stalwart ready for acceptance tests.

  export TF_ACC=1
  export STALWART_URL=${URL}
  export STALWART_USER=${ADMIN_USER}
  export STALWART_PASSWORD=${ADMIN_PASSWORD}

  make testacc

Remove it with: docker rm -f ${NAME}
EOF
