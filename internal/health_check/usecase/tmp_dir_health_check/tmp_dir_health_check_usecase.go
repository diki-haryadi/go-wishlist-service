package tmpDirHealthCheckUseCase

import (
	"path/filepath"
	"runtime"

	"golang.org/x/sys/unix"

	healthCheckDomain "github.com/diki-haryadi/go-micro-template/internal/health_check/domain"
	"github.com/diki-haryadi/ztools/config"
)

type useCase struct{}

func NewUseCase() healthCheckDomain.TmpDirHealthCheckUseCase {
	return &useCase{}
}

func (uc *useCase) Check() bool {
	if !config.IsProdEnv() {
		return true
	}

	_, callerDir, _, ok := runtime.Caller(0)
	if !ok {
		return false
	}

	tmpDir := filepath.Join(filepath.Dir(callerDir), "../../../..", "tmp")

	return unix.Access(tmpDir, unix.W_OK) == nil
}
