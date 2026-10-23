# 🎯 학습 목표 (Learning Objectives)
포인터 연산자(`&`, `*`)의 문법을 익히고, 변수의 메모리 주소를 확인하며, 포인터를 통해 값의 주소를 참조하고 수정하는 메커니즘을 이해합니다.

# 💡 핵심 포인트 (Key Points)
- **주소 연산자 (`&`)**: 변수의 메모리 주소를 추출하는 방법 학습.
- **역참조 연산자 (`*`)**: 포인터가 가 Menc를 가리키는 메모리 위치의 실제 값에 접근하고 수정하는 방법 학습.
- **포인터 타입 (`*T`)**: 특정 타입의 메모리 주소를 저장할 수 있는 포인터 변수의 선언 및 타입 이해.
- **포인터와 값의 차이**: 변수 자체의 값(Value)과 메모리 주소(Address)의 개념적 차이 명확화.

# 🛠️ 검증 시나리오 (Verification Scenario)
1. **`example.go` 구현 (Programmer)**:
    - **Case 1 (Address Extraction)**: `&` 연산자를 사용하여 정수 변수의 메모리 주소를 출력하고, 출력된 주소가 실제 변수의 위치임을 확인.
    - **Case 2 (Dereferencing)**: 포인터 변수를 생성하고, `*` 연산자를 사용하여 포인터가 가리키는 원본 변수의 값을 변경하는 실험.
    - **Case 3 (Pointer Manipulation)**: 함수에 변수의 주소를 전달하여, 함수 내부에서 외부 변수의 값을 수정하는(Pass-by-pointer) 패턴 구현.
    - **Case 4 (Pointer Type Assignment)**: `int` 타입의 주소를 담는 `*int` 타입 변수를 생성하고, 다른 포인터 변수로부터 주소를 복int하는 실험.
2. **실행 및 검증 (Debugger)**:
    - `go mod` 및 `vendor` 설정.
    - `go run . -mod=vendor`를 통해 주소 값 출력 및 역참조를 통한 값 변경이 의도대로 작동하는지 확인.
    - 교차 컴파일(Linux/Windows) 바이너리 생성 및 검증.
3. **심화 분석 작성 (Documentation)**:
    - `explanation.md`에 **"포인터를 통한 값의 수정"**과 **"메모리 주소 참조"**의 차이점, 그리고 왜 포인터를 사용하는지(효율성, 공유 원천)에 대해 기술.

# 📁 작업 디렉토리 구조 (Directory Structure)
```text
/workspace/GoLang/study/GoBasic/Ex17/
├── content.md            # 학습 목표 및 핵심 포인트
├── example.go            # 포인터 연산자 및 주소 참조 실험 코드
├── explanation.md       # 포인터의 원리 및 메모리 접근 심화 분석
├── go.mod                # 모듈 정의
└── (artifacts)           # 컴파일된 바이너리
```