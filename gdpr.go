/*
 * Copyright (c) 2023 AccelByte Inc. All Rights Reserved
 * This program is made available under the terms of the MIT License.
 */

package gdprsdk

import (
	gdprGrpc "github.com/AccelByte/accelbyte-gdpr-go-sdk/pkg/grpc"
	"github.com/AccelByte/accelbyte-gdpr-go-sdk/pkg/object"
	pb "github.com/AccelByte/accelbyte-gdpr-go-sdk/pkg/pb"
	"google.golang.org/grpc"
)

func NewGdprGrpc() *GdprGrpc {
	return &GdprGrpc{
		gdprServiceServer: gdprGrpc.NewGDPRServiceServer(),
	}
}

type GdprGrpc struct {
	gdprServiceServer *gdprGrpc.GDPRServiceServer
}

func (sdk GdprGrpc) RegisterGRPC(server *grpc.Server) {
	pb.RegisterGDPRServer(server, sdk.gdprServiceServer)
}

func (sdk GdprGrpc) SetDataGenerationHandler(handler object.DataGenerationHandler) {
	sdk.gdprServiceServer.DataGenerationHandler = handler
}

func (sdk GdprGrpc) SetDataDeletionHandler(handler object.DataDeletionHandler) {
	sdk.gdprServiceServer.DataDeletionHandler = handler
}

func (sdk GdprGrpc) SetDataRestrictionHandler(handler object.DataRestrictionHandler) {
	sdk.gdprServiceServer.DataRestrictionHandler = handler
}

func (sdk GdprGrpc) SetPlatformAccountClosureHandler(handler object.PlatformAccountClosureHandler) {
	sdk.gdprServiceServer.PlatformAccountClosureHandler = handler
}
