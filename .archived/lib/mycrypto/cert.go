package mycrypto

import (
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"errors"
	"math/big"
	random "math/rand"
	"time"
)

// https://blog.csdn.net/u010846177/article/details/54357239

// GenSignselfCertificate 生成自签名证书
// func GenSignselfCertificate(req *x509.CertificateRequest, publickey, privKey interface{}, maxPath int, days time.Duration) ([]byte, error) {
func GenSignselfCertificate(publickey, privKey interface{}, maxPath int, days time.Duration, isECC bool) ([]byte, error) {
	req := &x509.CertificateRequest{
		Subject: pkix.Name{
			Country:            []string{"CN"},
			Organization:       []string{"ONETHING"},
			OrganizationalUnit: []string{"IP"},
			Locality:           []string{"SZ"},
			Province:           []string{"GD"},
			StreetAddress:      []string{"SA"},
			PostalCode:         []string{"PC"},
			CommonName:         "TCKMS",
		},
	}
	template := &x509.Certificate{
		SerialNumber:          big.NewInt(random.Int63n(time.Now().Unix())),
		Subject:               req.Subject,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(days * 24 * time.Hour),
		BasicConstraintsValid: true,
		IsCA:                  true,
		SignatureAlgorithm:    x509.SHA256WithRSA, // 签名算法选择SHA256WithRSA
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign | x509.KeyUsageDataEncipherment,
		SubjectKeyId:          []byte{1, 2, 3},
	}
	if maxPath > 0 { //如果长度超过0则设置了 最大的路径长度
		template.MaxPathLen = maxPath
	}
	if isECC {
		template.SignatureAlgorithm = x509.ECDSAWithSHA256
	}

	cert, err := x509.CreateCertificate(rand.Reader, template, template, publickey, privKey)
	if err != nil {
		return nil, errors.New("SignselfCert fail")
	}
	//err = util.EncodePemFile(fileName, "CERTIFICATE", cert)
	//if err != nil {
	//	return err
	//}
	return cert, nil
}

// GenCertificate 生成非自签名证书
// func GenCertificate(req *x509.CertificateRequest, parentCert *x509.Certificate, pubKey, parentPrivKey interface{}, isCA bool, days time.Duration) ([]byte, error) {
func GenCertificate(parentCert *x509.Certificate, pubKey, parentPrivKey interface{}, isCA bool, days time.Duration) ([]byte, error) {
	req := &x509.CertificateRequest{
		Subject: pkix.Name{
			Country:            []string{"CN"},
			Organization:       []string{"ONETHING"},
			OrganizationalUnit: []string{"IP"},
			Locality:           []string{"SZ"},
			Province:           []string{"GD"},
			StreetAddress:      []string{"SA"},
			PostalCode:         []string{"PC"},
			CommonName:         "TCKMS",
		},
	}
	template := &x509.Certificate{
		SerialNumber: big.NewInt(random.Int63n(time.Now().Unix())),
		Subject:      req.Subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(days * 24 * time.Hour),
		// ExtKeyUsage: []x509.ExtKeyUsage{ //额外的使用
		//  x509.ExtKeyUsageClientAuth,
		//  x509.ExtKeyUsageServerAuth,
		// },
		//
		SignatureAlgorithm: x509.SHA256WithRSA,
	}
	if isCA {
		template.BasicConstraintsValid = true
		template.IsCA = true
	}

	cert, err := x509.CreateCertificate(rand.Reader, template, parentCert, pubKey, parentPrivKey)
	if err != nil {
		return nil, errors.New("Certificate fail")
	}
	//err = util.EncodePemFile(fileName, "CERTIFICATE", cert)
	//if err != nil {
	//	return err
	//}
	return cert, nil
}

// ==========================================================

// EncodeCsr 生成证书请求
func EncodeCsr(country, organization, organizationlUnit, locality, province, streetAddress, postallCode []string, commonName string, priv interface{}) ([]byte, error) {
	req := &x509.CertificateRequest{
		Subject: pkix.Name{
			Country:            country,
			Organization:       organization,
			OrganizationalUnit: organizationlUnit,
			Locality:           locality,
			Province:           province,
			StreetAddress:      streetAddress,
			PostalCode:         postallCode,
			CommonName:         commonName,
		},
	}

	data, err := x509.CreateCertificateRequest(rand.Reader, req, priv)
	if err != nil {
		return nil, err
	}
	//err = util.EncodePemFile(fileName, "CERTIFICATE REQUEST", data)
	return data, nil
}

//DecodeCsr 解析CSRpem文件
func DecodeCsr(data []byte) (*x509.CertificateRequest, error) {
	//data, err := util.DecodePemFile(fileName)
	//if err != nil {
	//	return nil, err
	//}

	req, err := x509.ParseCertificateRequest(data)
	return req, err
}
