# Generate a CA

Generate a key

```
	openssl genrsa -out ca.key 4096
```

Create a CSR

```
	openssl req -new -key ca.key -out ca.csr -subj '/CN=dns-edit'
```

Create an extension file - `ca.ext`

```
	basicConstraints = critical, CA:TRUE, pathlen:0
	keyUsage = critical, cRLSign, digitalSignature, keyCertSign
	subjectAltName = DNS:dns-edit
```

Sign the cerificate

```
	openssl x509 -req -in ca.csr -out ca.crt -signkey ca.key -days 365 -sha256 -extfile ca.ext
```

# Generate client cert
Generate a client key

```
	openssl genrsa -out webhook.key 4096
```

Create a configuration file - `webhook.cnf`

```
	[ req ]
	distinguished_name = dn
	req_extensions = req_ext
	x509_extensios = x509_ext

	[ dn ]
	emailAddress = admin@webhook.com

	[ req_ext ]
	basicConstraints = CA:FALSE
	keyUsage = nonRepudiation, digitalSignature, keyEncipherment
	extendedKeyUsage = serverAuth, clientAuth, codeSigning

	[ x509_ext ]
	basicConstraints = CA:FALSE
	keyUsage = nonRepudiation, digitalSignature, keyEncipherment
	extendedKeyUsage = serverAuth, clientAuth, codeSigning
```
		
Generate a CSR

```
	openssl req -new -key webhook.key -out webhook.crt -config webhook.cnf
```

Sign the request with the CA cert and key

```
	openssl x509 -req -in webhook.csr -out webhook.crt \
		-CA ca.crt -CAkey ca.key -CAcreateserial \
		-days 1000 -sha256 -extfile webhook.ext
```
