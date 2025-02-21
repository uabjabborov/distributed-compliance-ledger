# PKI types

#    plain ones
starport scaffold --module pki type CertificateIdentifier subject subjectKeyId 
starport scaffold --module pki type Certificate pemCert serialNumber issuer authorityKeyId rootSubject rootSubjectKeyId isRoot:bool owner subject subjectKeyId
#starport scaffold --module pki type CertificateInfo subject subjectKeyId serialNumber issuer authorityKeyId rootSubject rootSubjectKeyId isRoot:bool owner 

#    messages
starport scaffold --module pki message ProposeAddX509RootCert cert --signer signer
starport scaffold --module pki message ApproveAddX509RootCert subject subjectKeyId --signer signer
starport scaffold --module pki message AddX509Cert cert --signer signer
starport scaffold --module pki message ProposeRevokeX509RootCert subject subjectKeyId --signer signer
starport scaffold --module pki message ApproveRevokeX509RootCert subject subjectKeyId --signer signer
starport scaffold --module pki message RevokeX509Cert subject subjectKeyId --signer signer
starport scaffold --module pki message RejectAddX509RootCert cert --signer signer

# CRUD data types
starport scaffold --module pki map ApprovedCertificates certs:strings --index subject,subjectKeyId --no-message
starport scaffold --module pki map ProposedCertificate pemCert serialNumber owner approvals:strings --index subject,subjectKeyId --no-message
starport scaffold --module pki map ChildCertificates certIds:strings --index issuer,authorityKeyId --no-message
starport scaffold --module pki map ProposedCertificateRevocation  approvals:strings --index subject,subjectKeyId --no-message
starport scaffold --module pki map RevokedCertificates certs:strings --index subject,subjectKeyId --no-message
starport scaffold --module pki map UniqueCertificate present:bool --index issuer,serialNumber --no-message
starport scaffold --module pki single ApprovedRootCertificates certs:strings --no-message
starport scaffold --module pki single RevokedRootCertificates certs:strings --no-message
starport scaffold --module pki map ApprovedCertificatesBySubject subjectKeyIds:strings --index subject --no-message
starport scaffold --module pki map RejectedCertificate pemCert serialNumber owner approvals:strings --index subject,subjectKeyId --no-message
#starport scaffold --module pki map AllProposedCertificates --index subject,subjectKeyId --no-message

# Allow colons (:) in subject ID part in REST URLs
# TODO: need to copy the generated query.pb.gw.go into the correct folder
protoc -I "proto" -I "third_party/proto" --grpc-gateway_out=logtostderr=true,allow_colon_final_segments=true:. "--gocosmos_out=plugins=interfacetype+grpc,Mgoogle/protobuf/any.proto=github.com/cosmos/cosmos-sdk/codec/types:." proto/pki/query.proto 
