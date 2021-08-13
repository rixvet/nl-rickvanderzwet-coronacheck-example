Intro
=====
View attribute details of Dutch QR Code (prefixed NL2:) produced by CoronaCheck App


Setup (optional)
================
Note: Certicates already included for convience, to replicate the process


- Download 'Staat der Nederlanden EV Root CA' root certificate:
```
	# Note: Convience shortcut, get yourself a trusted/verified copy
	# somehow, one can never sure if the right is presented this way :-)
	# Also check with https://zoek.officielebekendmakingen.nl/stcrt-2011-527.html
	$ curl http://cert.pkioverheid.nl/EVRootCA.cer > EVRootCA.cer
```

- Convert 'Staat der Nederlanden EV Root CA' to format usable with verify command
```
	$ openssl x509 -in ./EVRootCA.cer -inform DER -out ./EVRootCA.pem
```


- Download VWS signing public keys:
```
	$ curl https://holder-api.coronacheck.nl/v4/holder/public_keys > public_keys
```


- Split payload and signature:
```
	$ cat public_keys | jq -r .payload | base64 -d > public_keys.payload
	$ cat public_keys | jq -r .signature | base64 -d > public_keys.signature
```


- Verify downloaded content:
```
	$ openssl cms  -verify -inform DER -in ./public_keys.signature \
	    -content ./public_keys.payload -purpose any \
	    -CAfile ./EVRootCA.pem  -CApath /var/empty > /dev/null
```


- [only if validation is OK] Export/Parse VWS:
```
	$ cat public_keys.payload | jq -r '.cl_keys | .[] | \
	     select(.id == "VWS-CC-1") | .public_key' | base64 -d > VWS-CC-1.xml
	$ cat public_keys.payload | jq -r '.cl_keys | .[] | \
	     select(.id == "VWS-CC-2") | .public_key' | base64 -d > VWS-CC-2.xml
```


Usage
=====
- Scan NL QR code and store resulting string:
```
	$ cat <<'EOF' > ./nl2-qrcode.txt
	EOF
```
- Execute to view certificate attributes:
```
	$ go run main.go ./nl2-qrcode.txt
```

Misc
====
The folder `testdata` contains some sample certificates which are signed by dummy 'testPk' issuer
```
$ go run main.go testdata/qrcode1.sample 
{
  "birthDay": "20",
  "birthMonth": "10",
  "firstNameInitial": "A",
  "isPaperProof": "0",
  "isSpecimen": "0",
  "lastNameInitial": "R",
  "validForHours": "24",
  "validFrom": "1628510400"
}
```
