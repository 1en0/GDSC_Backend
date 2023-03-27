package service

import (
	"github.com/pkg/errors"
	"hello-run/constant"
	"hello-run/dao"
	"strings"
)

func GetRecordVo(record *dao.Record) *RecordVo {
	if record == nil {
		return nil
	}
	return &RecordVo{
		Id:         record.Id,
		RiskType:   record.RiskType.String(),
		Title:      record.Title,
		Content:    record.Content,
		Screenshot: record.Screenshot,
		CreatedAt:  record.CreatedAt,
	}
}

func GetRecordVoList(recordList []dao.Record) []RecordVo {
	if recordList == nil {
		return nil
	}
	var recordVoList []RecordVo
	for _, record := range recordList {
		recordVoList = append(recordVoList, *GetRecordVo(&record))
	}
	return recordVoList
}

func GetRecordListByRoomId(roomId int64) ([]RecordVo, error) {
	recordList, err := dao.GetRecordListByRoomId(roomId)
	if err != nil {
		return nil, err
	}
	return GetRecordVoList(recordList), nil
}

func Str2RiskType(riskTypeStr string) (constant.RiskType, error) {
	var riskType constant.RiskType
	riskTypeStr = strings.ToLower(riskTypeStr)
	switch riskTypeStr {
	case "low":
		riskType = constant.Low
	case "medium":
		riskType = constant.Medium
	case "high":
		riskType = constant.High
	default:
		return 0, errors.New("Invalid parameter: risk_type")
	}
	return riskType, nil
}

func CreateRecord(userId string, roomId int64, riskTypeStr string, title string, content string, screenshot string) (*RecordVo, error) {
	riskType, err := Str2RiskType(riskTypeStr)
	if err != nil {
		return nil, err
	}
	record, err := dao.CreateRecord(userId, roomId, riskType, title, content, screenshot)
	if err != nil {
		return nil, err
	}
	return GetRecordVo(record), nil
}

func UpdateRecord(recordId int64, riskTypeStr string, title string, content string, screenshot string) (*RecordVo, error) {
	var riskType constant.RiskType
	if len(riskTypeStr) == 0 {
		riskType = constant.Low //Default
	} else {
		var err error
		riskType, err = Str2RiskType(riskTypeStr)
		if err != nil {
			return nil, err
		}
	}
	//check if the record exists
	_, err := dao.GetRecordById(recordId)
	if err != nil {
		return nil, err
	}
	recordId, err = dao.UpdateRecord(recordId, riskType, title, content, screenshot)
	if err != nil {
		return nil, err
	}
	record, _ := dao.GetRecordById(recordId)
	return GetRecordVo(record), nil
}

func DeleteRecord(recordId int64) error {
	_, err := dao.GetRecordById(recordId)
	if err != nil {
		return err
	}
	return dao.DeleteRecordById(recordId)
}
