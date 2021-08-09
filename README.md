Intro
=====
View attribute details of Dutch QR Code (prefixed NL2:) produced by CoronaCheck App


Setup
=====
- Download VWS signing public keys:
	$ curl https://holder-api.coronacheck.nl/v4/holder/public_keys > public_keys
	$ # TODO: Validate authenticity of file

- Export/Parse VWS:
	$ cat public_keys | jq -r .payload | base64 -d  | jq -r '.cl_keys | .[] | select(.id == "VWS-CC-1") | .public_key' | base64 -d > VWS-CC-1.xml
	$ cat public_keys | jq -r .payload | base64 -d  | jq -r '.cl_keys | .[] | select(.id == "VWS-CC-2") | .public_key' | base64 -d > VWS-CC-2.xml



Usage
=====
- Scan NL QR code and store resulting string:
	$ cat <<'EOF' > ./nl2-qrcode.txt
	EOF

- Execute to view certificate attributes:
	$ go run main.go ./nl2-qrcode.txt


Misc
====
The folder 'testdata' contains some sample certificates which are signed by dummy 'testPk' issuer
