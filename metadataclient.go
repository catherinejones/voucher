package voucher

// MetadataClient is an interface that represents something that communicates
// with the Metadata server.
type MetadataClient interface {
	CanAttest() bool
	NewPayloadBody(ImageData) (string, error)
	GetMetadata(ImageData, MetadataType) ([]Occurrence, error)
	AddAttestationToImage(ImageData, AttestationPayload) (Occurrence, error)
}
