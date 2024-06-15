

token="${1:-"0000"}"

username="${2:-"username"}"

host="${3:-"sub.domain.com"}"

response=$( curl -sSL "http://0.0.0.0:8080/describeuser?token=${token}&username=${username}&host=${host}" )

echo "response: ${response}"
