# qnin-sp

```
openssl req -x509 -sha256 -nodes -days 3650 -newkey rsa:2048 -subj /CN=sp.example.org -keyout sp.example.org.key -out sp.example.org.crt
openssl req -x509 -sha256 -nodes -days 3650 -newkey rsa:2048 -subj /CN=idp.example.com -keyout idp.example.com.key -out idp.example.com.crt
cp idp.example.com.* ./proxy/cert/
cp idp.example.com.* ./saml-idp/cert/
cp sp.example.org.* ./proxy/cert/
```