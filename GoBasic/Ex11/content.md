# 🎯 학습 목표 (Learning Objectives)
Go 언어의 유일한 반복문인 `for` 문의 다양한 활용 패턴(표준형, 조건형, 무한 루프)을 익히고, `range`를 이용한 컬렉션 순회법을 학습합니다.

# 💡 핵심 포인트 (Key Points)
- **표준 `for` 루프 (C-style)**: 초기화, 조건식, 증감식(`i++`)을 사용하는 전통적인 방식.
- **조건 기반 `for` 루프 (While-style)**: 초기화식 없이 조건식만 사용하여 특정 조건이 만족되는 동안 반복하는 방식.
    - Go에는 `while` 키워드가 없으며, `for` 문 하나로 이 기능을 수행합니다.
- **무한 루프 (Infinite Loop)**: 조건식 없이 `for { ... }`로 작성하여 무한히 실행되는 구조와 `break`를 통한 탈출법.
- **`for range` 루턴 (Iteration)**: 슬라이스(Slice), 배열(Array), 맵(Map), 문자열 등을 안전하고 편리하게 순회하는 방법 (인덱스와 값 추출).
- **루프 제어**: `break`(탈출)와 `continue`(다음 반복으로 건너뛰기)의 사용 사례.

# 🛠️ 검증 시나리오 (Verification Scenario)
1. **`example.go` 구현 (Programmer)**:
    - **Case 1**: 표준 `for` 루프를 이용한 숫자 카운팅 (1부터 N까지).
    - **Case 2**: `for` 문을 `while` 루프처럼 사용하여 특정 조건 도달 시까지 반복.
    - **Case 3**: `for range`를 사용하여 슬라이스의 각 요소를 순회하며 출력.
    - **Case 4**: `break`와 `continue`를 포함하여 특정 수치에서 루프를 멈추거나 건너뛰는 로직.
2. **실행 및 검증 (Debugger)**:
    - `go mod` 설정 및 `vendor` 생성.
    - `go run . -mod=vendor`를 통해 모든 반복 패턴이 의도한 횟수와 로직대로 출력되는지 확인.
    - Linux/Windows용 교차 컴파일 및 바이너리 생성 확인.
3. **심화 분석 작성 (Documentation)**:
    - `explanation.md`에 Go가 왜 `while`을 없애고 `for`로 통합했는지(언어의 단순성 측면)와 `range`의 효율성을 기술.

# 📁 작업 디렉토리 구조 (Directory Structure)
```text
/workspace/GoLang/study/GoBasic/Ex11/
├── content.md            # 학습 목표 및 핵심 포인트
├── example.go            # for 루프의 모든 형태 실험 코드
├── explanation.md       # Go 반복문의 단순성과 range 활용 심화 분석
├── go.mod                # 모듈 정의
└── (artifacts)           # 컴파일된 바이너리
```