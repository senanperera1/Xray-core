package mobile

import (
    "github.com/xtls/xray-core/core"
)

var instance *core.Instance

// Start Xray with JSON config
func Start(configJSON string) error {
    conf, err := core.LoadConfigFromJson([]byte(configJSON))
    if err != nil {
        return err
    }

    inst, err := core.New(conf)
    if err != nil {
        return err
    }

    instance = inst
    return instance.Start()
}

// Stop Xray
func Stop() error {
    if instance == nil {
        return nil
    }
    err := instance.Close()
    instance = nil
    return err
}
