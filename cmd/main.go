package main

import "notepad/pkg/logging"

func main() {

	logger := logging.GetLogger("trace")
	logger.Info("start")
}
