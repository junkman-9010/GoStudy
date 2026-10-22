# 🎯 학습 목표 (Learning Objectives)
Go 언어에서 함수가 일급 객체(First-class Object)로서 어떻게 값처럼 취급되는지 이해하고, 함수를 인자로 전달하거나 결과로 반환하는 고차 함수(Higher-order Functions)의 활용법을 익힙니다.

# 💡 핵심 포인트 (Key Points)
- **함수 할당 (Function Assignment)**: 함수를 변수에 저장하고 호출하는 기본 구조 이해.
- **고차 함수 (Higher-order Functions)**:
    - **함수를 인자로 받는 함수**: 특정 로직을 주입(Injection)하는 방식 학습.
    - **함수를 반환하는 함수**: 함수를 통해 새로운 로직을 생성하는 패턴 확인.
- **함수 타입 정의 (Function Type Definition)**: `type FilterFunc func(int) bool`과 같이 함수 자체를 새로운 타입으로 정의하여 코드의 가독성과 재사용성을 높이는 방법 학습.
- **콜백(Callback) 패턴**: 로직의 유연성을 극대화하는 콜백 메커니즘 이해.

# 🛠️ 검증 시나리오 (Verification Scenario)
1. **`example.go` 구현 (Programmer)**:
    - **Case 1 (Assignment)**: 익명 함수를 변수에 할당하여 호출하는 기초 테스트.
    - **Case 2 (Function as Argument)**: 정수 슬라이스를 받아 특정 조건(짝수/홀수)에 맞는 요소만 필터링하는 `filter` 함수 구현.
    - **Case 3 (Function as Return Value)**: 함수를 인자로 받아 새로운 로직을 가진 함수를 반환하는 패턴 구현.
    - **Case 4 (Typed Function)**: 함수 타입을 정의하여 `FilterFunc` 형태의 복잡한 시그니처를 깔끔하게 관리하는 예제 구현.
2. **실행 및 검증 (Debugger)**:
    - `go mod` 및 `vendor` 설정.
    - `go run . -mod=vendor`를 통해 필터링 로직과 함수 전달이 의도한 대로 동작하는지 확인.
    - 교차 컴파일(Linux/Windows) 바이너리 생성 및 검증.
3. **심화 분석 작성 (Documentation)**:
    - `explanation.md`에 함수 타입(Function Type) 정의의 이점과 고차 함수가 코드의 결합도(Coupling)를 낮추는 원리에 대해 기술.

# 📁 작업 디렉토리 구조 (Directory Structure)
```text
/workspace/GoLang/study/GoBasic/Ex16/
├── content.md            # 학습 목표 및 핵심 포인트
├── example.go            # 함수 타입 및 고차 함수 실험 코드
├── explanation.md       # 함수 타입 정의 및 고차 함수 활용 심화 분석
├── go.mod                # 모듈 정의
└── (artifacts)           # 컴파일된 바이너리
```