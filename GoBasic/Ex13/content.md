# 🎯 학습 목표 (Learning Objectives)
함수에 매개변수를 전달하는 다양한 방법과, Go 특유의 다중 반환값을 활용하여 로직을 분리하고 에러를 안정적으로 전달하는 방법을 익힙니다.

# 💡 핵심 포인트 (Key Points)
- **함수 정의와 호출**: `func Name(param type) returnType` 구조 이해.
- **다중 반환값 (Multiple Return Values)**: `(int, error)` 또는 `(string, bool)`과 같이 여러 값을 동시에 반환하여 상태와 데이터를 함께 전달하는 패턴.
- **매개변수 전달 방식 (Pass by Value)**: Go의 기본 함수 호출은 인자의 **복사본**을 전달한다는 점을 실험을 통해 확인.
- **에러 반환 패턴**: `return nil, err`와 같이 에러를 반환값의 두 번째 요소로 사용하는 Go의 표준 에러 처리 관례(Convention) 학습.

# 🛠️ 검증 시나리오 (Verification Scenario)
1. **`example.go` 구현 (Programmer)**:
    - **Case 1 (Basic)**: 기본적인 정수 덧셈 함수와 단일 반환값 확인.
    - **Case 2 (Multiple Returns)**: 숫자를 입력받아 몫과 나머지를 동시에 반환하는 함수 구현.
    - **Case 3 (Error Pattern Simulation)**: 입력값이 음수일 경우 에러를 함께 반래하는 `(int, error)` 패턴 구현.
    - **Case 4 (Slice Parameter)**: 슬라이스를 매개변수로 받아 합계를 계산하는 함수 구현.
2. **실행 및 검증 (Debugger)**:
    - `go mod` 설정 및 `vendor` 생성.
    - `go run . -mod=vendor`를 통해 다중 반환값이 의도한 대로(몫, 나머지, 에러 상태) 출력되는지 확인.
    - 교차 컴파일(Linux/Windows) 바이너리 생성 및 검증.
3. **심화 분석 작성 (Documentation)**:
    - `explanation.md`에 다중 반환값이 Go의 에러 처리 철학에 어떻게 기여하는지, 그리고 값이 복사되어 전달되는 특성이 미치는 영향에 대해 기술.

# 📁 작업 디렉토리 구조 (Directory Structure)
```text
/workspace/GoLang/study/GoBasic/Ex13/
├── content.md            # 학습 목표 및 핵심 포인트
├── example.go            # 함수 매개변수 및 다중 반환값 실험 코드
├── explanation.md       # 함수 호출 메커니즘 및 다중 반환값 심화 분석
├── go.mod                # 모듈 정의
└── (artifacts)           # 컴파일된 바이너리
```