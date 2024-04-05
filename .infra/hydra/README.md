## hydra

Let's create the OAuth 2.0 Client:

```
client=$(docker compose exec hydra \
    hydra create client \
    --endpoint http://127.0.0.1:4445/ \
    --format json \
    --grant-type client_credentials)
# We parse the JSON response using jq to get the client ID and client secret:
client_id=$(echo $client | jq -r '.client_id')
client_secret=$(echo $client | jq -r '.client_secret')
echo 'client_id='$(echo $client | jq -r '.client_id')
echo 'client_secret='$(echo $client | jq -r '.client_secret')
```

Let's perform the client credentials grant:

```
docker compose exec hydra \
  hydra perform client-credentials \
  --endpoint http://127.0.0.1:4444/ \
  --client-id "$client_id" \
  --client-secret "$client_secret"
```

Let's perform token introspection on that token. Make sure to copy the token you just got and not the dummy value.

```
docker compose exec hydra \
  hydra introspect token \
  --format json-pretty \
  --endpoint http://127.0.0.1:4445/ \
  UDYMha9TwsMBejEvKfnDOXkhgkLsnmUNYVQDklT5bD8.ZNpuNRC85erbIYDjPqhMwTinlvQmNTk_UvttcLQxFJY
```

Next, we will perform the OAuth 2.0 Authorization Code Grant. For that, we must first create a client that's capable of performing that grant:

```
docker compose exec hydra \
    hydra create client \
    --endpoint http://127.0.0.1:4445 \
    --grant-type authorization_code,refresh_token \
    --response-type code,id_token \
    --format json \
    --scope openid --scope email --scope profile \
    --redirect-uri http://localhost:5173/auth-callback

code_client_id=$(echo $code_client | jq -r '.client_id')
code_client_secret=$(echo $code_client | jq -r '.client_secret')
```

The following command starts a server that serves an example web application.

```
docker compose exec hydra \
    hydra perform authorization-code \
    --client-id $code_client_id \
    --client-secret $code_client_secret \
    --endpoint http://127.0.0.1:4444/ \
    --port 5555 \
    --scope openid --scope offline
```