

username="${1:-"username"}"

password="${2:-"password"}"

response=$( curl -sSL "http://0.0.0.0:8080/login?username=${username}&password=${password}" )

echo "response: ${response}"
