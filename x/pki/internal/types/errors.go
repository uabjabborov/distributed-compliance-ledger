package types

import (
	// nolint:goimports
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	Codespace sdk.CodespaceType = ModuleName

	CodeCertificateAlreadyExists         sdk.CodeType = 401
	CodeCertificateDoesNotExist          sdk.CodeType = 402
	CodeProposedCertificateAlreadyExists sdk.CodeType = 403
	CodeProposedCertificateDoesNotExist  sdk.CodeType = 404
	CodeInappropriateCertificateType     sdk.CodeType = 405
	CodeInvalidCertificate               sdk.CodeType = 406
)

func ErrCertificateAlreadyExists(issuer string, serialNumber string) sdk.Error {
	return sdk.NewError(Codespace, CodeCertificateAlreadyExists,
		fmt.Sprintf("X509 certificate associated with the combination of "+
			"issuer=%v and serialNumber=%v already exists on the ledger", issuer, serialNumber))
}

func ErrCertificateDoesNotExist(subject string, subjectKeyID string) sdk.Error {
	return sdk.NewError(Codespace, CodeCertificateDoesNotExist,
		fmt.Sprintf("No X509 certificate associated with the "+
			"combination of subject=%v and subjectKeyID=%v on the ledger", subject, subjectKeyID))
}

func ErrProposedCertificateAlreadyExists(subject string, subjectKeyID string) sdk.Error {
	return sdk.NewError(Codespace, CodeProposedCertificateAlreadyExists,
		fmt.Sprintf("Proposed X509 root certificate associated with the combination "+
			"of subject=%v and subjectKeyID=%v already exists on the ledger", subject, subjectKeyID))
}

func ErrProposedCertificateDoesNotExist(subject string, subjectKeyID string) sdk.Error {
	return sdk.NewError(Codespace, CodeProposedCertificateDoesNotExist,
		fmt.Sprintf("No proposed X509 root certificate associated "+
			"with the combination of subject=%v and subjectKeyID=%v on the ledger. "+
			"The cerificate either does not exists or already approved.", subject, subjectKeyID))
}

func ErrInappropriateCertificateType(error interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeInappropriateCertificateType, fmt.Sprintf("%v", error))
}

func ErrCodeInvalidCertificate(error interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeInvalidCertificate, fmt.Sprintf("%v", error))
}
