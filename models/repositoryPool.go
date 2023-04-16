package models

import (
	"time"
	"xorm.io/xorm"
)

type RepositoryPool struct {
	Id         int
	Identity   string
	Hash       string
	Name       string
	Ext        string
	Size       int64
	Path       string
	Type       string
	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
	DeleteTime time.Time `xorm:"deleted"`
}

func (p RepositoryPool) TableName() string {
	return "repository_pool"
}

// GetHashByRepositoryPool 根据hash查询文件存储池中是否存在相同的文件
func (p RepositoryPool) GetHashByRepositoryPool(hash string, sql *xorm.Engine) (*RepositoryPool, error) {
	_, err := sql.Where("hash=?", hash).Get(&p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// Insert 保存文件存储池数据
func (p RepositoryPool) Insert(sql *xorm.Engine) (int64, error) {
	return sql.Insert(&p)
}

// GetByIdentity 根据identity查询资源
func (p RepositoryPool) GetByIdentity(identity string, engine *xorm.Engine) (*RepositoryPool, error) {
	_, err := engine.Where("identity = ?", identity).Get(&p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
