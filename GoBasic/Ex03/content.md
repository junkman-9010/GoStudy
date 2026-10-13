# 🎯 학습 목표 (Learning Objectives)
Go 프로그램의 가장 기초적인 시작점인 `package main`의 구조를 이해하고, 소스 코드가 어떻게 실행 가능한 바이너리로 변환되는지(`go run` vs `go build`) 기초 메커니즘을 습득합니다.

# 💡 핵심 포인트 (Key Points)
- **`package main`의 중요성**: 실행 가능한 프로그램(Executable)을 만들기 위한 필수 선언.
- **`import` 구문**: 외부 패키지(예: `fmt`)를 가져와 기능을 확장하는 방법.
- **`func main()`**: 프로그램 실행의 시작점(Entry Point) 역할.
- **실행 방식의 차이**: 
    - `go run`: 컴파일과 실행을 한 번에 수행하는 개발용 방식.
    - `go build`: 영구적인 실행 파일을 생성하는 배인용 방식.