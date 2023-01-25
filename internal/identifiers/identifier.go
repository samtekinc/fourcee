package identifiers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Identifier struct {
	ResourceType ResourceType
	Version      IdentifierVersion
	Handle       []byte
}

type IdentifierVersion byte

const (
	IdentifierVersion0             IdentifierVersion = 0
	identifierVersion0HandleLength int               = 8
	CurrentIdentifierVersion       IdentifierVersion = IdentifierVersion0
)

func (i *Identifier) String() string {
	return fmt.Sprintf("%s-%d%x", i.ResourceType, i.Version, i.Handle)
}

func (i *Identifier) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return &types.AttributeValueMemberS{Value: i.String()}, nil
}

func (i *Identifier) UnmarshalDynamoDBAttributeValue(av types.AttributeValue) error {
	avS, ok := av.(*types.AttributeValueMemberS)
	if !ok {
		return nil
	}

	identifier, err := IdentifierFromString(avS.Value)
	if err != nil {
		return err
	}

	*i = *identifier
	return nil
}

func (i *Identifier) MarshalJSON() ([]byte, error) {
	return []byte(i.String()), nil
}

func (i *Identifier) UnmarshalJSON(data []byte) error {
	identifier, err := IdentifierFromString(string(data))
	if err != nil {
		return err
	}
	*i = *identifier
	return nil
}

func IdentifierFromString(s string) (*Identifier, error) {
	parts := strings.SplitN(s, "-", 2)
	if len(parts) != 2 {
		return nil, &InvalidIdentifierError{wrapped: fmt.Errorf("missing '-': %q", s)}
	}

	var resourceType ResourceType
	switch parts[0] {
	case string(ResourceTypeOrgDimension):
		resourceType = ResourceTypeOrgDimension
	default:
		return nil, &InvalidIdentifierError{wrapped: fmt.Errorf("invalid resource type: %q", parts[0])}
	}

	if len(parts[1]) < 1 {
		return nil, &InvalidIdentifierError{wrapped: fmt.Errorf("missing version: %q", s)}
	}

	version := parts[1][0]
	switch version {
	case '0':
		handle := parts[1][1:]
		handleData := make([]byte, identifierVersion0HandleLength)
		i, err := hex.Decode(handleData, []byte(handle))
		if err != nil {
			return nil, &InvalidIdentifierError{wrapped: fmt.Errorf("invalid handle: %q", handle)}
		}
		if i != identifierVersion0HandleLength {
			return nil, &InvalidIdentifierError{wrapped: fmt.Errorf("invalid handle: %q", handle)}
		}
		return &Identifier{
			ResourceType: resourceType,
			Version:      IdentifierVersion0,
			Handle:       handleData,
		}, nil
	default:
		return nil, &InvalidIdentifierError{wrapped: fmt.Errorf("invalid version: %q", version)}
	}
}

func newRandomHandle(version IdentifierVersion) ([]byte, error) {
	switch version {
	case IdentifierVersion0:
		randBytes := make([]byte, identifierVersion0HandleLength)
		_, err := rand.Read(randBytes)
		if err != nil {
			return nil, &IdentifierGeneratorError{wrapped: err}
		}

		return randBytes, nil
	default:
		return nil, &IdentifierGeneratorError{wrapped: fmt.Errorf("invalid version: %q", version)}
	}
}

func NewIdentifier(resourceType ResourceType) (*Identifier, error) {
	handle, err := newRandomHandle(CurrentIdentifierVersion)
	if err != nil {
		return nil, err
	}
	return &Identifier{
		ResourceType: resourceType,
		Version:      CurrentIdentifierVersion,
		Handle:       handle,
	}, nil
}
