# 🎯 학습 목표 (Learning Objectives)
Go 언어의 `const` 선언 방식을 익히고, `iota` 식별자를 활용하여 연속된 정수 값을 자동으로 생성하는 강력한 패턴을 이해하여 코드의 유지보수성을 높입니다.

# 💡 핵심 포인트 (Key Points)
- **`const` 키워드 사용법**:
    - 런타임이 아닌 컴파일 타임에 결정되는 불변(Immutable) 값의 특성 이해.
    - `var`와 `const`의 근본적인 차이점(메모리 할당 및 재할당 가능 여부).
- **`iota`를 이용한 자동 열거형(Auto-enumeration)**:
    - `iota`의 동작 원리: `const` 블록 내에서 0부터 시작하여 순차적으로 증가하는 카운터 역할.
    - `iota`를 활용한 상태(Status), 역할(Role), 요일(Day) 등 순차적 상수 생성 패턴.
- **상수 그룹화 및 복합 연산**:
    - `const (...)` 블록을 통한 관련 상수들의 논리적 그룹화.
    - `iota + 1` 또는 `iota * 10`과 같이 `iota`를 활용한 수식 연산 기법.

# 🛠️ 검증 시나리오 (Verification Scenario)
1. **`example.go` 구현 (Programmer)**:
    - `const` 블록을 생성하여 1) 단순 상수, 2) `iota`를 이용한 기본 열거형, 3) `iota`에 변형을 가한(예: `iota + 1`) 상수를 구현.
    - `fmt.Printf`를 사용하여 각 상수의 값이 의도한 대로(0, 1, 2...) 출력되는지 확인.
2. **실행 및 빌드 검증 (Debugger)**:
    - `go mod init ex07-const` 및 `go mod vendor` 실행.
    - `go run . -mod=vendor`를 통해 `iota` 패턴의 출력 결과가 수식과 일치하는지 검증.
    - `go build`를 통해 Linux/Windows용 바이너리 생성 및 실행 확인.
3. **심화 분석 작성 (Documentation)**:
    - `explanation.md`에 `iota`의 동작 메커니즘과 `const` 사용 시의 이점(컴파일 타임 최적화)을 심도 있게 기술.

# 📁 작업 디렉토리 구조 (Directory Structure)
```text
/workspace/GoLang/study/GoBasic/Ex07/
├── content.md            # 학습 목표 및 핵심 포인트
├── example.go            # iota 및 const 활용 코드
├── explanation.md       # iota 메커니즘 및 상수 활용법 심화 분석
├── go.mod                # 모듈 정의
└── (artifacts)           # 컴파일된 바이너리 (Linux/Windows)
```