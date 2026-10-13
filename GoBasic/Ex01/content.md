# 🎯 학습 목표 (Learning Objectives)
Go 1.12+ 이후 도입된 Go Modules 시스템을 이해하고, `go mod`를 통한 의존성 관리 및 `go mod vendor`를 이용한 프로젝트 격리(Hermetic Build) 방법을 습득합니다.

# 💡 핵심 포인트 (Key Points)
- **Go Modules의 도입**: `GOPATH` 기반 방식에서 탈피하여 프로젝트 단위의 독립적인 개발 환경 구축.
- **`go mod tidy`**: 사용하지 않는 의존성을 제거하고, 필요한 패지키를 `go.mod`에 자동 등록하는 최적화 과정.
- **`go mod vendor`**: 외부 패키지를 프로젝트 내부(`vendor/` 폴더)에 포함시켜 네트워크 연결 없이도 빌드가 가능한 환경(Hermetic Build) 구축.
- **의존성 무결성**: `go.sum`을 통한 패키지 변조 방지 및 보안 검증.