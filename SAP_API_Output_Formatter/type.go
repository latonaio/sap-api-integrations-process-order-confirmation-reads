package sap_api_output_formatter

type ProductionOrderConfirmation struct {
	ConnectionKey      string `json:"connection_key"`
	Result             bool   `json:"result"`
	RedisKey           string `json:"redis_key"`
	Filepath           string `json:"filepath"`
	APISchema          string `json:"api_schema"`
	ManufacturingOrder string `json:"production_order"`
	Deleted            bool   `json:"deleted"`
}

type Confirmation struct {
			ConfirmationGroup              string      `json:"ConfirmationGroup"`
			ConfirmationCount              string      `json:"ConfirmationCount"`
			OrderID                        string      `json:"OrderID"`
			OrderOperation                 string      `json:"OrderOperation"`
			OrderSuboperation              string      `json:"OrderSuboperation"`
			OrderType                      string      `json:"OrderType"`
			OrderOperationInternalID       string      `json:"OrderOperationInternalID"`
			ConfirmationText               string      `json:"ConfirmationText"`
			Language                       string      `json:"Language"`
			Material                       string      `json:"Material"`
			OrderPlannedTotalQty           string      `json:"OrderPlannedTotalQty"`
			ProductionUnit                 string      `json:"ProductionUnit"`
			FinalConfirmationType          string      `json:"FinalConfirmationType"`
			IsFinalConfirmation            bool        `json:"IsFinalConfirmation"`
			OpenReservationsIsCleared      bool        `json:"OpenReservationsIsCleared"`
			IsReversed                     bool        `json:"IsReversed"`
			IsReversal                     bool        `json:"IsReversal"`
			APIConfHasNoGoodsMovements     bool        `json:"APIConfHasNoGoodsMovements"`
			OrderConfirmationRecordType    string      `json:"OrderConfirmationRecordType"`
			ConfirmationEntryDate          string      `json:"ConfirmationEntryDate"`
			ConfirmationEntryTime          string      `json:"ConfirmationEntryTime"`
			EnteredByUser                  string      `json:"EnteredByUser"`
			LastChangeDate                 string      `json:"LastChangeDate"`
			LastChangedByUser              string      `json:"LastChangedByUser"`
			ConfirmationExternalEntryDate  string      `json:"ConfirmationExternalEntryDate"`
			ConfirmationExternalEntryTime  string      `json:"ConfirmationExternalEntryTime"`
			EnteredByExternalUser          string      `json:"EnteredByExternalUser"`
			ExternalSystemConfirmation     string      `json:"ExternalSystemConfirmation"`
			Plant                          string      `json:"Plant"`
			WorkCenterTypeCode             string      `json:"WorkCenterTypeCode"`
			WorkCenter                     string      `json:"WorkCenter"`
			Personnel                      string      `json:"Personnel"`
			TimeRecording                  string      `json:"TimeRecording"`
			EmployeeWageType               string      `json:"EmployeeWageType"`
			EmployeeWageGroup              string      `json:"EmployeeWageGroup"`
			BreakDurationUnit              string      `json:"BreakDurationUnit"`
			BreakDurationUnitISOCode       string      `json:"BreakDurationUnitISOCode"`
			BreakDurationUnitSAPCode       string      `json:"BreakDurationUnitSAPCode"`
			ConfirmedBreakDuration         string      `json:"ConfirmedBreakDuration"`
			EmployeeSuitability            string      `json:"EmployeeSuitability"`
			NumberOfEmployees              string      `json:"NumberOfEmployees"`
			PostingDate                    string      `json:"PostingDate"`
			ConfirmedExecutionStartDate    string      `json:"ConfirmedExecutionStartDate"`
			ConfirmedExecutionStartTime    string      `json:"ConfirmedExecutionStartTime"`
			ConfirmedSetupEndDate          string      `json:"ConfirmedSetupEndDate"`
			ConfirmedSetupEndTime          string      `json:"ConfirmedSetupEndTime"`
			ConfirmedProcessingStartDate   string      `json:"ConfirmedProcessingStartDate"`
			ConfirmedProcessingStartTime   string      `json:"ConfirmedProcessingStartTime"`
			ConfirmedProcessingEndDate     string      `json:"ConfirmedProcessingEndDate"`
			ConfirmedProcessingEndTime     string      `json:"ConfirmedProcessingEndTime"`
			ConfirmedTeardownStartDate     string      `json:"ConfirmedTeardownStartDate"`
			ConfirmedTeardownStartTime     string      `json:"ConfirmedTeardownStartTime"`
			ConfirmedExecutionEndDate      string      `json:"ConfirmedExecutionEndDate"`
			ConfirmedExecutionEndTime      string      `json:"ConfirmedExecutionEndTime"`
			ConfirmationUnit               string      `json:"ConfirmationUnit"`
			ConfirmationUnitISOCode        string      `json:"ConfirmationUnitISOCode"`
			ConfirmationUnitSAPCode        string      `json:"ConfirmationUnitSAPCode"`
			ConfirmationYieldQuantity      string      `json:"ConfirmationYieldQuantity"`
			ConfirmationScrapQuantity      string      `json:"ConfirmationScrapQuantity"`
			VarianceReasonCode             string      `json:"VarianceReasonCode"`
			OpWorkQuantityUnit1            string      `json:"OpWorkQuantityUnit1"`
			WorkQuantityUnit1ISOCode       string      `json:"WorkQuantityUnit1ISOCode"`
			WorkQuantityUnit1SAPCode       string      `json:"WorkQuantityUnit1SAPCode"`
			OpConfirmedWorkQuantity1       string      `json:"OpConfirmedWorkQuantity1"`
			NoFurtherOpWorkQuantity1IsExpd bool        `json:"NoFurtherOpWorkQuantity1IsExpd"`
			OpWorkQuantityUnit2            string      `json:"OpWorkQuantityUnit2"`
			WorkQuantityUnit2ISOCode       string      `json:"WorkQuantityUnit2ISOCode"`
			WorkQuantityUnit2SAPCode       string      `json:"WorkQuantityUnit2SAPCode"`
			OpConfirmedWorkQuantity2       string      `json:"OpConfirmedWorkQuantity2"`
			NoFurtherOpWorkQuantity2IsExpd bool        `json:"NoFurtherOpWorkQuantity2IsExpd"`
			OpWorkQuantityUnit3            string      `json:"OpWorkQuantityUnit3"`
			WorkQuantityUnit3ISOCode       string      `json:"WorkQuantityUnit3ISOCode"`
			WorkQuantityUnit3SAPCode       string      `json:"WorkQuantityUnit3SAPCode"`
			OpConfirmedWorkQuantity3       string      `json:"OpConfirmedWorkQuantity3"`
			NoFurtherOpWorkQuantity3IsExpd bool        `json:"NoFurtherOpWorkQuantity3IsExpd"`
			OpWorkQuantityUnit4            string      `json:"OpWorkQuantityUnit4"`
			WorkQuantityUnit4ISOCode       string      `json:"WorkQuantityUnit4ISOCode"`
			WorkQuantityUnit4SAPCode       string      `json:"WorkQuantityUnit4SAPCode"`
			OpConfirmedWorkQuantity4       string      `json:"OpConfirmedWorkQuantity4"`
			NoFurtherOpWorkQuantity4IsExpd bool        `json:"NoFurtherOpWorkQuantity4IsExpd"`
			OpWorkQuantityUnit5            string      `json:"OpWorkQuantityUnit5"`
			WorkQuantityUnit5ISOCode       string      `json:"WorkQuantityUnit5ISOCode"`
			WorkQuantityUnit5SAPCode       string      `json:"WorkQuantityUnit5SAPCode"`
			OpConfirmedWorkQuantity5       string      `json:"OpConfirmedWorkQuantity5"`
			NoFurtherOpWorkQuantity5IsExpd bool        `json:"NoFurtherOpWorkQuantity5IsExpd"`
			OpWorkQuantityUnit6            string      `json:"OpWorkQuantityUnit6"`
			WorkQuantityUnit6ISOCode       string      `json:"WorkQuantityUnit6ISOCode"`
			WorkQuantityUnit6SAPCode       string      `json:"WorkQuantityUnit6SAPCode"`
			OpConfirmedWorkQuantity6       string      `json:"OpConfirmedWorkQuantity6"`
			NoFurtherOpWorkQuantity6IsExpd bool        `json:"NoFurtherOpWorkQuantity6IsExpd"`
			BusinessProcessEntryUnit       string      `json:"BusinessProcessEntryUnit"`
			BusProcessEntrUnitISOCode      string      `json:"BusProcessEntrUnitISOCode"`
			BusProcessEntryUnitSAPCode     string      `json:"BusProcessEntryUnitSAPCode"`
			BusinessProcessConfirmedQty    string      `json:"BusinessProcessConfirmedQty"`
			NoFurtherBusinessProcQtyIsExpd bool        `json:"NoFurtherBusinessProcQtyIsExpd"`
}

