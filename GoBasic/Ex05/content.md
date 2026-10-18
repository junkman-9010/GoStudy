# 🎯 학습 목표 (Learning Objectives)
Go 언어에서 변수를 선언하고 초기화하지 않았을 때 자동으로 할당되는 **제로 값(Zero Value)**의 개념을 이해하고, 각 데이터 타입별 기본값이 무엇인지 확인합니다.

# 💡 핵심 포인트 (Key Points)
- **제로 값의 정의**: 변수 선언 시 개발자가 명시적으로 값을 할당하지 않아도 컴파일러가 보장하는 기본 상태.
- **타입별 제로 값 확인**:
    - `int`, `float64`: `0`
    - `string`: `""` (빈 문자열)
    - `bool`: `false`
    - `pointer`, `slice`, `map`, `interface`, `channel`: `nil`
- **안전성(Safety)**: 제로 값이 프로그래밍 시 발생할 수 있는 '쓰레기 값(Garbage Value)'으로 인한 런타임 에러를 어떻게 방지하는지 이해.

# 🛠️ 검증 시나리오 (Verification Scenario)
1. **`example.go` 구현 (Programmer)**:
    - 다양한 타입(`int`, `string`, `bool`, `float64`)의 변수를 선언하되, **초기화하지 않은 상태**로 작성.
    - `fmt.Println`을 통해 각 변수의 출력값이 예상된 제로 값과 일치하는지 확인.
2. **실행 및 빌드 검증 (Debugger)**:
    - `go run . -mod=vendor`를 통해 실행 결과 확인.
    - `go build`를 통해 Linux/Windows용 바이너리 생성 및 실행 확인.
3. **심화 분석 작성 (Documentation)**:
    - `explanation.md`에 제로 값이 메모리 초기화와 안전성에 미치는 영향 기술.

# 📁 작업 디렉토리 구조 (Directory Structure)
```text
/workspace/GoLang/study/GoBasic/Ex05/
├── content.md            # 학습 목표 및 핵심 포인트
├── example.go            # 제로 값 확인을 위한 실험 코드
├── explanation.md       # 제로 값의 메커니즘 및 안전성 심화 분석
├── go.mod                # 모듈 정의
└── (artifacts)           # 컴파일된 바이너리 (Linux/Windows)
```