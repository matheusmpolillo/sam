package presenter

import (
	"net/http"

	"github.com/goinfinite/os/src/domain/dto"
	internalDbInfra "github.com/goinfinite/os/src/infra/internalDatabase"
	"github.com/goinfinite/os/src/presentation/service"
	uiHelper "github.com/goinfinite/os/src/presentation/ui/helper"
	"github.com/goinfinite/os/src/presentation/ui/page"
	"github.com/labstack/echo/v4"
)

type AccountsPresenter struct {
	accountService *service.AccountService
}

func NewAccountsPresenter(
	persistentDbSvc *internalDbInfra.PersistentDatabaseService,
	trailDbSvc *internalDbInfra.TrailDatabaseService,
) *AccountsPresenter {
	return &AccountsPresenter{
		accountService: service.NewAccountService(persistentDbSvc, trailDbSvc),
	}
}

func (presenter *AccountsPresenter) Handler(c echo.Context) error {
	responseOutput := presenter.accountService.Read(
		map[string]interface{}{
			"shouldIncludeSecureAccessPublicKeys": true,
		},
	)
	if responseOutput.Status != service.Success {
		return nil
	}

	typedOutputBody, assertOk := responseOutput.Body.(dto.ReadAccountsResponse)
	if !assertOk {
		return nil
	}

	pageContent := page.AccountsIndex(typedOutputBody.Accounts)
	return uiHelper.Render(c, pageContent, http.StatusOK)
}
