# ipfs-go-client
This is a simple client for IPFS - add and receive files to/from ipfs


# Usage 

To send a file to a recipient (an email will be sent to the recipient with the file hash) - 

```
go run sender.go "email" "filepath"
```

The recipient can download the file by using (hash will be in the email) - 

```
go run receiver.go "hash" "filename-to-save"
```
