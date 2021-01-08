package main

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func mockLogEmail(text string) error {
    return nil
}

func setTestEmailLogger() *EmailLogger {
    emailLogger := createEmailLogger(mockLogEmail)
    return emailLogger
}

func TestEmailLogging(t *testing.T) {
    // emailLogger := setTestEmailLogger()
    emailLogger := createEmailLogger(mockLogEmail)

    err := emailLogger.LogEmail("text")

    assert.Nil(t, err)
}
