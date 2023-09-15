package soda

import (
	"net/url"
	"strconv"
)

type Query struct {
	StartIndex   int
	PageSize     int
	Search       bool
	Distributed  bool
	Name         string
	Description  string
	ClassId      string
	Lang         string
	LangFallback bool
	AllVersions  bool
	CountOnly    bool
	Format       string
}

func (query *Query) NextPage() Query {
	return Query{
		StartIndex:   query.StartIndex + query.PageSize,
		PageSize:     query.PageSize,
		Search:       query.Search,
		Distributed:  query.Distributed,
		Name:         query.Name,
		Description:  query.Description,
		ClassId:      query.ClassId,
		Lang:         query.Lang,
		LangFallback: query.LangFallback,
		AllVersions:  query.AllVersions,
		CountOnly:    query.CountOnly,
		Format:       query.Format,
	}
}

func (query *Query) String() string {
	params := url.Values{}
	if query.StartIndex != 0 {
		params.Add("startIndex", strconv.Itoa(query.StartIndex))
	}
	if query.PageSize != 500 {
		params.Add("pageSize", strconv.Itoa(query.PageSize))
	}
	if query.Search {
		params.Add("search", "true")
	}
	if query.Distributed {
		params.Add("distributed", "true")
	}
	if query.LangFallback {
		params.Add("langFallback", "true")
	}
	if query.AllVersions {
		params.Add("allVersions", "true")
	}
	if query.CountOnly {
		params.Add("countOnly", "true")
	}
	if len(query.Name) > 0 {
		params.Add("name", query.Name)
	}
	if len(query.Description) > 0 {
		params.Add("description", query.Description)
	}
	if len(query.ClassId) > 0 {
		params.Add("classId", query.ClassId)
	}
	if len(query.Lang) > 0 {
		params.Add("lang", query.Lang)
	}
	if len(params) > 0 {
		return "?" + params.Encode()
	} else {
		return ""
	}
}

func DefaultQuery() Query {
	return Query{
		StartIndex: 0,
		PageSize:   500,
	}
}
