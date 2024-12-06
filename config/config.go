package config

func Init() {
    ConnectDatabase()
}

//Get new log instace
func GetLogger(p string) *Logger {
    logger := NewLogger(p)
    return logger
}
