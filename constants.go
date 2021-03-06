package smpp34

const (
	// SMPP Protocol Version
	SMPP_VERSION = 0x34

	// Max PDU size to minimize some attack vectors
	MAX_PDU_SIZE = 1024 * 4 // 4KB

	// Sequence number start/end
	SEQUENCE_NUM_START = 0x00000001
	SEQUENCE_NUM_END   = 0x7FFFFFFF
)

const (
	// ESME Error Constants
	ESME_ROK              CMDStatus = 0x00000000 // OK!
	ESME_RINVMSGLEN       CMDStatus = 0x00000001 // Message Length is invalid
	ESME_RINVCMDLEN       CMDStatus = 0x00000002 // Command Length is invalid
	ESME_RINVCMDID        CMDStatus = 0x00000003 // Invalid Command ID
	ESME_RINVBNDSTS       CMDStatus = 0x00000004 // Incorrect BIND Status for given command
	ESME_RALYBND          CMDStatus = 0x00000005 // ESME Already in Bound State
	ESME_RINVPRTFLG       CMDStatus = 0x00000006 // Invalid Priority Flag
	ESME_RINVREGDLVFLG    CMDStatus = 0x00000007 // Invalid Registered Delivery Flag
	ESME_RSYSERR          CMDStatus = 0x00000008 // System Error
	ESME_RINVSRCADR       CMDStatus = 0x0000000A // Invalid Source Address
	ESME_RINVDSTADR       CMDStatus = 0x0000000B // Invalid Dest Addr
	ESME_RINVMSGID        CMDStatus = 0x0000000C // Message ID is invalid
	ESME_RBINDFAIL        CMDStatus = 0x0000000D // Bind Failed
	ESME_RINVPASWD        CMDStatus = 0x0000000E // Invalid Password
	ESME_RINVSYSID        CMDStatus = 0x0000000F // Invalid System ID
	ESME_RCANCELFAIL      CMDStatus = 0x00000011 // Cancel SM Failed
	ESME_RREPLACEFAIL     CMDStatus = 0x00000013 // Replace SM Failed
	ESME_RMSGQFUL         CMDStatus = 0x00000014 // Message Queue Full
	ESME_RINVSERTYP       CMDStatus = 0x00000015 // Invalid Service Type
	ESME_RINVNUMDESTS     CMDStatus = 0x00000033 // Invalid number of destinations
	ESME_RINVDLNAME       CMDStatus = 0x00000034 // Invalid Distribution List name
	ESME_RINVDESTFLAG     CMDStatus = 0x00000040 // Destination flag is invalid
	ESME_RINVSUBREP       CMDStatus = 0x00000042 // Invalid 'submit with replace' request
	ESME_RINVESMCLASS     CMDStatus = 0x00000043 // Invalid esm_class field data
	ESME_RCNTSUBDL        CMDStatus = 0x00000044 // Cannot Submit to Distribution List
	ESME_RSUBMITFAIL      CMDStatus = 0x00000045 // submit_sm or submit_multi failed
	ESME_RINVSRCTON       CMDStatus = 0x00000048 // Invalid Source address TON
	ESME_RINVSRCNPI       CMDStatus = 0x00000049 // Invalid Source address NPI
	ESME_RINVDSTTON       CMDStatus = 0x00000050 // Invalid Destination address TON
	ESME_RINVDSTNPI       CMDStatus = 0x00000051 // Invalid Destination address NPI
	ESME_RINVSYSTYP       CMDStatus = 0x00000053 // Invalid system_type field
	ESME_RINVREPFLAG      CMDStatus = 0x00000054 // Invalid replace_if_present flag
	ESME_RINVNUMMSGS      CMDStatus = 0x00000055 // Invalid number of messages
	ESME_RTHROTTLED       CMDStatus = 0x00000058 // Throttling error
	ESME_RINVSCHED        CMDStatus = 0x00000061 // Invalid Scheduled Delivery Time
	ESME_RINVEXPIRY       CMDStatus = 0x00000062 // Invalid message validity period (Expiry time)
	ESME_RINVDFTMSGID     CMDStatus = 0x00000063 // Predefined Message Invalid or Not Found
	ESME_RX_T_APPN        CMDStatus = 0x00000064 // ESME Receiver Temporary App Error Code
	ESME_RX_P_APPN        CMDStatus = 0x00000065 // ESME Receiver Permanent App Error Code
	ESME_RX_R_APPN        CMDStatus = 0x00000066 // ESME Receiver Reject Message Error Code
	ESME_RQUERYFAIL       CMDStatus = 0x00000067 // Query_sm request failed
	ESME_RINVOPTPARSTREAM CMDStatus = 0x000000C0 // Error in the optional part of the PDU Body
	ESME_ROPTPARNOTALLWD  CMDStatus = 0x000000C1 // Optional Parameter not allowed
	ESME_RINVPARLEN       CMDStatus = 0x000000C2 // Invalid Parameter Length
	ESME_RMISSINGOPTPARAM CMDStatus = 0x000000C3 // Expected Optional Parameter missing
	ESME_RINVOPTPARAMVAL  CMDStatus = 0x000000C4 // Invalid Optional Parameter Value
	ESME_RDELIVERYFAILURE CMDStatus = 0x000000FE // Delivery Failure (used for data_sm_resp)
	ESME_RUNKNOWNERR      CMDStatus = 0x000000FF // Unknown Error
)

