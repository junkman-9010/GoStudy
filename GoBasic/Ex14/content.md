# 🎯 학습 목표 (Learning Objectives)
가변 인자(Variadic Parameters) 함수의 정의 방법과 호출법을 익히고, 슬라이스(Slice)와 가변 인자 간의 관계를 이해하여 유연한 함수 설계 능력을 배양합니다.

# 💡 핵심 포인트 (Key Points)
- **가변 인자 문법 (`...type`)**: 함수 매개변수 앞에 `...`을 붙여 정의하는 방식 이해.
- **슬라이스로의 변환**: 함수 내부에서 가변 인자가 어떻게 `slice` 타입으로 취슬되는지 분석.
- **명시적 슬라이스 전달 vs 가변 인자 전달**: `[]int{...}`와 `1, 2, 3...` 방식의 차이점 확인.
- **인자 개수의 유연성**: 0개의 인자를 전달했을 때의 동작(빈 슬라이스) 확인.
- **매개변수 위치 제약**: 일반 매개변수 뒤에 가변 인자 매개변수가 와야 하는 규칙 학습.

# 🛠️ 검증 시나리오 (Verification Scenario)
1. **`example.go` 구현 (Programmer)**:
    - **Case 1 (Basic)**: 여러 정수를 받아 그 합을 반환하는 `sumAll(nums ...int) int` 구현.
    - **Case 2 (Zero Args)**: 인자를 하나도 전달하지 않았을 때의 결과(0 반int) 확인.
    - **Case 3 (Slice Unpacking)**: 기존에 존재하는 슬라이스를 `...` 연산자를 사용하여 가변 인자 함수에 전달하는 법(`sumAll(mySlice...)`) 구현.
    - **Case 4 (Mixed Params)**: 특정 고정 인자(예: `prefix string`)와 가변 인자를 혼합하여 구성한 함수 구현.
2. **실행 및 검증 (Debugger)**:
    - `go mod` 및 `vendor` 설정.
    - `go run . -mod=vendor`를 통해 각 케이스의 출력값이로 정확한지 확인.
    - 교차 컴파일(Linux/Windows) 바이너리 생성 및 검증.
3. **심화 분석 작성 (Documentation)**:
    - `explanation.md`에 가변 인자가 내부적으로 슬라이스로 동작하는 원리와 `...` 연산자의 역할에 대해 기술.

# 📁 작업 디렉토리 구조 (Directory Structure)
```text
/workspace/GoLang/study/GoBasic/Ex14/
├── content.md            # 학습 목표 및 핵심 포인트
├── example.go            # 가변 인자 및 슬라이스 언패킹 실험 코드
├── explanation.md       # 가변 인자 문법 및 슬라이스 변환 심화 분석
├── go.mod                # 모듈 정의
└── (artifacts)           # 컴파일된 바이너리
```