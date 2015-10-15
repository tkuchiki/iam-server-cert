# iam-server-cert
(List|Upload|Update|Delete) the server certificates for the AWS account

# Installation

Download from https://github.com/tkuchiki/alp/releases

# Usage

```
$ ./iam-server-cert --help
usage: iam-server-cert [<flags>] <command> [<args> ...]

(List|Upload|Update|Delete) the server certificates for the AWS account

Flags:
  --help                         Show context-sensitive help (also try --help-long and --help-man).
  --server-certificate-name=SERVER-CERTIFICATE-NAME
                                 The name for the server certificate
  --certificate-body=CERTIFICATE-BODY
                                 The contents of the public key certificate in PEM-encoded format
  --private-key=PRIVATE-KEY      The contents of the private key in PEM-encoded format
  --certificate-chain=CERTIFICATE-CHAIN
                                 The contents of the certificate chain
  --certificate-body-path=CERTIFICATE-BODY-PATH
                                 Path to public key certificate
  --private-key-path=PRIVATE-KEY-PATH
                                 Path to private key
  --certificate-chain-path=CERTIFICATE-CHAIN-PATH
                                 Path to certificate chain
  --access-key=ACCESS-KEY        AWS access key
  --access-secret=ACCESS-SECRET  AWS secret key
  --credentials=CREDENTIALS      Credential file
  --profile="default"            Use a specific profile from your credential file
  --timezone="UTC"               Timezone
  --version                      Show application version.

Commands:
  help [<command>...]
    Show help.

  list [<flags>]
    List SSL certificates

  delete
    Delete SSL certificate

  upload [<flags>]
    Upload SSL certificate

  update [<flags>]
    Update SSL certificate
```

## Sub commands

### List

```
$ ./iam-server-cert list --help
usage: iam-server-cert list [<flags>]

List SSL certificates

Flags:
  --help                         Show context-sensitive help (also try --help-long and --help-man).
  --server-certificate-name=SERVER-CERTIFICATE-NAME
                                 The name for the server certificate
  --certificate-body=CERTIFICATE-BODY
                                 The contents of the public key certificate in PEM-encoded format
  --private-key=PRIVATE-KEY      The contents of the private key in PEM-encoded format
  --certificate-chain=CERTIFICATE-CHAIN
                                 The contents of the certificate chain
  --certificate-body-path=CERTIFICATE-BODY-PATH
                                 Path to public key certificate
  --private-key-path=PRIVATE-KEY-PATH
                                 Path to private key
  --certificate-chain-path=CERTIFICATE-CHAIN-PATH
                                 Path to certificate chain
  --access-key=ACCESS-KEY        AWS access key
  --access-secret=ACCESS-SECRET  AWS secret key
  --token=TOKEN                  Session token
  --credentials=CREDENTIALS      Credential file
  --profile="default"            Use a specific profile from your credential file
  --time-offset=0                Time offsets
  --version                      Show application version.
  --marker=MARKER                Paginating results and only after you receive a response indicating that the results are truncated
  --max-items=MAX-ITEMS          The total number of items to return
  --path-prefix="/"              The path prefix for filtering the results

Subcommands:
```

#### Example

