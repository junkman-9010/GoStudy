# 🎯 학습 목표 (Learning Objectives)
Go 언어의 4대 기본 타입(`int`, `float64`, `string`, `bool`)의 사용법을 익히고, 각 타입의 특성에 맞는 데이터 처리 및 `fmt` 패키지의 서식 지정자(Verb)를 사용하여 데이터를 정확하게 출력하는 능력을 기릅니다.

# 💡 핵심 포인트 (Key Points)
- **데이터 타입별 특징 이해**:
    - `int`: 정수형 데이터 처리 및 정수 연산.
    - `float64`: 부동 소수점 데이터의 정밀도 및 소수점 제어.
    - `string`: 문자열 데이터의 생성 및 기본적인 문자열 포맷팅.
    - `bool`: 논리값(`true`, `false`)을 이용한 조건문의 기초.
- **`fmt` 패키지의 서식 지정자(Verbs) 마스터**:
    - `%v`: 기본 포맷 (Default format).
    - `%d`: 정수형 (`int`).
    - `%.2f`: 소수점 자릿수 제한 (`float64`).
    - `%s`: 문자열 (`string`).
    - `%t`: 불리언 (`bool`).
- **데이터 타입 간의 관계**: 기본적인 타입 변환(Type Conversion)의 필요성 인지.

# 🛠️ 검증 시나리오 (Verification Scenario)
1. **`example.go` 구현 (Programmer)**:
    - 각 타입(`int`, `float64`, `string`, `bool`)의 변수를 선언하고 의미 있는 값(예: 나이, 가격, 이름, 활성화 여부)을 할당.
    - `fmt.Printf`를 사용하여 각 변수를 **서로 다른 서식 지정자**로 출력하여 타입별 포맷팅이 정확히 적용되는지 확인.
2. **실행 및 빌드 검증 (Debugger)**:
    - `go mod init ex06-types` 및 `go mod vendor`를 통한 환경 구축.
    - `go run . -mod=vendor`를 통해 출력 결과가 의도한 포맷(예: 소수점 2자리 등)으로 나오는지 검증.
    - `go build`를 통해 Linux 및 Windows용 바이너리 생성 및 실행 확인.
3. **심화 분석 작성 (Documentation)**:
    - `explanation.md`에 각 타입의 메모리 크기(단순 개념)와 서식 지정자 활용 팁을 정리.

# 📁 작업 디렉토리 구조 (Directory Structure)
```text
/workspace/GoLang/study/GoBasic/Ex06/
├── content.md            # 학습 목표 및 핵심 포인트
├── example.go            # 타입별 서식 지정 변수 출력 코드
├── explanation.md       # 데이터 타입별 서식 지정자 및 활용법 심화 분석
├── go.mod                # 모듈 정의
└── (artifacts)           # 컴파일된 바이너리 (Linux/Windows)
```