package main

import (
  "context"
  "fmt"
  "time"
  "log"
  "github.com/jackc/pgx/v5"
  "github.com/xuri/excelize/v2"
)

func main() {
    loc, err := time.LoadLocation("Pacific/Auckland")
    if err != nil {
      log.Fatal(err)
    }

    now := time.Now().In(loc)
    weekday := int(now.Weekday())
    if weekday == 0 {
      weekday = 7
    }
    monday := now.AddDate(0, 0, -(weekday - 1))
    weekStart := time.Date(
      monday.Year(),
      monday.Month(),
      monday.Day(),
      5, 0, 0, 0,
      loc,
    )

    shiftStart := now.Add(-12 * time.Hour)

    beginning := "dayshift"
    if now.Hour() == 5 {
      beginning = "nightshift"
    }

    shiftFileName := fmt.Sprintf(
      "shift-%s-%s-%s.xlsx",
      now.Format("Monday"),
      beginning,
      shiftStart.Format("2-January-2006"),
    )
    weekFileName := fmt.Sprintf(
      "Monday-to-%s-%s-week-start-%s.xlsx",
      now.Format("Monday"),
      beginning,
      weekStart.Format("2-January-2006"),
    )

    runReport(shiftStart, now, shiftFileName)
    runReport(weekStart, now, weekFileName)
}

func runReport(start time.Time, end time.Time, fileName string) {
  fmt.Println("between ", start, " and ", end)
  fmt.Println("saving as ", fileName)
    ctx:= context.Background()
    conn, err := pgx.Connect(ctx,
    "postgres://admin:password123@localhost:5432/eventlog",
  )
  if err != nil {
    log.Fatal(err)
  } 
  defer conn.Close(ctx)

  f := excelize.NewFile()

  for sheetnumber := 1; sheetnumber <= 3; sheetnumber++ {

    query :=fmt.Sprintf(`
      SELECT
      product,
      to_timestamp(entered) AS entered,
      delta_seconds,
      to_timestamp(exited) AS exited
      FROM cbw%dpocket16
      WHERE exited BETWEEN %d AND %d
      ORDER BY exited DESC
    `, sheetnumber, start.Unix(), end.Unix())

    rows, err := conn.Query(ctx, query)
    if err != nil {
      log.Fatal(err)
    }
    defer rows.Close()

    sheet := fmt.Sprintf("sheet%d", sheetnumber)
    f.NewSheet(sheet)
    cols := rows.FieldDescriptions()
    fmt.Println("cols: ", cols)
    headers := make([]string, len(cols))

    for i, c := range cols {
      headers[i] = string(c.Name)
      fmt.Println("writing header: ", headers[i])
      cell, _ := excelize.CoordinatesToCellName(i+1, 1)
      f.SetCellValue(sheet, cell, headers[i])
    }

    rowNum := 2 
    for rows.Next() {
      values, err := rows.Values()
      if err != nil {
	log.Fatal(err)
      }
      for colNum, v := range values {
	cell, _ := excelize.CoordinatesToCellName(colNum+1, rowNum)
	switch t := v.(type) {
	case time.Time:
	  f.SetCellValue(sheet, cell, t)
	  f.SetCellStyle(sheet, cell, cell, styleDateTime(f))
	default:
	  f.SetCellValue(sheet, cell, v)
	}
      }
      // computed columns
      deltaSeconds, ok := values[2].(int32)
      //fmt.Println("deltaSeconds: ", deltaSeconds, ok, values[2])
      //fmt.Printf("type=%T, value=%v\n", values[2], values[2])
      if !ok {
	deltaSeconds = 0
      }
      waitTime := deltaSeconds - 150
      //fmt.Println("waitTime: ", waitTime)
      if waitTime < 0 {
	waitTime = 0
      }
      cell, _ := excelize.CoordinatesToCellName(len(values)+1, rowNum)
      f.SetCellValue(sheet, cell, waitTime)
      // end computed columns

      rowNum++
    }
    if err := rows.Err(); err != nil {
      log.Fatal(err)
    }

    durationCol := findColumn(headers, "delta_seconds")
    if durationCol == -1 {
      log.Fatal("delta_seconds not found")
    }

    fmt.Println("finished")

    // formatting (optional)
    lastDataRow := rowNum -1
    colLetter, _ := excelize.ColumnNumberToName(durationCol + 1)
    rangeRef := fmt.Sprintf("%s2:%s%d", colLetter, colLetter, lastDataRow)

    style, err := f.NewConditionalStyle(&excelize.Style{
      Fill: excelize.Fill{
	Type: "pattern",
	Color: []string{"#FFCCCC"},
	Pattern: 1,
      },
    })
    if err != nil {
      log.Fatal(err)
    }
    
    err = f.SetConditionalFormat(sheet, rangeRef, []excelize.ConditionalFormatOptions{
      {
	Type: "cell",
	Criteria: ">",
	Value: "300",
	Format: &style,
      },
    })
    if err != nil {
      log.Fatal(err)
    }
  } 


  if err := f.SaveAs(fileName); err != nil {
    log.Fatal(err)
  }

}

func styleDateTime(f *excelize.File) int {
  styleID, err := f.NewStyle(&excelize.Style{
    NumFmt: 22,
  })
  if err != nil {
    log.Fatal(err)
  }
  return styleID
}

func findColumn(headers []string, name string) int {
  for i, h := range headers {
    if h == name {
      return i
    }
  }
  return -1
}
