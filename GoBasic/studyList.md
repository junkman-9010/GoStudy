# 🚀 GoLang 마스터 클래스: 상세 학습 커리큘럼

이 문서는 Go 언어의 기초부터 심화 단계까지의 체계적인 학습을 위한 로드맵입니다.

## 🟢 1단계: 기초 (The Fundamentals)
*가장 기본이 되는 문법과 Go의 동작 원리를 이해하는 단계*

### 1.1 Go 환경 구성 및 워크스페이스 (Environment & Workspace)
- [x] 1.1.1 Go SDK 설치 및 환경 변수(`GOROOT`, `GOPATH`) 설정 (dir: GoBasic/Ex00)
- [x] 1.1.2 `go mod`를 이용한 모듈 시스템 이해 (dir: GoBasic/Ex01)
- [x] 1.1.3 Go 컴파일 과정 이해 (dir: GoBasic/Ex02)
- [x] 1.1.4 Hello World 작성 및 패키지(`package main`)의 역할 이해 (dir: GoBasic/Ex03)

### 1.2 변수, 상수 및 기본 타입 (Data Types & Syntax)
- [x] 1.2.1 변수 선언 방식 (`var`, `:=`) (dir: GoBasic/Ex04)
- [x] 1.2.2 제로 값(Zero Values)의 개념과 중요성 (dir: GoBasic/Ex05)
- [x] 1.2.3 기본 타입 마스터: `int`, `float64`, `string`, `bool` (dir: GoBasic/Ex06)
- [x] 1.2.4 상수(`const`) 및 `iota`를 이용한 자동 열거형 생성 (dir: GoBasic/Ex07)
- [x] 1.2.5 타입 변환(Type Conversion)과 명시적 형변환의 필요성 (dir: GoBasic/Ex08)

### 1.3 제어 흐름 (Control Flow)
- [x] 1.3.1 조건문: `if` (조건부 실행), `else if`, `else` (dir: GoBasic/Ex09)
- [x] 1.3.2 `switch` 문: 조건 기반 분기 및 `fallthrough` (dir: GoBasic/Ex10)
- [x] 1.3.3 `for` 루프: Go의 유일한 반복문 (dir: GoBasic/Ex11)
- [x] 1.3.4 `defer` 키워드: 함수 종료 시 실행 보장 (dir: GoBasic/Ex12)

### 1.4 함수와 클로저 (Functions & Closures)
- [x] 1.4.1 함수의 매개변수와 다중 반환값(Multiple Return Values) (dir: GoBasic/Ex13)
- [x] 1.4.2 가변 인자 함수(Variadic Functions) 구현 (dir: GoBasic/Ex14)
- [x] 1.4.3 익명 함수(Anonymous Functions)와 클로저(Closures)의 메모리 동작 (dir: GoBasic/Ex15)
- [x] 1.4.4 함수 타입(First-class Functions)으로서의 특징 (dir: GoBasic/Ex16)

### 1.5 포인터와 메모리 기초 (Pointers & Memory)
- [x] 1.5.1 포인터 연산자: 주소 연산자(`&`)와 역참조 연산자(`*`) (dir: GoBasic/Ex17)
- [x] 1.5.2 포인터 조작 및 추적 (dir: GoBasic/Ex18)
- [x] 1.5.3 Nil 포인터와 안전한 메모리 접근법 (dir: GoBasic/Ex19)
- [ ] 1.5.4 값 전달(Pass by Value) vs 참조 전달(Pass by Pointer)의 비용 비교 (dir: pointers_memory)

### 1.6 구조체와 인터페이스 (Structs & Interfaces)
- [ ] 1.6.1 구조체(`struct`) 정의 및 필드 선언 (dir: structs_interfaces)
- [ ] 1.6.2 구조체 임베딩(Embedding)을 통한 상속 효과 구현 (dir: structs_interfaces)
- [ ] 1.6.3 인터페이스(`interface`) 정의 및 암시적 구현(Implicit Implementation) (dir: structs_interfaces)
- [ ] 1.6.4 타입 단언(Type Assertion) 및 타입 스위치(Type Switch) (dir: structs_interfaces)

### 1.7 컬렉션 데이터 구조 (Collections)
- [ ] 1.7.1 배열(Array)의 고정 크기 특성 이해 (dir: collections_data)
- [ ] 1.7.2 슬라이스(Slice) 마스터: `len`, `cap`, `append`, `re-slicing` (dir: collections_data)
- [ ] 1.7.3 맵(Map) 활용: `make`, 데이터 삽력/삭제, 존재 여부 확인 (dir: collections_data)
- [ ] 1.7.4 `range`를 이용한 반복문 순회 테크닉 (dir: collections_data)

