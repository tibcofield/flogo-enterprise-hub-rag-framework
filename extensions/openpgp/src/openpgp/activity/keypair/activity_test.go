package keypair

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ProtonMail/go-crypto/openpgp"
	"github.com/ProtonMail/go-crypto/openpgp/armor"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

const entityName = `MyName`
const entityComment = `some comment`
const entityEmail = `email@localhost`

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&MyActivity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestEval(t *testing.T) {

	act := &MyActivity{}

	tc := test.NewActivityContext(act.Metadata())

	tc.SetInput("name", entityName)
	tc.SetInput("comment", entityComment)
	tc.SetInput("email", entityEmail)
	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	assert.True(t, done)

	var actualPublicKey = tc.GetOutput("publicKey")
	fmt.Println("publicKey    : ", actualPublicKey)

	var actualPrivateKey = tc.GetOutput("privateKey")
	fmt.Println("privatekey    : ", actualPrivateKey)
}


func entityFromPrivateKey(armoredKey string) (*openpgp.Entity, error) {
	keyring, err := openpgp.ReadArmoredKeyRing(strings.NewReader(armoredKey))
	if err != nil {
		return nil, fmt.Errorf("read armored key ring: %w", err)
	}
	if len(keyring) == 0 {
		return nil, fmt.Errorf("no keys found")
	}

	entity := keyring[0]
	if entity.PrivateKey == nil {
		return nil, fmt.Errorf("no private key found — this is a public key only")
	}

	return entity, nil
}


func entityFromPublicKey(armoredKey string) (*openpgp.Entity, error) {
	// Decode the armor block
	block, err := armor.Decode(strings.NewReader(armoredKey))
	if err != nil {
		return nil, fmt.Errorf("armor decode: %w", err)
	}

	// ReadKeyRing handles both single keys and multi-key blocks
	keyring, err := openpgp.ReadKeyRing(block.Body)
	if err != nil {
		return nil, fmt.Errorf("read keyring: %w", err)
	}
	if len(keyring) == 0 {
		return nil, fmt.Errorf("no keys found in armored block")
	}

	return keyring[0], nil
}

