package csv

import (
	"encoding/csv"
	"io"

	"github.com/khulnasoft/chk-licenses/chklicenses"
)

type Presenter struct {
	resultStream <-chan chklicenses.LicenseResult
}

func NewPresenter(results <-chan chklicenses.LicenseResult) Presenter {
	return Presenter{
		resultStream: results,
	}
}

func (p Presenter) Present(target io.Writer) error {
	writer := csv.NewWriter(target)
	for result := range p.resultStream {
		if err := writer.Write([]string{result.Library, result.URL, result.Type, result.License}); err != nil {
			return err
		}
	}
	writer.Flush()
	return writer.Error()
}
