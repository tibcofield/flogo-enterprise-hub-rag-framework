/*
 * Copyright © 2017. TIBCO Software Inc.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */
package encrypt

import (
	"fmt"
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&MyActivity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestEval(t *testing.T) {

	act := &MyActivity{}

	tc := test.NewActivityContext(act.Metadata())

	tc.SetInput("name", "MyName")
	tc.SetInput("comment", "some comment")
	tc.SetInput("email", "email@localhost.com")
	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	assert.True(t, done)
	
	var publicKey = tc.GetOutput("publicKey")
	fmt.Println("publicKey    : ", publicKey)

	var privateKey = tc.GetOutput("publicKey")
	fmt.Println("publicKey    : ",privateKey)
}
