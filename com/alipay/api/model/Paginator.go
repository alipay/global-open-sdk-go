package model

import (

)




type Paginator struct {
        CurrentPage int32 `json:"currentPage,omitempty"`
        PageSize int32 `json:"pageSize,omitempty"`
        TotalPage int32 `json:"totalPage,omitempty"`
        TotalCount int32 `json:"totalCount,omitempty"`
}








