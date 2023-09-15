package soda

import (
	"fmt"
	"io"
)

func (client *Client) GetDataStocks() (*DataStockList, error) {
	stocks := DataStockList{}
	if err := client.get("/datastocks", &stocks); err != nil {
		return nil, err
	} else {
		return &stocks, nil
	}
}

func (client *Client) findDataStock(idOrName string) (*DataStock, error) {
	stocks, err := client.GetDataStocks()
	if err != nil {
		return nil, err
	}
	for i := range stocks.DataStocks {
		s := &stocks.DataStocks[i]
		if s.ShortName == idOrName || s.UUID == idOrName {
			client.stock = s.UUID
			return s, nil
		}
	}
	err = fmt.Errorf("could not find data stock: %s", idOrName)
	return nil, err
}

func (client *Client) ExportDataStock(idOrName string, writer io.Writer) error {
	stock, err := client.findDataStock(idOrName)
	if err != nil {
		return err
	}
	path := "/datastocks/" + stock.UUID + "/export"
	return client.handleBody(path, func(body io.Reader) error {
		_, err := io.Copy(writer, body)
		return err
	})
}

func (client *Client) ExportDataStockCSV(idOrName string, writer io.Writer) error {
	stock, err := client.findDataStock(idOrName)
	if err != nil {
		return err
	}
	path := "/datastocks/" + stock.UUID + "/exportCSV"
	return client.handleBody(path, func(body io.Reader) error {
		_, err := io.Copy(writer, body)
		return err
	})
}