### 1.8 에러 처리 (Error Handling)
- [ ] 1.8.1 `error` 인터페이스 이해 (dir: error_handling)
- [ ] 1.8.2 `panic`과 `recover`를 이용한 예외 처리 (dir: error_handling)
- [ ] 1.8.3 에러 래핑(Error Wrapping) 및 `%w` 사용법 (dir: error_handling)

## 🟡 2단계: 중급 (Concurrency & Engineering)
*Go의 강력한 무기인 동시성을 다루고 실무용 패턴을 익히는 단계*

### 2.1 동시성 기초 (Concurrency Essentials)
- [ ] 2.1.1 고루틴(Goroutines)의 경량 스레드 특징 및 스케줄링 (dir: concurrency_basic)
- [ ] 2.1.2 [unbuffered] 채널(Channels): 데이터 통신 및 동기화 (dir: concurrency_basic)
- [ ] 2.1.3 [buffered] 채널: 버퍼 크기에 따른 블로킹 동작 차이 (dir: concurrency_basic)
- [ ] 2.1.4 `select` 문: 여러 채널의 이벤트를 기다리는 비차단(Non-blocking) 로직 (dir: concurrency_basic)

### 2.2 동시성 심화 (Advanced Concurrency)
- [ ] 2.2.1 `context` 패키지: 취소 신호(Cancel), 타임아웃(Timeout) 전파 (dir: concurrency_advanced)
- [ ] 2.2.2 `sync.WaitGroup`: 고루틴 작업 그룹 동기화 (dir: concurrency_advanced)
	- [ ] 2.2.3 `sync.Mutex` & `sync.RWMutex`: 공유 자원 보호 (dir: concurrency_advanced)
- [ ] 2.2.4 `select` 문 및 Context 패키지 (dir: concurrency_advanced)

### 2.3 실무 동시성 패턴 (Concurrency Patterns)
- [ ] 2.3.1 Worker Pool 패턴 (dir: concurrency_patterns)
- [ ] 2.3.2 Pipeline 패턴 (dir: concurrency_patterns)
- [ ] 2.3.3 Fan-in / Fan-out 패턴 (dir: concurrency_patterns)

### 2.4 모듈 및 테스트 (Modules & Testing)
- [ ] 2.4.1 `go mod`를 이용한 의존성 관리 (dir: module_management)
- [ ] 2.4.2 `testing` 패키지: Unit Test 작성법 (dir: testing_bench)
- [ ] 2.4.3 Table-Driven Tests (dir: testing_bench)
- [ ] 2.4.4 `go benchmark`를 이용한 성능 측정 (dir: testing_bench)

## 🔴 3단계: 고급 및 에코시스템 (Advanced & Ecosystem)
*네트워킹, 데이터베이스, 외부 시스템 통합을 배우는 단계*

### 3.1 리플렉션 및 메타 프로그래밍 (Reflection)
- [ ] 3.1.1 `reflect` 패키지 활용 (dir: reflection_type)
- [ ] 3.1.2 타입 단언 및 타입 스위칭 (dir: reflection_type)

### 3.2 네트워크 및 API (Networking & API)
- [ ] 3.2.1 `net/http`를 이용한 HTTP 서버 구축 (dir: http_server)
- [ ] 3.2.2 HTTP Client 및 REST API 연동 (dir: http_server)

### 3.3 데이터베이스 및 통신 (Data & Communication)
- [ ] 3.3.1 SQL 및 NoSQL 데이터베이스 연동 (dir: database_integration)
- [ ] 3.3.2 gRPC 및 Protobuf 통신 (dir: grpc_proto)

### 3.4 고급 패턴 및 시스템 (Advanced System)
- [ ] 3.4.1 프로덕션용 동시성 패턴 (dir: production_concurrency)
- [ ] 3.4.2 배포 및 도커화 (dir: deployment)
- [ ] 3.4.3 성능 최적화 및 프로파일링 (dir: optimization)

## 🟣 4단계: 시스템 및 특수 목적 (Specialized)
- [ ] 4.1.1 시스템 프로그래밍 기초 (dir: system_programming)
- [ ] 4.1.2 임베디드 및 IoT 개발 (dir: embedded_iot)
- [ ] 4.1.3 고도화된 성능 최적화 구현 (dir: high_performance)
