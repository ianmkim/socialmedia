package main

import (
    "github.com/Kamva/mgm/v2"
    "time"
    "strconv"
    "strings"
)

func IndexArr(arr []float64, index int) float64{
    return arr[index]
}

func UtsToTime(ts int64) string{
    tm := time.Unix(ts / 1000, 0)
    year, month, day := tm.Date()
    return strconv.Itoa(year)  + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)
}

func ReplaceSpaces(name string) string {
    return strings.ReplaceAll(name, " ", "_")
}

func FindId(model mgm.DefaultModel) string {
    ph := model.IDField.ID.Hex()
    return ph
}
