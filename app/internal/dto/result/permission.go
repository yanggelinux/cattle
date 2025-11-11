package result

import "github.com/yanggelinux/cattle/internal/store/model"

type PermissionRet struct {
	ID          int64  `json:"id"`
	ParentID    int64  `json:"parentID"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Uri         string `json:"uri"`
	Method      string `json:"method"`
	Project     string `json:"project"`
	PermType    int8   `json:"permType"`
	IsEnabled   int8   `json:"isEnabled"`
	Sort        int64  `json:"sort"`
	UpdatedTime string `json:"updatedTime"`
	CreatedTime string `json:"createdTime"`
}

type PermissionResult struct {
	Total   int64          `json:"total"`
	RetList []*ResTreeNode `json:"retList"`
}

type PermTreeNode struct {
	*model.Permission
	Children []*PermTreeNode `json:"children,omitempty"`
}

type ResTreeNode struct {
	*PermissionRet
	Children []*ResTreeNode `json:"children,omitempty"`
}

type PermTreeData struct {
	RolePermIDList  []int64         `json:"rolePermIDList"`
	AllPermTreeData []*PermTreeNode `json:"allPermTreeData"`
}
