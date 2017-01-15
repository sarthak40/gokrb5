package asnAppTag

const (
	Ticket = 1
	Authenticator = 2
	EncTicketPart = 3
	ASREQ = 10
	TGSREQ = 12
	ASREP = 11
	TGSREP = 13
	APREQ = 14
	APREP = 15
	KRBSafe = 20
	KRBPriv = 21
	KRBCred = 22
	EncASRepPart = 25
	EncTGSRepPart = 26
	EncAPRepPart = 27
	EncKrbPrivPart  = 28
	EncKrbCredPart = 29
	KRBError = 30
)

//TODO review if we want to consolidate with the MsgTypes in the dictionary