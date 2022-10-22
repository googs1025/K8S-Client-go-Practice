package main

import (
	"bytes"
	"log"
	"os/exec"
)

// 结合kustomize K8s运维级开发:客户端插件开发实战、命令行可视化 第8课时

func kustomize(path string) string {

	cmd := exec.Command("kustomize", "kuild", path  )
	 var res bytes.Buffer
	cmd.Stdout = &res

	if err := cmd.Run(); err != nil {
		log.Println(err)
		return ""
	}

	return res.String()

}


func main() {

}
