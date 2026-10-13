# 🚀 GoLang Mastery Repository

Welcome to the **GoLang Mastery Repository**. This repository is designed as a highly structured, automated, and iterative learning environment for mastering the Go programming language. 

The learning process is driven by a **Master Roadmap (`studyList.md`)** and orchestrated by an **Automated Learning Agent (Cron-based)**.

---

## 🛠 Project Architecture

The repository follows a hierarchical structure that scales from high-level curriculum topics down to granular, executable code examples.

### 📂 Directory Hierarchy

```text
/workspace/GoLang/study/
├── studyList.md                  # 📜 The Master Roadmap (The single source of truth)
├── [module_name]/                # 📦 Module folder (e.g., basics_setup, variables_types)
│   └── ExXX/                     # 🧪 Exercise unit (e.g., Ex00, Ex01, ...)
│       ├── content.md            # 📖 Theory & Purpose (What we are learning and why)
│       ├── example.go            # 💻 Implementation (The actual Go code)
│       └── explanation.txt       # 📝 Detailed Technical Breakdown
└── README.md                     # ℹ️ Project Overview (You are here)
```

### 📑 Component Descriptions

1.  **`studyList.md` (The Master Roadmap)**:
    - Contains the complete, step-by-step curriculum.
    - Each item specifies a `dir:` (the module folder) and a completion status `[DONE]`.
    - This file acts as the task queue for the automated agent.

2.  **`[module_name]` (The Learning Stage)**:
    - A directory corresponding to a specific topic in the roadmap (e.g., `basics_setup`).
    - Grouping all exercises related to a single phase of learning.

3.  **`ExXX` (The Exercise Unit)**:
    - The smallest unit of learning, organized by index (`Ex00`, `Ex01`, etc.).
    - Each unit focuses on one specific micro-topic or technical concept.

4.  **`content.md` (The Context)**:
    - Explains the **Objective** of the exercise.
    - Outlines the **Logic** and the technical challenges involved.
    - Servs as the documentation for the programmer/learner.

5.  **`example.go` (The Implementation)**:
    - A functional, clean, and idiable Go program that demonstrates the concept.
    - Must be runnable and pass all local tests.

6.  **`explanation.txt` (The Deep Dive)**:
    - Provides an in-depth technical analysis of the code.
    - Covers edge cases, memory implications, and performance considerations.

---

## 🤖 Automated Learning Workflow

This repository utilizes a **Cron-based Agent** (`Go_Lang_Study_Scheduler`) that performs the following loop every 3 hours:

1.  **Scan**: Reads `studyList.md` to find the first item without a `[DONE]` tag.
2.  **Generate**: 
    - Creates the `ExXX` directory.
    - Writes the theoretical `content.md`.
    - Writes the implementation `example.go`.
    - Writes the technical `explanation.txt`.
3.  **Verify**: Executes the `example.go` code to ensure it runs correctly.
4.  **Commit**: Updates `studyList.md` with the `[DONE]` tag upon successful verification.

---

## 🚀 Getting Started

To manually trigger a learning session or add a new module, simply update the `studyList.md` and ensure the `dir:` path is valid. The agent will take care of the rest.

**Current Status:** 
- 🟢 **Environment**: Configured
- 🟢 **Roadmap**: Defined
- 🟡 **Learning Agent**: Active (Scheduled)
