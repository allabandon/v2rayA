package tools

import (
	"V2RayA/models"
	"io/ioutil"
	"os"
	"os/exec"
)

func RestartV2rayService() (err error) {
	_, err = exec.Command("service", "v2ray", "restart").CombinedOutput()
	if err != nil{
		_, err = exec.Command("systemctl", "restart", "v2ray").Output()
	}
	return
}

func StopV2rayService() (err error) {
	_, err = exec.Command("service", "v2ray", "stop").CombinedOutput()
	if err != nil{
		_, err = exec.Command("systemctl", "stop", "v2ray").Output()
	}
	return
}

//TODO: EnableV2rayService
func DisableV2rayService() (err error) {
	_, err = exec.Command("update-rc.d", "v2ray", "disable").CombinedOutput()
	if err != nil{
		_, err = exec.Command("systemctl", "disable", "v2ray").Output()
	}
	return
}

func WriteV2rayConfig(content []byte) (err error) {
	return ioutil.WriteFile("/etc/v2ray/config.json", content, os.ModeAppend)
}

func IsV2RayRunning() bool {
	out, err := exec.Command("sh", "-c", "service v2ray status|head -n 5|grep running").CombinedOutput()
	if err != nil {
		out, err = exec.Command("sh", "-c", "systemctl status v2ray|head -n 5|grep running").Output()
	}
	return err == nil && len(out) > 0
}
func UpdateV2RayConfig(vmessInfo *models.VmessInfo) (err error) {
	//读配置，转换为v2ray配置并写入
	tmpl := models.NewTemplate()
	err = tmpl.FillWithVmessInfo(*vmessInfo)
	if err != nil {
		return
	}
	err = WriteV2rayConfig(tmpl.ToConfigBytes())
	if err != nil {
		return
	}
	return RestartV2rayService()
}