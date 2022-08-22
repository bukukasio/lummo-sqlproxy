# Install lummo-sqlproxy

## MacOS

```bash
brew install bukukasio/tools/lummo-sqlproxy
```

## Ubuntu

```bash
curl -s https://raw.githubusercontent.com/bukukasio/lummo-sqlproxy/master/scripts/install.sh | bash
```

# Usage

**To connect a database instance** <br />
![gif directory](assets/connect.gif) <br />

```bash
lummo-sqlproxy connect --env=<env> --port=<port-number>
```

By default `env=dev` and `port=5432` and are optional <br />
<br />
**To disconnect instance** <br />
![gif directory](assets/disconnect.gif) <br />

```bash
lummo-sqlproxy disconnect
```

<br />

**For all commands** <br />

```bash
lummo-sqlproxy --help
```

 <br />

# Prerequisites:

- [cloud_sql_proxy](https://bukukas.atlassian.net/wiki/spaces/TD/pages/538148955/How+to+connect+CloudSQL)
- [gcloud](https://cloud.google.com/sdk/docs/install)

### Install prerequisites in single click (WIP)

```bash
curl -s https://raw.githubusercontent.com/bukukasio/lummo-sqlproxy/master/scripts/install_prerequisites.sh | bash
```
