# 🔍 심화 기술 분석 (Technical Deep Dive): Go First-class Functions

Go 언어에서 함수는 **일급 객체(First-class Object)**로 취급됩니다. 이는 함수가 숫자나 문자열 같은 일반적인 데이터 타입과 동일하게 취급되어, 프로그램의 흐름을 매우 유연하게 제어할 수 있음을 의미합니다.

### 1. 일급 객체(First-class Object)의 3가지 핵심 특징

Go에서 함수가 일급 객체라는 것은 다음 세 가지 작업이 가능함을 의미합니다.

1.  **함수의 변수 할당 (Function Assignment)**:
    함수를 변수에 저장할 수 있습니다. 이를 통해 함수를 마치 데이터처럼 전달하거나 저장할 수 있으며, 실행 시점에 로직을 바인딩할 수 있습니다.
    ```go
    add := func(a, b int) int { return a + b }
    ```

2.  **함수의 인자 전달 (Function as Argument)**:
    다른 함수에 함수의 실행 로직을 인자로 넘길 수 있습니다. 이를 **고차 함수(Higher-order Function)**라고 합니다. 이는 로직의 재사용성을 극대화하며, 특정 전략(Strategy)을 동적으로 주입하는 데 사용됩니다.

3.  **함수의 반환 (Function as Return Value)**:
    함수의 결과값으로 또 다른 함수를 반환할 수 있습니다. 이 과정에서 **클로저(Closure)**가 생성되어 함수 내부의 상태를 캡처하고 유지하는 강력한 패턴을 만들어냅니다.

### 2. 함수 타입 정의 (Function Type Definition)의 이점

함수 시그니처(매개변수와 반환값의 타입)가 복잡해지면 코드를 읽기 어려워집니다. Go에서는 `type` 키워드를 사용하여 함수 타입을 정의함으로써 가독성을 높일 수 있습니다.

*   **Before (복잡함)**: `func filter(nums []int, f func(int) bool) []int`
*   **After (명확함)**: 
    ```go
    type Predicate func(int) bool
    func filter(nums []int, p Predicate) []int
    ```
이렇게 정의된 `Predicate` 타입은 코드의 의도를 명확히 하고, 재사용 가능한 인터페이스 역할을 합니다.

### 3. 실무적 가치: 결합도(Coupling) 낮추기

고차 함수를 사용하면 **의존성 주입(Dependency Injection)**을 매우 가볍게 구현할 수 있습니다. 

*   **예시**: 데이터를 처리하는 함수는 "어떻게 처리할지(로직)"를 직접 알 필요가 없습니다. 단지 필터링 조건(함수)을 인자로 받아 호출하기만 하면 됩니다. 
*   **결과**: 특정 로직이 변경되어도 기존의 `filter` 엔진을 수정할 필요 없이, 새로운 함수(인자)만 넘겨주면 되므로 코드 간의 결합도가 낮아지고 유지보수가 용이해집니다.

---
**핵สาร 키워드**: `First-class Function`, `Higher-order Function`, `Function Type`, `Callback`, `Dependency Injection`