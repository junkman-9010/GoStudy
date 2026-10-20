# 🎯 학습 목표 (Learning Objectives)
`defer` 키워드의 동작 원리와 스택(Stack) 방식의 실행 순서를 이해하고, 리소스 누수를 방지하는 안전한 코딩 패턴을 익힙니다.

# 💡 핵심 포인트 (Key Points)
- **`defer`의 기본 동작**: 함수가 반환(return)되기 직전에 예약된 함수가 실행되는 메커용.
- **LIFO (Last-In, First-Out) 구조**: 여러 개의 `defer`가 연속 호출되었을 때, 가장 나중에 선언된 `defer`가 가장 먼저 실행되는 스택 구조 이해.
- **매개변수 평가 (Argument Evaluation)**: `defer` 문이 호출되는 시점에 인자(arguments)가 즉시 평가되어 고정되는 특성 분석.
- **에러 처리와의 결합**: 함수의 반환값(error)을 `defer`를 통해 안전하게 처리하거나 로그를 남기는 패턴.
- **주의사항**: 루프 내부에서의 `defer` 사용 시 발생할 수 있는 리소스 지연 해제 문제(메모리/파일 핸들 누수)와 해결책.

# 🛠️ 검증 시나리오 (Verification Scenario)
1. **`example.go` 구현 (Programmer)**:
    - **Case 1**: 단순 `defer`를 이용한 실행 순서 확인 (기본 출력).
    - **Case 2**: 다중 `defer`를 통한 LIFO(스ubstack) 실행 순서 검증.
    - **Case 3**: `defer` 내 인자 평가 테스트.
    - **Case 4**: 실무 패턴 모사 (리소스 해제 로직).
2. **실행 및 검증 (Debugger)**:
    - `go mod` 설정 및 `vendor` 생성.
    - `go run . -mod=vendor`를 통해 LIFO 순서와 인자 고정값이 의도한 대로 출력되는지 검증.
    - 교차 컴파일(Linux/Windows) 바이너리 생성 확인.
3. **심화 분석 작성 (Documentation)**:
    - `explanation.md`에 `defer` 스택 구조와 인자 평가 시점의 중요성을 기술.

# 📁 작업 디렉토리 구조 (Directory Structure)
```text
/workspace/GoLang/study/GoBasic/Ex14/
├── content.md            # 학습 목표 및 핵심 포인트
├── example.go            # defer 스택 및 인자 평가 실험 코드
├── explanation.md       # defer의 동작 원리 및 스택/인자 평가 심화 분석
├── go.mod                # 모듈 정의
└── (artifacts)           # 컴파일된 바이너리
```