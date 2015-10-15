package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/Songmu/prompter"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"os"
	"time"
)

func AWSConfig(key, secret, token, file, profile string) (conf *aws.Config) {
	var creds *credentials.Credentials
	if file != "" {
		creds = credentials.NewSharedCredentials(file, profile)
	}
	if key != "" && secret != "" {
		creds = credentials.NewStaticCredentials(key, secret, token)
	}

	conf = &aws.Config{
		Credentials: creds,
	}

	return conf
}

func ReadFile(fpath string) (body string, err error) {
	var f *os.File
	var b []byte

	f, err = os.Open(fpath)
	if err != nil {
		return body, err
	}

	b, err = ioutil.ReadAll(f)
	if err != nil {
		return body, err
	}

	return string(b), err
}

func CertificateBody(cb, cbPath string) (body string, err error) {
	if cbPath != "" {
		body, err = ReadFile(cbPath)
	} else {
		body = cb
	}

	return body, err
}

func PrivateKey(pk, pkPath string) (body string, err error) {
	if pkPath != "" {
		body, err = ReadFile(pkPath)
	} else {
		body = pk
	}

	return body, err
}

func CertificateChain(cc, ccPath string) (body string, err error) {
	if ccPath != "" {
		body, err = ReadFile(ccPath)
	} else {
		body = cc
	}

	return body, err
}

func UploadServerCertificateInput(cBody, cChain, path, pKey, sCert string) (uploadSCInput *iam.UploadServerCertificateInput) {
	uploadSCInput = &iam.UploadServerCertificateInput{
		CertificateBody:       aws.String(cBody),
		CertificateChain:      aws.String(cChain),
		Path:                  aws.String(path),
		PrivateKey:            aws.String(pKey),
		ServerCertificateName: aws.String(sCert),
	}

	return uploadSCInput
}

func UpdateServerCertificateInput(newPath, newSCert, sCert string) (updateSCInput *iam.UpdateServerCertificateInput) {
	updateSCInput = &iam.UpdateServerCertificateInput{
		NewPath:                  aws.String(newPath),
		NewServerCertificateName: aws.String(newSCert),
		ServerCertificateName:    aws.String(sCert),
	}

	return updateSCInput
}

func ListServerCertificatesInput(marker string, mexItems int, pathPrefix string) (listSCInput *iam.ListServerCertificatesInput) {
	listSCInput = &iam.ListServerCertificatesInput{}

	if marker != "" {
		listSCInput.Marker = aws.String(marker)
	}

	maxItems64 := int64(*maxItems)

	if maxItems64 > 0 {
		listSCInput.MaxItems = aws.Int64(maxItems64)
	}

	if pathPrefix != "" {
		listSCInput.PathPrefix = aws.String(pathPrefix)
	}

	return listSCInput
}

func RetrieveNames(list *iam.ListServerCertificatesOutput) (names []string) {
	for _, metadata := range list.ServerCertificateMetadataList {
		names = append(names, *metadata.ServerCertificateName)
	}
	return names
}

func Timezone(offset int) (location *time.Location) {
	location = time.FixedZone("", offset*60*60)
	return location
}

func FormatedTime(t time.Time, zone *time.Location) (fTime string) {
	return t.In(zone).Format("2006-01-02 15:04:05 -0700")
}

func OutputList(list *iam.ListServerCertificatesOutput, offset int) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Arn", "Path", "ServerCertificateId", "ServerCertificateName", "UploadDate", "Expiration"})
	zone := Timezone(offset)
	for _, metadata := range list.ServerCertificateMetadataList {
		uploadDate := FormatedTime(*metadata.UploadDate, zone)
		expiration := FormatedTime(*metadata.Expiration, zone)
		data := []string{
			*metadata.Arn, *metadata.Path, *metadata.ServerCertificateId, *metadata.ServerCertificateName,
			fmt.Sprint(uploadDate), fmt.Sprint(expiration),
		}
		table.Append(data)
	}
	table.Render()
}

