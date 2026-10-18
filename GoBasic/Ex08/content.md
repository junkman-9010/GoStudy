# 🎯 학습 목표 (Learning Objectives)
Go 언어의 엄격한 타입 시스템을 이해하고, 정수형 간의 변환(상위 $\to$ 하위, 하위 $\to$ 상위) 및 실수형과 정수형 간의 명시적 형변환 방법을 익힙니다.

# 💡 핵심 포인트 (Key Points)
- **명시적 형변환 (Explicit Conversion)**: `T(v)` 문법을 사용하여 타입을 변환하는 방법.
- **데이터 손실 주의 (Data Loss/Overflow)**:
    - 큰 타입에서 작은 타입으로 변환 시 발생하는 데이터 절단(Truncation) 현상.
    - 실수형에서 정수형으로 변환 시 소수점 아래 자릿수 손실.
- **타입 안전성 (Type Safety)**: Go가 왜 암시적 변환(Implicit Conversion)을 허용하지 않는지에 대한 이유.

# 🛠️ 검증 시나리오 (Verification Scenario)
1. **`example.go` 구현 (Programmer)**:
    - `float64` $\to$ `int` (소수점 손실 확인)
    - `int` (large) $\to$ `int8` (Overflow/Truncation 확인)
    - `float64` $\to$ `int` $\to$ `float64` (반올림이 아닌 버림 발생 확인)
2. **실행 및 검증 (Debugger)**:
    - `go mod vendor` 환경에서 실행 결과가 의도한 데이터 손실 패턴을 출력하는지 확인.
    - Linux/Windows 바이너리 교차 컴phi 컴파일 및 존재 여부 확인.

# 📁 작업 디렉토리 구조 (Directory Structure)
```text
/workspace/GoLang/study/GoBasic/Ex08/
├── content.md            # 학습 목표 및 핵심 포인트
├── example.go            # 타입 변환 실험 코드
├── explanation.md       # 데이터 손실 및 Overflow 심화 분석
├── go.mod                # 모듈 정의
└── (artifacts)           # 컴파일된 바이너리
```