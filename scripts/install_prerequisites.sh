#!/bin/bash
if which cloud_sql_proxy &>/dev/null
then
{    
    echo "cloud_sql_proxy is installed"
}
else
{    
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    {
        wget https://dl.google.com/cloudsql/cloud_sql_proxy.linux.amd64 -O cloud_sql_proxy
    }
    elif [[ "$OSTYPE" == "darwin"* ]]; then
    {
        if [ "${arch_name}" = "x86_64" ]; then
        {
            curl -o cloud_sql_proxy https://dl.google.com/cloudsql/cloud_sql_proxy.darwin.amd64
        }
        else
        {
            curl -o cloud_sql_proxy https://dl.google.com/cloudsql/cloud_sql_proxy.darwin.arm64
        }
        fi
    }
    fi
    chmod +x cloud_sql_proxy
    sudo mv cloud_sql_proxy /usr/local/bin/
}
fi    