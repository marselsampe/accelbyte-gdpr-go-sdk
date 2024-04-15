# AccelByte GDPR Go SDK

GDPR SDK for integrating Go services with AGS (AccelByte Gaming Services) GDPR service.

This GDPR SDK could be used by participant services to integrate into AGS GDPR workflow.
There are 4 GDPR workflow that this GDPR SDK supported:
1. Right to data portability
2. Right to erasure (right to be forgotten)
3. Right to restrict processing
4. Right to erase 3rd party account information

The participant services will hook their _**concrete go function**_ of 3 functionalities above into this GDPR SDK.

Under the hood, this GDPR SDK was using gRPC protocol for communication with AGS GDPR service:

```
+---------------+              +-----------------+
|     AGS       |     gRPC     | Your Go Service |
| GDPR Service -+------------->|  (gRPC Server)  |
| (gRPC Client) |              |                 |
+---------------+              +-----------------+
```

## Usage

### Install GDPR SDK

```
go get -u github.com/AccelByte/accelbyte-gdpr-go-sdk
```

### Initialize GDPR SDK

Import GDPR SDK references:
```go
import (
    "github.com/AccelByte/accelbyte-gdpr-go-sdk"
    "github.com/AccelByte/accelbyte-gdpr-go-sdk/pkg/object"
)
```
Create a new GDPR SDK gRPC server:
```go
yourGrpcServer := grpc.NewServer()
	
// initialize GDPR SDK gRPC
gdprSDK := gdprsdk.NewGdprGrpc()
gdprSDK.RegisterGRPC(yourGrpcServer)

lis, err := net.Listen("tcp", ":8081")
if err = yourGrpcServer.Serve(lis); err != nil {
    fmt.Errorf("%v", err)
}
```

### Hook your GDPR functionalities into GDPR SDK

```go
// register data generation handler (workflow: Right to data portability)
gdprSDK.SetDataGenerationHandler(func(namespace, userID string, isPublisherNamespace bool) (*object.DataGenerationResult, error) {
    logrus.Info("collecting user data...")
	
    // your implementation here...
	
    // example result
    return &object.DataGenerationResult{
        Data: map[string][]byte{
            "module1": []byte("{\"data\":\"lorem ipsum dolor sit amet\"}"),
            "module2": []byte("[\"lorem\",\"ipsum\",\"dolor\",\"sit\",\"amet\"]"),
        },
    }, nil
})

// register data deletion handler (workflow: Right to erasure) 
gdprSDK.SetDataDeletionHandler(func(namespace, userID string, isPublisherNamespace bool) error {
    logrus.Info("deleting user data...")
	
    // your implementation here...
	
    return nil
})

// register data restriction handler (workflow: Right to restrict processing) 
gdprSDK.SetDataRestrictionHandler(func(namespace, userID string, restrict, isPublisherNamespace bool) error {
    logrus.Info("restrict processing user data...")

    // your implementation here...

    return nil
})

// register platform account closure handler (workflow: Right to erase 3rd party account information) 
gdprSDK.SetPlatformAccountClosureHandler(func(platform, platformUserID string, accounts []*registered_v1.AccountInfo) error {
logrus.Info("deleting platform account info...")
// your implementation here...

return nil
})
```