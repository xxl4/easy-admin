package service

import (
	"errors"
	"fmt"

	"github.com/nicelizhi/easy-admin-core/sdk/runtime"
	"github.com/nicelizhi/easy-admin-core/sdk/service"

	"gorm.io/gorm"

	"github.com/nicelizhi/easy-admin/app/admin/models"
	"github.com/nicelizhi/easy-admin/app/admin/service/dto"
	"github.com/nicelizhi/easy-admin/common/actions"
	cDto "github.com/nicelizhi/easy-admin/common/dto"
	"github.com/nicelizhi/easy-admin/common/global"
)

type SysApi struct {
	service.Service
}

// GetPage 获取SysApi列表
func (e *SysApi) GetPage(c *dto.SysApiGetPageReq, p *actions.DataPermission, list *[]models.SysApi, count *int64) error {
	var err error
	var data models.SysApi

	orm := e.Orm.Debug().Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		)
	if c.Type != "" {
		qType := c.Type
		if qType == "暂无" {
			qType = ""
		}
		if global.Driver == "postgres" {
			orm = orm.Where("type = ?", qType)
		} else {
			orm = orm.Where("`type` = ?", qType)
		}

	}
	err = orm.Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("Service GetSysApiPage error:%s", err)
		return err
	}
	return nil
}

// Get 获取SysApi对象with id
func (e *SysApi) Get(d *dto.SysApiGetReq, p *actions.DataPermission, model *models.SysApi) *SysApi {
	var data models.SysApi
	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("the object being viewed does not exist or does not have permission to view it")
		e.Log.Errorf("Service GetSysApi error:%s", err)
		_ = e.AddError(err)
		return e
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		_ = e.AddError(err)
		return e
	}
	return e
}

// Update 修改SysApi对象
func (e *SysApi) Update(c *dto.SysApiUpdateReq, p *actions.DataPermission) error {
	var model = models.SysApi{}
	db := e.Orm.Debug().First(&model, c.GetId())
	if db.RowsAffected == 0 {
		return errors.New("do not have permission to update this data")
	}
	c.Generate(&model)
	db = e.Orm.Save(&model)
	if err := db.Error; err != nil {
		e.Log.Errorf("Service UpdateSysApi error:%s", err)
		return err
	}

	return nil
}

// Remove 删除SysApi
func (e *SysApi) Remove(d *dto.SysApiDeleteReq, p *actions.DataPermission) error {
	var data models.SysApi

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveSysApi error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("no right to delete this data")
	}
	return nil
}

// CheckStorageSysApi 创建SysApi对象
func (e *SysApi) CheckStorageSysApi(c *[]runtime.Router) error {
	for _, v := range *c {
		err := e.Orm.Debug().Where(models.SysApi{Path: v.RelativePath, Action: v.HttpMethod}).
			Attrs(models.SysApi{Handle: v.Handler}).
			FirstOrCreate(&models.SysApi{}).Error
		if err != nil {
			err := fmt.Errorf("Service CheckStorageSysApi error: %s \r\n ", err.Error())
			return err
		}
	}
	return nil
}
