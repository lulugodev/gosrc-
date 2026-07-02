package compiler

import (
        "os"
        "os/exec"
)

func CC() string {
        if cc := os.Getenv("CC"); cc != "" {
                return cc
        }

        if _, err := exec.LookPath("clang"); err == nil {
                return "clang"
        }

        if _, err := exec.LookPath("gcc"); err == nil {
                return "gcc"
        }

        return "cc"
}

func CXX() string {
        if cxx := os.Getenv("CXX"); cxx != "" {
                return cxx
        }

        if _, err := exec.LookPath("clang++"); err == nil {
                return "clang++"
        }

        if _, err := exec.LookPath("g++"); err == nil {
                return "g++"
        }

        return "c++"
}

func Env() []string {
        return append(
                os.Environ(),
                "CC="+CC(),
                "CXX="+CXX(),
        )
}