var (
	iamSrvCert            = kingpin.New("iam-server-cert", "(List|Upload|Update|Delete) the server certificates for the AWS account")
	serverCertificateName = iamSrvCert.Flag("server-certificate-name", "The name for the server certificate").String()
	certificateBody       = iamSrvCert.Flag("certificate-body", "The contents of the public key certificate in PEM-encoded format").String()
	privateKey            = iamSrvCert.Flag("private-key", "The contents of the private key in PEM-encoded format").String()
	certificateChain      = iamSrvCert.Flag("certificate-chain", "The contents of the certificate chain").String()
	certificateBodyPath   = iamSrvCert.Flag("certificate-body-path", "Path to public key certificate").String()
	privateKeyPath        = iamSrvCert.Flag("private-key-path", "Path to private key").String()
	certificateChainPath  = iamSrvCert.Flag("certificate-chain-path", "Path to certificate chain").String()
	accessKey             = iamSrvCert.Flag("access-key", "AWS access key").String()
	accessSecret          = iamSrvCert.Flag("access-secret", "AWS secret key").String()
	token                 = iamSrvCert.Flag("token", "Session token").String()
	credentialsPath       = iamSrvCert.Flag("credentials", "Credential file").String()
	profile               = iamSrvCert.Flag("profile", "Use a specific profile from your credential file").Default("default").String()

	iamSrvCertList = iamSrvCert.Command("list", "List SSL certificates")
	marker         = iamSrvCertList.Flag("marker", "Paginating results and only after you receive a response indicating that the results are truncated").String()
	maxItems       = iamSrvCertList.Flag("max-items", "The total number of items to return").Int()
	pathPrefix     = iamSrvCertList.Flag("path-prefix", "The path prefix for filtering the results").Default("/").String()
	timeOffset     = iamSrvCertList.Flag("time-offset", "Time offsets").Default("0").Int()

	iamSrvCertDelete = iamSrvCert.Command("delete", "Delete SSL certificate")

	iamSrvCertUpload = iamSrvCert.Command("upload", "Upload SSL certificate")
	uploadPath       = iamSrvCertUpload.Flag("path", "The path for the server certificate").Default("/").String()

	iamSrvCertUpdate         = iamSrvCert.Command("update", "Update SSL certificate")
	newPath                  = iamSrvCertUpdate.Flag("new-path", "The new path for the server certificate").String()
	newServerCertificateName = iamSrvCertUpdate.Flag("new-server-certificate-name", "The new name for the server certificate").String()

	iamSvc *iam.IAM
)

func main() {
	iamSrvCert.Version("0.0.1")
	subCmd, err := iamSrvCert.Parse(os.Args[1:])

	iamSvc = iam.New(AWSConfig(*accessKey, *accessSecret, *token, *credentialsPath, *profile))
	log := logrus.New()

	switch subCmd {
	case "upload":
		cBody, err := CertificateBody(*certificateBody, *certificateBodyPath)
		if err != nil {
			log.Fatal(err)
		}

		pKey, err := PrivateKey(*privateKey, *privateKeyPath)
		if err != nil {
			log.Fatal(err)
		}

		cChain, err := CertificateChain(*certificateChain, *certificateChainPath)
		if err != nil {
			log.Fatal(err)
		}

		uploadSCInput := UploadServerCertificateInput(cBody, cChain, *uploadPath, pKey, *serverCertificateName)

		_, err = iamSvc.UploadServerCertificate(uploadSCInput)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Uploaded", *serverCertificateName)
	case "update":
		updateSCInput := UpdateServerCertificateInput(*newPath, *newServerCertificateName, *serverCertificateName)

		_, err = iamSvc.UpdateServerCertificate(updateSCInput)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Updated from", *serverCertificateName, "to", *newServerCertificateName)
	case "list":
		var lscOut *iam.ListServerCertificatesOutput
		lscOut, err = iamSvc.ListServerCertificates(ListServerCertificatesInput(*marker, *maxItems, *pathPrefix))
		if err != nil {
			log.Fatal(err)
		}

		OutputList(lscOut, *timeOffset)
	case "delete":
		var lscOut *iam.ListServerCertificatesOutput
		lscOut, err = iamSvc.ListServerCertificates(ListServerCertificatesInput(*marker, *maxItems, *pathPrefix))
		if err != nil {
			log.Fatal(err)
		}

		names := RetrieveNames(lscOut)

		var cert string
		if *serverCertificateName != "" {
			cert = *serverCertificateName
		} else {
			cert = (&prompter.Prompter{
				Choices:    names,
				UseDefault: false,
				Message:    "Choose the server certificate you want to delete.",
				IgnoreCase: true,
			}).Prompt()
		}

		_, err = iamSvc.DeleteServerCertificate(&iam.DeleteServerCertificateInput{ServerCertificateName: aws.String(cert)})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Deleted", cert)
	}
}
