package main

type LogFunction func(string) error

type EmailLogger struct {
    LogEmail LogFunction
}

func NewEmailLogger(logEmail LogFunction) *EmailLogger {
    return &EmailLogger{LogEmail: logEmail}
}


