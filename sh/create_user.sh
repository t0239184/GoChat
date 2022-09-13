#/bin/sh
ACCOUNT=$1
PASSWORD=$2

if [[ -z $ACCOUNT ]] || [[ -z $PASSWORD ]]; then
    echo "Usage: $0 <account> <password>"
    exit 1
fi

curl -X POST -v http://localhost:8080/api/v1/user -H 'Content-Type: application/json' -d  "{\"account\":\"$ACCOUNT\",\"password\":\"$PASSWORD\"}"
