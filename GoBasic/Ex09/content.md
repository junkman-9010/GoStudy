# 🎯 학습 목표 (Learning Objectives)
Go 언어의 핵심 제어 흐름인 조건문(`if`)의 문법을 익히고, 복잡한 논리 구조를 `if-else if-else` 체인을 통해 구현하는 방법을 학습합니다. 특히 Go만의 독특한 조건문 문법(초기화 문장 포함)을 이해하는 데 중점을 둡니다.

# 💡 핵심 포인트 (Key Points)
- **기본 조건문 구조**: `if condition { ... }`의 기본 문법과 중괄호(`{}`) 사용 규칙.
- **`if-else if-else` 체인**: 여러 조건을 순차적으로 검차하여 적절한 실행 블록을 선택하는 로직.
- **조건문 내 초기화 문장 (If with Initialization)**: 
    - `if initialization; condition { ... }` 패턴.
    - 조건문 블록 안에서만 유효한 변수 범위를 활용한 깔끔한 코드 작성법.
- **논리 연산자 활용**: `&&` (AND), `||` (OR), `!` (NOT)을 사용한 복합 조건 구성.

# 🛠️ 검증 시나리오 (Verification Scenario)
1. **`example.go` 구현 (Programmer)**:
    - **Case 1**: 단순 `if-else` (예: 숫자의 양수/음수 판별).
    - **Case 2**: `if-else if-else` (예: 점수에 따른 학점 판별: A, B, C, F).
    - **Case 3**: **초기화 문장이 포함된 `if`** (예: 함수 호출 결과를 변수에 담음과 동시에 결과가 `err != nil`인지 확인하는 패턴).
2. **실행 및 검증 (Debugger)**:
    - `go mod` 설정 및 `vendor` 생성.
    - `go run . -mod=vendor`를 실행하여 각 조건 케이스가 의도한 로직대로 출력되는지 확인.
    - Linux/Windows용 교차 컴파일 확인.
3. **심화 분석 작성 (Documentation)**:
    - `explanation.md`에 Go의 `if` 문법이 왜 다른 언어보다 간결하고 안전한지(변수 스크프 제한 측면) 기술.

# 📁 작업 디렉토리 구조 (Directory Structure)
```text
/workspace/GoLang/study/GoBasic/Ex09/
├── content.md            # 학습 목표 및 핵심 포인트
├── example.go            # 조건문 및 초기화 문장 실험 코드
├── explanation.md       # if 문법의 특징 및 스코프 심화 분석
├── go.mod                # 모듈 정의
└── (artifacts)           # 컴파일된 바이너리
```