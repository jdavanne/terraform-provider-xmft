package provider

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// openssl req -x509 -newkey rsa:1024 -keyout key.pem -out cert.pem -sha256 -days 3650 -nodes -subj "/C=XX/ST=StateName/L=CityName/O=CompanyName/OU=CompanySectionName/CN=CommonNameOrHostname"
func GenSelfSignedCert(keySize int) (string, error) {
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(2019),
		Subject: pkix.Name{
			Organization:  []string{"Company, INC."},
			Country:       []string{"US"},
			Province:      []string{""},
			Locality:      []string{"San Francisco"},
			StreetAddress: []string{"Golden Gate Bridge"},
			PostalCode:    []string{"94016"},
		},
		IPAddresses: []net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().AddDate(10, 0, 0),
		// IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
		SignatureAlgorithm:    x509.SHA256WithRSA,
	}

	caPrivKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return "", err
	}
	caBytes, err := x509.CreateCertificate(rand.Reader, ca, ca, &caPrivKey.PublicKey, caPrivKey)
	if err != nil {
		return "", err
	}
	caPEM := new(bytes.Buffer)
	pem.Encode(caPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caBytes,
	})

	/*caPrivKeyPEM := new(bytes.Buffer)
	pem.Encode(caPrivKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(caPrivKey),
	})*/

	caPEMStr := caPEM.String()
	return caPEMStr, nil
}

func TestNewSTCertificateResource(t *testing.T) {
	t.Parallel()
	r := time.Now().Format("-2006-01-02_15-04-05") + "-" + fmt.Sprint(time.Now().UnixNano())
	resourceType := "xmft_st_certificate"
	resourceName := "cert1"
	name := "cert1" + r

	account := "account1" + r

	cert, err := GenSelfSignedCert(1024)
	if err != nil {
		t.Fatalf("Failed to generate self-signed certificate: %v", err)
	}
	cert2, err := GenSelfSignedCert(2048)
	if err != nil {
		t.Fatalf("Failed to generate self-signed certificate: %v", err)
	}

	localConfig := `
resource "xmft_st_account" "account1" {
		provider    = xmft.st1
		name        = "` + account + `"
		#type        = "user"
		uid         = "1000"	
		gid         = "1000"
		home_folder = "/files/account1"
		user = {
		  	name =  "` + account + `"
		  	password_credentials = {
				password = "zouzou"
		  	}
		}
}
	`

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + localConfig + `
					resource "` + resourceType + `" "` + resourceName + `" {
						provider         = xmft.st1
						name             = "` + name + `"
						account 		 = xmft_st_account.account1.name
						type             = "x509"
						usage            = "login"
						#overwrite        = true
						content          = "` + strings.ReplaceAll(cert, "\n", "\\n") + `"
						
					}
					`,
				Check: resource.ComposeAggregateTestCheckFunc(
					checkResources([]byte(`
					resource "`+resourceType+`" "`+resourceName+`" {
						#provider        = xmft.st1
						name             = "`+name+`"
						#account 		 = xmft_st_account.account1.name
						type             = "x509"
						usage            = "login"
						#overwrite       = true
						validity_period  = "3652"
						sign_algorithm    = "SHA256WITHRSA"
						key_algorithm	 = "RSA"
						key_size		  = "1024"
						subject          = "O=\"Company, INC.\", OID.2.5.4.17=94016, STREET=Golden Gate Bridge, L=San Francisco, ST=, C=US"
					}
					`)),
					resource.TestCheckResourceAttrSet(resourceType+"."+resourceName, "last_updated"),
				),
			},
			/*// ImportState testing
			{
				ResourceName:      "xmft_staccount.account1",
				ImportState:       true,
				ImportStateVerify: true,
				// The last_updated attribute does not exist in the HashiCups
				// API, therefore there is no value for it during import.
				ImportStateVerifyIgnore: []string{"last_updated"},
			},*/
			// Update and Read testing
			{
				Config: providerConfig + localConfig + `
				resource "` + resourceType + `" "` + resourceName + `" {
					    provider         = xmft.st1
						name             = "` + name + `"
						account 		 = xmft_st_account.account1.name
						type             = "x509"
						usage            = "login"
						#overwrite        = true
						content          = "` + strings.ReplaceAll(cert2, "\n", "\\n") + `"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					checkResources([]byte(`
					resource "`+resourceType+`" "`+resourceName+`" {
						#provider       = xmft.st1
						name             = "`+name+`"
						#account 		 = xmft_st_account.account1.name
						type             = "x509"
						usage            = "login"
						key_size		  = "2048"
						#overwrite        = true
					}
					`)),
					resource.TestCheckResourceAttrSet(resourceType+"."+resourceName, "last_updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