type MaterialMovements struct {
			ConfirmationGroup          string      `json:"ConfirmationGroup"`
			ConfirmationCount          string      `json:"ConfirmationCount"`
			MaterialDocument           string      `json:"MaterialDocument"`
			MaterialDocumentItem       string      `json:"MaterialDocumentItem"`
			MaterialDocumentYear       string      `json:"MaterialDocumentYear"`
			OrderType                  string      `json:"OrderType"`
			OrderID                    string      `json:"OrderID"`
			OrderItem                  string      `json:"OrderItem"`
			ManufacturingOrderCategory string      `json:"ManufacturingOrderCategory"`
			Material                   string      `json:"Material"`
			Plant                      string      `json:"Plant"`
			Reservation                string      `json:"Reservation"`
			ReservationItem            string      `json:"ReservationItem"`
			StorageLocation            string      `json:"StorageLocation"`
			ProductionSupplyArea       string      `json:"ProductionSupplyArea"`
			Batch                      string      `json:"Batch"`
			InventoryValuationType     string      `json:"InventoryValuationType"`
			GoodsMovementType          string      `json:"GoodsMovementType"`
			GoodsMovementRefDocType    string      `json:"GoodsMovementRefDocType"`
			InventoryUsabilityCode     string      `json:"InventoryUsabilityCode"`
			InventorySpecialStockType  string      `json:"InventorySpecialStockType"`
			SalesOrder                 string      `json:"SalesOrder"`
			SalesOrderItem             string      `json:"SalesOrderItem"`
			WBSElementExternalID       string      `json:"WBSElementExternalID"`
			Supplier                   string      `json:"Supplier"`
			Customer                   string      `json:"Customer"`
			ReservationIsFinallyIssued bool        `json:"ReservationIsFinallyIssued"`
			IsCompletelyDelivered      bool        `json:"IsCompletelyDelivered"`
			ShelfLifeExpirationDate    string      `json:"ShelfLifeExpirationDate"`
			ManufactureDate            string      `json:"ManufactureDate"`
			StorageType                string      `json:"StorageType"`
			StorageBin                 string      `json:"StorageBin"`
			MaterialDocumentItemText   string      `json:"MaterialDocumentItemText"`
			EntryUnit                  string      `json:"EntryUnit"`
			EntryUnitISOCode           string      `json:"EntryUnitISOCode"`
			EntryUnitSAPCode           string      `json:"EntryUnitSAPCode"`
			QuantityInEntryUnit        string      `json:"QuantityInEntryUnit"`
}  
