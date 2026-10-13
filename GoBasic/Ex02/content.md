# 🎯 학습 목표 (Learning Objectives)
Go 컴파일러의 동작 원리를 이해하고, `GOOS`와 `GOARCH` 환경 변수를 활용하여 다양한 플랫폼(Windows, Linux)용 바이너리를 생성하는 **Cross-Compilation(교차 컴파일)** 기술을 습득합니다.

# 💡 핵심 포인트 (Key Points)
- **Compilation Stages**: Lexical Analysis $\to$ Parsing $\to$ Type Checking $\to$ SSA $\to$ Machine Code Generation.
- **Go Toolchain & Cross-Compilation**:
    - **`GOOS` & `GOARCH`**: 대상 운영체제(windows, linux)와 아키텍처(amd64, arm64)를 지정하는 환경 변수 활용법.
    - **Command Reference**: 
        - **Linux용**: `GOOS=linux GOARCH=amdint go build -o build_linux`
        - **Windows용**: `GOOS=windows GOARCH=amd64 go build -o build_windows.exe`
- **Go Tool 활용**: `go tool compile -S` (Assembly Dump) 및 `go tool objdump`를 통한 기계어 수준의 코드 분석.