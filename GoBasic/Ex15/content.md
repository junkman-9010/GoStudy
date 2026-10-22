# 🎯 학습 목표 (Learning Objectives)
익명 함수의 정의와 사용법을 익히고, 클로저가 어떻게 주변 환경의 변수를 캡처하여 상태를 유지하는지 메커니즘을 실험을 통해 이해합니다.

# 💡 핵심 포인트 (Key Points)
- **익명 함수 (Anonymous Functions)**: 이름 없이 정의되고 변수에 할당되거나 즉시 실행(IIFE)되는 함수의 특징 이해.
- **클로저 (Closures)**: 함수가 선언될 당시의 주변 변수(Lexical Scope)를 캡처하여, 함수 실행 종료 후에도 해당 상태를 유지하는 원리 분석.
- **상태 캡슐화 (State Encapsulation)**: 클로저를 활용하여 전역 변수 없이 특정 데이터의 상태를 안전하게 은인하고 관리하는 패턴 학습.
- **고차 함수 (Higher-Order Functions)**: 함수를 인자로 전달하거나, 함수를 결과로 반환하는 함수의 활용 사례 확인.

# 🛠️ 검증 시나리오 (Verification Scenario)
1. **`example.go` 구현 (Programmer)**:
    - **Case 1 (IIFE)**: 익명 함수를 즉시 실행하여 초기화 로직을 수행하는 패턴 구현.
    - **Case 2 (Function as Variable)**: 익명 함수를 변수에 할당하고 호출하여 함수의 일급 객체 특성 확인.
    - **Case 3 (Basic Closure)**: 외부 카운터 변수를 캡처하여 호출될 때마다 값이 증가하는 `counter` 클로저 구현.
    - **Case 4 (Function Factory)**: 특정 배수를 곱하는 함수를 생성하는 `makeMultiplier(factor int) func(int) int` 패턴 구현.
2. **실행 및 검증 (Debugger)**:
    - `go mod` 및 `vendor` 설정.
    - `go run . -mod=vendor`를 통해 클로저의 상태 유지(State Persistence)가 정상적인지 확인.
    - 교차 컴파일(Linux/Windows) 바이너리 생성 및 검증.
3. **심화 분석 작성 (Documentation)**:
    - `explanation.md`에 클로저의 **Lexical Scoping**과 **Escape Analysis**(힙 할당 유도)에 대해 기술.

# 📁 작업 디렉토리 구조 (Directory Structure)
```text
/workspace/GoLang/study/GoBasic/Ex15/
├── content.md            # 학습 목표 및 핵심 포인트
├── example.go            # 익명 함수 및 클로저 실험 코드
├── explanation.md       # 클로저의 스코프 및 변수 캡처 심화 분석
├── go.mod                # 모듈 정의
└── (artifacts)           # 컴파일된 바이너리
```