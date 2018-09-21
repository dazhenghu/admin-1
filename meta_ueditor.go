package admin

import (
    "github.com/qor/qor/resource"
    "github.com/qor/qor"
    "github.com/qor/qor/utils"
)

type UeditorConfig struct {
    AssetManager         *Resource
    DisableHTMLSanitizer bool
    Plugins              []UeditorPlugin
    Settings             map[string]interface{}
    CtrlId               string
    metaConfig
}

// RedactorPlugin register redactor plugins into rich editor
type UeditorPlugin struct {
    Name   string
    Source string
}

func (ueditorConfig *UeditorConfig) ConfigureQorMeta(metaor resource.Metaor) {
    if meta, ok := metaor.(*Meta); ok {
        meta.Type = "ueditor"

        if !ueditorConfig.DisableHTMLSanitizer {
            setter := meta.GetSetter()
            meta.SetSetter(func(resource interface{}, metaValue *resource.MetaValue, context *qor.Context) {
                metaValue.Value = utils.HTMLSanitizer.Sanitize(utils.ToString(metaValue.Value))
                setter(resource, metaValue, context)
            })
        }

        if ueditorConfig.Settings == nil {
            ueditorConfig.Settings = map[string]interface{}{}
        }

        plugins := []string{"source"}
        for _, plugin := range ueditorConfig.Plugins {
            plugins = append(plugins, plugin.Name)
        }
        ueditorConfig.Settings["plugins"] = plugins
    }
}
