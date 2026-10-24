# 🎯 학습 목표 (Learning Objectives)
포인터를 이용한 변수 값의 직접적인 수정 과정을 추적하고, 함수 호출 시 포인터 전달(Pass-by-pointer)이 원본 데이터에 미치는 영향을 실험하여 메모리 상의 데이터 흐름을 이해합니다.

# 💡 핵심 포인트 (Key Points)
- **포인터 추적 (Pointer Tracing)**: 변수의 주소값이 변하지 않더라도, 포인터를 통해 해당 주소의 *값*이 어떻게 변하는지 관찰.
- **Pass-by-Pointer (참조에 의한 전달)**: 함수에 값의 복사본이 아닌 주소를 전달하여, 함수 호출 범위를 벗어나서도 원본 변수의 상태를 영구적으로 변경하는 메커int 메커니즘 학습.
- **메모리 주소의 불변성 확인**: 변수가 선언된 후 메모리 주소 자체는 고정되어 있으며, 오직 그 주소에 담긴 내용물(Value)만 바뀐다는 점을 확인.
- **간접 참조의 중첩 (Indirect Reference)**: 포인터 변수 자체의 주소를 또 다른 포인터에 저장하는 기초적인 개념 탐색.

# 🛠️ 검증 시나리오 (Verification Scenario)
1. **`example.go` 구현 (Programmer)**:
    - **Case 1 (Value Trace)**: 변수 `A`의 주소를 출력한 후, 포인터를 통해 값을 변경하고 다시 주소와 값을 출력하여 주소는 그대로인 것을 확인.
    - **Case 2 (Function via Pointer)**: `increment(n *int)` 함수를 구현하여, 함수 호출 후 `main`에 있는 원본 변수 값이 증가했는지 검증.
    - **Case 3 (Pointer to Pointer - 기초)**: `**int` 타입의 개념을 맛보기 위해, 포인터 변수의 주소를 함수에 전달하여 포인터가 가인키는 대상 자체를 바꾸는 실험.
    - **Case 4 (Memory Mapping Visualization)**: 출력 결과에 `[Address] -> [Value]` 형태의 로그를 남겨 메모리 맵을 시각적으로 추적할 수 있도록 구현.
2. **실행 및 검증 (Debugger)**:
    - `go mod` 및 `vendor` 설정.
    - `go run . -mod=vendor`를 통해 함수 호출 전/후의 메모리 주소와 값의 변화가 로그와 일치하는지 검증.
    - 교차 컴파일(Linux/Windows) 바이너리 생성 및 검증.
3. **심화 분석 작성 (Documentation)**:
    - `explanation.md`에 **"왜 포인터를 사용하면 원본을 수정할 수 있는가?"**에 대한 메모리 레이아웃 관점의 분석을 작성.

# 📁 작업 디렉토리 구조 (Directory Structure)
```text
/workspace/GoLang/study/GoBasic/Ex18/
├── content.md            # 학습 목표 및 핵심 포인트
├── example.go            # 포인터 추적 및 함수를 통한 값 수정 실험 코드
├── explanation.md       # 포인터 전달(Pass-by-pointer)의 메커니즘 심화 분석
├── go.mod                # 모듈 정의
└── (artifacts)           # 컴파일된 바이너리
```