const (
	// PDU Types
	GENERIC_NACK          CMDId = 0x80000000
	BIND_RECEIVER         CMDId = 0x00000001
	BIND_RECEIVER_RESP    CMDId = 0x80000001
	BIND_TRANSMITTER      CMDId = 0x00000002
	BIND_TRANSMITTER_RESP CMDId = 0x80000002
	QUERY_SM              CMDId = 0x00000003
	QUERY_SM_RESP         CMDId = 0x80000003
	SUBMIT_SM             CMDId = 0x00000004
	SUBMIT_SM_RESP        CMDId = 0x80000004
	DELIVER_SM            CMDId = 0x00000005
	DELIVER_SM_RESP       CMDId = 0x80000005
	UNBIND                CMDId = 0x00000006
	UNBIND_RESP           CMDId = 0x80000006
	REPLACE_SM            CMDId = 0x00000007
	REPLACE_SM_RESP       CMDId = 0x80000007
	CANCEL_SM             CMDId = 0x00000008
	CANCEL_SM_RESP        CMDId = 0x80000008
	BIND_TRANSCEIVER      CMDId = 0x00000009
	BIND_TRANSCEIVER_RESP CMDId = 0x80000009
	OUTBIND               CMDId = 0x0000000B
	ENQUIRE_LINK          CMDId = 0x00000015
	ENQUIRE_LINK_RESP     CMDId = 0x80000015
	SUBMIT_MULTI          CMDId = 0x00000021
	SUBMIT_MULTI_RESP     CMDId = 0x80000021
	ALERT_NOTIFICATION    CMDId = 0x00000102
	DATA_SM               CMDId = 0x00000103
	DATA_SM_RESP          CMDId = 0x80000103
)

const (

	// SMSC Available Messaging Modes (bits 1-0)
	ESM_MESSAGING_MODE_DEFAULT       = "smsc_mode_default"       // xxxxxx00
	ESM_MESSAGING_MODE_DATAGRAM      = "smsc_mode_datagram"      // xxxxxx01
	ESM_MESSAGING_MODE_FORWARD       = "smsc_mode_forward"       // xxxxxx10
	ESM_MESSAGING_MDOE_STORE_FORWARD = "smsc_mode_store_forward" // xxxxxx11

	// SMSC Supported Message Type (bits 5-2)
	ESM_MESSAGE_TYPE_DEFAULT      = "smsc_message_type_default"    // xx0000xx
	ESM_MESSAGE_TYPE_DELIVERY_ACK = "smsc_message_type_auto_ack"   // xx0010xx
	ESM_MESSAGE_TYPE_MANUAL_ACK   = "smsc_message_type_manual_ack" // xx0100xx

	// SMSC GSM Network Specific Features (bits 7-6)
	ESM_GSM_FEATURE_DEFAULT    = "smsc_gsm_feature_default"    // 00xxxxxx
	ESM_GSM_FEATURE_UDHI       = "smsc_gsm_feature_udhi"       // 01xxxxxx
	ESM_GSM_FEATURE_REPLY      = "smsc_gsm_feature_reply"      // 10xxxxxx
	ESM_GSM_FEATURE_UDHI_REPLY = "smsc_gsm_feature_udhi_reply" // 11xxxxxx

)

