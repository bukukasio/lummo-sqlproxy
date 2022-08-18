# Usage

**To connect a database instance** <br />
![gif directory](assets/connect.gif) <br />
`lummo-sqlproxy connect --env=<env> --port=<port-number>` <br />
By default `env=dev` and `port=5432` and are optional <br />
<br />
**To disconnect instance** <br />
![gif directory](assets/disconnect.gif) <br />
`lummo-sqlproxy disconnect` <br />
<br />
**For all commands** <br />
`lummo-sqlproxy --help` <br />

# Install

To setup run: `GOPRIVATE="github.com/bukukasio/*" go install github.com/bukukasio/lummo-sqlproxy@latest`

## Prerequisites:

- [cloud_sql_proxy](https://bukukas.atlassian.net/wiki/spaces/TD/pages/538148955/How+to+connect+CloudSQL)
- [gcloud](https://cloud.google.com/sdk/docs/install)