```
$ ./iam-server-cert list
+------------------------------------------------------------------+-----------+-----------------------+-----------------------+---------------------------+---------------------------+
|                               ARN                                |   PATH    |  SERVERCERTIFICATEID  | SERVERCERTIFICATENAME |        UPLOADDATE         |        EXPIRATION         |
+------------------------------------------------------------------+-----------+-----------------------+-----------------------+---------------------------+---------------------------+
| arn:aws:iam::xxxxxxxxxxxx:server-certificate/cert1               | /         | XXXXXXXXXXXXXXXXXXXXX | cert1                 | 2015-11-09 08:03:39 +0000 | 2016-10-15 02:18:29 +0000 |
| arn:aws:iam::xxxxxxxxxxxx:server-certificate/foo/bar/cert2       | /foo/bar/ | XXXXXXXXXXXXXXXXXXXXX | cert2                 | 2015-11-15 03:05:50 +0000 | 2016-10-15 02:18:29 +0000 |
+------------------------------------------------------------------+-----------+-----------------------+-----------------------+---------------------------+---------------------------+

$ ./iam-server-cert list --path-prefix /foo/bar/
+---------------------------------------------------------------+-----------+-----------------------+-----------------------+---------------------------+---------------------------+
|                              ARN                              |   PATH    |  SERVERCERTIFICATEID  | SERVERCERTIFICATENAME |        UPLOADDATE         |        EXPIRATION         |
+---------------------------------------------------------------+-----------+-----------------------+-----------------------+---------------------------+---------------------------+
| arn:aws:iam::xxxxxxxxxxxx:server-certificate/foo/bar/cert2    | /foo/bar/ | XXXXXXXXXXXXXXXXXXXXX | cert2                 | 2015-11-15 03:05:50 +0000 | 2016-10-15 02:18:29 +0000 |
+---------------------------------------------------------------+-----------+-----------------------+-----------------------+---------------------------+---------------------------+

$ ./iam-server-cert list --time-offset 9
+------------------------------------------------------------------+-----------+-----------------------+-----------------------+---------------------------+---------------------------+
|                               ARN                                |   PATH    |  SERVERCERTIFICATEID  | SERVERCERTIFICATENAME |        UPLOADDATE         |        EXPIRATION         |
+------------------------------------------------------------------+-----------+-----------------------+-----------------------+---------------------------+---------------------------+
| arn:aws:iam::xxxxxxxxxxxx:server-certificate/cert1               | /         | XXXXXXXXXXXXXXXXXXXXX | cert1                 | 2015-11-09 17:03:39 +0900 | 2016-10-15 11:18:29 +0900 |
| arn:aws:iam::xxxxxxxxxxxx:server-certificate/foo/bar/cert2       | /foo/bar/ | XXXXXXXXXXXXXXXXXXXXX | cert2                 | 2015-11-15 12:05:50 +0900 | 2016-10-15 11:18:29 +0900 |
+------------------------------------------------------------------+-----------+-----------------------+-----------------------+---------------------------+---------------------------+

$ ./iam-server-cert list --max-items 1
+------------------------------------------------------------------+------+-----------------------+-----------------------+---------------------------+---------------------------+
|                               ARN                                | PATH |  SERVERCERTIFICATEID  | SERVERCERTIFICATENAME |        UPLOADDATE         |        EXPIRATION         |
+------------------------------------------------------------------+------+-----------------------+-----------------------+---------------------------+---------------------------+
| arn:aws:iam::xxxxxxxxxxxx:server-certificate/cert1               | /    | XXXXXXXXXXXXXXXXXXXXX | cert1                 | 2015-11-09 08:03:39 +0000 | 2016-10-15 02:18:29 +0000 |
+------------------------------------------------------------------+------+-----------------------+-----------------------+---------------------------+---------------------------+

```

### Upload

```
$ ./iam-server-cert upload --help
usage: iam-server-cert upload [<flags>]

Upload SSL certificate

Flags:
  --help                         Show context-sensitive help (also try --help-long and --help-man).
  --server-certificate-name=SERVER-CERTIFICATE-NAME
                                 The name for the server certificate
  --certificate-body=CERTIFICATE-BODY
                                 The contents of the public key certificate in PEM-encoded format
  --private-key=PRIVATE-KEY      The contents of the private key in PEM-encoded format
  --certificate-chain=CERTIFICATE-CHAIN
                                 The contents of the certificate chain
  --certificate-body-path=CERTIFICATE-BODY-PATH
                                 Path to public key certificate
  --private-key-path=PRIVATE-KEY-PATH
                                 Path to private key
  --certificate-chain-path=CERTIFICATE-CHAIN-PATH
                                 Path to certificate chain
  --access-key=ACCESS-KEY        AWS access key
  --access-secret=ACCESS-SECRET  AWS secret key
  --token=TOKEN                  Session token
  --credentials=CREDENTIALS      Credential file
  --profile="default"            Use a specific profile from your credential file
  --version                      Show application version.
  --path="/"                     The path for the server certificate

Subcommands:
```

