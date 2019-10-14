package server

import (
  "github.com/asaskevich/govalidator"
  "github.com/develar/errors"
  "net/http"
  "strconv"
)

func (t *StatsServer) handleReportRequest(request *http.Request) ([]byte, error) {
  urlQuery, err := parseQuery(request)
  if err != nil {
    return nil, err
  }

  product, machine, _, err := getProductAndMachine(urlQuery)
  if err != nil {
    return nil, err
  }

  generatedTimeValue := urlQuery.Get("generatedTime")
  if len(generatedTimeValue) == 0 {
    return nil, NewHttpError(400, "The generatedTime parameter is required")
  } else if !govalidator.IsNumeric(generatedTimeValue) {
    return nil, NewHttpError(400, "The generatedTime parameter must be numeric")
  }

  generatedTime, err := strconv.ParseInt(generatedTimeValue, 10, 64)
  if err != nil {
    return nil, NewHttpError(400, "The generatedTime parameter is not correct")
  }

  var rawReport []byte
  err = t.db.QueryRow("select raw_report from report where product = ? and machine = ? and generated_time = ?", product, machine, generatedTime).Scan(&rawReport)
  if err != nil {
    return nil, errors.WithStack(err)
  }
  return rawReport, nil
}
