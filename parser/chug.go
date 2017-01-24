package parser

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/lager/chug"
)

func entry(raw []byte) (entry chug.Entry) {
	copiedBytes := make([]byte, len(raw))
	copy(copiedBytes, raw)
	entry = chug.Entry{
		IsLager: false,
		Raw:     copiedBytes,
	}

	rawString := string(raw)
	idx := strings.Index(rawString, "{")
	if idx == -1 {
		return
	}

	var lagerLog lager.LogFormat
	decoder := json.NewDecoder(strings.NewReader(rawString[idx:]))
	err := decoder.Decode(&lagerLog)
	if err != nil {
		return
	}

	entry.Log, entry.IsLager = convertLagerLog(lagerLog)

	return
}

func convertLagerLog(lagerLog lager.LogFormat) (chug.LogEntry, bool) {
	timestamp, err := strconv.ParseFloat(lagerLog.Timestamp, 64)

	if err != nil {
		return chug.LogEntry{}, false
	}

	data := lagerLog.Data

	var logSession string
	dataSession, ok := lagerLog.Data["session"]
	if ok {
		logSession, ok = dataSession.(string)
		if !ok {
			return chug.LogEntry{}, false
		}
		delete(lagerLog.Data, "session")
	}

	return chug.LogEntry{
		Timestamp: time.Unix(0, int64(timestamp*1e9)),
		LogLevel:  lagerLog.LogLevel,
		Source:    lagerLog.Source,
		Message:   lagerLog.Message,
		Session:   logSession,

		Data: data,
	}, true
}
