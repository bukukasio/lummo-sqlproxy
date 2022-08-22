#!/bin/bash
curl -s https://raw.githubusercontent.com/bukukasio/lummo-sqlproxy/master/scripts/install_prerequisites.sh | bash
LATEST_RELEASE=$(curl -L -s -H 'Accept: application/json' https://github.com/bukukasio/lummo-sqlproxy/releases/latest)
LATEST_VERSION_TAG=$(echo $LATEST_RELEASE | sed -e 's/.*"tag_name":"\([^"]*\)".*/\1/')
LATEST_VERSION="${LATEST_VERSION_TAG:1}"
ARTIFACT_URL=https://github.com/bukukasio/lummo-sqlproxy/releases/download/${LATEST_VERSION_TAG}/lummo-sqlproxy_${LATEST_VERSION}_Linux_x86_64.tar.gz
wget -qc  $ARTIFACT_URL -P /tmp && sudo tar -xvf /tmp/lummo-sqlproxy_${LATEST_VERSION}_Linux_x86_64.tar.gz -C /usr/local/bin/ lummo-sqlproxy >/dev/null 2>&1
lummo-sqlproxy connect --help
