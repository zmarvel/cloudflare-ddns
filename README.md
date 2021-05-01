# cloudflare-ddns

Update Cloudflare DNS records at a fixed interval via the Cloudflare API.


## Building

```
go build
```


## Running

First, create `cloudflare-ddns.json`. See `example-config.json` for an example.
You will need a Cloudflare API token (this is different from a Cloudflare API
key). Go to

1. My Profile
2. API Tokens
3. Create Token
   - Read permission for Zone.Zone
   - Write permission for Zone.DNS
4. Copy the token into the config file

Then, just run:

```
./cloudflare-ddns
```