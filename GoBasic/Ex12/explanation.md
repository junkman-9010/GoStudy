# 🔍 심화 기술 분석 (Technical Deep Dive): Go `defer` Statement

Go 언어의 `defer` 문은 함수의 종료 시점에 특정 작업을 예약하는 강력한 도구입니다. 함수의 반환(return)이나 패닉(panic) 발생 시 실행을 보장하므로, 자원 해제 및 에러 처리 로직의 안정성을 확보하는 데 핵심적인 역할을 합니다.

### 1. 핵심 동작 메커니즘 (Core Mechanisms)

#### ① LIFO (Last-In, First-Out) 구조: 스택(Stack) 방식의 실행
`defer` 문은 호출될 때마다 실행 스택에 쌓입니다. 함수가 종료될 때 스택의 가장 위에 있는(즉, 가장 최근에 선록된) `defer`부터 차례대로 실행됩니다. 이는 마치 접시를 쌓아 올린 후 위에서부터 하나씩 꺼내는 것과 같습니다.

**실험 사례:**
```go
defer fmt.Println("First")
defer fmt.Println("Second")
// 출력 결과:
// Second
// First
```

#### ② 인자 평가 (Argument Evaluation)
`defer` 문이 호출되는 시점에 전달된 인자(Arguments)들은 **즉시 평가(Evaluate)**되어 고정됩니다. `defer`가 실제로 실행되는 시점의 변수 값이 아니라, `defer` 문이 작성된 그 시점의 값이 복사되어 저장됩니다. 이는 프로그래머가 의도하지 않은 부수 효과(Side effect)를 방지하는 데 매우 중요합니다.

**실험 사례:**
```go
x := 1
defer fmt.Println("Value of x:", x)
x = 10
// 출력 결과:
// Value of x: 1 (호출 시점의 x 값이 고정됨)
```

### 2. 실무 활용 패턴 (Practical Usage Patterns)

#### ① 자원 해제 (Resource Cleanup)
파일 핸들, 네트워크 소켓, 뮤텍스(Mutex) 잠금 해제 등 반드시 닫거나 해제해야 하는 자원을 `defer`로 관리하면, 함수 내 어디에서 `return`이 발생하든 자원 누수를 방도할 수 있습니다.

```go
f, err := os.Open("data.txt")
if err != nil {
    return err
}
defer f.Close() // 함수 종료 시 안전하게 파일 닫기 보장
```

#### ② 에러 처리 및 로깅 (Error Handling & Logging)
함수가 종료되기 직전에 에러 상태를 확인하거나, 작업의 시작과 끝을 로그로 남기는 데 활용됩니다.

### 3. 주의사항 (Best Practices & Pitfalls)

- **루프 내 `defer` 사용 주의**: `for` 루프 내부에서 `defer`를 과도하게 사용하면, 루프가 종료될 때까지 자원이 해제되지 않고 스택에 계속 쌓여 메모리 부족이나 파일 핸들 부족 현상을 초래할 수 있습니다. (이 경우 별도의 함수로 분리하여 스코프를 제한해야 합니다.)
- **인자 값 변경에 유의**: 인자 값이 즉시 평가된다는 특성을 이해하지 못하면, 로직 오류를 범할 수 있습니다.

---
**요약:** `defer`는 **LIFO** 방식으로 실행되며, **호출 시점의 인자값**을 저장합니다. 자원 관리의 안전성을 극대화하는 Go의 핵심 문법입니다.