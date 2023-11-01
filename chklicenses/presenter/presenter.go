package presenter

import (
	"io"

	"github.com/khulnasoft/chk-licenses/chklicenses"
	"github.com/khulnasoft/chk-licenses/chklicenses/presenter/csv"
	"github.com/khulnasoft/chk-licenses/chklicenses/presenter/json"
	"github.com/khulnasoft/chk-licenses/chklicenses/presenter/text"
)

type Presenter interface {
	Present(io.Writer) error
}

func GetPresenter(option Option, results <-chan chklicenses.LicenseResult) Presenter {
	switch option {
	case CSVPresenter:
		return csv.NewPresenter(results)
	case JSONPresenter:
		return json.NewPresenter(results)
	case TextPresenter:
		return text.NewPresenter(results)

	default:
		return nil
	}
}
