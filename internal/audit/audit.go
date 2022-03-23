package audit

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"os"
	"time"
)

var auditDir string

func Init(d string) {
	auditDir = d
	if _, err := os.Stat(d); os.IsNotExist(err) {
		if err := os.Mkdir(auditDir, 0755); err != nil {
			zap.L().With(zap.Error(err)).Fatal("Failed to create audit directory")
		}
	}

	auditFile := getAuditFile()
	_, err := os.Stat(auditFile)
	if os.IsNotExist(err) {
		file, err := os.Create(auditFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}

	zap.L().With(zap.String("file", auditFile)).Info("Audit Initialised")
}

func getAuditFile() string {
	return fmt.Sprintf("%s/%s.log", auditDir, time.Now().Format("2006-02-01"))
}
