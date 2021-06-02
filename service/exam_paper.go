package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateExam_paper
//@description: 创建Exam_paper记录
//@param: exam_paper model.Exam_paper
//@return: err error

func CreateExam_paper(exam_paper model.Exam_paper) (err error) {
	global.GVA_DB.AutoMigrate(&exam_paper)  //创建数据库

	err = global.GVA_DB.Create(&exam_paper).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteExam_paper
//@description: 删除Exam_paper记录
//@param: exam_paper model.Exam_paper
//@return: err error

func DeleteExam_paper(exam_paper model.Exam_paper) (err error) {
	err = global.GVA_DB.Delete(&exam_paper).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteExam_paperByIds
//@description: 批量删除Exam_paper记录
//@param: ids request.IdsReq
//@return: err error

func DeleteExam_paperByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Exam_paper{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateExam_paper
//@description: 更新Exam_paper记录
//@param: exam_paper *model.Exam_paper
//@return: err error

func UpdateExam_paper(exam_paper model.Exam_paper) (err error) {
	err = global.GVA_DB.Save(&exam_paper).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetExam_paper
//@description: 根据id获取Exam_paper记录
//@param: id uint
//@return: err error, exam_paper model.Exam_paper

func GetExam_paper(id uint) (err error, exam_paper model.Exam_paper) {
	err = global.GVA_DB.Where("id = ?", id).First(&exam_paper).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetExam_paperInfoList
//@description: 分页获取Exam_paper记录
//@param: info request.Exam_paperSearch
//@return: err error, list interface{}, total int64

func GetExam_paperInfoList(info request.Exam_paperSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Exam_paper{})
    var exam_papers []model.Exam_paper
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Paper_title != "" {
        db = db.Where("`paper_title` LIKE ?","%"+ info.Paper_title+"%")
    }
    if info.Paper_status != 0 {
        db = db.Where("`paper_status` = ?",info.Paper_status)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&exam_papers).Error
	return err, exam_papers, total
}