var (

	// ESM_MESSAGING_MODES List of available messaging modes
	ESM_MESSAGING_MODES = map[string]string{
		ESM_MESSAGING_MODE_DEFAULT:       "00", // xxxxxx00
		ESM_MESSAGING_MODE_DATAGRAM:      "01", // xxxxxx01
		ESM_MESSAGING_MODE_FORWARD:       "10", // xxxxxx10
		ESM_MESSAGING_MDOE_STORE_FORWARD: "11", // xxxxxx11
	}

	// ESM_MESSAGE_TYPES List of available message types
	ESM_MESSAGE_TYPES = map[string]string{
		ESM_MESSAGE_TYPE_DEFAULT:      "0000", // xx0000xx
		ESM_MESSAGE_TYPE_DELIVERY_ACK: "0010", // xx0010xx
		ESM_MESSAGE_TYPE_MANUAL_ACK:   "0100", // xx0100xx
	}

	ESM_GSM_FEATURES = map[string]byte{
		ESM_GSM_FEATURE_DEFAULT:    0x00, // 00xxxxxx
		ESM_GSM_FEATURE_UDHI:       0x40, // 01xxxxxx
		ESM_GSM_FEATURE_REPLY:      0x80, // 10xxxxxx
		ESM_GSM_FEATURE_UDHI_REPLY: 0xc0, // 11xxxxxx
	}
)

const (
	// FIELDS
	SYSTEM_ID               = "system_id"
	PASSWORD                = "password"
	SYSTEM_TYPE             = "system_type"
	INTERFACE_VERSION       = "interface_version"
	ADDR_TON                = "addr_ton"
	ADDR_NPI                = "addr_npi"
	ADDRESS_RANGE           = "address_range"
	SERVICE_TYPE            = "service_type"
	SOURCE_ADDR_TON         = "source_addr_ton" // type of number (originator)
	SOURCE_ADDR_NPI         = "source_addr_npi" // number plan indicator (originator)
	SOURCE_ADDR             = "source_addr"
	DEST_ADDR_TON           = "dest_addr_ton"
	DEST_ADDR_NPI           = "dest_addr_npi"
	DESTINATION_ADDR        = "destination_addr"
	ESM_CLASS               = "esm_class"
	ESM_MESSAGE_MODE        = "esm_message_mode"
	ESM_MESSAGE_TYPE        = "esm_message_type"
	ESM_GSM_NETWORK_TYPE    = "esm_gsm_network_type"
	PROTOCOL_ID             = "protocol_id"
	PRIORITY_FLAG           = "priority_flag"
	SCHEDULE_DELIVERY_TIME  = "schedule_delivery_time"
	VALIDITY_PERIOD         = "validity_period"
	REGISTERED_DELIVERY     = "registered_delivery"
	REPLACE_IF_PRESENT_FLAG = "replace_if_present_flag"
	DATA_CODING             = "data_coding"
	SM_DEFAULT_MSG_ID       = "sm_default_msg_id"
	SM_LENGTH               = "sm_length"
	SHORT_MESSAGE           = "short_message"
	MESSAGE_ID              = "message_id"
	FINAL_DATE              = "final_date"
	MESSAGE_STATE           = "message_state"
	ERROR_CODE              = "error_code"
)

const (
	// Used for source/destination type of number
	TON_UNKNOWN           = 0 // 00000000
	TON_INTERNATIONAL     = 1 // 00000001
	TON_NATIONAL          = 2 // 00000010
	TON_NETWORK_SPECIFIC  = 3 // 00000011
	TON_SUBSCRIBER_NUMBER = 4 // 00000100
	TON_ALPHANUMERIC      = 5 // 00000101
	TON_ABBREVIATED       = 6 // 00000110

	// Used for source/destination numeric plan indicator NPI
	NPI_UNKNOWN     = 0  // 00000000
	NPI_ISDN        = 1  // 00000001 - E.163 - E.164
	NPI_DATA        = 3  // 00000011 - X.121
	NPI_TELEX       = 4  // 00000100 - F.69
	NPI_LAND_MOBILE = 6  // 00000110 - E.212
	NPI_NATIONAL    = 8  // 00001000
	NPI_PRIVATE     = 9  // 00001001
	NPI_ERMES       = 10 // 00001010
	NPI_IP          = 14 // 00001110
	NPI_WAP         = 18 // 00010010
)
