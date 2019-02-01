# ip-director
Direct IPs from different geographical locations to different URLs.

# Server installation

go get -u github.com/gorilla/mux

go get github.com/oschwald/geoip2-golang

go get github.com/stretchr/testify/assert

go build ipDirector.go

./ipDirector

# Client usage

In directInfo.txt, the second URL is to direct IPs in mainland China, the first URL is to direct the rest of IPs.

You can revise these URLs to what you want.

_Direct through current visitor's IP:_

http://127.0.0.1:8012/ipToCountry

_Direct through a submitted IP for test:_

http://127.0.0.1:8012/ipToCountrySubmitAddr/123.116.97.159


# Notes

Current codes were deployed in cloudflare CDN enviroument, 

so I setup X-real-ip in nginx to get the real IP. If you 

are not in cloudflare enviroument, please comment this line:

clinetIP = r.Header.Get("X-real-ip")

and then remove comment on this line:

//clinetIP, _, _ = net.SplitHostPort(r.RemoteAddr)
