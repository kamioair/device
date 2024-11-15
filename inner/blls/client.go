package blls

import (
	"device/inner/config"
	"device/inner/daos"
	"device/inner/models"
	"encoding/json"
	"fmt"
	"github.com/kamioair/qf/qdefine"
	"strconv"
	"strings"
	"sync"
)

type Client struct {
	lock *sync.RWMutex
}

func NewClient() *Client {
	return &Client{
		lock: &sync.RWMutex{},
	}
}

// NewDeviceCode 生成新的设备码
func (c *Client) NewDeviceCode() (string, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	// 获取指定名称的ID
	id, err := daos.IdDao.GetCondition("type = ?", "ClientId")
	if err != nil {
		return "", err
	}

	// 写入数据库
	if id == nil {
		id = &daos.ClientId{
			Type:  "ClientId",
			Value: config.Config.StartId,
		}
	} else {
		id.Value++
	}
	err = daos.IdDao.Save(id)
	if err != nil {
		return "", err
	}

	// 写入客户端信息表
	info := &daos.ClientInfo{
		DbFull:  qdefine.DbFull{Id: id.Value},
		Name:    "",
		Modules: "",
	}
	err = daos.InfoDao.Save(info)
	if err != nil {
		return "", err
	}

	// 格式化客户端ID并返回
	idStr := fmt.Sprintf("%s%0*d", config.Config.IdPrefix, config.Config.IdLength, id.Value)
	return idStr, nil
}

// GetDeviceList 获取所有客户端列表
func (c *Client) GetDeviceList(key string) ([]models.ClientInfo, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	list, err := daos.InfoDao.GetAll()
	if err != nil {
		return nil, err
	}

	// 将数据库內容转为对外输出的內容
	finals := make([]models.ClientInfo, 0)
	for _, v := range list {
		info := models.ClientInfo{
			ClientId:   v.Id,
			ClientName: v.Name,
		}
		if key == "" {
			finals = append(finals, info)
		} else {
			if strings.Contains(strconv.Itoa(int(v.Id)), key) ||
				strings.Contains(v.Name, key) {
				finals = append(finals, info)
			}
		}
	}
	return finals, nil
}

func (c *Client) KnockDoor(info map[string]string) (any, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	// 将当前连接的模块信息，更新到客户端表中
	id, err := strconv.Atoi(info["DeviceCode"])
	if err == nil {
		find, _ := daos.InfoDao.GetModel(uint64(id))
		if find != nil {
			modules := make([]map[string]string, 0)
			_ = json.Unmarshal([]byte(find.Modules), &modules)

			// 不存在添加，反之更新
			exist := false
			for _, v := range modules {
				if v["Name"] == info["ModuleName"] {
					v["Desc"] = info["ModuleDesc"]
					v["Version"] = info["Version"]
					exist = true
					break
				}
			}
			if exist == false {
				modules = append(modules, map[string]string{
					"Name":    info["ModuleName"],
					"Desc":    info["ModuleDesc"],
					"Version": info["Version"],
				})
			}
			str, _ := json.Marshal(modules)
			find.Modules = string(str)
			// 更新数据库
			_ = daos.InfoDao.Save(find)
		}
	}

	return "", nil
}