#### Example

```
$ ./iam-server-cert list
+------------------------------------------------------------------+-----------+-----------------------+-----------------------+---------------------------+---------------------------+
|                               ARN                                |   PATH    |  SERVERCERTIFICATEID  | SERVERCERTIFICATENAME |        UPLOADDATE         |        EXPIRATION         |
+------------------------------------------------------------------+-----------+-----------------------+-----------------------+---------------------------+---------------------------+
| arn:aws:iam::xxxxxxxxxxxx:server-certificate/cert1               | /         | XXXXXXXXXXXXXXXXXXXXX | cert1                 | 2015-11-09 08:03:39 +0000 | 2016-10-15 02:18:29 +0000 |
| arn:aws:iam::xxxxxxxxxxxx:server-certificate/foo/bar/cert2       | /foo/bar/ | XXXXXXXXXXXXXXXXXXXXX | cert2                 | 2015-11-15 03:05:50 +0000 | 2016-10-15 02:18:29 +0000 |
+------------------------------------------------------------------+-----------+-----------------------+-----------------------+---------------------------+---------------------------+

$ ./iam-server-cert upload --server-certificate-name cert3 --certificate-body-path /path/to/cert-body --private-key-path /path/to/pkey --certificate-chain-path /path/to/cert-chain --path "/hoge/"
Uploaded cert3

$ ./iam-server-cert list
+------------------------------------------------------------------+-----------+-----------------------+-----------------------+---------------------------+---------------------------+
|                               ARN                                |   PATH    |  SERVERCERTIFICATEID  | SERVERCERTIFICATENAME |        UPLOADDATE         |        EXPIRATION         |
+------------------------------------------------------------------+-----------+-----------------------+-----------------------+---------------------------+---------------------------+
| arn:aws:iam::xxxxxxxxxxxx:server-certificate/cert1               | /         | XXXXXXXXXXXXXXXXXXXXX | cert1                 | 2015-11-09 08:03:39 +0000 | 2016-10-15 02:18:29 +0000 |
| arn:aws:iam::xxxxxxxxxxxx:server-certificate/foo/bar/cert2       | /foo/bar/ | XXXXXXXXXXXXXXXXXXXXX | cert2                 | 2015-11-15 03:05:50 +0000 | 2016-10-15 02:18:29 +0000 |
| arn:aws:iam::xxxxxxxxxxxx:server-certificate/hoge/cert3          | /hoge/    | XXXXXXXXXXXXXXXXXXXXX | cert3                 | 2015-12-21 11:22:33 +0000 | 2016-11-21 22:11:00 +0000 |
+------------------------------------------------------------------+-----------+-----------------------+-----------------------+---------------------------+---------------------------+
```

### Update

```
$ ./iam-server-cert update --help
usage: iam-server-cert update [<flags>]

Update SSL certificate

Flags:
  --help                         Show context-sensitive help (also try --help-long and --help-man).
  --server-certificate-name=SERVER-CERTIFICATE-NAME
                                 The name for the server certificate
  --certificate-body=CERTIFICATE-BODY
                                 The contents of the public key certificate in PEM-encoded format
  --private-key=PRIVATE-KEY      The contents of the private key in PEM-encoded format
  --certificate-chain=CERTIFICATE-CHAIN
                                 The contents of the certificate chain
  --certificate-body-path=CERTIFICATE-BODY-PATH
                                 Path to public key certificate
  --private-key-path=PRIVATE-KEY-PATH
                                 Path to private key
  --certificate-chain-path=CERTIFICATE-CHAIN-PATH
                                 Path to certificate chain
  --access-key=ACCESS-KEY        AWS access key
  --access-secret=ACCESS-SECRET  AWS secret key
  --token=TOKEN                  Session token
  --credentials=CREDENTIALS      Credential file
  --profile="default"            Use a specific profile from your credential file
  --version                      Show application version.
  --new-path=NEW-PATH            The new path for the server certificate
  --new-server-certificate-name=NEW-SERVER-CERTIFICATE-NAME
                                 The new name for the server certificate

Subcommands:
```

