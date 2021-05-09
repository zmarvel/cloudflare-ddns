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


## Installation

No need to run this as root. Let's first add a user.

```
sudo useradd --system --shell /bin/false --user-group cloudflare-ddns
```

Next, copy the `cloudflare-ddns` executable to somewhere that the new user can execute it, for example

```
sudo cp cloudflare-ddns /usr/local/bin
sudo chmod 755 /usr/local/bin/cloudflare-ddns
```

Copy the config somewhere that **only** the cloudflare-ddns user (or group) can read it (see
`example-config.json`). Take care to restrict the permissions since the file contains your private
API token!

```
sudo mv cloudflare-ddns.json /etc
sudo chown cloudflare-ddns:cloudflare-ddns /etc/cloudflare-ddns.json
sudo chmod 400 /etc/cloudflare-ddns.json
```

Create a systemd unit. See `cloudflare-ddns.service` as a starting point. If you picked a different
username or a different path, you may have to change things around.

```
sudo cp cloudflare-ddns.service /etc/systemd/system
sudo systemctl daemon-reload
sudo systemctl enable cloudflare-ddns
sudo systemctl start cloudflare-ddns
```
