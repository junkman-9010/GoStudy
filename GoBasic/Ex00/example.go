package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	fmt.Println("=== Go 환경 설치 및 설정 검증 ===")

	// 1. OS 및 아키텍처 확인
	fmt.Printf("운영체제: %s\n", runtime.GOOS)
	fmt.Printf("아arrchitecture: %s\n", runtime.GOARCH)

	// 2. Go 버전 확인 (go version 명령어 실행)
	cmdVersion := exec.Command("go", "version")
	versionOutput, err := cmdVersion.Output()
	if err != nil {
		fmt.Printf("❌ Go 버전 확인 실패: %v\n", err)
	} else {
		fmt.Printf("Go 버전: %s", string(versionOutput))
	}

	// 3. 환경 변수 확인 (GOROOT, GOPATH, PATH 상의 GO 관련 경로)
	fmt.Println("\n--- 주요 환경 변수 확인 ---")
	
	envVars := []string{"GOROOT", "GOPATH", "PATH"}
	for _, v := range envVars {
		val := os.Getenv(v)
		if val != "" {
			fmt.Printf("%s: %s\n", v, val)
		} else {
			fmt.Printf("%s: (설정되지 않음)\n", v)
		}
	}

	// 4. 최종 검증 결과
	fmt.Println("\n=== 검증 완료 ===")
	if err == nil && os.Getenv("GOROOT") != "" {
		fmt.Println("✅ Go 환경이 성공적으로 인식되었습니다.")
	} else {
		fmt.Println("⚠️  환경 변수 설정을 다시 확인하세요.")
	}
}