#### Example

```
$ ./iam-server-cert list
+------------------------------------------------------------------+-----------+-----------------------+-----------------------+---------------------------+---------------------------+
|                               ARN                                |   PATH    |  SERVERCERTIFICATEID  | SERVERCERTIFICATENAME |        UPLOADDATE         |        EXPIRATION         |
+------------------------------------------------------------------+-----------+-----------------------+-----------------------+---------------------------+---------------------------+
| arn:aws:iam::xxxxxxxxxxxx:server-certificate/cert1               | /         | XXXXXXXXXXXXXXXXXXXXX | cert1                 | 2015-11-09 08:03:39 +0000 | 2016-10-15 02:18:29 +0000 |
| arn:aws:iam::xxxxxxxxxxxx:server-certificate/foo/bar/cert2       | /foo/bar/ | XXXXXXXXXXXXXXXXXXXXX | cert2                 | 2015-11-15 03:05:50 +0000 | 2016-10-15 02:18:29 +0000 |
+------------------------------------------------------------------+-----------+-----------------------+-----------------------+---------------------------+---------------------------+

$ ./iam-server-cert update --server-certificate-name cert1 --new-path /hoge/ --new-server-certificate-name cert3
Updated from test-ssl to cert3

$ ./iam-server-cert list
+------------------------------------------------------------------+-----------+-----------------------+-----------------------+------------------+--------+---------------------------+
|                           ARN                                    |  PATH     |  SERVERCERTIFICATEID  | SERVERCERTIFICATENAME |        UPLOADDATE         |        EXPIRATION         |
+------------------------------------------------------------------+-----------+-----------------------+-----------------------+---------------------------+---------------------------+
| arn:aws:iam::xxxxxxxxxxxx:server-certificate/hoge/cert3          | /hoge/    | XXXXXXXXXXXXXXXXXXXXX | cert3                 | 2015-11-15 03:05:50 +0000 | 2016-10-15 02:18:29 +0000 |
| arn:aws:iam::xxxxxxxxxxxx:server-certificate/foo/bar/cert2       | /foo/bar/ | XXXXXXXXXXXXXXXXXXXXX | cert2                 | 2015-11-15 03:05:50 +0000 | 2016-10-15 02:18:29 +0000 |
+---------------------------------------------------------+--------+-----------+-----------------------+-----------------------+---------------------------+---------------------------+
```

### Delete

```
$ ./iam-server-cert delete --help
usage: iam-server-cert delete

Delete SSL certificate

Flags:
  --help                         Show context-sensitive help (also try --help-long and --help-man).
  --server-certificate-name=SERVER-CERTIFICATE-NAME
                                 The name for the server certificate
  --certificate-body=CERTIFICATE-BODY
                                 The contents of the public key certificate in PEM-encoded format
  --private-key=PRIVATE-KEY      The contents of the private key in PEM-encoded format
  --certificate-chain=CERTIFICATE-CHAIN
                                 The contents of the certificate chain
  --certificate-body-path=CERTIFICATE-BODY-PATH
                                 Path to public key certificate
  --private-key-path=PRIVATE-KEY-PATH
                                 Path to private key
  --certificate-chain-path=CERTIFICATE-CHAIN-PATH
                                 Path to certificate chain
  --access-key=ACCESS-KEY        AWS access key
  --access-secret=ACCESS-SECRET  AWS secret key
  --token=TOKEN                  Session token
  --credentials=CREDENTIALS      Credential file
  --profile="default"            Use a specific profile from your credential file
  --version                      Show application version.

Subcommands:
```

#### Example

```
$  ./iam-server-cert delete
Choose the server certificate you want to delete. (cert1/cert2): cert1[Enter]
Deleted cert1

$ ./iam-server-cert delete --server-certificate-name cert1
Deleted cert1
```
