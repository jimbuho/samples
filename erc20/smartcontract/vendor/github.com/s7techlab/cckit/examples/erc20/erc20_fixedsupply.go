package erc20

import (
	// "golang.org/x/crypto/sha3"
	// "crypto/elliptic"
	// "encoding/hex"
	// "math/big"
	// "reflect"
	"fmt"
	// "crypto/sha256"

	"github.com/pkg/errors"
	"github.com/s7techlab/cckit/extensions/owner"
	"github.com/s7techlab/cckit/router"
	p "github.com/s7techlab/cckit/router/param"
)

const SymbolKey = `symbol`
const NameKey = `name`
const TotalSupplyKey = `totalSupply`
const IsInitKey = `isInit`

func NewErc20FixedSupply() *router.Chaincode {
	r := router.New(`erc20fixedSupply`).Use(p.StrictKnown).

		// Chaincode init function, initiates token smart contract with token symbol, name and totalSupply
		Init(invokeInitFixedSupply,p.String(`init`),  p.String(`symbol`), p.String(`name`), p.Int(`totalSupply`)).

		// Get token symbol
		Query(`symbol`, querySymbol).

		// Get token name
		Query(`name`, queryName).

		// Get the total token supply
		Query(`totalSupply`, queryTotalSupply).

		//  get account balance
		// Query(`balanceOf`, queryBalanceOf, p.String(`mspId`), p.String(`certId`)).
		Query(`balanceOf`, queryBalanceOf, p.String(`publicKey`)).

		//Send value amount of tokens
		// Invoke(`transfer`, invokeTransfer, p.String(`toMspId`), p.String(`toCertId`), p.Int(`amount`)).
		Invoke(`transfer`, invokeTransfer, p.String(`toPublicKey`), p.Int(`amount`)).

		// Allow spender to withdraw from your account, multiple times, up to the _value amount.
		// If this function is called again it overwrites the current allowance with _valu
		// Invoke(`approve`, invokeApprove, p.String(`spenderMspId`), p.String(`spenderCertId`), p.Int(`amount`)).
		Invoke(`approve`, invokeApprove, p.String(`spenderPublicKey`), p.Int(`amount`)).

		//    Returns the amount which _spender is still allowed to withdraw from _owner]
		// Query(`allowance`, queryAllowance, p.String(`ownerMspId`), p.String(`ownerCertId`),
		// 	p.String(`spenderMspId`), p.String(`spenderCertId`)).
		Query(`allowance`, queryAllowance, p.String(`ownerPublicKey`), p.String(`spenderPublicKey`)).

		// Send amount of tokens from owner account to another
		// Invoke(`transferFrom`, invokeTransferFrom, p.String(`fromMspId`), p.String(`fromCertId`),
		// 	p.String(`toMspId`), p.String(`toCertId`), p.Int(`amount`))
		Invoke(`transferFrom`, invokeTransferFrom, p.String(`fromPublicKey`), p.String(`toPublicKey`), p.Int(`amount`))

	return router.NewChaincode(r)
}

func invokeInitFixedSupply(c router.Context) (interface{}, error) {
	ownerIdentity, err := owner.SetFromCreator(c)
	if err != nil {
		return nil, errors.Wrap(err, `set chaincode owner`)
	}

	if (c.State().get(IsInitKey)) {
		return ownerIdentity, nil
	} else {
		if err := c.State().Insert(IsInitKey, true); err != nil {
			return nil, err
		}
	}

	// save token configuration in state
	if err := c.State().Insert(SymbolKey, c.ParamString(`symbol`)); err != nil {
		return nil, err
	}

	if err := c.State().Insert(NameKey, c.ParamString(`name`)); err != nil {
		return nil, err
	}

	if err := c.State().Insert(TotalSupplyKey, c.ParamInt(`totalSupply`)); err != nil {
		return nil, err
	}

	//get publicKeyString
	// pKey := ownerIdentity.GetPublicKey()
	// s := reflect.ValueOf(pKey).Elem()
	// curveData := s.Field(0).Interface().(elliptic.Curve)
	// publicKeyBytes := elliptic.Marshal(curveData, s.FieldByName("X").Interface().(*big.Int), s.FieldByName("Y").Interface().(*big.Int))
	// publicKeyString := sha256.Sum256([]byte(""))
	// p := sha3.Sum256(publicKeyBytes)
	// fmt.Println("printing public key: ")
	// fmt.Printf("%x", publicKeyString)
	// fmt.Println("printing public key 3: ")
	// fmt.Printf("%x", p)
	// str := fmt.Sprintf("%x", p)
	// fmt.Println("p is: ", string(p[:]), str)
	// publicKeyBytesHex := make([]byte, hex.EncodedLen(len(publicKeyBytes)))
	// hex.Encode(publicKeyBytesHex, publicKeyBytes)

	// ID as public key
	pKey := ownerIdentity.GetID()
	fmt.Println("pKay: ", pKey)
	// set token owner initial balance
	if err := setBalance(c, pKey, c.ParamInt(`totalSupply`)); err != nil {
		return nil, errors.Wrap(err, `set owner initial balance`)
	}

	return ownerIdentity, nil
}
