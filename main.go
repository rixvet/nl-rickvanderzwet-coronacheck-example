package main

import (
	"github.com/go-errors/errors"
	"github.com/minvws/nl-covid19-coronacheck-idemix/verifier"
	"github.com/privacybydesign/gabi"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)


func PrettyPrint(v interface{}, created int64) (err error) {
	data := map[string]interface{}{
		"created": created,
		"attributes": v,
	}
	
	b, err := json.MarshalIndent(data, "", "  ")
	if err == nil {
	fmt.Println(string(b))
	}
	return
}

func main() {
	if (len(os.Args[1:]) < 1) {
		fmt.Println("Usage: ", os.Args[0], "<NL2_certfile>")
		os.Exit(128)
	}
	certFile := os.Args[1]

	/* Read QR Code (base45 encoded) */
	proofPrefixedRaw, err := ioutil.ReadFile(certFile)

	/* Strip newlines in file */
	delCRLF := func (r rune) rune {
		switch {
		case r == '\n':
			return -1
		case r == '\r':
			return -1
		}
		return r
	}
	proofPrefixed := bytes.Map(delCRLF, proofPrefixedRaw)

	if err != nil {
		panic(err)
	}

	/* Verify */
	v := createVerifier2()
	verifiedCred, err := v.VerifyQREncoded(proofPrefixed)
	if err != nil {
		fmt.Println("Could not verify disclosed credential:", err.Error())
	}

	/* Display */
	PrettyPrint(verifiedCred.Attributes, verifiedCred.DisclosureTimeSeconds)
}


func createVerifier2() *verifier.Verifier {
	return verifier.New(holderFindIssuerPk2)
}

func holderFindIssuerPk2(issuerPkId string) (*gabi.PublicKey, error) {
	issuerPks := map[string]*gabi.PublicKey{
		"testPk": testPk,
		"VWS-CC-1": vwsCc1Pk,
		"VWS-CC-2": vwsCc2Pk,
	}
	issuerPk, ok := issuerPks[issuerPkId]
	if !ok {
		return nil, errors.Errorf("Could not find public key id (%s) chosen by issuer", issuerPkId)
	}

	return issuerPk, nil
}

var testPk, _ = gabi.NewPublicKeyFromFile("./testPk.xml")
var vwsCc1Pk, _ = gabi.NewPublicKeyFromFile("./VWS-CC-1.xml")
var vwsCc2Pk, _ = gabi.NewPublicKeyFromFile("./VWS-CC-2.xml")
