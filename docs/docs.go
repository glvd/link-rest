// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/show": {
            "get": {
                "description": "get all data info from server",
                "produces": [
                    "application/json"
                ],
                "summary": "Show data inf",
                "parameters": [
                    {
                        "type": "string",
                        "description": "give your selected page",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Paginator"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.Media"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.File": {
            "type": "object",
            "properties": {
                "key_hash": {
                    "description": "秘钥",
                    "type": "string"
                },
                "key_path": {
                    "description": "秘钥",
                    "type": "string"
                },
                "m3u8_hash": {
                    "description": "切片Hash",
                    "type": "string"
                },
                "m3u8_index": {
                    "description": "M3U8名",
                    "type": "string"
                },
                "m3u8_path": {
                    "description": "切片地址",
                    "type": "string"
                },
                "poster_hash": {
                    "description": "海报Hash",
                    "type": "string"
                },
                "poster_path": {
                    "description": "海报地址",
                    "type": "string"
                },
                "root_hash": {
                    "description": "跟索引",
                    "type": "string"
                },
                "source_hash": {
                    "description": "原片Hash",
                    "type": "string"
                },
                "source_path": {
                    "description": "原片地址",
                    "type": "string"
                },
                "thumb_hash": {
                    "description": "缩略图Hash",
                    "type": "string"
                },
                "thumb_path": {
                    "description": "缩略图地址",
                    "type": "string"
                }
            }
        },
        "model.Info": {
            "type": "object",
            "properties": {
                "alias": {
                    "description": "别名，片名",
                    "type": "object",
                    "$ref": "#/definitions/model.StringArray"
                },
                "caption": {
                    "description": "字幕",
                    "type": "string"
                },
                "director": {
                    "description": "导演",
                    "type": "string"
                },
                "episode": {
                    "description": "集数",
                    "type": "string"
                },
                "format": {
                    "description": "输出格式：3D，2D,VR(VR格式：Half-SBS：左右半宽,Half-OU：上下半高,SBS：左右全宽)",
                    "type": "string"
                },
                "group": {
                    "description": "分组",
                    "type": "string"
                },
                "index": {
                    "description": "索引",
                    "type": "string"
                },
                "intro": {
                    "description": "简介",
                    "type": "string"
                },
                "language": {
                    "description": "语言",
                    "type": "string"
                },
                "length": {
                    "description": "时长",
                    "type": "string"
                },
                "media_type": {
                    "description": "类型：film，FanDrama",
                    "type": "string"
                },
                "producer": {
                    "description": "生产商",
                    "type": "string"
                },
                "publisher": {
                    "description": "发行商",
                    "type": "string"
                },
                "release_date": {
                    "description": "发行日期",
                    "type": "string"
                },
                "role": {
                    "description": "主演",
                    "type": "object",
                    "$ref": "#/definitions/model.StringArray"
                },
                "sample": {
                    "description": "样板图",
                    "type": "object",
                    "$ref": "#/definitions/model.StringArray"
                },
                "season": {
                    "description": "季",
                    "type": "string"
                },
                "series": {
                    "description": "系列",
                    "type": "string"
                },
                "sharpness": {
                    "description": "清晰度",
                    "type": "string"
                },
                "systematics": {
                    "description": "分级",
                    "type": "string"
                },
                "tags": {
                    "description": "标签",
                    "type": "object",
                    "$ref": "#/definitions/model.StringArray"
                },
                "total_episode": {
                    "description": "总集数",
                    "type": "string"
                },
                "uncensored": {
                    "description": "有码,无码",
                    "type": "boolean"
                },
                "video_no": {
                    "description": "编号",
                    "type": "string"
                }
            }
        },
        "model.Media": {
            "type": "object",
            "properties": {
                "file": {
                    "type": "object",
                    "$ref": "#/definitions/model.File"
                },
                "info": {
                    "type": "object",
                    "$ref": "#/definitions/model.Info"
                }
            }
        },
        "model.Paginator": {
            "type": "object",
            "properties": {
                "current_page": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "first_page_url": {
                    "type": "string"
                },
                "from": {
                    "type": "integer"
                },
                "last_page": {
                    "type": "integer"
                },
                "last_page_url": {
                    "type": "string"
                },
                "next_page_url": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "per_page": {
                    "type": "integer"
                },
                "prev_page_url": {
                    "type": "string"
                },
                "to": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "model.StringArray": {
            "type": "array",
            "items": {
                "type": "string"
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:18080",
	BasePath:    "/api/v0",
	Schemes:     []string{},
	Title:       "Swagger Example API",
	Description: "This is a sample server Petstore server.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
