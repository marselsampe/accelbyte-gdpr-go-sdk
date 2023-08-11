/*
 * Copyright (c) 2023 AccelByte Inc. All Rights Reserved
 * This program is made available under the terms of the MIT License.
 */

package object

type DataGenerationHandler func(namespace, userID string) (*DataGenerationResult, error)
type DataDeletionHandler func(namespace, userID string) error

// DataGenerationResult contains result from DataGenerationHandler implementation
type DataGenerationResult struct {

	// The DataGenerationHandler implementation should put the data result into this "Data" property.
	// This "Data" property was in form of Map<ModuleId, byte[]>, this allows the DataGenerationHandler
	// implementation to categorized multiple data based on the modules they have.
	//
	// Example:
	//  Map<ModuleId, byte[]> = {
	//    entitlement: entitlementData,
	//    wallet: walletData,
	//    transaction: transactionData
	//  }
	Data map[string][]byte
}
